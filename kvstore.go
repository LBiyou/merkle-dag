package merkledag

// KVStore 接口定义了键值存储器的基本操作
type KVStore interface {
	// Has 方法用于检查存储器中是否存在指定键的值
	Has(key []byte) (bool, error)

	// Put 方法用于将键值对存储到存储器中
	Put(key, value []byte) error

	// Get 方法用于从存储器中获取指定键对应的值
	Get(key []byte) ([]byte, error)

	// Delete 方法用于删除存储器中指定键的值
	Delete(key []byte) error
}
