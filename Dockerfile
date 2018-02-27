FROM golang

ADD ./run.sh /usr/local/bin/
ADD ./api.exe /usr/local/bin/
ADD ./web.exe /usr/local/bin/
ADD ./web/index.html /var/www/html/index.html

ENTRYPOINT /usr/local/bin/run.sh

EXPOSE 8088
