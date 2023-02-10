export MSYS_NO_PATHCONV=1
rm -rf api
rm -rf models
docker run --rm -v  ${PWD}/../../../:/local openapitools/openapi-generator-cli:v6.0.0 generate -i /local/openapi/commerce-merged.swagger.json -g typescript-axios -o /local/web/packages/commerce-api -c /local/web/packages/commerce-api/tools/config.json
