FROM golang

ADD ./web.exe /usr/local/bin/
ADD ./web/index.html /var/www/html/index.html


ENTRYPOINT /usr/local/bin/web.exe

EXPOSE 8089 
