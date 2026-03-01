# 1. 빌드 스테이지
FROM node:22.18.0-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# 2. 실행 스테이지
FROM node:22.18.0-alpine AS runner
WORKDIR /app

ENV NODE_ENV = production

# standalone 모드에 필요한 파일들만 복사
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

RUN mkdir -p /usr/share/next-resources && chmod 777 /usr/share/next-resources

EXPOSE 3001
CMD ["node", "server.js"]