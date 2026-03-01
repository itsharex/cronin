package server_in

import (
	"cron/internal/basic/config"
	"fmt"
	"net/http"
)

func InitHttp(conf *config.SystemConf) {
	if conf == nil || conf.HttpPort == "" {
		return
	}
	mux := http.NewServeMux()
	// 设置路由
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {})
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/Health", health)
	mux.Handle("/metrics", metricsHandle())

	// 启动服务
	if err := http.ListenAndServe(":"+conf.HttpPort, mux); err != nil {
		panic(fmt.Errorf("system http server listen error: %w", err))
	}
}
