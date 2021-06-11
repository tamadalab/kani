# :crab: kani

A tool for supporting git/GitHub operation for the novice developers.

[![License](https://img.shields.io/badge/License-CC0--1.0-blue?logo=spdx)](https://creativecommons.org/publicdomain/zero/1.0/)
[![Version](https://img.shields.io/badge/Version-1.0.0-blue.svg)](https://github.com/tamadalab/kani/releases/tag/v1.0.0)
[![DOI](https://zenodo.org/badge/285447906.svg)](https://zenodo.org/badge/latestdoi/285447906)

[![tamada/brew/kani](https://img.shields.io/badge/Homebrew-tamadalab%2Fbrew%2Fkani-green?logo=homebrew)](https://github.com/tamadalab/homebrew-brew)

[![Discussion](https://img.shields.io/badge/GitHub-Discussion-orange?logo=GitHub)](https://github.com/tamadalab/kani/discussions)

## :speaking_head: Overview

Learning the git and GitHub operation is difficult for novice programmers.
Since even programming is a pretty complicated task, it becomes challenging to consider various kinds of operation except programming, such as good commit, pull requests, etc.
However, the recent development process ordinary manages product versions with SCM such as `git`.
Therefore, it requires the novices should learn them as soon as possible.

Then, we focus on the commits' timing, and the proposed system recommends `commit` at the suitable timing.
Also, the recommendation text contains a simple help message for `git`.
From above two manners, we expect the novices accustomed `git`.

The proposed system, named `kani`, records the terminal's executed commands and runs the recommendation engine by each recoding.
The recommendation engine analyzes the histories of the executed commands with their status codes and statuses of `git`.

## :speech_balloon: Description

### :wind_chime: Features

* Records the executed command in your terminal in your repository,
* Runs analysis scripts by each command execution,
    * Recommends the timing of `git commit`, and `git push`.
* Shows simple usage of `git commit` and `git push`.

### :checkered_flag: Goal of `kani`

`kani` recommends the operations of git and GitHub from behaviors of novices in the terminal.
Therefore, the goal of `kani` is to become unnecessary for novices, like walking assist instrument for babies.
For this, `kani` wants to measure git/GitHub operation maturity rates (GOMR) for the novices.
GOMR would measure by the following methods.

* Matching rates with the behaviors of experts (recording beforehand)
* Following rates of Git flow, GitHub flow, or GitLab flow.


## :anchor: Install

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

## :runner: Usage

### :shoe: Initialization

At first, write the following snippet in your login shell initializer (`.bashrc`, or `.zshrc`)

```sh
eval "$(kani init -)"
```

### :running_shirt_with_sash: Setup repository

* Choose your `git` repositories to run `kani`,
* Change directory to your local `git` repository,
* Run `kani init` command for preparing the `kani` running on the repository.

### :coffee: Stop `kani` temporarily in a particular repository

If you want to stop recoding command temporary, run `kani disable` (`kani` creates file `disable` on `.kani` directory).
Then, want to start recoding again, run `kani enable`.

### :broken_heart: Stop `kani`  permanently in a particular repository

If typing `kani deinit` in a particular repository, kani never records anymore.
After that, typing `kani init`, `kani` starts recoding again.

## :smile: About

### :scroll: License

[![CC0-1.0](https://img.shields.io/badge/License-CC0--1.0-blue?logo=spdx)](https://creativecommons.org/publicdomain/zero/1.0/)

> The person who associated a work with this deed has dedicated the work to the public domain by waiving all of his or her rights to the work worldwide under copyright law, including all related and neighboring rights, to the extent allowed by law.
>
> You can copy, modify, distribute and perform the work, even for commercial purposes, all without asking permission. See Other Information below.

### :page_with_curl: Cite `kani`



### :jack_o_lantern: Logo

![logo](https://tamadalab.github.io/kani/images/kani.png)

This image comes from https://freesvg.org/crab-image ([Public Domain](https://creativecommons.org/licenses/publicdomain/))

### :name_badge: Project name comes from?

We do not know.
Note that `kani` means crab in Japanese.

### :man_office_worker: Developers :woman_office_worker:

* [Arisa Masuda](https://github.com/ma-sa321)
* [Haruaki Tamada](https://github.com/tamada)

### :books: References

* Lassi Haaranen, and Teemu Lehtinen, "Teaching Git on the Side: Version Control System as a Course Platform," Proc. the 2015 ACM Conference on In novation and Technology in Computer Science Education, pp. 87–92, DOI: 10.1145/2729094.2742608, June 2015, https://dl.acm.org/doi/10.1145/2729094.2742608.
* Ville Isomöttönen, and Michael Cochez, "Challenges and Confusions in Learning Version Control with Git," Proc. Information and Communication Technologies in Education, Research, and Industrial Applications (ICTERI 2014), pp. 178-193, DOI: 10.1007/978-3-319-13206-8_9, November 2014, https://link.springer.com/chapter/10.1007/978-3-319-13206-8_9.
* 井上 拓海, 小島 遥一郎, 藤原 賢二, 井垣 宏, "版管理システム利用時のソフトウェア開発フロー遵守状況可視化手法の検討", 信学技法, No.SS2017-55, January 2018.
