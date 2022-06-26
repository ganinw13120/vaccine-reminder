package service

import "vaccine-reminder/model"

type vaccineService struct {
}

type VaccineService interface {
	Webhook(model.WebhookPayload)
}

func NewVaccineService() *vaccineService {
	return &vaccineService{}
}

func (s vaccineService) Webhook(model.WebhookPayload) {

}
