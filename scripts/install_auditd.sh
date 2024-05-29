#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Function to print usage
usage() {
    echo "Usage: $0 [-y]"
    echo "  -y  Assume yes to all prompts and run non-interactively."
    exit 1
}

# Check for non-interactive flag
ASSUME_YES=0
while getopts ":y" opt; do
    case ${opt} in
        y )
            ASSUME_YES=1
            ;;
        \? )
            usage
            ;;
    esac
done

# Update the package list
echo "Updating package list..."
apt-get update

# Install auditd
if [ $ASSUME_YES -eq 1 ]; then
    echo "Installing auditd non-interactively..."
    apt-get install -y auditd audispd-plugins
else
    echo "Installing auditd..."
    apt-get install auditd audispd-plugins
fi

# Enable and start the auditd service
echo "Enabling and starting auditd service..."
systemctl enable auditd
systemctl start auditd

# Check the status of the auditd service
echo "Checking auditd service status..."
systemctl status auditd

echo "auditd has been installed and started successfully."
