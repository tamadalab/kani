#!/bin/sh

script_dir=$(dirname $0)

function initialize_db() {
    echo "CREATE TABLE histories ( \
  id       INTEGER PRIMARY KEY, \
  datetime TEXT    DEFAULT CURRENT_TIMESTAMP, \
  command  TEXT    NOT NULL, \
  status   INTEGER NOT NULL, \
  branch   TEXT, \
  revision TEXT  \
);" |  sqlite3 $1
}

function store_db() {
    echo "INSERT INTO histories (command, status, branch, revision) \
  VALUES ( \
    '$2', \
    $3,     \
    \"$4\", \
    \"$5\" \
  );" | sqlite3 $1
}

if [[ ! -f $1 ]] ; then
    initialize_db $1
fi

store_db "$@"
