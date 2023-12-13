package elasticsearch

import (
	"context"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// DeleteIndexIfExists ...
func (cli *Client) DeleteIndexIfExists(index string) error {
	exists, err := cli.ExistsIndex(index)
	if err != nil {
		return err
	}

	if !exists {
		return nil
	}

	return cli.DeleteIndex(index)
}

// ExistsIndex ...
func (cli *Client) ExistsIndex(index string) (bool, error) {
	res, err := cli.es.Indices.Exists([]string{index})
	if err != nil {
		return false, err
	}

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

// DeleteIndex ...
func (cli *Client) DeleteIndex(index string) error {
	res, err := cli.es.Indices.Delete([]string{index})
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("[%s] error in deleting index %s", res.Status(), index)
	}

	return nil
}

// CreateIndex ...
func (cli *Client) CreateIndex(ctx context.Context, index string) error {
	req := esapi.IndicesCreateRequest{
		Index: index,
	}

	res, err := req.Do(ctx, cli.es)
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("[%s] error in creating index %s", res.Status(), index)
	}

	if err = res.Body.Close(); err != nil {
		return err
	}

	return nil
}
