FROM ubuntu:latest
LABEL authors="beppo"

ENTRYPOINT ["top", "-b"]