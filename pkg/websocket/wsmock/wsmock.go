// Package wsmock implements mock websocket server
package wsmock

//go:generate noice -t errors.toml

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

type PeriodicServerConfig struct {
	MsgPeriod time.Duration
	MsgText   string
}

type PeriodicServer struct {
	cfg PeriodicServerConfig

	srv *httptest.Server
}

func NewPeriodicServer(cfg PeriodicServerConfig) *PeriodicServer {
	var ret PeriodicServer
	ret.cfg = cfg
	ret.srv = httptest.NewServer(http.HandlerFunc(ret.periodicHandler))
	return &ret
}

func (p *PeriodicServer) URL() *url.URL {
	serverURL, _ := url.Parse(p.srv.URL)
	serverURL.Scheme = "ws"
	return serverURL
}

type EchoServer struct {
	srv *httptest.Server
}

func NewEchoServer() *EchoServer {
	var ret EchoServer
	ret.srv = httptest.NewServer(http.HandlerFunc(ret.echoHandler))
	return &ret
}

func (s *EchoServer) URL() *url.URL {
	serverURL, _ := url.Parse(s.srv.URL)
	serverURL.Scheme = "ws"
	return serverURL
}
