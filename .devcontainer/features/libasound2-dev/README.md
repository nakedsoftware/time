# ALSA Development Library Feature

This devcontainer feature installs the `libasound2-dev` package, which provides development files for the ALSA (Advanced Linux Sound Architecture) library.

## Usage

Add this feature to your `devcontainer.json`:

```json
{
    "features": {
        "./features/libasound2-dev": {}
    }
}
```

## What it installs

- `libasound2-dev` - Development files for the ALSA library

## Description

The ALSA library provides audio and MIDI functionality to applications. The development package includes headers and libraries needed to compile applications that use ALSA.

This is commonly needed for:
- Audio processing applications
- Music software development
- Sound-related Go packages that use CGO
- Applications that need low-level audio access

## Supported platforms

- Ubuntu/Debian-based containers

## Version

1.0.0