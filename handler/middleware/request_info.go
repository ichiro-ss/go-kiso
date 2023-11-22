package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Log struct {
	TimeStamp time.Time `json:"time_stamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func RequestInfo(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		st := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()

		// fmt.Println(end.UnixNano(), "-", st.UnixNano(), "/", int64(time.Millisecond))
		var latency = (end.UnixNano() - st.UnixNano()) / int64(time.Millisecond)
		var requestLog = Log{
			TimeStamp: time.Now(),
			Latency:   latency,
			Path:      r.URL.Path,
		}
		h = PutOsNameOnContext(h)
		val, ok := r.Context().Value(k).(string)
		if ok {
			requestLog.OS = val
		}
		res, _ := json.Marshal(requestLog)
		fmt.Println(string(res))
	}
	return http.HandlerFunc(fn)
}
