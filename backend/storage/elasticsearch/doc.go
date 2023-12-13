package elasticsearch

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// Search ...
func (cli *Client) Search(ctx context.Context, index, query string) (string, error) {
	resp, err := cli.es.Search(
		cli.es.Search.WithIndex(index),
		cli.es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("[%s] error in searching in index %s", resp.StatusCode, index)
	}

	body := resp.Body

	bs, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	if err = body.Close(); err != nil {
		return "", err
	}

	return string(bs), nil
}

// BatchIndexDocuments ...
func (cli *Client) BatchIndexDocuments(ctx context.Context, index string, docs []string) error {
	for i, doc := range docs {
		if err := cli.IndexDocument(ctx, index, fmt.Sprintf("%d", i), doc); err != nil {
			return err
		}
	}
	return nil
}

// IndexDocument ...
func (cli *Client) IndexDocument(ctx context.Context, index string, docID, doc string) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: docID,
		Body:       strings.NewReader(doc),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, cli.es)
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf(res.String())
	}

	if err = res.Body.Close(); err != nil {
		return err
	}

	return nil
}
