#! /bin/sh

# Find the project directory.
# This script traverses from the current directory to the root directory and finds the .git directory.
# If the script found the .git directory, the found directory is the root of the project directory.

function find_project_directory() {
    dir=$1
    while [[ $dir != "/" ]]
    do
        if [[ -d $dir/.git ]]
        then
            echo $dir
            return 0
        fi
        dir=$(dirname $dir)
    done
    return 1
}

find_project_directory $PWD
