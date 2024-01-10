package domain

// Message message struct
type Message struct {
	Message string `json:"message"`
}

// NewMessage is function for create a new message
func NewMessage(err ...error) *Message {
	if len(err) > 0 {
		return &Message{
			Message: err[0].Error(),
		}
	}

	return &Message{
		Message: "success",
	}
}
