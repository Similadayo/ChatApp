package chat

// Message represents a chat message.
type Message struct {
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}

//NewMessage creates a new Message
func NewMessage(userName, content string) *Message {
	return &Message{
		UserName: userName,
		Content:  content,
	}
}
