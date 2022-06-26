package service

import (
	"fmt"
	"regexp"
	"strings"
	"vaccine-reminder/model"
	"vaccine-reminder/repository"
)

type vaccineService struct {
	sheetRepository repository.SheetRepository
}

type VaccineService interface {
	Webhook(req model.WebhookPayload)
}

func NewVaccineService(
	sheetRepository repository.SheetRepository,
) *vaccineService {
	return &vaccineService{
		sheetRepository: sheetRepository,
	}
}

func (s vaccineService) Webhook(req model.WebhookPayload) {
	regex := regexp.MustCompile("(.*)( {1})(.*)( {1})([0-3][0-9]/[0-1][0-9]/[0-9]{4})")
	if !regex.Match([]byte(req.Events[0].Message.Text)) {
		fmt.Printf("Regex for %s mismatch\n", req.Events[0].Message.Text)
		return
	} else {
		fmt.Printf("Regex match %s from %s\n", req.Events[0].Message.Text, req.Events[0].Source.UserID)
	}
	personNameSlice := strings.Split(req.Events[0].Message.Text, " ")[:2]
	personName := fmt.Sprintf("%s %s", personNameSlice[0], personNameSlice[1])
	logs, err := s.sheetRepository.GetAllUserLog()
	if err != nil {
		fmt.Println(err)
		return
	}
	isDuplicate := false
	for _, v := range logs {
		if v.UserId == req.Events[0].Source.UserID {
			if v.PersonName == personName {
				isDuplicate = true
				break
			}
		}
	}
	if isDuplicate {
		err := s.sheetRepository.UpdateUser(req.Events[0].Source.UserID, personName, strings.Split(req.Events[0].Message.Text, " ")[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("update user success")
	} else {
		err := s.sheetRepository.InsertUser(req.Events[0].Source.UserID, personName, strings.Split(req.Events[0].Message.Text, " ")[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("insert user success")
	}
}
