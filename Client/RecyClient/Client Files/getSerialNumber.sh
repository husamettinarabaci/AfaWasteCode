#!/bin/bash
cat /proc/cpuinfo | perl -n -e '/^Serial\s*:\s([0-9a-f]{16})$/ && print "$1\n"'
