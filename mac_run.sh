#!/bin/bash

# Save the current directory
OPENDIR=$(pwd)

# Launch API Users
(cd users && go run ./cmd/main.go) &

# Launch API Musics

(cd musics && go run ./cmd/main.go) &

# Kill ports 5173, 8080, 8081

kill -9 $(lsof -t -i:5173) &

kill -9 $(lsof -t -i:8080) &

kill -9 $(lsof -t -i:8081) &

# Launch Flask
(cd flask_base && export PYTHONPATH=$PYTHONPATH:$pwd && python3 src/app.py) &

# Launch Frontend
(cd tp_middleware_front-main && npm run dev) &
