FROM alpine

ENV TERM=xterm-256color

COPY gloomberg /

ENTRYPOINT ["/gloomberg"]
