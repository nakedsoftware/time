// Copyright 2025 Naked Software, LLC
//
// License Agreement
//
// Version 1.0.0 (September 28, 2025)
//
// This Naked Time License Agreement ("Agreement") is a legal agreement between
// you ("Licensee") and Naked Software, LLC ("Licensor") for the use of the
// Naked Time software product ("Software"). By using the Software, you agree
// to be bound by the terms of this Agreement.
//
// 1. Grant of License
//
// Licensor grants Licensee a non-exclusive, non-transferable,
// non-sublicensable license to use the Software for non-commercial,
// educational, or other non-production purposes. Licensee may not use the
// Software for any commercial purposes without purchasing a commercial license
// from Licensor.
//
// 2. Commercial Use
//
// To use the Software for commercial purposes, Licensee must purchase a
// commercial license from Licensor. A commercial license allows Licensee to
// use the Software in production environments, build their own version, and
// add custom features or bug fixes. Licensee may not sell the Software or any
// derivative works to others.
//
// 3. Derivative Works
//
// Licensee may create derivative works of the Software for their own use,
// provided that they maintain a valid commercial license. Licensee may not
// sell or distribute derivative works to others. Any derivative works must
// include a copy of this Agreement and retain all copyright notices.
//
// 4. Sharing and Contributions
//
// Licensee may share their changes or bug fixes to the Software with others,
// provided that such changes are made freely available and not sold. Licensee
// is encouraged to contribute their bug fixes back to Licensor for inclusion
// in the Software.
//
// 5. Restrictions
//
// Licensee may not:
//
// - Use the Software for any commercial purposes without a valid commercial
//   license.
// - Sell, sublicense, or distribute the Software or any derivative works.
// - Remove or alter any copyright notices or proprietary legends on the
//   Software.
//
// 6. Termination
//
// This Agreement is effective until terminated. Licensor may terminate this
// Agreement at any time if Licensee breaches any of its terms. Upon
// termination, Licensee must cease all use of the Software and destroy all
// copies in their possession.
//
// 7. Disclaimer of Warranty
//
// The Software is provided "as is" without warranty of any kind, express or
// implied, including but not limited to the warranties of merchantability,
// fitness for a particular purpose, and noninfringement. In no event shall
// Licensor be liable for any claim, damages, or other liability, whether in
// an action of contract, tort, or otherwise, arising from, out of, or in
// connection with the Software or the use or other dealings in the Software.
//
// 8. Limitation of Liability
//
// In no event shall Licensor be liable for any indirect, incidental, special,
// exemplary, or consequential damages (including, but not limited to,
// procurement of substitute goods or services; loss of use, data, or profits;
// or business interruption) however caused and on any theory of liability,
// whether in contract, strict liability, or tort (including negligence or
// otherwise) arising in any way out of the use of the Software, even if
// advised of the possibility of such damage.
//
// 9. Governing Law
//
// This Agreement shall be governed by and construed in accordance with the
// laws of the jurisdiction in which Licensor is located, without regard to its
// conflict of law principles.
//
// 10. Entire Agreement
//
// This Agreement constitutes the entire agreement between the parties with
// respect to the Software and supersedes all prior or contemporaneous
// understandings regarding such subject matter.
//
// By using the Software, you acknowledge that you have read this Agreement,
// understand it, and agree to be bound by its terms and conditions.

package pomodoro

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/gosuri/uilive"
)

//go:embed alarm.mp3
var alarmSound []byte

// The default duration of a pomodoro is 25 minutes.
const defaultPomodoroDuration = 25 * time.Minute

// The time remaining should be updated every 250 milliseconds. This ensures
// that the timer display is smooth and responsive without being too
// distracting.
const updateFrequency = 250 * time.Millisecond

// pomodoroCancelled is returned as an error by the timer function if the
// process is terminated or the user cancels by pressing Ctrl-C.
type pomodoroCancelled struct{}

func (e pomodoroCancelled) Error() string {
	return "The pomodoro was cancelled"
}

// Run starts the Pomodoro timer and runs a pomodoro to completion.
func Run(ctx context.Context) (bool, error) {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	// Display the initial time. This is mostly for visual effect so that the
	// user sees the timer start off at the full time briefly before the time
	// starts counting down.
	startTime := time.Now()
	endTime := startTime.Add(defaultPomodoroDuration)
	err := updateTimer(writer, startTime, endTime)
	if err != nil {
		return false, err
	}

	signalContext, signalStop := signal.NotifyContext(
		ctx,
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer signalStop()

	durationContext, cancel := context.WithDeadline(signalContext, endTime)
	err = timer(durationContext, writer, endTime)
	cancel()
	if err != nil {
		if errors.Is(err, pomodoroCancelled{}) {
			_, err := fmt.Fprintln(writer, "The pomodoro was cancelled.")
			if err != nil {
				return false, err
			}

			return false, nil
		}

		return false, err
	}

	err = pomodoroCompleted(writer)
	if err != nil {
		// Return true because the pomodoro completed successfully even though
		// there was an error during the completion notification.
		return true, err
	}

	return true, nil
}

func timer(ctx context.Context, writer *uilive.Writer, endTime time.Time) error {
	c := time.Tick(updateFrequency)
	for next := range c {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if errors.Is(err, context.DeadlineExceeded) {
				return nil
			}

			if errors.Is(err, context.Canceled) {
				return pomodoroCancelled{}
			}

			return err

		default:
			err := updateTimer(writer, next, endTime)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func updateTimer(writer *uilive.Writer, time, endTime time.Time) error {
	timeRemaining := endTime.Sub(time)
	minutes := max(int(timeRemaining.Minutes()), 0)
	seconds := max(int(timeRemaining.Seconds())%60, 0)
	_, err := fmt.Fprintf(
		writer,
		"Pomodoro time remaining: %d:%02d\n",
		minutes,
		seconds,
	)
	return err
}

func pomodoroCompleted(writer *uilive.Writer) error {
	// playAlarm will play the alarm sound asynchronously. It returns a channel
	// that will be closed when the sound is finished playing, and a function
	// that can be called to close the sound stream. While the alarm sound is
	// playing, we can display the completion message and show the desktop
	// notification concurrently.

	done, closeStream, err := playAlarm()
	if err != nil {
		return err
	}

	defer closeStream()

	_, err = fmt.Fprintln(writer, "The pomodoro is complete!")
	if err != nil {
		return err
	}

	err = showDesktopNotification()
	if err != nil {
		return err
	}

	<-done
	return nil
}

func playAlarm() (done chan bool, close func(), err error) {
	done = make(chan bool)

	streamer, format, err := mp3.Decode(
		io.NopCloser(bytes.NewReader(alarmSound)),
	)
	if err != nil {
		return
	}

	close = func() {
		_ = streamer.Close()
	}

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return
	}

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	return
}

func showDesktopNotification() error {
	return beeep.Notify(
		"Pomodoro Complete",
		"Congratulations, your pomodoro is complete! You should now take a break before starting another pomodoro.",
		"",
	)
}
