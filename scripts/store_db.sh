#!/bin/sh

script_dir=$(dirname $0)

function initialize_db() {
    echo "CREATE TABLE histories ( \
  id          INTEGER PRIMARY KEY, \
  datetime    TEXT    DEFAULT CURRENT_TIMESTAMP, \
  command     TEXT    NOT NULL, \
  status_code INTEGER NOT NULL, \
  branch      TEXT, \
  revision    TEXT, \
  shell       TEXT  \
);" |  sqlite3 $1
}

function store_db() {
    currentShell=$SHELL # how to get executing shell?
    # echo $0 returns script name, therefore, not suitable.
    echo "INSERT INTO histories (command, status_code, branch, revision, shell) \
  VALUES ( \
    '$2', \
    $3,     \
    \"$4\", \
    \"$5\", \
    \"$currentShell\"
  );" | sqlite3 $1
}

if [[ ! -f $1 ]] ; then
    initialize_db $1
fi

store_db "$@"
