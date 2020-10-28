# Git操作補助ツールkani

### テスト用テンプレートリポジトリとGitHubClassroomの部屋
https://github.com/tmdlab2020TestTeam/testTemplate

https://classroom.github.com/classrooms/73564814-tmdlab2020testteam-githubclassroom/assignments/tmdlab-test-team

# kani

初学者のgit学習を支援するためのプロジェクト．


* `scripts` 以下に分析用のスクリプトを準備する．
    * 本番環境では，`/usr/local/Cellar/kani/scripts` 以下にインストールされる．
* `git kani init`で対象プロジェクトを`kani`での分析対象とする．
* `git kani disable/enable` で分析の一時停止/再開を行う．
* `git kani deinit` で対象プロジェクトを分析対象から外す．

* zsh のフックには，`kani/scripts` 以下のスクリプトを絶対パスで指定すると良いと思います．
    * `git kani init` を実行すると，`analyses` ディレクトリの内容が `PROJECT_ROOT/.kani/analyses`にコピーされますが，コピーせず，`kani/scripts`内のスクリプトを呼び出せば良いかと思います．

* `PROJECT_ROOT/.kani` には，分析結果のデータを収集日を添えて入れておくと良いのではないかと思います．

## 使い方

### プロジェクトでの初期設定

```sh
$ git kani init
```

* 上記のコマンドで `.git` と同じディレクトリに `.kani` ディレクトリが作成される．
    * `.kani` ディレクトリには，`analyses`ディレクトリにあるスクリプトがコピーされる．
    * 将来的に，`analyses`は不要かもしれない
* `$HOME/.config/kani/projects` にプロジェクトのパスが追記される．

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

## ヘルパースクリプト

* `find-project-dir.sh` 現在のディレクトリからプロジェクトのルートを取得するスクリプト．
* `is-target-project.sh` 現在のディレクトリのプロジェクトがkaniの分析対象かどうかを判定するスクリプト．
    * 終了ステータス（`$?`）で判断する．詳細は，スクリプトの冒頭を確認してください．

## インストール方法

```sh
$ brew tap tamadalab/brew
$ brew install kani
```

上記コマンドにて，以下のようなディレクトリが作成される．

```sh
/usr/local/Celler/kani
├── README.md
├── analyses
│   ├── analyses.py
│   └── guide_commit.txt
├── bin
│   └── git-kani
└── scripts # ユーティリティスクリプト
    ├── chpwd_hook.sh
    ├── periodic_hook.sh
    ├── precmd_hook.sh
    ├── preexec_hook.sh
    ├── find-project-dir.sh
    └── is-target-project.sh
```

`zshrc.txt` に書かれていた内容は，`git kani init -` で出力するようにしました．
そのため，`~/.zshrc` の最後に，次の1行を追加すればOKにするようにしました．

```sh
eval $(git kani init -)
```
