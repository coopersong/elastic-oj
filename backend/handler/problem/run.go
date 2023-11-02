package problem

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	Pass = "PASS"
	Fail = "FAIL"
)

// Submit ...
type Submit struct {
	ProblemID      string `json:"ProblemID"`
	SubmittedQuery string `json:"SubmittedQuery""`
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
	problem, err := h.db.GetProblemByID(submit.ProblemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var (
		standardResult []byte
		realResult     []byte
	)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		standardResult = h.callES(ctx, problem.ESIndex, problem.StandardQuery)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		realResult = h.callES(ctx, problem.ESIndex, submit.SubmittedQuery)
	}()

	wg.Wait()

	if len(standardResult) == 0 || len(realResult) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "no data in es response",
		})
		return
	}

	standardResultStr := string(standardResult)
	realResultStr := string(realResult)
	if standardResultStr[strings.Index(standardResultStr, "hits"):] == realResultStr[strings.Index(realResultStr, "hits"):] {
		ctx.JSON(http.StatusOK, gin.H{
			"message": Pass,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": Fail,
	})
}

func (h *Handler) callES(ctx *gin.Context, index string, query string) []byte {
	resp, err := h.es.Search(
		h.es.Search.WithIndex(index),
		h.es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return nil
	}

	if resp.IsError() {
		if resp.StatusCode >= 500 {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": resp.StatusCode,
			})
			return nil
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": resp.StatusCode,
		})
		return nil
	}

	body := resp.Body
	defer body.Close()

	result, err := io.ReadAll(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return nil
	}

	return result
}
