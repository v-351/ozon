package service

type memoryStorage struct {
	rawURL   map[string]string
	shortURL map[string]string
}

func NewMemoryStorage(n int) *memoryStorage {
	return &memoryStorage{
		make(map[string]string, n),
		make(map[string]string, n),
	}
}

func (m *memoryStorage) Put(raw string, short string) {
	m.rawURL[raw] = short
	m.shortURL[short] = raw
}

func (m *memoryStorage) GetRaw(short string) (string, bool) {
	val, ok := m.shortURL[short]
	return val, ok
}

func (m *memoryStorage) GetShort(raw string) (string, bool) {
	val, ok := m.rawURL[raw]
	return val, ok
}

func (m *memoryStorage) Close() {}
