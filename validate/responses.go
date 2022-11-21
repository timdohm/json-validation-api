package validate

type Response struct {
	Action string `json:"action"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type ResponseWithMessage struct {
	Action  string `json:"action"`
	ID      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewReponse(action string, id string, status string) *Response {
	return &Response{action, id, status}
}

func NewReponseWithMessage(action string, id string, status string, message string) *ResponseWithMessage {
	return &ResponseWithMessage{action, id, status, message}
}
