package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"vaccine-reminder/model"
)

var baseUrl = "https://script.google.com/macros/s/AKfycbzMY3AEAIS4iM4xKxEi4j0w1f2pr-ZJZZwTOy-i3RSOfkzW33VMFK4Niv-TJAEt4sY_/exec"
var method string = "GET"

type sheetRepository struct {
}

type SheetRepository interface {
	GetAllFiles() ([]*model.Files, error)
	GetAllUserLog() ([]*model.UserLog, error)
	UpdateUser(userId, personName, birth string) error
	InsertUser(userId, personName, birth string) error
	AddUserNotification(userId, personName, notification string) error
}

func NewSheetRepository() *sheetRepository {
	return &sheetRepository{}
}

func generateUrl(action, queryParam string) string {
	urls := fmt.Sprintf("%s?action=%s%s", baseUrl, action, queryParam)
	fmt.Println(urls)
	return urls
}

func (r sheetRepository) GetAllUserLog() ([]*model.UserLog, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, generateUrl("getLogs", ""), nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response []model.UserLogResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	logs := make([]*model.UserLog, len(response))
	for i, v := range response {
		birth, err := time.Parse("02/01/2006", v.Birth)
		if err != nil {
			return nil, err
		}
		logs[i] = &model.UserLog{
			UserId:       v.UserId,
			PersonName:   v.PersonName,
			Birth:        birth,
			Notification: strings.Split(v.Notifications, ","),
		}
	}

	return logs, nil
}

func (r sheetRepository) GetAllFiles() ([]*model.Files, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, generateUrl("getFiles", ""), nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response []*model.Files
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r sheetRepository) UpdateUser(userId, personName, birth string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, generateUrl("updateUser", fmt.Sprintf("&userId=%s&personName=%s&birth=%s", url.QueryEscape(userId), url.QueryEscape(personName), url.QueryEscape(birth))), nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (r sheetRepository) AddUserNotification(userId, personName, notification string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, generateUrl("addNotification", fmt.Sprintf("&userId=%s&personName=%s&notifications=%s", url.QueryEscape(userId), url.QueryEscape(personName), url.QueryEscape(notification))), nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (r sheetRepository) InsertUser(userId, personName, birth string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, generateUrl("insertUser", fmt.Sprintf("&userId=%s&personName=%s&birth='%s", url.QueryEscape(userId), url.QueryEscape(personName), url.QueryEscape(birth))), nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
