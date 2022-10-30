package tokencon

import (
	"log"
	"net/http"
	"sync"
	"test_go/pkg/auth"
	"time"
)

var (
	mutex sync.RWMutex
	data  = make(map[*http.Request]map[string]*auth.ResultClaims)
	datat = make(map[*http.Request]int64)
)

func Set(r *http.Request, key string, val *auth.ResultClaims) {
	mutex.Lock()

	if data[r] == nil {
		data[r] = make(map[string]*auth.ResultClaims)
		datat[r] = time.Now().Unix()
	}
	data[r][key] = val
	log.Println(data[r][key])
	mutex.Unlock()
}

func Get(r *http.Request, key string) *auth.ResultClaims {
	mutex.RLock()
	log.Println(data[r])
	if ctx := data[r]; ctx != nil {
		value := ctx[key]
		mutex.RUnlock()
		return value
	}
	mutex.RUnlock()
	return nil
}
