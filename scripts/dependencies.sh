#!/bin/bash -eu

go mod download

tools=()
while IFS='' read -r line; do tools+=("$line"); done < <(sed -En 's/[[:space:]]+_ "(.*)"/\1/p' tools/tools.go)
for tool in "${tools[@]}"; do
  echo "run: go install tool ($tool)"
  go get "$tool"
done

go mod tidy
