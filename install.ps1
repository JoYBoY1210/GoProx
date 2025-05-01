$BinaryUrl = "https://github.com/JoYBoY1210/GoProx/releases/download/v1.0.0/GoProxSetup.exe"
$InstallDir = "$env:USERPROFILE\AppData\Local\Programs\GoProx"
$BinaryPath = "$InstallDir\GoProx.exe"

if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

Invoke-WebRequest -Uri $BinaryUrl -OutFile $BinaryPath
Write-Output "`n✅ Installed GoProx.exe to $InstallDir"


$UserPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)
if ($UserPath -notlike "*$InstallDir*") {
    [System.Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", [System.EnvironmentVariableTarget]::User)
    Write-Output "➡️  Added '$InstallDir' to your PATH."
} else {
    Write-Output "ℹ️  '$InstallDir' already in PATH."
}

Write-Output "`nYou can now run the proxy tool directly from the terminal with 'GoProx'. Please restart your terminal if it doesn't work immediately."
