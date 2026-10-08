@echo off
REM ============================================
REM golyadmin integrated deployment build (Windows)
REM Output to build\:
REM   golyadmin.exe     executable
REM   dist\             frontend pages
REM   config.pro.yaml   production config
REM   media\            upload dir (empty)
REM Deploy: upload ALL contents of build\ to server
REM Run: golyadmin.exe start -c config.pro.yaml
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

echo [3/4] Building backend (Windows exe)...
cd backend
go build -ldflags "-s -w" -o ..\build\golyadmin.exe main.go
if errorlevel 1 (
    echo Backend build FAILED
    pause
    exit /b 1
)

echo [4/4] Done.
echo.
echo Deploy artifacts in build\:
dir /b "%~dp0build"
echo.
echo Upload ALL contents of build\ to the server, then run:
echo   golyadmin.exe start -c config.pro.yaml
pause
