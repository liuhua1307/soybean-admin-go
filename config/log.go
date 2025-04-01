package config

import "soybean-admin-go/utils/log"

var Logger log.Logger

func init() {
	Logger = log.NewSlogLogger("./logs")
}
