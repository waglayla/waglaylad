#!/bin/bash
rm -rf /tmp/waglaylad-temp

NUM_CLIENTS=128
waglaylad --devnet --appdir=/tmp/waglaylad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
waglaylad_PID=$!
waglaylad_KILLED=0
function killwaglayladIfNotKilled() {
  if [ $waglaylad_KILLED -eq 0 ]; then
    kill $waglaylad_PID
  fi
}
trap "killwaglayladIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $waglaylad_PID

wait $waglaylad_PID
waglaylad_EXIT_CODE=$?
waglaylad_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "waglaylad exit code: $waglaylad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $waglaylad_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
