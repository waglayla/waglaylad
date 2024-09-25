#!/bin/bash
rm -rf /tmp/waglaylad-temp

waglaylad --devnet --appdir=/tmp/waglaylad-temp --profile=6061 --loglevel=debug &
waglaylad_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $waglaylad_PID

wait $waglaylad_PID
waglaylad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "waglaylad exit code: $waglaylad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $waglaylad_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
