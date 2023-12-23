@echo off
SET OPENDIR=%cd%

:: Launch API Ratings
start "API Ratings" cmd /k "cd ratings_api && go run ./cmd/main.go"

:: Launch Flask
start "Flask" cmd /k "cd flask_base && set PYTHONPATH=%PYTHONPATH%;%cd% && python src/app.py"

:: Launch Frontend
start "Frontend" cmd /k "cd tp_middleware_front-main && npm run dev"
