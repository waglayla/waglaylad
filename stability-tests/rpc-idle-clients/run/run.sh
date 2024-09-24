#!/bin/bash
rm -rf /tmp/waglaylad-temp

NUM_CLIENTS=128
waglaylad --devnet --appdir=/tmp/waglaylad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
pyipad_PID=$!
pyipad_KILLED=0
function killpyipadIfNotKilled() {
  if [ $pyipad_KILLED -eq 0 ]; then
    kill $pyipad_PID
  fi
}
trap "killpyipadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $pyipad_PID

wait $pyipad_PID
pyipad_EXIT_CODE=$?
pyipad_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "waglaylad exit code: $pyipad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $pyipad_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
