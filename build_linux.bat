@echo off
REM ============================================
REM golyadmin integrated deployment build (Linux amd64, cross-compile on Windows)
REM Output to build\:
REM   golyadmin         Linux executable (chmod +x)
REM   dist\             frontend pages
REM   config.pro.yaml   production config
REM   media\            upload dir (empty)
REM   restart.sh        start/restart script
REM   stop.sh           stop script
REM Deploy: upload ALL contents of build\ to server
REM Run: chmod +x golyadmin && ./golyadmin start -c config.pro.yaml
REM ============================================
cd /d "%~dp0"

echo [1/4] Building frontend (npm run build)...
cd web
call npm run build
if errorlevel 1 (
    echo Frontend build FAILED
    pause
    exit /b 1
)

echo [2/4] Preparing build directory...
cd /d "%~dp0"
if exist build rmdir /s /q build
mkdir build\media
robocopy web\dist build\dist /E /NFL /NDL /NJH /NJS >nul
if errorlevel 8 (
    echo Copy dist FAILED
    pause
    exit /b 1
)
copy /y backend\config.pro.yaml build\ >nul
copy /y backend\restart.sh build\ >nul
copy /y backend\stop.sh build\ >nul

echo [3/4] Cross-compiling backend (Linux amd64)...
cd backend
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -ldflags "-s -w" -o ..\build\golyadmin main.go
if errorlevel 1 (
    echo Backend build FAILED
    pause
    exit /b 1
)
set GOOS=
set GOARCH=
set CGO_ENABLED=

echo [4/4] Done.
echo.
echo Deploy artifacts in build\:
dir /b "%~dp0build"
echo.
echo Upload ALL contents of build\ to the server, then run:
echo   chmod +x golyadmin ^&^& ./golyadmin start -c config.pro.yaml
pause
