#!/bin/bash
# golyadmin stop - kills ONLY the instance running from this directory
cd "$(dirname "$0")"

PID=""
if [ -f golyadmin.pid ]; then
    PID=$(cat golyadmin.pid)
else
    for p in $(pgrep -x golyadmin); do
        if [ "$(readlink /proc/$p/cwd)" = "$(pwd)" ]; then
            PID=$p
        fi
    done
fi

if [ -n "$PID" ] && kill "$PID" 2>/dev/null; then
    echo "stop golyadmin (pid $PID) success"
else
    echo "golyadmin not running in this directory"
fi
