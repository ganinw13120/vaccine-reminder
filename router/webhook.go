package router

import (
	"net/http"
	"vaccine-reminder/handler"

	"github.com/gin-gonic/gin"
)

type webhookRouter struct {
	app            *gin.Engine
	webhookHandler handler.WebhookHandler
}

func NewWebhookRouter(
	app *gin.Engine,
	webhookHandler handler.WebhookHandler,
) *webhookRouter {
	return &webhookRouter{
		app:            app,
		webhookHandler: webhookHandler,
	}
}

func (r webhookRouter) InitRouter() {

	r.app.POST("/webhook", r.webhookHandler.Webhook)
	r.app.GET("/cron", r.webhookHandler.CronJob)

	r.app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
