#!/bin/bash

# This script is used to install third party server in a docker container. such as mysql, redis, rabbitmq, etc.


docker run --name mysql -p 3306:3306 \
-v d:/docker/mysql/conf:/etc/mysql/conf.d \
-v d:/docker/mysql/logs:/etc/mysql/logs \
-v d:/docker/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
-d mysql:5.7 --default-authentication-plugin=mysql_native_password


docker run -it -p 4040:4040 pyroscope/pyroscope:latest server