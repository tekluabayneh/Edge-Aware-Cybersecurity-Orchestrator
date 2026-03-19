[Setup]
AppName=Agent Orchestrator
AppVersion={#AppVersion}
AppPublisher=Your Name / Company
AppPublisherURL=https://github.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator
DefaultDirName={sd}\AgentOrchestrator
DefaultGroupName=Agent Orchestrator
PrivilegesRequired=admin
OutputBaseFilename=Agent_Setup_{#Arch}
OutputDir=Output_{#Arch}
Compression=lzma
SolidCompression=yes
WizardStyle=modern
UninstallDisplayIcon={app}\agent.exe
UninstallFilesDir={app}\unins

[Files]
; The "Source" is where ISCC finds the file on the GitHub runner
; The "DestDir" is {app} (C:\Program Files\...) on YOUR computer
Source: "{#SourcePath}"; DestDir: "{app}"; DestName: "agent.exe"; Flags: ignoreversion

; Optional: add other files from project root if needed
; Source: "README.md"; DestDir: "{app}"; Flags: ignoreversion
; Source: "LICENSE"; DestDir: "{app}"; Flags: ignoreversion

[Dirs]
; Folder with full user permissions (your original requirement)
Name: "{app}\internal\register"; Permissions: users-full

[Icons]
Name: "{group}\Agent Orchestrator"; Filename: "{app}\agent.exe"
Name: "{group}\Uninstall Agent Orchestrator"; Filename: "{uninstallexe}"
Name: "{autodesktop}\Agent Orchestrator"; Filename: "{app}\agent.exe"; Tasks: desktopicon

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Code]
// Optional: small Pascal script to make uninstall cleaner (optional but nice)
procedure CurStepChanged(CurStep: TSetupStep);
begin
  if CurStep = ssPostInstall then
  begin
  end;
end;

