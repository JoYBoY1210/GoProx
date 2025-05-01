
$downloadUrl = "https://github.com/JoYBoY1210/GoProx/releases/download/v1.0.0/GoProxSetup.exe"

$destination = "$env:USERPROFILE\GoProx.exe"


Write-Host "Downloading GoProx.exe..."
Invoke-RestMethod -Uri $downloadUrl -OutFile $destination


if (Test-Path $destination) {
    Write-Host "GoProx.exe downloaded successfully to $destination."
   
    $env:Path += ";$destination"
    Write-Host "GoProx has been added to your system PATH."
} else {
    Write-Host "Failed to download GoProx.exe."
}
