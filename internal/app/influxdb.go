package app

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/wanhello/omgind/pkg/global"
)

func InitInfluxDB() (influxdb2.Client, func(), error) {
	cfg := global.CFG.InfluxDB

	fmt.Println(" ------- influx db dsn ", cfg.DSN())

	client := influxdb2.NewClient(cfg.DSN(), cfg.Token)

	return client, func() {
		client.Close()
	}, nil
}
