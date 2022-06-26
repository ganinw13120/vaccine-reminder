package main

import (
	"vaccine-reminder/handler"
	"vaccine-reminder/router"
	"vaccine-reminder/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	vaccineService := service.NewVaccineService()
	webhookHandler := handler.NewWebhookHandler(vaccineService)
	router := router.NewWebhookRouter(r, webhookHandler)
	router.InitRouter()

	r.Run(":8000")
}
