package context

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			//log.Println("store error:", err)
			logger.Warn("store error", "error", err)
			return
		}
		fmt.Fprint(w, data)
	}
}
