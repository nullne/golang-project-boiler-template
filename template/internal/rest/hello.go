package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"{{GoModule}}/internal/domain"
)

type helloRequest struct{
	Who string `json:"who"`
}

type helloResponse struct{
	Who string `json:"who"`
	Days int `json:"days"`
}

func fromHello(h *domain.Hello) helloResponse{
	return helloResponse{
		Who: h.Who,
		Days: h.Days,
	}
}

// Hello godoc
// @Summary Say hello
// @Tags hello
// @Param  who path string true "who"
// @Success 200 {object} helloResponse
// @Router /hello/{who}  [get]
func (h *handler) Hello(ctx *gin.Context) {
	who := ctx.Param("who")
	data, err := h.{{ServiceName}}.GetHello(context.Background(), who)
	if err != nil {
		log.Err(err).Msg("get hello")
		return
	}
	ctx.JSON(http.StatusOK, fromHello(data))
}
