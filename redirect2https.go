// Package redirect2https is a plugin.
package redirect2https

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	permanent bool `yaml:"permanent"`
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
	if req.Header.Get("X-Forwarded-Proto") == "http" {
		host := req.Header.Get("X-Forwarded-Host")
		path := req.URL.String()
		port := req.Header.Get("X-Forwarded-Port")
		if port != "" && port != "443" {
			http.Redirect(rw, req, fmt.Sprintf("https://%s:%s%s", host, port, path), http.StatusMovedPermanently)
		} else {
			http.Redirect(rw, req, fmt.Sprintf("https://%s%s", req.Host, req.URL.Path), http.StatusMovedPermanently)
		}
		return
	} else {
		a.next.ServeHTTP(rw, req)
		return
	}
}
