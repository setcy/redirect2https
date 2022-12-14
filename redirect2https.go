// Package redirect2https is a plugin.
package redirect2https

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	permanent bool
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		permanent: false,
	}
}

// Server a Server plugin.
type Server struct {
	config *Config
	next   http.Handler
	name   string
}

// New created a new Server plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Server{
		config: config,
		next:   next,
		name:   name,
	}, nil
}

func (a *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Scheme == "https" {
		a.next.ServeHTTP(rw, req)
		return
	} else if req.URL.Scheme == "http" {
		if a.config.permanent {
			rw.WriteHeader(http.StatusMovedPermanently)
		} else {
			rw.WriteHeader(http.StatusFound)
		}
		req.URL.Scheme = "https"
		rw.Header().Set("Location", req.URL.String())
		return
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
