@echo off
echo =====================================
echo Checking Docker Status
echo =====================================
echo.

docker version >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] Docker is not installed!
    pause
    exit /b 1
)

docker ps >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] Docker Desktop is not running!
    echo.
    echo Please start Docker Desktop:
    echo 1. Open Docker Desktop from Start Menu
    echo 2. Wait for it to fully start
    echo 3. Run this script again
    echo.
    echo Starting Docker Desktop for you...
    start "" "C:\Program Files\Docker\Docker\Docker Desktop.exe"
    echo.
    echo Please wait for Docker to start, then run this script again.
    pause
    exit /b 1
)

echo [OK] Docker is running!
echo.
echo Checking existing images...
docker images | findstr "gocv paddle myid"
echo.
echo You can now run: docker-compose build
pause