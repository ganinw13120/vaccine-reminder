package main

import (
	"vaccine-reminder/handler"
	"vaccine-reminder/repository"
	"vaccine-reminder/router"
	"vaccine-reminder/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sheetRepository := repository.NewSheetRepository()

	vaccineService := service.NewVaccineService(sheetRepository)
	// go func() {
	// 	vaccineService.CronJob()
	// 	c := cron.New()
	// 	c.AddFunc("0 8,10,12,15,17 * * *", func() { vaccineService.CronJob() })
	// }()
	webhookHandler := handler.NewWebhookHandler(vaccineService)
	router := router.NewWebhookRouter(r, webhookHandler)
	router.InitRouter()

	r.Run(":8000")
}
