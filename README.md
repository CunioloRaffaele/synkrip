# SynkRip

<p align="center">
    <img src="https://github.com/CunioloRaffaele/synkrip/blob/main/build/appicon.png?raw=true" alt="SynkRip logo" width="300">
</p>

#### A powerful desktop application that seamlessly syncs and manages your music libraries from multiple platforms and services. 🎵 ✨

![OS](https://img.shields.io/badge/OS-MacOs/Windows-cyan?style=for-the-badge)
![language](https://img.shields.io/badge/lang-go-blueviolet?style=for-the-badge)
![License](https://img.shields.io/github/license/CunioloRaffaele/synkrip?style=for-the-badge)
<br>
![LastCommit](https://img.shields.io/github/last-commit/CunioloRaffaele/synkrip/main)
![Commits](https://img.shields.io/github/commit-activity/w/CunioloRaffaele/synkrip)
![Issues](https://img.shields.io/bitbucket/issues/CunioloRaffaele/synkrip?color=yellow)

## Features

- **Automatic Library Management**:<br>Create, open, and scan music libraries for better organization, with automatic synchronization between online playlists and local folders
- **Spotify download and sync**:<br>Fetch playlists and synchronize them with your local library
- **YouTube Music download and sync**:<br>Fetch playlists and synchronize them with your local library 
- **SoundCloud download and sync**:<br>![TODO](https://img.shields.io/badge/status-TODO-red)<br>Fetch playlists and synchronize them with your local library
- **Apple Music download and sync**:<br>![TODO](https://img.shields.io/badge/status-TODO-red)<br>Fetch playlists and synchronize them with your local library
- **Cross-Platform Support**:<br>Works on macOS and Windows
- **User-Friendly Experience**:<br>Unlike command-line tools, SynkRip offers an intuitive graphical interface that makes music library management accessible to everyone

## Installation

1. 📥   **Download the App**:
    - Visit the [Releases Page](https://github.com/CunioloRaffaele/synkrip/releases) to download the latest version available for your operating system.
    <br>**macOS**: Download the `.dmg` file
    <br>**Windows**: Download the `.exe` file

2. 🛠️  **Install the App**:
    - **macOS**: Copy SynkRip from the mounted .dmg into your Applications folder
    - **Windows**: Run the installer and follow the simple setup wizard

3. 🚀  **Launch SynkRip**:
    - Open the app from your Applications folder (macOS) or Start Menu (Windows)
    <br>
> [!WARNING]
> When launching for the first time, you may encounter an "unknown developer" warning. Right-click the app and select "Open" to bypass this warning, or adjust your security settings in System Preferences.
> This happens even if the app is signed since macOS requires apps to be notarized by Apple to bypass the "unknown developer" warning.

## Usage

- **Create a New Library**: Set up a new music library and add playlists that need to be synchronized.
- **Open an Existing Library**: Load a previously created library and manage or update its contents.

## Disclaimers
**LEGAL DISCLAIMER AND TERMS OF USE**

SynkRip is provided for personal, non-commercial use only. By using this application, you acknowledge and agree to the following:

1. **Copyright Compliance**: You are solely responsible for ensuring that you have the legal right to download, store, and use any music files managed through this application. This includes having proper licenses, subscriptions, or ownership rights to the content.

2. **No Endorsement of Piracy**: The developers of SynkRip do not endorse, encourage, or condone piracy, copyright infringement, or any form of intellectual property theft.

3. **User Responsibility**: Users must comply with all applicable local, national, and international copyright laws and regulations. Any unauthorized downloading or distribution of copyrighted material is strictly prohibited.

4. **Service Terms**: Users must also comply with the terms of service of any third-party platforms (Spotify, YouTube Music, SoundCloud, Apple Music, etc.) when accessing content through their services.

5. **Limitation of Liability**: The developers disclaim any responsibility for users' actions or any legal consequences arising from the misuse of this application.

6. **Indemnification**: By using SynkRip, you agree to indemnify and hold harmless the developers from any claims, damages, or legal issues arising from your use of the application.

**USE THIS SOFTWARE AT YOUR OWN RISK AND ENSURE COMPLIANCE WITH ALL APPLICABLE LAWS.**
<br><br>
This project uses my private API keys to fetch playlists from Spotify but, these keys, are subject to rate limits and restrictions. To ensure uninterrupted usage, it is highly recommended that you generate and use your own API keys. Instructions for setting up your own API keys can be found in the respective service's developer documentation.
<br>
For YouTube Music, SoundCloud, and Apple Music, I am utilizing the privte API endpoints of the respective services. This means that the app may not work as expected if these endpoints change or are deprecated. I am not responsible for any issues that may arise from using these private APIs.
<br>
<br><br>
This project was created for fun during my free time and has not been meticulously polished or professionally maintained. Use it as-is, and feel free to contribute or improve it!

## License

SynkRip is open-source under the Apache License Version 2.0

## Tech Stack

- [Wails](https://wails.io/) - Desktop application framework
- [Vue.js](https://vuejs.org/) - Frontend framework
- [yt-dlp](https://github.com/yt-dlp/yt-dlp) - Media downloading for YoutTube
- [ffmpeg](https://github.com/FFmpeg/FFmpeg) - Media processing


## Contact

- **Author**: Cuniolo Raffaele
- **GitHub**: [CunioloRaffaele](https://github.com/CunioloRaffaele)

Thank you for using SynkRip! ❤️
