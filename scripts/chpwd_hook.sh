#! /bin/sh

script_dir=$(dirname $0)

${script_dir}/is-target-project.sh
if [[ $? -ne 0 ]]; then
    # echo "is not target project"
    exit 0
fi

. $script_dir/init_envs.sh

echo "hooked chpwd on project $KANI_PROJECT_DIR"
