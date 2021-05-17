package mock

func NewWW() *ww {
	return &ww{
		data: make([]byte, 0),
	}
}

type ww struct {
	data []byte
}

func (w *ww) Write(p []byte) (n int, err error) {
	w.data = append(make([]byte, 0, len(p)), p...)
	return len(p), nil
}

func (w *ww) Read() []byte {
	return w.data
}

func (w *ww) Sync() error { return nil }
