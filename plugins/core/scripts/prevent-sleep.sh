#!/bin/bash

# Prevent Mac from sleeping while Claude Code is running (for up to 1 hour)
# Kill any previously running caffeinate process started by this script

if [ -f /tmp/claude_caffeinate.pid ]; then
    old_pid=$(cat /tmp/claude_caffeinate.pid)
    if ps -p "$old_pid" > /dev/null 2>&1; then
        # Ensure the process is caffeinate (full command line check)
        if ps -p "$old_pid" -o args= | grep -q '^caffeinate'; then
            kill "$old_pid" 2>/dev/null
        fi
    fi
    rm -f /tmp/claude_caffeinate.pid
fi

nohup caffeinate -i -t 3600 > /dev/null 2>&1 &
echo $! > /tmp/claude_caffeinate.pid

