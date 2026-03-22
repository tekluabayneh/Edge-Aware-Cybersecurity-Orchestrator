param(
    [string]$Version,
    [string]$Arch,
    [string]$Binary,
    [string]$Dist
)

Write-Host "=== INSTALLER BUILD START [$Arch] ==="
Write-Host "Binary: $Binary"
Write-Host "Dist Target: $Dist"

# 1. Resolve paths relative to project root (../ from agent dir)
$RootDir = Resolve-Path (Join-Path $PSScriptRoot "..")
$AgentExe = Join-Path $RootDir "agent.exe"
$IssFile = Join-Path $RootDir "installer.iss"

# 2. Copy binary to root for ISCC
Write-Host "Copying binary to root: $AgentExe"
Copy-Item -Path $Binary -Destination $AgentExe -Force

# 3. Run ISCC
Write-Host "Running ISCC: $IssFile"
& "iscc.exe" "/DAppVersion=$Version" "/DArch=$Arch" $IssFile

if ($LASTEXITCODE -ne 0) {
    Write-Error "ISCC failed with exit code $LASTEXITCODE"
    exit 1
}

# 4. Locate Output
$OutputDir = Join-Path $RootDir "Output_$Arch"
$OutputFile = Join-Path $OutputDir "Agent_Setup_$Arch.exe"

Write-Host "Checking output: $OutputFile"
if (-not (Test-Path $OutputFile)) {
    Write-Error "OUTPUT NOT FOUND: $OutputFile"
    exit 1
}

# 5. Ensure Dist directory exists and copy
$DistPath = Join-Path $RootDir $Dist
if (-not (Test-Path $DistPath)) {
    New-Item -ItemType Directory -Force -Path $DistPath | Out-Null
}

$DistFile = Join-Path $DistPath "agent_installer_$Version_$Arch.exe"
Write-Host "Copying to Dist: $DistFile"
Copy-Item -Path $OutputFile -Destination $DistFile -Force

if (-not (Test-Path $DistFile)) {
    Write-Error "INSTALLER NOT IN DIST: $DistFile"
    exit 1
}

# 6. Cleanup
Remove-Item -Path $AgentExe -Force

Write-Host "=== INSTALLER BUILD COMPLETE [$Arch] ==="
