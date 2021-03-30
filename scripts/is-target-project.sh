#! /bin/sh

# This script determines if the current directory is the sub-directory in the target project of kani.
# The status code of this script means that:
#     0: the current directory is the sub-directory in the target project.
#     1: the current directory is target project, but disable flag is on.
#     2: the current directory is not target project.
#     3: .git directory not found.

GREP=grep
# GREP=rg
SCRIPT_DIR=$(dirname $0)

function is_git_dir_not_found() {
    if [[ $1 == "" ]]
    then
        exit 3
    fi
}

function is_target_dir() {
    project_dir=$1
    if [[ ! -e "$HOME/.config/kani/projects" ]] # ~/.config/kani/projects not exist
    then
        exit 2
    fi

<<<<<<< HEAD
    # determine if ~/.config/kani/projects contains project_dir.
    grep --silent $project_dir $HOME/.config/kani/projects
    if [[ $? -ne 0 ]] # not found.
=======
    # does include project_dir in ~/.config/kani/projects?
    grep --silent $project_dir $HOME/.config/kani/projects
    if [[ $? -ne 0 ]] # not exist
>>>>>>> master
    then
        exit 2
    fi
}

function is_enable_kani() {
    project_dir=$1
<<<<<<< HEAD
    # determine if $project_dir/.kani/disable exist.
=======
    # does exists $project_dir/.kani/disable?
>>>>>>> master
    if [[ -d $project_dir/.kani/disable ]]
    then
        exit 1
    else
        exit 0
    fi
}

project_dir=$($SCRIPT_DIR/find-project-dir.sh $PWD)

is_git_dir_not_found $project_dir
is_target_dir $project_dir
is_enable_kani $project_dir
