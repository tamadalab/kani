---
title: ":house: Home"
---

[![License](https://img.shields.io/badge/License-CC0--1.0-blue?logo=spdx)](https://creativecommons.org/publicdomain/zero/1.0/)
[![Version](https://img.shields.io/badge/Version-1.1.0-blue.svg)](https://github.com/tamadalab/kani/releases/tag/v1.0.0)

[![tamada/brew/wildcat](https://img.shields.io/badge/Homebrew-tamadalab%2Fbrew%2Fkani-green?logo=homebrew)](https://github.com/tamadalab/homebrew-brew)

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



## :speaking_head: 概要（Japanese version of [Overview](#-overview)）

プログラミング初学者にとって，Gitの操作の学習はハードルが高い．
プログラミングだけでも大変であるのに，プログラミング以外のことも考える必要があるためである．
しかし，昨今の開発では，gitなどを用いたバージョン管理は当たり前のことであり，初学者もできるだけ早い段階で身につけることが望ましい．
そこで，適切なタイミングで git 操作を推薦したり，git操作の簡易ヘルプを自動で出してくれる環境を提案する．
これにより，作業のタイミングや作業内容を推薦に従って操作することで，git操作に慣れてもらうことを目指す．

そのために，提案システム`kani`は`git`リポジトリ内での CUI 操作を記録しておき，記録ごとに推薦エンジンを走らせる．
推薦エンジンは，これまでのコマンドの実行履歴（ステータスコードやgitのステータスを含む）から適切な`git`の操作を推薦する．
