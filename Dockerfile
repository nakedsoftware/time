# Use distroless base image as base
FROM gcr.io/distroless/base-debian12:latest

# Create a non-root user
USER 65534:65534

# Set working directory
WORKDIR /opt/nakedtime

# Copy the appropriate architecture binary
ARG TARGETARCH
COPY --chown=65534:65534 out/time-${TARGETARCH} bin/time

# Set the entrypoint to run the time command
ENTRYPOINT ["bin/time"]