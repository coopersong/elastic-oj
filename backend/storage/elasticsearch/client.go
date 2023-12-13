package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Client struct {
	es *elasticsearch.Client
}

// NewClient ...
func NewClient() *Client {
	cli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
		},
	})
	if err != nil {
		return nil
	}

	return &Client{
		es: cli,
	}
}
