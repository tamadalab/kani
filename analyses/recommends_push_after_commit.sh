#! /bin/sh

function recommends_push_ja() {
    echo "全てのファイルがコミットされたため，\"git push\"を実行し，変更内容をサーバに送りましょう．"
    exit 1
}

function recommends_push() {
    echo "Recommend \"git push\", since all files are committed, but not pushed."
    exit 1
}

# count not committed, and not staged files.
fileCount=$(git status -s | grep -v \?\? | wc -l)

if [[ $fileCount -ne 0 ]]; then
    # not recommends push, since non-committed files exist.
    exit 0
fi

# get the current branch name.
currentBranch=$(git rev-parse --abbrev-ref HEAD)

# get remote branch name
remoteBranch=$(git branch -r --format "%(refname)" | grep $currentBranch)
if [[ remoteBranch == "" ]]; then
    # the user does not push the current branch, yet.
    recommends_push_ja
fi

# find the commit count of not pushed.
commitCount=$(git log --oneline $remoteBranch..$currentBranch | wc -l | xargs)

if [[ commitCount -ne 0 ]]; then
    # this repository has not pushed commits.
    recommends_push_ja
fi

exit 0
