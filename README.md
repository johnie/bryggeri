# 🍼 Bryggeri

> Bundle install useful packages with [Homebrew](http://brew.sh/)

## Installation

```console
 $ curl -fsSL git.io/bryggeri | sh
```

All this installation script does is download the `bryggeri` script, make it an executable, and copy it to your `$PATH (/usr/local/bin)`. For copying to your `$PATH`, it may be required to enter your password. If there is a better way to do this, please send in a pull request.

`bryggeri` relies solely on `brew` - if you don't already have it installed just paste the following into the terminal:

```console
 $ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

## Usage

```console
  Usage: bryggeri [options]

  Options:
      install           Install packages
      install --casks   Install casks
      -V, --version     Output version
      -L, --list        Output the packages
      -h, --help        This message
      --                End of options
```

### Formulaes

- [ansible](https://formulae.brew.sh/formula/ansible): An open-source automation tool for configuration management, application deployment, and task automation.
- [bat](https://formulae.brew.sh/formula/bat): A cat clone with syntax highlighting and Git integration.
- [chrome-cli](https://formulae.brew.sh/formula/chrome-cli): A command-line tool for controlling Google Chrome.
- [coreutils](https://formulae.brew.sh/formula/coreutils): GNU core utilities (including "ls", "cp", "mv", and more) for macOS.
- [curl](https://formulae.brew.sh/formula/curl): Command-line tool for transferring data with URLs.
- [entr](https://formulae.brew.sh/formula/entr): Event notify test runner.
- [eza](https://formulae.brew.sh/formula/eza): A modern replacement for "ls" with additional features.
- [ffmpeg](https://formulae.brew.sh/formula/ffmpeg): A complete solution to record, convert, and stream audio and video.
- [fx](https://formulae.brew.sh/formula/fx): Command-line JSON processing tool.
- [fzf](https://formulae.brew.sh/formula/fzf): Command-line fuzzy finder.
- [git-delta](https://formulae.brew.sh/formula/git-delta): A viewer for Git and diff output.
- [git-extras](https://formulae.brew.sh/formula/git-extras): Small Git utilities.
- [git](https://formulae.brew.sh/formula/git): Distributed version control system.
- [go](https://formulae.brew.sh/formula/go): The Go programming language.
- [httpie](https://formulae.brew.sh/formula/httpie): Command-line HTTP client.
- [gh](https://formulae.brew.sh/formula/gh): Command-line tool that wraps Git in GitHub.
- [imagemagick](https://formulae.brew.sh/formula/imagemagick): A software suite to create, edit, and compose bitmap images.
- [mas](https://formulae.brew.sh/formula/mas): A command-line interface for the Mac App Store.
- [ncdu](https://formulae.brew.sh/formula/ncdu): NCurses Disk Usage, a disk usage analyzer with an ncurses interface.
- [neovim](https://formulae.brew.sh/formula/neovim): Neovim, a highly configurable, extensible text editor.
- [nmap](https://formulae.brew.sh/formula/nmap): Network exploration tool and security scanner.
- [nvm](https://formulae.brew.sh/formula/nvm): Node Version Manager, a tool for managing multiple versions of Node.js.
- [openssl](https://formulae.brew.sh/formula/openssl): SSL/TLS cryptography library.
- [python](https://formulae.brew.sh/formula/python): Interpreted, high-level, general-purpose programming language.
- [rust](https://formulae.brew.sh/formula/rust): The Rust programming language.
- [rclone](https://formulae.brew.sh/formula/rclone): Rsync for cloud storage.
- [ripgrep](https://formulae.brew.sh/formula/ripgrep): Recursively search your current directory for a regex pattern.
- [speedtest-cli](https://formulae.brew.sh/formula/speedtest-cli): Command-line interface for testing internet bandwidth.
- [tmux](https://formulae.brew.sh/formula/tmux): Terminal multiplexer.
- [tre-command](https://formulae.brew.sh/formula/tre-command): Display directories as trees.
- [wget](https://formulae.brew.sh/formula/wget): Internet file retriever.
- [youtube-dl](https://formulae.brew.sh/formula/youtube-dl): A command-line program to download videos from YouTube and other sites.

### Casks

- [adobe-creative-cloud](https://formulae.brew.sh/cask/adobe-creative-cloud): Adobe Creative Cloud, a suite of creative software.
- [appcleaner](https://formulae.brew.sh/cask/appcleaner): AppCleaner, a utility to uninstall unwanted apps.
- [bartender](https://formulae.brew.sh/cask/bartender): Bartender, a tool to manage your menu bar items.
- [brave-browser](https://formulae.brew.sh/cask/brave-browser): A web browser focused on privacy.
- [docker](https://formulae.brew.sh/cask/docker): Docker, a platform for developing, shipping, and running applications.
- [dropbox](https://formulae.brew.sh/cask/dropbox): Dropbox, a cloud storage and file synchronization service.
- [figma](https://formulae.brew.sh/cask/figma): Figma, a collaborative interface design tool.
- [firefox](https://formulae.brew.sh/cask/firefox): Mozilla Firefox, a popular web browser.
- [google-chrome](https://formulae.brew.sh/cask/google-chrome): Google Chrome, a web browser by Google.
- [iterm2](https://formulae.brew.sh/cask/iterm2): iTerm2, a terminal emulator for macOS.
- [karabiner-elements](https://formulae.brew.sh/cask/karabiner-elements): Karabiner-Elements, a keyboard customizer for macOS.
- [keepingyouawake](https://formulae.brew.sh/cask/keepingyouawake): KeepingYouAwake, a utility to prevent your Mac from sleeping.
- [raycast](https://formulae.brew.sh/cask/raycast): Raycast, an app launcher for developers.
- [rectangle](https://formulae.brew.sh/cask/rectangle): Rectangle, a window manager for macOS.
- [spotify](https://formulae.brew.sh/cask/spotify): Spotify, a popular music streaming service.
- [visual-studio-code](https://formulae.brew.sh/cask/visual-studio-code): Visual Studio Code, a code editor by Microsoft.
- [warp](https://formulae.brew.sh/cask/warp): Warp, a fast and modern Terminal.

### Contributing

Contributions are always welcome, from a typo in the README to a request to add another package in the list to a completely new feature itself.

Fork the code, make a new branch, and send in a pull request.

### Licence

[MIT](licence) © Johnie Hjelm
