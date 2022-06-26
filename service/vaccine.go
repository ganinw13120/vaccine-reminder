package service

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
	"vaccine-reminder/model"
	"vaccine-reminder/repository"
	"vaccine-reminder/util"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type vaccineService struct {
	sheetRepository repository.SheetRepository
}

type VaccineService interface {
	Webhook(req model.WebhookPayload)
	CronJob()
}

func NewVaccineService(
	sheetRepository repository.SheetRepository,
) *vaccineService {
	return &vaccineService{
		sheetRepository: sheetRepository,
	}
}

func (s vaccineService) CronJob() {
	logs, err := s.sheetRepository.GetAllUserLog()
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := s.sheetRepository.GetAllFiles()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, log := range logs {
		year, month, _, _, _, _ := util.Diff(log.Birth, time.Now().AddDate(543, 0, 0))
		monthDiff := (year * 12) + month
		for _, file := range files {
			isSent := false
			for _, v := range log.Notification {
				if fmt.Sprint(file.Id) == v {
					isSent = true
				}
			}
			if isSent {
				continue
			}
			if file.Month == monthDiff {
				fmt.Println("Processing notification")
				/// Update sheet
				s.sheetRepository.AddUserNotification(log.UserId, log.PersonName, fmt.Sprint(file.Id))
				/// Push Message
				err := s.pushImageMessage(log.UserId, file.Url)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	fmt.Println("Scheule complete")
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
	s.pushTextMessage(req.Events[0].Source.UserID, "บันทึกข้อมูลสำเร็จ")
}

func (s vaccineService) pushImageMessage(userId, imgUrl string) error {
	client := &http.Client{}
	bot, err := linebot.New("1a13854d8d764f63bb8a35309c240a5a", "juxBi5xsAE9T9+CjJf0PJlqUjyCWStF1GP9Zt/gJ+49PhBPrQKIVQvQWRALPZ6dOINzgMoIjcx8+GVI0oP+TY4kaBg7Kh9VjdQmkPcYqnhApbMMZ3QqCP+R1Hi5va+nFqHQ8PxS58YjQ/EvQaJcurAdB04t89/1O/w1cDnyilFU=", linebot.WithHTTPClient(client))
	if err != nil {
		return err
	}
	img := linebot.NewImageMessage(imgUrl, imgUrl)
	_, err = bot.PushMessage(userId, img).Do()
	return err
}

func (s vaccineService) pushTextMessage(userId, text string) error {
	client := &http.Client{}
	bot, err := linebot.New("1a13854d8d764f63bb8a35309c240a5a", "juxBi5xsAE9T9+CjJf0PJlqUjyCWStF1GP9Zt/gJ+49PhBPrQKIVQvQWRALPZ6dOINzgMoIjcx8+GVI0oP+TY4kaBg7Kh9VjdQmkPcYqnhApbMMZ3QqCP+R1Hi5va+nFqHQ8PxS58YjQ/EvQaJcurAdB04t89/1O/w1cDnyilFU=", linebot.WithHTTPClient(client))
	if err != nil {
		return err
	}
	txt := linebot.NewTextMessage(text)
	_, err = bot.PushMessage(userId, txt).Do()
	return err
}
