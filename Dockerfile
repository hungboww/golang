FROM golang:1.18 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]
# REDIS
#REDIS_VERSION=6.2-alpine
#REDIS_SECRET_SESSION_KEY=REDIS_SECRET_SESSION_KEY
#REDISDB_PORT=6379
#REDISDB_HOST=localhost
#REDIS_EXPIRE_TIME=3600
#REDIS_PASSWORD=
#BROKER_URL='redis://' + REDISDB_HOST + ':' + REDISDB_PORT + '/0'
#
##WEB
#DATABASE_NAME=golang
#DATABASE_USER=golang
#DATABASE_PASSWORD=root
#DATABASE_HOST=db
#DATABASE_PORT=3306
#
##MYSQL
#MYSQL_TCP_PORT=3306
#MYSQL_DATABASE=blog
#MYSQL_ROOT_PASSWORD=root
#MYSQL_PASSWORD=root
#MYSQL_NAME=root
#MYSQL_VERSION=5.7.38-oracle
#
##CELERY
#DEBUG=1
#SECRET_KEY=11111111112222222222222
#DJANGO_ALLOWED_HOSTS=*
#CELERY_BROKER=redis://redis:6379/0
#CELERY_BACKEND=redis://redis:6379/0
#
#PORT=3000
#SECRET=MIIBIjANBgkqhkiG9w0BAQEFAAOCAxxxxxx
#DB=root:root@tcp(172.22.0.4:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local
#DB2=root:root@tcp(172.20.0.3:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local
