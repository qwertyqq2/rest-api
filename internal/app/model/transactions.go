package model

type Transaction struct {
	ID         int `json:"id"`
	SenderID   int `json:"idSender"`
	ReceiverID int `json:"idReceiver"`
	Value      int `json:"value"`
}
