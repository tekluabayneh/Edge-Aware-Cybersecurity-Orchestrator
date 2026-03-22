param(
    [string]$Version,
    [string]$Arch,
    [string]$Binary,
    [string]$Dist
)

Write-Host "=== INSTALLER BUILD START [$Arch] ==="
Write-Host "Binary: $Binary"
Write-Host "Dist: $Dist"

# 1. Copy binary to root for ISCC (ISCC often expects relative paths)
$RootDir = Split-Path -Parent $PSScriptRoot
$AgentExe = Join-Path $RootDir "agent.exe"

Write-Host "Copying binary to root..."
Copy-Item -Path $Binary -Destination $AgentExe -Force

# 2. Run ISCC
$IssFile = Join-Path $RootDir "installer.iss"
Write-Host "Running ISCC..."
& "iscc.exe" "/DAppVersion=$Version" "/DArch=$Arch" $IssFile

if ($LASTEXITCODE -ne 0) {
    Write-Error "ISCC failed with exit code $LASTEXITCODE"
    exit 1
}

# 3. Locate Output
$OutputDir = Join-Path $RootDir "Output_$Arch"
$OutputFile = Join-Path $OutputDir "Agent_Setup_$Arch.exe"

Write-Host "Checking output: $OutputFile"
if (-not (Test-Path $OutputFile)) {
    Write-Error "OUTPUT NOT FOUND: $OutputFile"
    exit 1
}

# 4. Copy to Dist
$DistFile = Join-Path $Dist "agent_installer_$Version_$Arch.exe"
Write-Host "Copying to Dist: $DistFile"
Copy-Item -Path $OutputFile -Destination $DistFile -Force

if (-not (Test-Path $DistFile)) {
    Write-Error "INSTALLER NOT IN DIST: $DistFile"
    exit 1
}

# 5. Cleanup
Remove-Item -Path $AgentExe -Force

Write-Host "=== INSTALLER BUILD COMPLETE [$Arch] ==="
