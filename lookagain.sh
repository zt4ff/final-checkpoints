#!/bin/bash

find . -type f -name "*.sh" \
    | xargs -n 1 basename \
    | cut -d '.' -f 1 \
    | sort -r
