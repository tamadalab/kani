echo "test2"

if [[ $* == *"git"*"status" ]]; then
  cmd="python3 ~/project/tmd/kani/.kani/test.py s"
  eval $cmd
fi
