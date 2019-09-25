package cache

// Cache - cache interface used to store and retrieve serialized data
type Cache interface {
	Set(key, value []byte) error
	Get(key []byte) ([]byte, error)
	GetAllValues() ([][]byte, error)
	GetAllEntries() (map[string][]byte, error)
	RecordsCount() (int, error)
	Delete(key []byte) error
	DeleteData() error
	GetAllKeys() (map[string]bool, error)
}
