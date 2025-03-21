@echo off
if "%1" == "up" (
    docker-compose --env-file .env.dev up --build -d
) else if "%1" == "down" (
    docker-compose down
) else if "%1" == "build" (
    docker-compose --env-file .env.dev build
) else (
    echo Invalid command. Use 'up', 'down', or 'build'.
)