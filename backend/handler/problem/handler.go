package problem

import (
	"github.com/elastic/go-elasticsearch/v8"

	"elastic-oj/storage"
)

// Handler ...
type Handler struct {
	db *storage.ProblemDao
	es *elasticsearch.Client
}

// NewHandler ...
func NewHandler() *Handler {
	db, err := storage.NewProblemDao()
	if err != nil {
		return nil
	}

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
		},
	})
	if err != nil {
		return nil
	}

	return &Handler{
		db: db,
		es: es,
	}
}
