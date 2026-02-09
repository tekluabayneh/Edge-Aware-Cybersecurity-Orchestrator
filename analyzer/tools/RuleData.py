DESTRUCTIVE_KEYWORDS = [
    "rm -rf",
    "delete /f",
    "del /f",
    "format ",
    "mkfs",
    "diskpart",
    "dd if=",
    "wipe",
    "shred",
    "erase",
    "destroy",
    "remove-all",
    "truncate"
]

# Persistence / auto-start
PERSISTENCE_KEYWORDS = [
    "startup",
    "autorun",
    "autostart",
    "onboot",
    "onstart",
    "register-service",
    "create service",
    "enable service",
    "install service",
    "schedule",
    "task scheduler",
    "cron",
    "launch",
    "daemon",
    "background"
]

# Download + remote execution
REMOTE_EXEC_KEYWORDS = [
    "download",
    "fetch",
    "curl",
    "wget",
    "invoke-webrequest",
    "http://",
    "https://",
    "| sh",
    "| bash",
    "| powershell",
    "exec",
    "execute",
    "run command",
    "shell"
]

# Obfuscation / hiding intent
OBFUSCATION_KEYWORDS = [
    "base64",
    "decode",
    "encode",
    "decrypt",
    "encrypt",
    "eval",
    "exec(",
    "obfuscate",
    "compress",
    "pack",
    "unpack",
    "hidden",
    "silent"
]

# Privilege escalation
PRIVILEGE_KEYWORDS = [
    "admin",
    "administrator",
    "root",
    "elevated",
    "sudo",
    "runas",
    "setuid",
    "grant-permission",
    "full-access",
    "system-level"
]

# Surveillance / spying behavior
SURVEILLANCE_KEYWORDS = [
    "keylog",
    "keystroke",
    "screen",
    "screenshot",
    "capture",
    "record",
    "monitor",
    "track",
    "spy",
    "watch",
    "observe"
]

# Network / backdoor behavior
NETWORK_KEYWORDS = [
    "connect",
    "socket",
    "bind",
    "listen",
    "remote",
    "reverse",
    "backdoor",
    "tunnel",
    "proxy",
    "relay",
    "command-and-control",
    "c2"
]

# Resource abuse / crypto mining
RESOURCE_ABUSE_KEYWORDS = [
    "mine",
    "miner",
    "hashrate",
    "cpu usage",
    "gpu usage",
    "worker",
    "pool",
    "background job"
]

# Optional: collect everything
ALL_CONTENT_KEYWORDS = {
    "destructive": DESTRUCTIVE_KEYWORDS,
    "persistence": PERSISTENCE_KEYWORDS,
    "remote_exec": REMOTE_EXEC_KEYWORDS,
    "obfuscation": OBFUSCATION_KEYWORDS,
    "privilege": PRIVILEGE_KEYWORDS,
    "surveillance": SURVEILLANCE_KEYWORDS,
    "network": NETWORK_KEYWORDS,
    "resource_abuse": RESOURCE_ABUSE_KEYWORDS,
}

