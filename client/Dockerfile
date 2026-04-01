# ---------- BUILD STAGE ----------
FROM node:24-alpine AS builder
WORKDIR /app
COPY package*.json ./
# Use clean-install to ensure no local junk interferes
RUN npm ci 
COPY . .

ARG VITE_BACKEND_BASE_URL

ENV VITE_BACKEND_BASE_URL=$VITE_BACKEND_BASE_URL

# This build NOW uses the base: "./" you added
RUN npm run build

# ---------- RUNNER STAGE ----------
FROM node:20-alpine AS runner
WORKDIR /app

# Copy the static files
COPY --from=builder /app/dist ./dist
# Copy the config so 'preview' sees the 'allowedHosts'
COPY --from=builder /app/vite.config.* ./
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/node_modules ./node_modules

EXPOSE 5173

# Force the preview to use the exact port and host
CMD ["npm", "run", "preview", "--", "--host", "0.0.0.0", "--port", "5173"]
