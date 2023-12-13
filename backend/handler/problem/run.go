package problem

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"elastic-oj/common"
	"elastic-oj/storage/mysql"
)

const (
	Pass = "PASS"
	Fail = "FAIL"
)

// Submit ...
type Submit struct {
	ProblemID      string `json:"ProblemID"`
	SubmittedQuery string `json:"SubmittedQuery"`
}

// RunProblem ...
func (h *Handler) RunProblem(ctx *gin.Context) {
	submit := &Submit{}
	err := ctx.ShouldBindJSON(submit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// load problem info
	problem, err := h.problemDao.GetProblemByID(submit.ProblemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	cases, err := h.caseDao.GetCasesByProblemID(submit.ProblemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, cs := range cases {
		result, err := h.runCase(ctx, submit, problem, cs)
		if err != nil {
			ctx.JSON(err.StatusCode(), gin.H{
				"message": err.Error(),
			})
			return
		}

		if result == Fail {
			ctx.JSON(http.StatusOK, gin.H{
				"message": Fail,
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": Pass,
	})
}

func (h *Handler) runCase(ctx *gin.Context, submit *Submit, problem *mysql.Problem, cs *mysql.Case) (string, *common.Error) {
	var docs []string
	if err := json.Unmarshal([]byte(cs.Docs), &docs); err != nil {
		return "", common.NewError(http.StatusInternalServerError, err.Error())
	}

	if err := h.prepareENV(ctx, problem.ESIndex, docs); err != nil {
		return "", err
	}

	standardResult, err := h.callES(ctx, problem.ESIndex, problem.StandardQuery)
	if err != nil {
		return "", err
	}

	realResult, err := h.callES(ctx, problem.ESIndex, submit.SubmittedQuery)
	if err != nil {
		return "", err
	}

	if len(standardResult) == 0 || len(realResult) == 0 {
		return "", common.NewError(http.StatusInternalServerError, "no data in es response")
	}

	if standardResult[strings.Index(standardResult, "hits"):] == realResult[strings.Index(realResult, "hits"):] {
		return Pass, nil
	}

	return Fail, nil
}

func (h *Handler) prepareENV(ctx context.Context, index string, docs []string) *common.Error {
	if err := h.es.DeleteIndexIfExists(index); err != nil {
		return common.NewError(http.StatusInternalServerError, err.Error())
	}

	if err := h.es.CreateIndex(ctx, index); err != nil {
		return common.NewError(http.StatusInternalServerError, err.Error())
	}

	if err := h.es.BatchIndexDocuments(ctx, index, docs); err != nil {
		return common.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (h *Handler) callES(ctx *gin.Context, index string, query string) (string, *common.Error) {
	resp, err := h.es.Search(ctx, index, query)
	if err != nil {
		return "", common.NewError(http.StatusInternalServerError, err.Error())
	}

	return resp, nil
}
