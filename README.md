<p align="center">
    <img src="docs/logo.png" height="250" alt="gopass Gopher by Vincent Leinweber, remixed from the Renée French original Gopher" title="gopass Gopher by Vincent Leinweber, remixed from the Renée French original Gopher" />
</p>

# Overview [![Build Status](https://img.shields.io/github/workflow/status/gopasspw/gopass/Build%20gopass/master)](https://github.com/gopasspw/gopass/actions/workflows/build.yml?query=branch%3Amaster) [![Packaging status](https://repology.org/badge/tiny-repos/gopass-gopasspw.svg)](https://repology.org/project/gopass-gopasspw/versions) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/gopasspw/gopass/blob/master/LICENSE) [![Github All Releases](https://img.shields.io/github/downloads/gopasspw/gopass/total.svg)](https://github.com/gopasspw/gopass/releases) [![Gopass Slack](https://img.shields.io/badge/%23gopass-Slack-brightgreen)](https://join.slack.com/t/gopassworkspace/shared_invite/zt-17jl74b5x-U1OUW4ts4AQ7eAf2V4QaaQ)

> The slightly more awesome standard UNIX password manager for teams.

Manage your credentials with easy. In a globally distributed team, on multiple devices or fully offline on an air gapped machine.

- **Works everywhere** - The same user experience on Linux, MacOS, *BSD or Windows
- **Built for teams** - Built from our experience working in distributed development teams
- **Full autonomy** - No network connectivity required, unless you want it

# How Does It Work?

By default your credentials are encrypted with GPG and versioned in git. This can be customized easily.
The primary interface is the command line. Making it an excellent choice for CLI fans, CI/CD systems or
anything you can hook it up with. Gopass can also integrate with your browser so you can largely avoid
the command line - if you want.

# Installation

## Necessary prerequisits for running `gopass`

`gopass` can operate without any dependencies but most users will want to use it with `gpg` and `git`.
An external editor is required if you want to be able to use `gopass edit`.

## Installation through package managers

### [Homebrew](https://brew.sh) (Linux/MacOS)

[![homebrew version](https://img.shields.io/homebrew/v/gopass)](https://github.com/Homebrew/homebrew-core/blob/master/Formula/gopass.rb)

```shell
brew install act
```

### Debian (Ubuntu, Debian, Raspbian, ...)

**Warning**: Do not install the `gopass` package for the official repositories. That is a completely different project that has no relation to us.

```shell
$ curl https://packages.gopass.pw/repos/gopass/gopass-archive-keyring.gpg | sudo tee /usr/share/keyrings/gopass-archive-keyring.gpg >/dev/null
$ cat << EOF | sudo tee /etc/apt/sources.list.d/gopass.sources
Types: deb
URIs: https://packages.gopass.pw/repos/gopass
Suites: stable
Architectures: all amd64 arm64 armhf
Components: main
Signed-By: /usr/share/keyrings/gopass-archive-keyring.gpg
EOF
$ sudo apt update
$ sudo apt install gopass gopass-archive-keyring
```

### Fedora / RedHat / CentOS

```shell
dnf install gopass
```

Note: You might need to run `dnf copr enable daftaupe/gopass` first.

### Arch Linux

```shell
pacmsn -S gopass
```

### Windows


```shell
# Chocolatey
choco install gpg4win
choco install gopass
# Alternatively
scopp install gopass
```

### FreeBSD / OpenBSD


```shell
cd /usr/ports/security/gopass
make install
```

### Alpine Linux

```shell
apk add gopass
```

## Other installation options

Please see [docs/setup.md](https://github.com/gopasspw/gopass/blob/master/docs/setup.md) for other options.

### From Source

```shell
go install github.com/gopasspw/gopass@latest
```

Note: `latest` is not a stable release. We recommend to only use released versions.

### Manual download

Download the [latest release](https://github.com/gopasspw/gopass/releases/latest) and add the binary to your PATH.

# Example usage

```shell
# Command structure
gopass [<command>] [options] [args]
# Shortcut for gopass show [<key>]
gopass [<key>]

# Enter the gopass REPL
gopass

# Find all entries matching the search string
gopass find github

# List your store
gopass ls

# List all mounts
gopass mounts

# List all recipients
gopass recipients

# Sync with all remotes
gopass sync

# Setup a new store
gopass setup
```

## First `gopass` run

```shell
$ gopass setup

   __     _    _ _      _ _   ___   ___
 /'_ '\ /'_'\ ( '_'\  /'_' )/',__)/',__)
( (_) |( (_) )| (_) )( (_| |\__, \\__, \
'\__  |'\___/'| ,__/''\__,_)(____/(____/
( )_) |       | |
 \___/'       (_)

🌟 Welcome to gopass!
🌟 Initializing a new password store ...
🌟 Configuring your password store ...
🎮 Please select a private key for encrypting secrets:
[0] gpg - 0xFEEDBEEF - John Doe <john.doe@example.org>
Please enter the number of a key (0-12, [q]uit) (q to abort) [0]: 0
❓ Do you want to add a git remote? [y/N/q]: y
Configuring the git remote ...
Please enter the git remote for your shared store []: git@gitlab.example.org:john/passwords.git
✅ Configured
```

Hint: `gopass setup` will use `gpg` encryption and `git` storage by default. Use `--crypto=age` to
use `age` encryption instead. Or you can opt out of using a versioned store with `--storage=fs`.

An existing store can be cloned with e.g. `gopass clone git@gitlab.example.org:john/passwords.git`.

## Screenshot

![screenshot](docs/showcase.png)

## Support

Please ask on [Slack](https://join.slack.com/t/gopassworkspace/shared_invite/zt-17jl74b5x-U1OUW4ts4AQ7eAf2V4QaaQ).

## Contributing

We welcome any contributions. Please see the [CONTRIBUTING.md](https://github.com/gopasspw/gopass/blob/master/CONTRIBUTING.md) for more information.

## Credit & License

gopass is licensed under the terms of the MIT license. You can find the complete text in `LICENSE`.

Please refer to the Git commit log for a complete list of contributors.
