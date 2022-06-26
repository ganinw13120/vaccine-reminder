package model

type WebhookPayload struct {
	Destination string         `json:"destination"`
	Events      []WebhookEvent `json:"events"`
}

type WebhookEvent struct {
	DeliveryContext DeliveryContext `json:"deliveryContext"`
	Message         Message         `json:"message"`
	Mode            string          `json:"mode"`
	ReplyToken      string          `json:"replyToken"`
	Type            string          `json:"type"`
	Source          Source          `json:"source"`
}

type DeliveryContext struct {
	IsReDelivery bool `json:"isRedelivery"`
}

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}
type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}
