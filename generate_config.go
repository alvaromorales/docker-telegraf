package main

import (
	"os"
	"text/template"
)

type ConfigParams struct {
	InfluxDBHostname         string
	InfluxDBPort             string
	DbName           string
	DbUsername       string
	DbPassword       string
	Interval         string
	Debug            bool
	CpuPlugin        bool
	DiskPlugin       bool
	DockerPlugin     bool
	IOPlugin         bool
	MemPlugin        bool
	MysqlPlugin      bool
	NetPlugin        bool
	PostgresqlPlugin bool
	RedisPlugin      bool
	SwapPlugin       bool
	SystemPlugin     bool
}

func main() {
	tmpl, err := template.ParseFiles("telegraf.toml.tmp")
	params := ConfigParams{
		os.Getenv("INFLUXDB_HOSTNAME"),
		os.Getenv("INFLUXDB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("INTERVAL"),
		os.Getenv("DEBUG") == "true",
		os.Getenv("CPU_PLUGIN") == "true",
		os.Getenv("DISK_PLUGIN") == "true",
		os.Getenv("DOCKER_PLUGIN") == "true",
		os.Getenv("IO_PLUGIN") == "true",
		os.Getenv("MEM_PLUGIN") == "true",
		os.Getenv("MYSQL_PLUGIN") == "true",
		os.Getenv("NET_PLUGIN") == "true",
		os.Getenv("POSTGRESQL_PLUGIN") == "true",
		os.Getenv("REDIS_PLUGIN") == "true",
		os.Getenv("SWAP_PLUGIN") == "true",
		os.Getenv("SYSTEM_PLUGIN") == "true",
	}

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, params)
	if err != nil {
		panic(err)
	}
}
