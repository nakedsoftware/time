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
