package client

import (
	"github.com/containerum/kube-client/pkg/rest"
	"github.com/containerum/kube-client/pkg/websocket/gorilla"
)

//TODO: Make Interface

//Client - rest client
type Client struct {
	Config
}

//User -
type User struct {
	Role string
}

// Config -- provides configuration for Client
// If APIurl or ResourceAddr is void,
// trys to get them from envvars
type Config struct {
	User     User
	RestAPI  rest.REST
	WSDialer *gorilla.Dialer
}

//NewClient -
func NewClient(config Config) (*Client, error) {
	client := &Client{
		Config: config,
	}
	return client, nil
}
