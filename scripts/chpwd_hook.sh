#! /bin/sh

script_dir=$(dirname $0)

${script_dir}/is-target-project.sh
if [[ $? -ne 0 ]]; then
    # echo "is not target project"
    exit 0
fi

project_dir=$(${script_dir}/find-project-dir.sh)
echo "hooked chpwd on project $project_dir"
