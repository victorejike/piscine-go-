#!/bin/bash
if [ - z "$HERO_ID" ]; then echo "Error: HERO_ID is not set" 
exit 1 
fi
result=$(curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives | gsub("\n"; "\\n")')
if [ -z "$result" ] || [ "$result" = "null" ]; then echo "No relatives found for HERO_ID=$HERO_ID"
exit 1
else
echo -n "$result"
fi
