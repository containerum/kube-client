package client

import (
	"net/url"
	"os"

	"git.containerum.net/ch/kube-client/pkg/rest"
)

//TODO: Make Interface

//Client - rest client
type Client struct {
	re rest.REST
	Config
	User User
}

//User -
type User struct {
	Role string
}

// Config -- provides configuration for Client
// If APIurl or ResourceAddr is void,
// trys to get them from envvars
type Config struct {
	User   User
	APIurl string
}

//NewClient -
func NewClient(config Config) (*Client, error) {
	var APIurl *url.URL
	var err error
	if config.APIurl == "" {
		APIurl, err = url.Parse(os.Getenv("API_URL"))
	} else {
		APIurl, err = url.Parse(config.APIurl)
	}
	if err != nil {
		return nil, err
	}
	config.APIurl = APIurl.String()
	client := &Client{
		re:     rest.NewResty(),
		Config: config,
		User:   config.User,
	}
	return client, nil
}
