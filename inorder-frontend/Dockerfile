FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm i
COPY . .
EXPOSE 3000
RUN npm install -g serve
RUN npm run build

FROM node:18-alpine
WORKDIR /app
COPY --from=0 /app/dist ./dist
RUN npm install -g serve
CMD ["serve", "-s", "dist"]