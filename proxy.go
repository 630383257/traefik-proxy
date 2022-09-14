package proxy

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Config the plugin configuration.
type Config struct {
	ProxyPath string `json:"proxypath,omitempty"`
	UrlParameter string `json:"urlparameter,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Demo a Demo plugin.
type Proxy struct {
	ProxyPath string
	UrlParameter string
	next     http.Handler
	name     string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.ProxyPath) == 0 {
		return nil, fmt.Errorf("ProxyPath cannot be empty")
	}
	if len(config.UrlParameter) == 0 {
		return nil, fmt.Errorf("UrlParameter cannot be empty")
	}

	return &Proxy{
		ProxyPath:config.ProxyPath,
		UrlParameter:config.UrlParameter,
		next:     next,
		name:     name,
	}, nil
}

func (a *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path==a.ProxyPath{
	urlStr,_ :=url.QueryUnescape(req.URL.Query().Get(a.UrlParameter))
	url,_ := url.Parse(urlStr)
	req.URL = url
	}
	a.next.ServeHTTP(rw, req)
}