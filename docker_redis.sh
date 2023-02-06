#!/bin/sh

docker run -d \
-h redis \
-e REDIS_PASSWORD=secret \
-v /redis_data/conf:/usr/local/etc/redis \
-p 6379:6379 \
--name redis \
--restart unless-stopped \
--network development \
-m 256m --memory-reservation=128m \
redis:6.2.5-alpine /bin/sh -c 'redis-server --appendonly yes --requirepass ${REDIS_PASSWORD}'
