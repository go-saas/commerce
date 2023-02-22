FROM swaggerapi/swagger-ui:latest
COPY ./openapi /usr/share/nginx/html/api
ENV URLS="[{url: 'api/commerce-merged.swagger.json', name: 'commerce'}]"
EXPOSE 8080