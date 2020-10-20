#! /bin/zsh

KANI_HOME=/usr/local/opt/kani

$KANI_HOME/scripts/is-target-project.sh
if [[ $? -ne 0 ]]
then
    exit 0
else
    pyc="python3 /usr/local/opt/kani/analyses/analyses.py"
    eval $pyc
fi



project_dir=$($KANI_HOME/scripts/find-project-dir.sh)
echo "hooked preexec on project $project_dir"
