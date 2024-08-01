FROM ubuntu:24.04

COPY bin /build-bin

RUN apt-get update \
    && apt-get install -y \
    make curl git \
    software-properties-common \
    jq sudo
RUN /build-bin/install-deps.sh
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /workspaces/eigenlayer-contracts
