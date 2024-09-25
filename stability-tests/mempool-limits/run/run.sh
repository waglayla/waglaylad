#!/bin/bash

APPDIR=/tmp/waglaylad-temp
waglaylad_RPC_PORT=29587

rm -rf "${APPDIR}"

waglaylad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${waglaylad_RPC_PORT}" --profile=6061 &
waglaylad_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${waglaylad_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $waglaylad_PID

wait $waglaylad_PID
waglaylad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "waglaylad exit code: $waglaylad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $waglaylad_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
