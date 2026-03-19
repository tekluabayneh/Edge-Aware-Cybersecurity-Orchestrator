[Setup]
AppName=Agent Orchestrator
AppVersion={#AppVersion}
DefaultDirName={autopf}\AgentOrchestrator
DefaultGroupName=Agent Orchestrator
PrivilegesRequired=admin
; This names the output file based on the architecture (amd64 or arm64)
OutputBaseFilename=Agent_Setup_{#Arch}
Compression=lzma
SolidCompression=yes

[Files]
; Using the full path sent by GoReleaser ({#SourcePath})
Source: "{#SourcePath}"; DestDir: "{app}"; Flags: ignoreversion

[Dirs]
; This creates the folder and gives the USER full control forever
Name: "{app}\internal\register"; Permissions: users-full

[Icons]
Name: "{group}\Agent Orchestrator"; Filename: "{app}\agent-orchestrator.exe"
Name: "{commondesktop}\Agent Orchestrator"; Filename: "{app}\agent-orchestrator.exe"
