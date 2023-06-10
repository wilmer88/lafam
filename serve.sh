#!/bin/sh

gin --port 5000 --path . --build ./src/server/ --i --all &

wait