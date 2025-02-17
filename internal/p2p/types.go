package p2p

type Message[T any] struct {
	MessageType MessageType `json:"msg_type"`
	RequestId   string      `json:"request_id"`
	DataType    string      `json:"data_type"`
	Data        T           `json:"data"`
}

type HeartbeatMessage struct {
	PeerID    string `json:"peer_id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"ts"`
}

type HandshakeMessage struct {
	PeerID    string `json:"peer_id"`
	Submitter string `json:"submitter"`
	Handshake string `json:"handshake"`
	Timestamp int64  `json:"ts"`
}

type MessageType int

const (
	MessageTypeUnknown = iota
	MessageTypeTssMsg
)
