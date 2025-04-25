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

# Find eligible builder and runner images on Docker Hub. We use Ubuntu/Debian
# instead of Alpine to avoid DNS resolution issues in production.
#
# https://hub.docker.com/r/hexpm/elixir/tags?page=1&name=ubuntu
# https://hub.docker.com/_/ubuntu?tab=tags
#
# This file is based on these images:
#
#   - https://hub.docker.com/r/hexpm/elixir/tags - for the build image
#   - https://hub.docker.com/_/debian?tab=tags&page=1&name=bullseye-20250407-slim - for the release image
#   - https://pkgs.org/ - resource for finding needed packages
#   - Ex: hexpm/elixir:1.18.3-erlang-27.3.3-debian-bullseye-20250407-slim
#
ARG DEBIAN_VERSION=bullseye-20250407-slim

ARG RUNNER_IMAGE="debian:${DEBIAN_VERSION}"

FROM ${RUNNER_IMAGE}

RUN apt-get update -y && \
  apt-get install -y libstdc++6 openssl libncurses5 locales ca-certificates \
  && apt-get clean && rm -f /var/lib/apt/lists/*_*

# Set the locale
RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && locale-gen

ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=en_US.UTF-8

WORKDIR "/app"
RUN chown nobody /app

# set runner ENV
ENV MIX_ENV="prod"

# Only copy the final release from the build stage
COPY --chown=nobody:root _build/${MIX_ENV}/rel/time ./

USER nobody

# If using an environment that doesn't automatically reap zombie processes, it is
# advised to add an init process such as tini via `apt-get install`
# above and adding an entrypoint. See https://github.com/krallin/tini for details
# ENTRYPOINT ["/tini", "--"]

CMD ["/app/bin/server"]
