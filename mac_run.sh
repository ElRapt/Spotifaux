#!/bin/bash

# Save the current directory
OPENDIR=$(pwd)

# Launch API Users
(cd users && go run ./cmd/main.go) &

# Launch API Musics

(cd musics && go run ./cmd/main.go) &

# Launch Flask
(cd flask_base && export PYTHONPATH=$PYTHONPATH:$pwd && python src/app.py) &

# Launch Frontend
(cd tp_middleware_front-main && npm run dev) &
