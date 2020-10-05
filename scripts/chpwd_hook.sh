#! /bin/zsh

KANI_HOME=/usr/local/opt/kani

$KANI_HOME/scripts/is-target-project.sh
if [[ $? -ne 0 ]]
then
    exit 0
fi

project_dir=$($KANI_HOME/scripts/find-project-dir.sh)
echo "hooked chpwd on project $project_dir"