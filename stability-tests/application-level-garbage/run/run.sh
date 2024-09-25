#!/bin/bash
rm -rf /tmp/waglaylad-temp

waglaylad --devnet --appdir=/tmp/waglaylad-temp --profile=6061 --loglevel=debug &
waglaylad_PID=$!
waglaylad_KILLED=0
function killwaglayladIfNotKilled() {
    if [ $waglaylad_KILLED -eq 0 ]; then
      kill $waglaylad_PID
    fi
}
trap "killwaglayladIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:12611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $waglaylad_PID

wait $waglaylad_PID
waglaylad_KILLED=1
waglaylad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "waglaylad exit code: $waglaylad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $waglaylad_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
