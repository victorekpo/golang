#!/bin/bash

while true; do
  go run main.go
  if [ $? -ne 0 ]; then
    echo "Error encountered. Exiting loop."
    break
  fi
done