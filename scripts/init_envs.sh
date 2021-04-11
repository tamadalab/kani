#! /bin/sh

if [ ! "${KANI_HOME:+foo}" ]; then
    export KANI_HOME=$(dirname $0)/..
fi

if [ ! "${KANI_PROJECT_DIR:+foo}" ]; then
    export KANI_PROJECT_DIR=$(${KANI_HOME}/scripts/find-project-dir.sh)
fi
