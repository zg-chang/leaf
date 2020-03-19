package network

type Processor interface {
	// must goroutine safe
	Route(msg interface{}, data []byte, userData interface{}) error
	// must goroutine safe
	Unmarshal(data []byte) (interface{}, error)
	// must goroutine safe
	Marshal(msg interface{}) ([][]byte, error)
}
