package router

import (
	"github.com/gin-gonic/gin"

	"elastic-oj/handler/problem"
)

func setupListProblems(r *gin.Engine) {
	if r == nil {
		return
	}

	r.GET("problems", problem.NewHandler().ListProblems)
}

func setupGetProblemByID(r *gin.Engine) {
	if r == nil {
		return
	}

	r.GET("problems/:problemID", problem.NewHandler().GetProblemByID)
}

func setupRunProblem(r *gin.Engine) {
	if r == nil {
		return
	}

	r.POST("problems/run", problem.NewHandler().RunProblem)
}
