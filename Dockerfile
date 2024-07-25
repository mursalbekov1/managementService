FROM ubuntu:latest
LABEL authors="mursa"

ENTRYPOINT ["top", "-b"]