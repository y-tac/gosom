FROM scratch

COPY server /server
COPY front  /front
COPY config.json /config.json

ENV GOROOT /usr/local/go

ENTRYPOINT ["/server"]
