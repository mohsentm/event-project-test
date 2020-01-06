FROM golang:latest

LABEL maintainer="Mohsen Hosseini <mail@mohsenhosseini.info>"

ENV INSTALL_DIR /app
ENV MIGRATE_VERSION v4.7.1
ENV MIGRATE_PLATFORM linux-amd64

WORKDIR $INSTALL_DIR

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/$MIGRATE_VERSION/migrate.$MIGRATE_PLATFORM.tar.gz | tar xvz
RUN mv migrate.linux-amd64  /usr/local/bin/migrate
RUN chmod +x /usr/local/bin/migrate

COPY ./ $INSTALL_DIR/

COPY ./build/docker/run.sh /usr/local/bin/run
RUN chmod +x /usr/local/bin/run

# Command to run the executable
CMD ["run"]
