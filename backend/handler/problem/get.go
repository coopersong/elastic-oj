package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProblemByID(ctx *gin.Context) {
	problemID := ctx.Param("problemID")
	problem, err := h.db.GetProblemByID(problemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result":  problem,
		"message": "success",
	})
}
