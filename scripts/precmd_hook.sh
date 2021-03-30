#! /bin/sh

script_dir=$(dirname $0)
PROJECT_DIR=$(${script_dir}/find-project-dir.sh)
# FAILURES_DIR=$PROJECT_DIR/.kani/failures_compilation

function find_previous_command() {
    prev_cmd=$PROJECT_DIR/.kani/prev_cmd
    if [[ -e $prev_cmd ]]; then
        cat $prev_cmd # read contents
        rm $prev_cmd  # after reading, remove it.
    fi
}

function store_db() {
    prevcmd="$1"
    statusCode=$2
    revision=$(git rev-parse HEAD)
    branch=$(git symbolic-ref HEAD)
    $script_dir/store_db.sh $PROJECT_DIR/.kani/kani.sqlite "$prevcmd" $statusCode $branch $revision
}

${script_dir}/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    prevcmd=$(find_previous_command)
<<<<<<< HEAD
    # echo "prev cmd: $prevcmd, status: $1" # for debugging, status code is $1
    echo $1 >> $PROJECT_DIR/.kani/test.log
    count=0 # initialize error count.
    if [[ $prevcmd =~ gcc* || $prevcmd =~ clang* && $1 -ne 0 ]]; then # count up if status code is non 0.
      echo "$prevcmd : $1" >> $FAILURES_DIR # counts failures
      echo "error : status $1" >> $PROJECT_DIR/.kani/test.log
  elif [[ $prevcmd =~ gcc* || $prevcmd =~ clang* && $1 -eq 0 && -e $FAILURES_DIR ]]; then # if error is corrected, record count of errors.
      count=$(wc -l $FAILURES_DIR)
      rm $FAILURES_DIR
    fi
    if [[ ! -e $PROJECT_DIR/.kani/disable ]]; then
      pyc="python3 $KANI_HOME/analyses/recommend.py $count"
      eval $pyc
    fi
=======
    # echo "prev cmd: \"$prevcmd\", status: $1" # (デバッグ用)終了ステータスは $1.
    store_db "$prevcmd" $1
>>>>>>> master
fi
