package client

// MemStorage stores runtime state of client
type MemStorage struct {
	Token string
}

// SetToken sets/updates token
func (m *MemStorage) SetToken(token string) {
	m.Token = token
}
