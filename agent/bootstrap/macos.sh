#!/bin/bash

APP_PATH="$(realpath ../bootstrap/linux.sh)"
PLIST="$HOME/Library/LaunchAgents/com.myagent.plist"

echo "Setting up persistence on macOS..."

cat > $PLIST <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
"http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.myagent</string>

    <key>ProgramArguments</key>
    <array>
        <string>$APP_PATH</string>
    </array>

    <key>RunAtLoad</key>
    <true/>

    <key>KeepAlive</key>
    <true/>
</dict>
</plist>
EOF

launchctl load $PLIST

echo "Persistence enabled on macOS."

