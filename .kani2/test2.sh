#!/bin/zsh
echo "test2"

if [[ $* == *"git"*"status" ]]; then
  cmd="python3 ~/project/tmd/kani/.kani/test.py s"
  eval $cmd
fi
if [[ $* == *"kani"*"status" ]]; then
  TOOL=$(<toolswitch.txt)
  echo "kaniの機能は現在$TOOLです"
fi
if [[ $* == *"kani"*"on" ]]; then
  echo "on" > toolswitch.txt 
  echo "kaniの機能をonにしました"
fi
if [[ $* == *"kani"*"off" ]]; then
  echo "off" > toolswitch.txt
  echo "kaniの機能をoffにしました" 
fi