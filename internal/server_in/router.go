package server_in

import (
	"cron/internal/basic/tracing"
	"cron/internal/biz"
	"fmt"
	"net/http"
)

// 安全检查
func health(w http.ResponseWriter, r *http.Request) {
	str := ""

	to := tracing.ObserveCheck()
	if to.SpanWriteQueueLen > 2000 {
		str += fmt.Sprintf("span log write queue overstock[%v];", to.SpanWriteQueueLen)
	}

	bo := biz.ObserveCheck()
	if !bo.CronStart {
		str += " cron not start;"
	}

	if str == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
	}
}
