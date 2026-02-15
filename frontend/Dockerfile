FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# Tailwindのコンパイルを含むビルドを実行
RUN npm run build

# CMD ["npm", "run", "start"]

FROM httpd:2.4-alpine

COPY ./my-httpd.conf /usr/local/apache2/conf/httpd.conf

COPY --from=builder /app/out /usr/local/apache2/htdocs/

EXPOSE 80