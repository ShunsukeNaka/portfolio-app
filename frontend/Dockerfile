FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# Tailwindのコンパイルを含むビルドを実行
RUN npm run build

# CMD ["npm", "run", "start"]

FROM httpd:2.4-alpine

RUN sed -i 's/#LoadModule proxy_module/LoadModule proxy_module/' /usr/local/apache2/conf/httpd.conf \
    && sed -i 's/#LoadModule proxy_http_module/LoadModule proxy_http_module/' /usr/local/apache2/conf/httpd.conf \
    && echo 'ProxyPass /api/ http://api:8080/' >> /usr/local/apache2/conf/httpd.conf \
    && echo 'ProxyPassReverse /api/ http://api:8080/' >> /usr/local/apache2/conf/httpd.conf \
    && echo 'ServerName localhost' >> /usr/local/apache2/conf/httpd.conf
    
    
COPY --from=builder /app/out /usr/local/apache2/htdocs/

EXPOSE 80