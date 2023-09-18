FROM ubuntu
COPY aoctl /bin/aoctl
RUN mkdir aoctl && \
    dpkg --add-architecture i386 && \
    apt update && \
WORKDIR /aoctl
ENTRYPOINT ["aoctl"]
