$AppPath = Resolve-Path "..\myagent.exe"

Write-Host "Setting up persistence on Windows..."

sc.exe create MyAgent binPath= "$AppPath" start= auto
sc.exe start MyAgent

Write-Host "Persistence enabled on Windows."
