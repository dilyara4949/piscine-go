curl   -s  "https://zero.academie.one/assets/superhero/all.json" > q
jq -r '.[] | select(.id == 170) | .name  ' q
jq -r '.[] | select(.id == 170) |  .powerstats.power  ' q
jq -r '.[] | select(.id == 170) | .appearance.gender' q