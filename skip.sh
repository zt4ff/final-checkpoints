#!/bin/bash

# ls: lists files in the current directory. It supports flags like:
# -l for detailed listing with permissions, size, and timestamps
# -a to show hidden files
# -1 to force one file per line
# It is often used to inspect directory structure and file metadata.


# grep: searches text using patterns. It supports:
# -i for case-insensitive search
# -r for recursive search
# -E for extended regex
# grep is powerful for filtering logs, configs, and command output.

# sed: streams and edits text. You can:
# replace text: sed 's/old/new/'
# delete lines: sed '/pattern/d'
# insert or append text around matches
# sed processes text without opening files in editors.

# jq: parses and transforms JSON. It supports:
# selecting fields: jq '.key'
# mapping arrays: jq '.[] | .name'
# pretty printing: jq '.'
# jq is essential whenever APIs return structured JSON.

# awk: processes and transforms text in columns. It supports:
# selecting fields: awk '{print $1}'
# conditional logic: awk '$3 > 100'
# full scripting with patterns and actions
# It is often used for log processing and automation.

ls -1 | sort | xargs ls -l | awk 'NR % 2 == 1'
