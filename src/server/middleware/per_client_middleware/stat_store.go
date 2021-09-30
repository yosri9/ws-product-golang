package per_client_middleware

import (
	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
	"log"
)

func statStore() limiter.Store {
	store, err := memorystore.New(&memorystore.Config{
		// Number of tokens allowed per interval.
		Tokens: statTokens,

		// Interval until tokens reset.
		Interval: statInterval,
	})
	if err != nil {
		log.Fatal(err)
	}
	return store
}
