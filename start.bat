@echo off
echo =====================================
echo ID Card Recognition System
echo =====================================
echo.

echo [1] Building Docker images...
docker-compose build
if %errorlevel% neq 0 (
    echo Build failed!
    pause
    exit /b 1
)

echo.
echo [2] Starting services...
docker-compose up -d
if %errorlevel% neq 0 (
    echo Failed to start services!
    pause
    exit /b 1
)

echo.
echo [3] Services are running!
echo.
echo Application: http://localhost:8080
echo PaddleOCR: http://localhost:8866
echo.
echo To stop: docker-compose down
echo.
pause