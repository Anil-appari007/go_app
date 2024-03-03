FROM node:20.11.1-alpine3.19 as build
WORKDIR /app
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

RUN npm ci
COPY . .

RUN npm run build



FROM nginx:alpine
COPY --from=build /app/.nginx/nginx.conf /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*

COPY --from=build /app/build/* .

ENTRYPOINT ["nginx", "-g", "daemon off;"]