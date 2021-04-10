#! /bin/sh

script_dir=$(dirname $0)
PROJECT_DIR=$(${script_dir}/find-project-dir.sh)
# FAILURES_DIR=$PROJECT_DIR/.kani/failures_compilation

function find_previous_command() {
    prev_cmd=$PROJECT_DIR/.kani/prev_cmd
    if [[ -e $prev_cmd ]]; then
        cat $prev_cmd # read contents
        rm $prev_cmd  # after reading, remove it.
    fi
}

function store_db() {
    prevcmd="$1"
    statusCode=$2
    revision=$(git rev-parse HEAD)
    branch=$(git symbolic-ref HEAD)
    $script_dir/../bin/kani store $PROJECT_DIR/.kani/kani.sqlite "$prevcmd" $statusCode $branch $revision
}

${script_dir}/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    prevcmd=$(find_previous_command)
    # echo "prev cmd: \"$prevcmd\", status: $1" # (デバッグ用)終了ステータスは $1.
    store_db "$prevcmd" $1
fi
