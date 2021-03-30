#! /bin/zsh

script_dir=$(dirname $0)
PROJECT_DIR=$($script_dir/find-project-dir.sh)

# .kani/prev_cmd に前回実行したコマンドを記録しておく．
function record_command() {
    if [[ $# -eq 0 ]]; then
        return
    fi
    echo "$@" > $PROJECT_DIR/.kani/prev_cmd
}


$script_dir/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    record_command $1
fi

# echo "hooked preexec on project $project_dir"
