package middleware

import (
	"context"
	"github.com/sethvargo/go-limiter/httplimit"
	"log"
	"net/http"
	"server/server/handler"
)

func StatMiddleware() (http.HandlerFunc, *httplimit.Middleware) {
	store := statStore()
	ctx := context.Background()

	// key is the unique value upon which you want to rate limit, like an IP or
	// MAC address.
	key := "127.0.0.1"
	tokens, remaining, reset, ok, err := store.Take(ctx, key)

	// tokens is the configured tokens (15 in this example).
	_ = tokens

	// remaining is the number of tokens remaining (14 now).
	_ = remaining

	// reset is the unix nanoseconds at which the tokens will replenish.
	_ = reset

	// ok indicates whether the take was successful. If the key is over the
	// configured limit, ok will be false.
	_ = ok

	// Here's a more realistic example:
	if !ok {
		panic(ok)
	}
	middleware, err := httplimit.NewMiddleware(store, httplimit.IPKeyFunc())
	if err != nil {
		log.Fatal(err)
	}

	finalHandler := http.HandlerFunc(handler.StatsHandler)

	return finalHandler, middleware
	return nil, nil
}
