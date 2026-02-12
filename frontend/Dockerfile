FROM node:20-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# Tailwindのコンパイルを含むビルドを実行
RUN npm run build

CMD ["npm", "run", "start"]
