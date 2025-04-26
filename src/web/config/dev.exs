# Copyright 2025 Naked Software, LLC
#
# Version: 1.0.0 (March 22, 2025)
#
# This Naked Time License Agreement ("Agreement") is a legal agreement between
# you ("Licensee") and Naked Software, LLC ("Licensor") for the use of the
# Naked Time software product ("Software"). By using the Software, you agree to
# be bound by the terms of this Agreement.
#
# 1. Grant of License
#
# Licensor grants Licensee a non-exclusive, non-transferable, non-sublicensable
# license to use the Software for non-commercial, educational, or other
# non-production purposes. Licensee may not use the Software for any commercial
# purposes without purchasing a commercial license from Licensor.
#
# 2. Commercial Use
#
# To use the Software for commercial purposes, Licensee must purchase a
# commercial license from Licensor. A commercial license allows Licensee to use
# the Software in production environments, build their own version, and add
# custom features or bug fixes. Licensee may not sell the Software or any
# derivative works to others.
#
# 3. Derivative Works
#
# Licensee may create derivative works of the Software for their own use,
# provided that they maintain a valid commercial license. Licensee may not
# sell or distribute derivative works to others. Any derivative works must
# include a copy of this Agreement and retail all copyright notices.
#
# 4. Sharing and Contributions
#
# Licensee may share their changes or bug fixes to the Software with others,
# provided that such changes are made freely available and not sold. Licensee
# is encouraged to contribute their bug fixes back to Licensor for inclusion in
# the Software.
#
# 5. Restrictions
#
# Licensee may not:
#
# - Use the Software for any commercial purposes without a valid commercial
#   license.
# - Sell, sublicense, or distribute the Software or any derivative works.
# - Remove or alter any copyright notices or proprietary legends on the
#   Software.
#
# 6. Termination
#
# This Agreement is effective until terminated. Licensor may terminate this
# Agreement at any time if Licensee breaches any of its terms. Upon
# termination, Licensee must cease all use of the Software and destroy all
# copies in their possession.
#
# 7. Disclaimer of Warranty
#
# The Software is provided "as is" without warranty of any kind, express or
# implied, including but not limited to the warranties of merchantability,
# fitness for a particular purpose, and noninfringement. In no event shall
# Licensor be liable for any claim, damages, or other liability, whether in an
# action of contract, tort, or otherwise, arising from, out of, or in
# connection with the Software or the use or other dealings in the Software.
#
# 8. Limitation of Liability
#
# In no event shall Licensor be liable for any indirect, incidental, special,
# exemplary, or consequential damages (including, but not limited to,
# procurement or substitute goods or services; loss of use, data, or profits;
# or business interruption) however caused and on any theory of liability,
# whether in contract, strict liability, or tort (including negligence or
# otherwise) arising in any way out of the use of the Software, even if advised
# of the possibility of such damage.
#
# 9. Governing Law
#
# This Agreement shall be governed by and construed in accordance with the laws
# of the jurisdiction in which Licensor is located, without regard to its
# conflict of law principles.
#
# 10. Entire Agreement
#
# This Agreement constitutes the entire agreement between the parties with
# respect to the Software and supersedes all prior or contemporaneous
# understandings regarding such subject matter.
#
# By using the Software, you acknowledge that you have read this Agreement,
# understand it, and agree to be bound by its terms and conditions.

import Config

# Configure your database
config :time, NakedTime.Repo,
  username: "nakedtime",
  password: "itsMyLittleS@cret123",
  hostname: "postgres",
  database: "time_dev",
  stacktrace: true,
  show_sensitive_data_on_connection_error: true,
  pool_size: 10

# For development, we disable any cache and enable
# debugging and code reloading.
#
# The watchers configuration can be used to run external
# watchers to your application. For example, we can use it
# to bundle .js and .css sources.
config :time, NakedTimeWeb.Endpoint,
  # Bind to 0.0.0.0 to expose the server to the docker host machine.
  # This makes make the service accessible from any network interface.
  # Change to `ip: {127, 0, 0, 1}` to allow access only from the server machine.
  http: [ip: {0, 0, 0, 0}, port: 4000],
  check_origin: false,
  code_reloader: true,
  debug_errors: true,
  secret_key_base: "z0t7bRtrf29bcEWQUtom5Soa3t7ah4J4Bka9lYFlBCskRXk171a/U0fzlYVcgnC9",
  watchers: [
    esbuild: {Esbuild, :install_and_run, [:time, ~w(--sourcemap=inline --watch)]},
    tailwind: {Tailwind, :install_and_run, [:time, ~w(--watch)]}
  ]

# ## SSL Support
#
# In order to use HTTPS in development, a self-signed
# certificate can be generated by running the following
# Mix task:
#
#     mix phx.gen.cert
#
# Run `mix help phx.gen.cert` for more information.
#
# The `http:` config above can be replaced with:
#
#     https: [
#       port: 4001,
#       cipher_suite: :strong,
#       keyfile: "priv/cert/selfsigned_key.pem",
#       certfile: "priv/cert/selfsigned.pem"
#     ],
#
# If desired, both `http:` and `https:` keys can be
# configured to run both http and https servers on
# different ports.

# Watch static and templates for browser reloading.
config :time, NakedTimeWeb.Endpoint,
  live_reload: [
    patterns: [
      ~r"priv/static/(?!uploads/).*(js|css|png|jpeg|jpg|gif|svg)$",
      ~r"priv/gettext/.*(po)$",
      ~r"lib/time_web/(controllers|live|components)/.*(ex|heex)$"
    ]
  ]

# Enable dev routes for dashboard and mailbox
config :time, dev_routes: true

# Do not include metadata nor timestamps in development logs
config :logger, :console, format: "[$level] $message\n"

# Set a higher stacktrace during development. Avoid configuring such
# in production as building large stacktraces may be expensive.
config :phoenix, :stacktrace_depth, 20

# Initialize plugs at runtime for faster development compilation
config :phoenix, :plug_init_mode, :runtime

config :phoenix_live_view,
  # Include HEEx debug annotations as HTML comments in rendered markup
  debug_heex_annotations: true,
  # Enable helpful, but potentially expensive runtime checks
  enable_expensive_runtime_checks: true

# Disable swoosh api client as it is only required for production adapters.
config :swoosh, :api_client, false
