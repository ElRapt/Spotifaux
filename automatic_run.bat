@echo off
SET OPENDIR=%cd%

:: Launch API Users
start "API Users" cmd /k "cd users/cmd && go run main.go"

:: Launch API Musics
start "API Musics" cmd /k "cd musics/cmd && go run main.go"

:: Launch API Ratings
start "API Ratings" cmd /k "cd ratings/cmd && go run main.go"

:: Launch Flask
start "Flask" cmd /k "cd flask_base && set PYTHONPATH=%PYTHONPATH%;%cd% && python src/app.py"