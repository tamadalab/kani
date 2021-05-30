---
title: ":anchor: Install"
---

### :megaphone: Notice

`kani` runs on [zsh](https://www.zsh.org/), and [bash](https://www.gnu.org/software/bash/) only.
Not work on other shells (`csh`, `ksh`, `fish`, ...).

### :pushpin: Requirements

* [rcaloras/bash-preexec](https://github.com/rcaloras/bash-preexec)
    * Runs `kani` on bash environment.

### :beer: Homebrew (macOS)

Type the following two commands, then `kani` will be installed into your `/usr/local/opt/kani`

```sh
brew tap tamadalab/brew
brew install kani
```

### :muscle: Install yourself

* Clone the `kani` repository from GitHub ([`tamadalab/kani`](https://github.com/tamadalab/kani)).
    * `git clone https://github.com/tamadalab/kani.git`
* Move to the cloned `kani` directory.
    * `cd kani`
* Compile the sources.
    * `make`
* Make distribution file.
    * `make dist`
* Extract the suitable archive file in `dist` directory in your system.
    * `tar xvfz dist/kani-1.0.0-windows-amd64.tar.gz -C INSTALL_DIR`
* Set `KANI_HOME` environment value to your `INSTALL_DIR`
    * `setenv KANI_HOME=${INSTALL_DIR}`
    * the above line should be written in your `.bashrc` or `.zshrc`.
