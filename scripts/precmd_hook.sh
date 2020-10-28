#! /bin/zsh

KANI_HOME=/usr/local/opt/kani
PROJECT_DIR=$($KANI_HOME/scripts/find-project-dir.sh)
FAILURES_DIR=$PROJECT_DIR/.kani/failures_compilation

function find_previous_command() {
    prev_cmd=$PROJECT_DIR/.kani/prev_cmd
    if [[ -e $prev_cmd ]]; then
        # cat $prev_cmd # 内容を確認する．
        rm $prev_cmd  # 読み出し後，削除する．
    fi
}

$KANI_HOME/scripts/is-target-project.sh
if [[ $? -ne 0 ]]; then
    exit 0
else
    prevcmd=$(find_previous_command)
    echo "$(date "+%H:%M:%S") $prevcmd" >> $PROJECT_DIR/.kani/test.log
    # echo "prev cmd: $prevcmd, status: $1" # (デバッグ用)終了ステータスは $1.
    count=0 # エラーのカウント
    if [[ $prevcmd =~ gcc* || $prevcmd =~ clang* && $1 -ne 0 ]]; then # 終了ステータスが正常0以外の時，回数をカウントする．
      # hogeの所gccに変更する．
      echo "$prevcmd : $1" >> $FAILURES_DIR # 失敗回数をカウントするように記述していく．
    elif [[ $prevcmd =~ gcc* || $prevcmd =~ clang* && $1 -eq 0 && -e $FAILURES_DIR ]]; then # エラーが直った場合，連続エラー回数を記録してファイルを削除
      count=$(wc -l $FAILURES_DIR)
      rm $FAILURES_DIR
    fi
    pyc="python3 $KANI_HOME/analyses/analyses.py $count"
    eval $pyc
fi
