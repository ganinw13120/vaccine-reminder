package handler

import (
	"vaccine-reminder/model"
	"vaccine-reminder/service"

	"github.com/gin-gonic/gin"
)

type webhookHandler struct {
	vaccineService service.VaccineService
}

type WebhookHandler interface {
	Webhook(c *gin.Context)
}

func NewWebhookHandler(
	vaccineService service.VaccineService,
) *webhookHandler {
	return &webhookHandler{
		vaccineService: vaccineService,
	}
}

func (w webhookHandler) Webhook(c *gin.Context) {
	var req model.WebhookPayload
	c.BindJSON(&req)
	w.vaccineService.Webhook(req)
}
