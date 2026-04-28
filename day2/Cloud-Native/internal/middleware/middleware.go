package middleware

import (
	"net/http"
)

func ApplyMiddlewares(h http.Handler) http.Handler {
	return h // Add middleware chaining here
}
