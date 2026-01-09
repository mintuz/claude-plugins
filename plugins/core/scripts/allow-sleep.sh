#!/bin/bash

# Re-enable Mac sleep by killing the caffeinate process and cleanup the PID file

if [ -f /tmp/claude_caffeinate.pid ]; then
    kill $(cat /tmp/claude_caffeinate.pid) 2>/dev/null
    rm /tmp/claude_caffeinate.pid
fi

