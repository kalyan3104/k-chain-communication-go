package testscommon

// WebsocketConnectionStub -
type WebsocketConnectionStub struct {
	OpenConnectionCalled func(url string) error
	ReadMessageCalled    func() (messageType int, payload []byte, err error)
	WriteMessageCalled   func(messageType int, data []byte) error
	IsOpenCalled         func() bool
	GetIDCalled          func() string
	CloseCalled          func() error
}

// IsOpen --
func (w *WebsocketConnectionStub) IsOpen() bool {
	if w.IsOpenCalled != nil {
		return w.IsOpenCalled()
	}

	return false
}

// OpenConnection -
func (w *WebsocketConnectionStub) OpenConnection(url string) error {
	if w.OpenConnectionCalled != nil {
		return w.OpenConnectionCalled(url)
	}

	return nil
}

// ReadMessage -
func (w *WebsocketConnectionStub) ReadMessage() (messageType int, payload []byte, err error) {
	if w.ReadMessageCalled != nil {
		return w.ReadMessageCalled()
	}

	return 0, nil, err
}

// WriteMessage -
func (w *WebsocketConnectionStub) WriteMessage(messageType int, data []byte) error {
	if w.WriteMessageCalled != nil {
		return w.WriteMessageCalled(messageType, data)
	}

	return nil
}

// GetID -
func (w *WebsocketConnectionStub) GetID() string {
	if w.GetIDCalled != nil {
		return w.GetIDCalled()
	}
	return ""
}

// Close -
func (w *WebsocketConnectionStub) Close() error {
	if w.CloseCalled != nil {
		return w.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (w *WebsocketConnectionStub) IsInterfaceNil() bool {
	return w == nil
}
