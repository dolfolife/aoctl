FROM ubuntu

LABEL org.opencontainers.image.authors="Rodolfo Sanchez <rodolfo2488@gmail.com>"
LABEL maintainer="Rodolfo Sanchez"
LABEL org.opencontainers.image.source="https://github.com/dolfolife/aoctl"

COPY aoctl /bin/aoctl
RUN mkdir aoctl && \
    dpkg --add-architecture i386 && \
    apt update
WORKDIR /aoctl
ENTRYPOINT ["aoctl"]
