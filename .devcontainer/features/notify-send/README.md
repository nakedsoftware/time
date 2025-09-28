# Desktop Notifications (notify-send)

This feature installs the `libnotify-bin` package, which provides the `notify-send` command for sending desktop notifications.

## Usage

After installation, you can use the `notify-send` command to send desktop notifications:

```bash
# Basic notification
notify-send "Hello" "This is a test notification"

# Notification with icon
notify-send -i info "Information" "This is an info notification"

# Notification with urgency level
notify-send -u critical "Critical" "This is a critical notification"
```

## What's Installed

- `libnotify-bin`: Package containing the `notify-send` utility

## Notes

- This feature requires a desktop environment or notification daemon to display notifications
- In development containers, notifications may not be visible unless X11 forwarding or similar is configured
- The feature is useful for scripts that need to send notifications to the host system

## Package Information

- **Package**: libnotify-bin
- **Description**: sends desktop notifications to a notification daemon
- **Homepage**: https://packages.ubuntu.com/noble/libnotify-bin