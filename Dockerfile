# Use Ubuntu slim base image for better glibc compatibility
FROM ubuntu:24.04

# Install runtime dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends libasound2t64 ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Create a non-root user
RUN groupadd -r nakedtime && useradd -r -g nakedtime nakedtime

# Set working directory
WORKDIR /opt/nakedtime

# Copy the appropriate architecture binary
ARG TARGETARCH
COPY --chown=nakedtime:nakedtime out/time-${TARGETARCH} bin/time

# Change to non-root user
USER nakedtime

# Set the entrypoint to run the time command
ENTRYPOINT ["bin/time"]