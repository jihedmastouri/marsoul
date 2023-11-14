package transport

type MessageType int8

const (
	SaveRq MessageType = iota + 1
	RetrRq
	Ping
)

type SaveRqPayload struct {
	FileName string `json:"0"`
	Size     int    `json:"1"`
	Replicas int8   `json:"2"`
	Region   string `json:"3"`
}

type RetrRqPayload struct {
	Id string `json:"0"`
}

type Response struct {
	Ok           bool   `json:"0"`
	FileNodeAddr string `json:"1"`
	Secret       string `json:"2"`
}

type SaveResponse struct {
	Response
}

type RetrResponse struct {
	Response
}
