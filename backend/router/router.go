package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter ...
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	setup(r)

	return r
}

func setup(r *gin.Engine) {
	setupListProblems(r)
	setupGetProblemByID(r)
	setupRunProblem(r)
}
