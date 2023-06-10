#!/bin/sh

ng serve &
gin --port 5000 --path . --build ./src/server/ --i --all &

wait