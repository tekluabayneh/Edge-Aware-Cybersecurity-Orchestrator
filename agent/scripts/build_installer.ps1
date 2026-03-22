param(
    [string]$Version,
    [string]$Arch,
    [string]$Binary
)

Write-Host "=== INSTALLER BUILD START [$Arch] ==="
Write-Host "Binary: $Binary"
Write-Host "Version: $Version"

# Resolve project root: script is at agent/scripts/, so root is ../
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$RootDir = Resolve-Path (Join-Path $ScriptDir "..")

# Paths relative to project root
$AgentExe = Join-Path $RootDir "agent.exe"
$IssFile = Join-Path $RootDir "installer.iss"
$OutputDir = Join-Path $RootDir "Output_$Arch"
$OutputFile = Join-Path $OutputDir "Agent_Setup_$Arch.exe"
$DistDir = Join-Path $RootDir "dist"
$DistFile = Join-Path $DistDir "agent_installer_$Version_$Arch.exe"

# 1. Copy built binary to root for ISCC
Write-Host "Copying binary to root: $AgentExe"
Copy-Item -Path $Binary -Destination $AgentExe -Force

# 2. Run ISCC
Write-Host "Running ISCC: $IssFile"
& "iscc.exe" "/DAppVersion=$Version" "/DArch=$Arch" $IssFile

if ($LASTEXITCODE -ne 0) {
    Write-Error "ISCC failed with exit code $LASTEXITCODE"
    exit 1
}

# 3. Verify output exists
Write-Host "Checking output: $OutputFile"
if (-not (Test-Path $OutputFile)) {
    Write-Error "OUTPUT NOT FOUND: $OutputFile"
    exit 1
}

# 4. Ensure dist directory exists
if (-not (Test-Path $DistDir)) {
    New-Item -ItemType Directory -Force -Path $DistDir | Out-Null
}

# 5. Copy installer to dist
Write-Host "Copying to dist: $DistFile"
Copy-Item -Path $OutputFile -Destination $DistFile -Force

if (-not (Test-Path $DistFile)) {
    Write-Error "INSTALLER NOT IN DIST: $DistFile"
    exit 1
}

# 6. Cleanup temporary agent.exe
Remove-Item -Path $AgentExe -Force

Write-Host "=== INSTALLER BUILD COMPLETE [$Arch] ==="
