# ---------- BUILD STAGE ----------
FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build


# ---------- RUNNER STAGE ----------
FROM node:20-alpine AS runner

WORKDIR /app

# copy built app
COPY --from=builder /app/dist ./dist

# copy deps + package
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

EXPOSE 5173

CMD ["npm", "run", "preview", "--", "--host", "--port", "5173"]
