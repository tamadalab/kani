#! /bin/zsh

KANI_HOME=/usr/local/opt/kani
PROJECT_DIR=$($KANI_HOME/scripts/find-project-dir.sh)

function find_previous_command() {
    prevcmd=$PROJECT_DIR/.kani2/prev_cmd
    if [[ -f $prev_cmd ]]; then
        cat $prev_cmd # 内容を確認する．
        rm $prev_cmd  # 読み出し後，削除する．
    fi
}

$KANI_HOME/scripts/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    prevcmd=$(find_previous_command)
    echo "prev cmd: $prevcmd, status: $1" # 終了ステータスは $1.
fi