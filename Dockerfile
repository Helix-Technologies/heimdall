FROM golang:latest

ARG HEIMDALL_DIR=/heimdall
ENV HEIMDALL_DIR=$HEIMDALL_DIR

RUN apt-get update -y && apt-get upgrade -y \
    && apt install build-essential git tini -y \
    && mkdir -p /heimdall

WORKDIR ${HEIMDALL_DIR}
COPY . .

RUN make install

COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh

ENV SHELL /bin/bash
EXPOSE 1317 26656 26657

ENTRYPOINT ["tini", "--"]
CMD ["/usr/local/bin/entrypoint.sh"]
