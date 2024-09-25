#!/bin/bash
rm -rf /tmp/waglaylad-temp

waglaylad --devnet --appdir=/tmp/waglaylad-temp --profile=6061 &
waglaylad_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:12611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $waglaylad_PID

wait $waglaylad_PID
waglaylad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "waglaylad exit code: $waglaylad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $waglaylad_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
