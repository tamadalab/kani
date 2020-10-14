#!/bin/sh

function access_analyses_file() {
  # dir=$1
  # echo $dir

  # cmd="ls -rtF | grep -v / | tail -n 1"
  # eval $cmd

  pyc="python3 /usr/local/opt/kani/analyses/analyses.py"
  eval $pyc

  # test=`find $dir -type f | ls -rtF | grep -v /`
  # echo $test
}

access_analyses_file $PWD

