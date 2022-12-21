package util

import (
	"github.com/golang/glog"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func GetGelfWriter(host, port string) gelf.Writer {
	graylogEndpoint := host + ":" + port
	glog.Infof("GELF UDP destination: %s", graylogEndpoint)
	gelfWriter, err := gelf.NewUDPWriter(graylogEndpoint)
	if err != nil {
		glog.Fatal(err)
	}

	return gelfWriter
}
