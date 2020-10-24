#! /bin/zsh

KANI_HOME=/usr/local/opt/kani
PROJECT_DIR=$($KANI_HOME/scripts/find-project-dir.sh)

function find_previous_command() {
    prev_cmd=$PROJECT_DIR/.kani2/prev_cmd
    if [[ -f $prev_cmd ]]; then
        cat $prev_cmd # 内容を確認する．
        rm $prev_cmd  # 読み出し後，削除する．
    fi
}

$KANI_HOME/scripts/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    prevcmd=$(find_previous_command)
    echo "prev cmd: $prevcmd, status: $1" # 終了ステータスは $1.
    if [[ $prevcmd =~ hoge* && $1 -ne 0 ]]; then # 終了ステータスが正常0以外の時，回数をカウントする．
      # hogeの所gccに変更する．
      echo "$prevcmd : $1" >> $PROJECT_DIR/.kani2/failures_compilation # 失敗回数をカウントするように記述していく．
    elif [[ $prevcmd =~ test* && $1 -eq 0 ]]; then
      rm $PROJECT_DIR/.kani2/failures_compilation
    fi
    pyc="python3 $KANI_HOME/analyses/analyses.py"
    eval $pyc
fi