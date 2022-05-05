FROM ubuntu:latest

RUN apt-get update && apt-get -y upgrade \
    && apt-get install -y mosquitto \
    && apt-get -y clean \
    && rm -rf /var/lib/apt/lists/*

CMD ["mosquitto"]
