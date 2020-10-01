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
$ brew install ma-sa321/kani/kani
```

上記コマンドにて，以下のようなディレクトリが作成される（予定）．

```sh
/usr/local/Celler/kani
├── README.md
├── bin
│   └── git-kani
├── scripts # ユーティリティスクリプト
│   ├── find-project-dir.sh
│   └── is-target-project.sh
├── analyses
│   └── 分析用のスクリプト
└── zshrc.txt
```

## 実装情報

### `git kani` を実現するには．

https://qiita.com/b4b4r07/items/6b76a5f969231e5e9748

`git-kani` コマンドを用意し，環境変数`PATH`の通ったところに置いておくと `git kani` コマンドが有効になる．
`git kani init` を実行することで，`$HOME/.conf/kani/projects` に`.git`ディレクトリのパスのリストを書き込んでおく．

開発時には，`PATH=bin:$PATH git kani hogehoge` で確認してね．

玉田の書きやすい言語のGoで書きました．．．
このような機能が欲しいなどあれば言ってもらえると追加します．

### `brew install ma-sa321/kani/kani` ができるようにするには．

`homebrew-kani` プロジェクトを作成し，`kani.rb` を用意する．
`kani.rb` の中身は今度書きます．
