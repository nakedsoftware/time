// Copyright 2025 Naked Software, LLC
//
// Version: 1.0.0 (March 22, 2025)
//
// This Naked Time License Agreement ("Agreement") is a legal agreement between
// you ("Licensee") and Naked Software, LLC ("Licensor") for the use of the
// Naked Time software product ("Software"). By using the Software, you agree to
// be bound by the terms of this Agreement.
//
// 1. Grant of License
//
// Licensor grants Licensee a non-exclusive, non-transferable, non-sublicensable
// license to use the Software for non-commercial, educational, or other
// non-production purposes. Licensee may not use the Software for any commercial
// purposes without purchasing a commercial license from Licensor.
//
// 2. Commercial Use
//
// To use the Software for commercial purposes, Licensee much purchase a
// commercial license from Licensor. A commercial license allows Licensee to use
// the Software in production environments, build their own version, and add
// custom features or bug fixes. Licensee may not sell the Software or any
// derivative works to others.
//
// 3. Derivative Works
//
// Licensee may create derivative works of the Software for their own use,
// provided that they maintain a valid commercial license. Licensee may not
// sell or distribute derivative works to others. Any derivative works must
// include a copy of this Agreement and retail all copyright notices.
//
// 4. Sharing and Contributions
//
// Licensee may share their changes or bug fixes to the Software with others,
// provided that such changes are made freely available and not sold. Licensee
// is encourage to contribute their bug fixes back to Licensor for inclusion in
// the Software.
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
// Licensor be liable for any claim, damages, or other liability, whether in an
// action of contract, tort, or otherwise, arising from, out of, or in
// connection with the Software or the use or other dealings in the Software.
//
// 8. Limitation of Liability
//
// In no event shall Licensor be liable for any indirect, incidental, special,
// exemplary, or consequential damages (including, but not limited to,
// procurement or substitute goods or services; loss of use, data, or profits;
// or business interruption) however caused and on any theory of liability,
// whether in contract, strict liability, or tort (including negligence or
// otherwise) arising in any way out of the use of the Software, even if advised
// of the possibility of such damage.
//
// 9. Governing Law
//
// This Agreement shall be governed by and construed in accordance with the laws
// of the jurisdiction in which Licensor is located, without regard to its
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

package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

const (
	logFormatKey = "log.format"
	logLevelKey  = "log.level"
)

const (
	jsonOutputFlag = "json-output"
	logLevelFlag   = "log-level"
)

const (
	formatJSON = "json"
	formatText = "text"
)

const (
	levelDebug = "debug"
	levelError = "error"
	levelInfo  = "info"
	levelWarn  = "warn"
)

var (
	jsonOutput bool
)

var rootCommand = &cobra.Command{
	Use:     "timeweb",
	Version: "1.0.0",
	Short:   "Naked Time Web Application Server",
	Long: `
The web application server implements an HTTP service that serves the files
and assets for the Naked Time web user experience and implements the APIs
that the web application will invoke to obtain data or execute application
services and actions.
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var level slog.Level
		logLevel := viper.GetString(logLevelKey)
		switch logLevel {
		case levelDebug:
			level = slog.LevelDebug
		case levelInfo:
			level = slog.LevelInfo
		case levelWarn:
			level = slog.LevelWarn
		case levelError:
			level = slog.LevelError
		default:
			return fmt.Errorf("unknown log level: %s", logLevel)
		}

		handlerOptions := slog.HandlerOptions{
			Level: level,
		}

		var handler slog.Handler
		if jsonOutput {
			handler = slog.NewJSONHandler(os.Stderr, &handlerOptions)
		} else {
			switch viper.GetString(logFormatKey) {
			case formatJSON:
				handler = slog.NewJSONHandler(os.Stderr, &handlerOptions)
			case formatText:
				handler = slog.NewTextHandler(os.Stderr, &handlerOptions)
			default:
				return fmt.Errorf("unknown log format: %s", logFormatKey)
			}
		}

		logger := slog.New(handler)
		slog.SetDefault(logger)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCommand.PersistentFlags().BoolVar(
		&jsonOutput,
		jsonOutputFlag,
		false,
		"output log events in JSON format",
	)

	viper.SetDefault(logFormatKey, formatText)

	viper.SetDefault(logLevelKey, levelInfo)
	rootCommand.PersistentFlags().String(
		logLevelFlag,
		levelInfo,
		"The level of logging output to see (debug, info, warn, error)",
	)
	_ = viper.BindPFlag(
		logLevelKey,
		rootCommand.PersistentFlags().Lookup(logLevelFlag),
	)
}
