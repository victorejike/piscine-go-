#!/bin/bash
find . -type f -name "*.sh" \
| xargs -n 1 basename \
| cut -d. -f1 \
| sort -r
