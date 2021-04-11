#! /bin/zsh

script_dir=$(dirname $0)

. $script_dir/init_envs.sh

# record the previous command to .kani/prev_cmd.
function record_command() {
    if [[ $# -eq 0 ]]; then
        return
    fi
    echo "$@" > $KANI_PROJECT_DIR/.kani/prev_cmd
    # echo "recoding: last update $datecmd"
}


$script_dir/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    record_command $1
fi

# echo "hooked preexec on project $project_dir"
