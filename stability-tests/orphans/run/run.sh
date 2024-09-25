#!/bin/bash
rm -rf /tmp/waglaylad-temp

waglaylad --simnet --appdir=/tmp/waglaylad-temp --profile=6061 &
waglaylad_PID=$!

sleep 1

orphans --simnet -alocalhost:12511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $waglaylad_PID

wait $waglaylad_PID
waglaylad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "waglaylad exit code: $waglaylad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $waglaylad_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
