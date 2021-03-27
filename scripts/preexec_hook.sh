#! /bin/zsh

KANI_HOME=/usr/local/opt/kani
PROJECT_DIR=$($KANI_HOME/scripts/find-project-dir.sh)

# record the previous command to .kani/prev_cmd.
function record_command() {
    if [[ $# -eq 0 ]]; then
        return
    fi
    echo "$@" > $PROJECT_DIR/.kani/prev_cmd
    datecmd=$(date "+%H:%M:%S")
    branch=$(git describe --contains --all HEAD)
    commitId=$(git rev-parse HEAD)
    echo -n "$datecmd,$@,$branch,$commitId" >> $PROJECT_DIR/.kani/test.log
    # echo "recoding: last update $datecmd"
}


$KANI_HOME/scripts/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    record_command $1
fi

# echo "hooked preexec on project $project_dir"
