#! /bin/sh

# 現在のディレクトリが kani の実行対象プロジェクトであるかを判断する．
# このコマンドを実行後の終了ステータス（`$?`）で判断する．
#     0なら対象プロジェクト，
#     1なら対象プロジェクトだけど実行をoffにしている状態，
#     2なら対象プロジェクトじゃないこと，
#     3なら.gitディレクトリが見つからないこと
# を表す．

GREP=grep
# GREP=rg
SCRIPT_DIR=$(dirname $0)

function is_git_dir_not_found() {
    if [[ $1 == "" ]]
    then
        exit 3
    fi
}

function is_target_dir() {
    project_dir=$1
    if [[ ! -e "$HOME/.config/kani/projects" ]] # ~/.config/kani/projects が存在しない場合
    then
        exit 2
    fi

    # does include project_dir in ~/.config/kani/projects?
    grep --silent $project_dir $HOME/.config/kani/projects
    if [[ $? -ne 0 ]] # not exist
    then
        exit 2
    fi
}

function is_enable_kani() {
    project_dir=$1
    # does exists $project_dir/.kani/disable?
    if [[ -d $project_dir/.kani/disable ]]
    then
        exit 1
    else
        exit 0
    fi
}

project_dir=$($SCRIPT_DIR/find-project-dir.sh $PWD)

is_git_dir_not_found $project_dir
is_target_dir $project_dir
is_enable_kani $project_dir
