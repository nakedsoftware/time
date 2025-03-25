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
	"github.com/nakedsoftware/time/src/web/service/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultAppPath string = "web"
	defaultHost string = ""
	defaultPort int16  = 80
)

const (
	appPathKey = "app.path"
	hostKey = "server.host"
	portKey = "server.port"
)

const (
	appPathFlag = "app-path"
	hostFlag = "host"
	portFlag = "port"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Starts the Naked Time web server to listen for requests",
	Long: `
The server command starts the web server that serves the Naked time web
application to users.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := &server.Server{
			Context: cmd.Context(),
			AppPath: viper.GetString(appPathKey),
			Host:    viper.GetString(hostFlag),
			Port:    int16(viper.GetInt(portKey)),
		}
		return server.Serve()
	},
}

func init() {
	rootCommand.AddCommand(serverCommand)

	viper.SetDefault(appPathKey, defaultAppPath)
	serverCommand.PersistentFlags().String(
		appPathFlag,
		defaultAppPath,
		"The path to the web application",
	)
	_ = viper.BindPFlag(
		appPathKey,
		serverCommand.PersistentFlags().Lookup(appPathFlag),
	)
	
	viper.SetDefault(hostKey, defaultHost)
	serverCommand.PersistentFlags().String(
		hostFlag,
		defaultHost,
		"The host or network interface to listen for requests on",
	)
	_ = viper.BindPFlag(
		hostKey,
		serverCommand.PersistentFlags().Lookup(hostFlag),
	)

	_ = viper.BindEnv(portKey, "APP_PORT")
	viper.SetDefault(portKey, defaultPort)
	serverCommand.PersistentFlags().Int16P(
		portFlag,
		"p",
		defaultPort,
		"The TCP/IP port to listen for requests on",
	)
	_ = viper.BindPFlag(
		portKey,
		serverCommand.PersistentFlags().Lookup(portFlag),
	)
}
