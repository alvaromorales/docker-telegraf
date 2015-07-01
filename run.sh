#!/usr/bin/env bash

echo "generating config"
cd /
go run /generate_config.go /telegraf.toml.tmp > /telegraf.toml
echo "ls -lh /"

/opt/influxdb/telegraf -config /telegraf.toml
