FROM golang

ADD ./web.exe /usr/local/bin/

ENTRYPOINT /usr/local/bin/web.exe

EXPOSE 8080
