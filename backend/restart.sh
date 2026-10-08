#!/bin/bash
# golyadmin restart - kills ONLY the instance running from this directory
# usage: sh restart.sh
cd "$(dirname "$0")"

# locate pid: pid file first, fallback to process whose cwd is this directory
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

if [ -n "$PID" ]; then
    echo "kill golyadmin (pid $PID, this directory only)"
    kill "$PID" 2>/dev/null
    sleep 1
fi

nohup ./golyadmin start -c config.pro.yaml >> access.log 2>&1 &
echo $! > golyadmin.pid
sleep 1

echo "run golyadmin success (pid $(cat golyadmin.pid))"
