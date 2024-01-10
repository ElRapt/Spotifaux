#!/bin/bash

OPENDIR=$(pwd)

# Launch API Users
gnome-terminal --tab --title="API Users" --working-directory="$OPENDIR" -- bash -c "cd ./users/ && go run ./cmd/main.go; exec bash"

# Launch Flask
gnome-terminal --tab --title="Flask" --working-directory="$OPENDIR" -- bash -c "cd ./flask_base/ && PYTHONPATH=$PYTHONPATH:$(pwd) python3 src/app.py; exec bash"

# Launch Frontend
gnome-terminal --tab --title="Frontend" --working-directory="$OPENDIR" -- bash -c "cd ./tp_middleware_front-main/ && npm run dev; exec bash"