#! /bin/sh

script_dir=$(dirname $0)

${script_dir}/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
fi

. $script_dir/init_envs.sh
echo "hooked periodic on project $KANI_PROJECT_DIR"
