# :crab: kani

初学者のためのGit/GitHub操作補助フレームワーク．

[![License](https://img.shields.io/badge/License-CC0--1.0-blue?logo=spdx)](https://creativecommons.org/publicdomain/zero/1.0/)
[![Version](https://img.shields.io/badge/Version-1.1.1-blue.svg)](https://github.com/tamadalab/kani/releases/tag/v1.1.1)
[![DOI](https://zenodo.org/badge/285447906.svg)](https://zenodo.org/badge/latestdoi/285447906)

[![tamada/brew/kani](https://img.shields.io/badge/Homebrew-tamadalab%2Fbrew%2Fkani-green?logo=homebrew)](https://github.com/tamadalab/homebrew-brew)

[![Discussion](https://img.shields.io/badge/GitHub-Discussion-orange?logo=GitHub)](https://github.com/tamadalab/kani/discussions)

## 本ツールは2020年度特別研究IIで作成

https://github.com/tamadalab/2020bthesis_masuda

### 実装機能概要
- コンパイル時にcommit差分を比較し，規定値以上の変更があった場合commitを促す．
- 連続したエラーが解消された時点で，commitを促す．
- commitを促す際，add commit pushの簡易にhelp表示する．

### 未修正点

- エラー情報を蓄積するDBが1つに集約されているため，ファイルAのエラーが修正されないまま別ファイルを実行すると，ファイルAのエラー情報の影響を受ける．
- コンパイルはgccとclangの場合と手動で制限している[該当ファイル](https://github.com/tamadalab/kani/blob/master/scripts/precmd_hook.sh)．
- .gitフォルダが上位ディレクトリにない場合，.git無いよとエラー出る(操作には影響ない)．
- 正直スパゲッティになってるファイルがある．


## 評価実験用

#### テンプレートリポジトリ

https://github.com/tmdlab2020TestTeam/testTemplate

#### GitHubClassroom[GitHub Classroomとは](http://takehiroman.hatenablog.com/entry/2016/03/31/135736)
https://classroom.github.com/classrooms/73564814-tmdlab2020testteam-githubclassroom/assignments/tmdlab-test-team

# kaniの機能詳細

### 分析の一時停止/再開

```sh
$ git kani disable/enable
```

* `disable` で `PROJECT_ROOT/.kani/disable` ファイルが作成される．
* `enable` を実行すると，`PROJECT_ROOT/.kani/disable` ファイルが削除される．

### 分析対象から外す

```sh
$ git kani deinit
```

* `$HOME/.config/kani/projects`からプロジェクトのパスが削除される．


## :anchor: Install

### :beer: Homebrew

```sh
$ brew tap tamadalab/brew
$ brew install kani
```

上記コマンドにて，以下のようなディレクトリが作成される．
Installing `kani` by Homebrew, the following directories are built.

```sh
/usr/local/Celler/kani
├── README.md
├── analyses
│   ├── recommend.py (どういった条件でcommitを促すか決める所)
│   └── commit_guide.txt (commitを促す際の文)
├── bin
│   └── git-kani
└── scripts # utilities (hook functions)
    ├── chpwd_hook.sh
    ├── periodic_hook.sh
    ├── precmd_hook.sh
    ├── preexec_hook.sh
    ├── find-project-dir.sh (現在のディレクトリからプロジェクトのルートを取得するスクリプト)
    └── is-target-project.sh (現在のディレクトリのプロジェクトがkaniの分析対象かどうかを判定するスクリプト)
```

### Initialize `kani` on your environment.

Write the following line into your `~/.zshrc`.

```sh
eval "$(git kani init -)"
```

hook関数については[この資料](https://qiita.com/mollifier/items/558712f1a93ee07e22e2)を参照してください．

### Enabling `kani` on your project.

Type the following command.

```sh
$ git kani init
```

* 上記のコマンドで `.git` と同じディレクトリに `.kani` ディレクトリが作成される．
    * `.kani` ディレクトリには，`analyses`ディレクトリにあるスクリプトがコピーされる．
* `$HOME/.config/kani/projects` にプロジェクトのパスが追記される．

* `PROJECT_ROOT/.kani` には，分析結果のデータを格納している．
