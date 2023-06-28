FROM node:lts as dependencies
WORKDIR /app
COPY fron-end/package.json fron-end/package-lock.json ./
RUN npm install --omit=dev

FROM node:lts as builder
WORKDIR /app
COPY fron-end/ ./
COPY --from=dependencies /app/node_modules ./node_modules
RUN npm run build

FROM node:lts as runner
WORKDIR /app
ENV NODE_ENV production
# If you are using a custom next.config.js file, uncomment this line.
# COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

EXPOSE 8080

CMD ["npm", "start"]
