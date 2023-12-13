package problem

import (
	"elastic-oj/storage/elasticsearch"
	"elastic-oj/storage/mysql"
)

// Handler ...
type Handler struct {
	problemDao *mysql.ProblemDao
	caseDao    *mysql.CaseDao
	es         *elasticsearch.Client
}

// NewHandler ...
func NewHandler() *Handler {
	problemDao, err := mysql.NewProblemDao()
	if err != nil {
		return nil
	}

	caseDao, err := mysql.NewCaseDao()
	if err != nil {
		return nil
	}

	es := elasticsearch.NewClient()
	if es == nil {
		return nil
	}

	return &Handler{
		problemDao: problemDao,
		caseDao:    caseDao,
		es:         es,
	}
}
