#!/bin/bash
#
# Feature: ALSA Development Library
# This script installs libasound2-dev package for ALSA development

set -e

# Update package lists
echo "Updating package lists..."
apt-get update

# Install libasound2-dev
echo "Installing libasound2-dev..."
apt-get install -y libasound2-dev

# Clean up
echo "Cleaning up..."
apt-get autoremove -y
apt-get autoclean

echo "libasound2-dev installation completed successfully!"