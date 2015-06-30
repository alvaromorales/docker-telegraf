FROM golang
MAINTAINER Alvaro Morales <alvarom@mit.edu>

# install dependencies
RUN apt-get update && \
  apt-get install -y curl

ENV INFLUXDB_HOSTNAME "influxdb"
ENV INFLUXDB_PORT  "8086"
ENV DB_NAME "grafana"
ENV DB_USERNAME "root"
ENV DB_PASSWORD "root"
ENV INTERVAL "10s"
ENV DEBUG **False**
ENV CPU_PLUGIN **True**
ENV CPU_PLUGIN **True**
ENV DISK_PLUGIN **True**
ENV DOCKER_PLUGIN **True**
ENV IO_PLUGIN **True**
ENV MEM_PLUGIN **True**
ENV MYSQL_PLUGIN **None**
ENV NET_PLUGIN **True**
ENV POSTGRESQL_PLUGIN **None**
ENV REDIS_PLUGIN **None**
ENV SWAP_PLUGIN **True**
ENV SYSTEM_PLUGIN **True**

ADD build.sh /build.sh
ADD generate_config.go /generate_config.go
ADD telegraf.toml.tmp /telegraf.toml.tmp

RUN chmod +x /build.sh
RUN /build.sh

CMD ["/opt/influxdb/telegraf", "-config", "/telegraf.toml"]
