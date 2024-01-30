curl   -s  "https://zero.academie.one/assets/superhero/all.json" > q
jq -r '.[] | select(.id == 70) | "\"\(.name)\""'  q