echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories
echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

apk update && apk upgrade -y

apk add go=1.23.5-r0
