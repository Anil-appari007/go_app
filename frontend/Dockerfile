FROM node:20.11.1-alpine3.19 as build
WORKDIR /app
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

RUN npm ci
COPY . .

RUN npm run build
RUN ls -lrt /app/build/
RUN ls -lrt /app/build/static



FROM nginx
COPY --from=build /app/.nginx/nginx.conf /etc/nginx/conf.d/default.conf

# WORKDIR /usr/share/nginx/html
# RUN rm -rf ./*

RUN rm -rf /usr/share/nginx/html

# COPY --from=build /app/build/* /usr/share/nginx/html/.
COPY --from=build /app/build /usr/share/nginx/html

RUN ls -lrt /usr/share/nginx/html/
RUN ls -lrt /usr/share/nginx/html/static
# ENTRYPOINT ["nginx", "-g", "daemon off;"]