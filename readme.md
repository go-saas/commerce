# Quick Start

### For Microservice

```
docker compose -f docker-compose.yml -f docker-compose.ms.yml -f docker-compose.kafka.yml -f docker-compose.tracing.yml pull && \
docker compose -f docker-compose.yml -f docker-compose.ms.yml -f docker-compose.kafka.yml -f docker-compose.tracing.yml up -d
```

Or with build
```
docker compose -f docker-compose.yml -f docker-compose.ms.yml -f docker-compose.kafka.yml  -f docker-compose.tracing.yml up -d --build
```