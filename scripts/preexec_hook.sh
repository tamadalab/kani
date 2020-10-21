#! /bin/zsh

KANI_HOME=/usr/local/opt/kani
PROJECT_DIR=$($KANI_HOME/scripts/find-project-dir.sh)

# .kani/prev_cmd に前回実行したコマンドを記録しておく．
function record_command() {
    if [[ $# -eq 0 ]]; then
        return
    fi
    echo "$@" > $PROJECT_DIR/.kani2/prev_cmd
}


$KANI_HOME/scripts/is-target-project.sh
if [[ $? -ne 0 ]]
then
    exit 0
else
    record_command $1
    pyc="python3 $KANI_HOME/analyses/analyses.py"
    eval $pyc
fi

# echo "hooked preexec on project $project_dir"
