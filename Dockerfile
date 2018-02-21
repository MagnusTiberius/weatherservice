FROM golang

ADD ./web.exe /usr/local/bin/
ADD ./web/index.html ~/index.html

ENTRYPOINT /usr/local/bin/web.exe

EXPOSE 8088
