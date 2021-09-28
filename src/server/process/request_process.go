package process

import (
	"math/rand"
	"net/http"
	"time"
)

var ()

func ProcessRequest(r *http.Request) error {
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return nil
}
