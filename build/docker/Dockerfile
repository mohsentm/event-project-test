FROM golang:latest

LABEL maintainer="Mohsen Hosseini <mail@mohsenhosseini.info>"

ENV INSTALL_DIR /app

WORKDIR $INSTALL_DIR

COPY ./ $INSTALL_DIR/

COPY ./build/docker/run.sh /usr/local/bin/run
RUN chmod +x /usr/local/bin/run

# Command to run the executable
CMD ["run"]
