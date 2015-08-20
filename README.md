# Bryggeri

> Bundle install useful packages with [Homebrew](http://brew.sh/)

## Installation

```console
 $ curl -fsSL git.io/bryggeri | sh
```

All this installation script does is download the `bryggeri` script, make it an executable, and copy it to your `$PATH (/usr/local/bin)`. For copying to your `$PATH`, it may require you to enter your password. If there is a better way to do this, please send in a pull request.

`bryggeri` relies solely on `brew` - if you don't already have it installed just paste the following into the terminal:

```console
 $ ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

## Usage

```console
  Usage: bryggeri [options]

  Options:
    -V, --version    Output version
    -L, --list       Output the packages
    -h, --help       This message.
    --
```

### Contributing

Contributions are always welcome, from a typo in the README to an request to add another package in the list to a completely new feature itself.

Fork the code, make a new branch, and send in a pull request.

### Licence

[MIT](licence) © Johnie Hjelm
