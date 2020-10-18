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

    # ~/.config/kani/projects に project_dir が存在するかを確認する．
    grep --silent $project_dir $HOME/.config/kani/projects 
    if [[ $? -ne 0 ]] # 存在しなかった
    then
        exit 2
    fi
}

function is_enable_kani() {
    project_dir=$1
    # $project_dir/.kani/disable が存在するかを確認する．
    if [[ -d $project_dir/.kani/disable ]]
    then
        exit 1
    fi
    /usr/local/opt/kani/analyses/analyses.sh # 対象プロジェクトの場合分析用shellを走らせる．
}

project_dir=$($SCRIPT_DIR/find-project-dir.sh $PWD)

is_git_dir_not_found $project_dir
is_target_dir $project_dir
is_enable_kani $project_dir

