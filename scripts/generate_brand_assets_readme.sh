#!/usr/bin/env bash
#
# update readme file in the assets/brand folder

# exit on error
set -e

IFS=$""

# get the repo root
ROOT="$(readlink -f ${BASH_SOURCE[0]%/*}/../)"

cd "$ROOT/assets/brand"

README_CONTENT=$(cat <<EOF
# Hasura Brand Assets

| Name | Asset |
| ---- | ----- |
EOF
)

for svg in *.svg; do
    if [[ "$svg" = *"white"* ]]; then
        BLACK='style="background-color: black;"'
    else
        BLACK=""
    fi
    README_CONTENT=$(cat <<EOF
$README_CONTENT
| $svg | <img src="$svg" width="150px" $BLACK /> |
EOF
)
done

echo $README_CONTENT > "$ROOT/assets/brand/README.md"
