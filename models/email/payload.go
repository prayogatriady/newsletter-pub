package email

type EmailPayload struct {
	RecipientList []EmailPayloadList `json:"recipient_list"`
}

type EmailPayloadList struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
