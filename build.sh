#!/usr/bin/env bash

curl -O http://get.influxdb.org/telegraf/telegraf_0.1.1_amd64.deb
dpkg -i telegraf_0.1.1_amd64.deb
cd /
go run /generate_config.go /telegraf.toml.tmp > /telegraf.toml
