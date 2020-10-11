#! /bin/sh

# project のディレクトリを検索して表示する．
# 現在のディレクトリからルートディレクトリにむけて遡って，.git ディレクトリを探す．
# 見つかったところがプロジェクトのルートディレクトリ（project ディレクトリ）

function find_project_directory() {
    dir=$1
    while [[ $dir != "/" ]]
    do
        if [[ -d $dir/.git ]]
        then
            echo $dir
            return 0
        fi
        dir=$(dirname $dir)
    done
    return 1
}

find_project_directory $PWD
