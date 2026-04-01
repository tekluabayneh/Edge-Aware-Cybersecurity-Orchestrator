#!/bin/sh
set -e

# 1. Log the URL we are using (check Render logs for this!)
echo "Target Backend URL: $VITE_BACKEND_BASE_URL"

# 2. Check if the dist folder actually exists
if [ ! -d "/app/dist" ]; then
  echo "ERROR: /app/dist folder not found!"
  exit 1
fi

# 3. Search and replace in ALL files inside /app/dist
# We use 'grep' first to see if the placeholder even exists
if grep -r "__VITE_REPLACE_ME__" /app/dist; then
  echo "Placeholder found! Replacing..."
  find /app/dist -type f -exec sed -i "s|__VITE_REPLACE_ME__|$VITE_BACKEND_BASE_URL|g" {} +
  echo "Replacement successful."
else
  echo "ERROR: Could not find __VITE_REPLACE_ME__ in /app/dist. Check your JS code!"
fi

exec "$@"
