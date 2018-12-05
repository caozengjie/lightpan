#!/usr/bin/env bash
#
docker rm -f yourfs-mysql
#docker run --restart=always --name yourfs-mysql -v /data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -p:33306:3306 -d mysql:5.7
docker rm -f yourfs-redis
docker run  --name yourfs-redis --net=bridge --restart=always -p 36379:6379 -d redis
docker ps |grep yourfs