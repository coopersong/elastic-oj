package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListProblems ...
func (h *Handler) ListProblems(ctx *gin.Context) {
	problems, err := h.problemDao.ListProblems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result":  problems,
		"message": "success",
	})
}
