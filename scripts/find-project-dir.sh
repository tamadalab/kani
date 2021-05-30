#! /bin/sh

# Find the project directory.
# This script traverses from the current directory to the root directory and finds the .git directory.
# If the script found the .git directory, the script prints the parent directory of .git as the root of the project directory, and returns 0.
# If not found, the script returns the status code 1.

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
