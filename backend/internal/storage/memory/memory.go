package memory

type memoryStorage struct{}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{}
}
