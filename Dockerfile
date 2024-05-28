FROM ubuntu:24.04

RUN apt-get update \
    && apt-get install -y make curl git software-properties-common jq \
    && add-apt-repository -y ppa:longsleep/golang-backports \
    && add-apt-repository -y ppa:ethereum/ethereum \
    && apt-get update \
    && apt-get install -y golang-1.22 ethereum=1.14.3+build29843+noble \
    && curl -L https://foundry.paradigm.xyz | bash \
    && /root/.foundry/bin/foundryup

RUN cp -R /root/.foundry/bin/* /usr/local/bin/

