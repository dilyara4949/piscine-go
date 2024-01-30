#!/bin/bash

curl -s "https://zero.academie.one/assets/superhero/all.json" > q
jq --argjson hero_id "$HERO_ID"  '.[] | select(.id == $hero_id) | .connections.relatives' q | tr -d '"'
# echo "uyih\n;lkdmfv;"