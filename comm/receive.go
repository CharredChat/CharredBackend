package comm

type ReceiveMessage struct {
	Op     uint64      `json:"op"`
	OpData interface{} `json:"opdata"`
}

type OpOne struct {
	Say string `json:"say"`
}
