#!/bin/sh
APP="heritago"

printf "\nBuilding: ${APP}\n"
time go build -o ../build/${APP} ../main.go
printf "\nBuilt: ${APP} size:"
ls -lah ../build | awk '{print $5}'
printf "\nDone building: ${APP}\n\n"