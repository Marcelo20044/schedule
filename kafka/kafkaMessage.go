package kafka

type Message struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}
