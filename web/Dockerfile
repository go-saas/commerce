FROM node:16.17.0 AS build
ARG ROUTE_BASE
ENV ROUTE_BASE $ROUTE_BASE

WORKDIR /usr/app

RUN npm install -g pnpm@7.9.0


COPY .npmrc ./
COPY pnpm-lock.yaml ./

RUN pnpm fetch

COPY . .

RUN pnpm install --prefer-offline
RUN pnpm run -r build
RUN pnpm run build

FROM nginx:alpine
COPY --from=build /usr/app/dist/ /usr/share/nginx/html/
COPY ./docker/nginx/nginx.conf /etc/nginx/conf.d/

EXPOSE 8080
