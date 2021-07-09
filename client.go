package opensearch

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"net/http"
)

// Client wraps original go-elasticsearch client and Opensearch API together
//
type Client struct {
	*elasticsearch.Client
	OSAPI *API
}

// Config wraps original go-elasticserach Config
//
type Config elasticsearch.Config

// NewDefaultClient returns Client with default configuration
//
func NewDefaultClient() (*Client, error) {
	// create default config
	cfg := Config{
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin",
		Password:  "admin",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// return elasticsearch client
	return NewClient(cfg)
}

// NewClient wraps original go-elasticsearch client creation
//
func NewClient(config Config) (*Client, error) {
	// create elasticsearch client
	client, err := elasticsearch.NewClient(elasticsearch.Config(config))
	if err != nil {
		return nil, err
	}

	// return opensearch client
	return &Client{
		Client: client,
		OSAPI:  NewAPI(client.Transport),
	}, nil
}
