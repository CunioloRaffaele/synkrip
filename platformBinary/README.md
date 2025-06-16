# Platform Binaries

This directory is designated for storing platform-specific binaries that are required for the application to function properly.

## Required Binaries

The following binaries should be placed in this directory with the platform suffix:

- `ffmpeg_<platform>` - FFmpeg binary for media processing
- `ffprobe_<platform>` - FFprobe binary for media analysis
- `yt-dlp_<platform>` - YouTube-DL binary for video downloading

## Naming Convention

All binaries must follow the naming convention of `<binary_name>_<platform>`, where:
- `<binary_name>` is one of: ffmpeg, ffprobe, or yt-dlp
- `<platform>` is the platform identifier (e.g., darwin, windows, linux)

## Examples

- For macOS: `ffmpeg_darwin`, `ffprobe_darwin`, `yt-dlp_darwin`
- For Windows: `ffmpeg_windows`, `ffprobe_windows`, `yt-dlp_windows`
- For Linux: `ffmpeg_linux`, `ffprobe_linux`, `yt-dlp_linux`

Make sure that the binaries you place here are compatible with the respective platforms and have executable permissions.