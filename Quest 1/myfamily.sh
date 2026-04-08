#!/bin/bash

curl -s https://01.tomorrow-school.ai/assets/superhero/all.json \
| jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id|tonumber)) | .connections.relatives'