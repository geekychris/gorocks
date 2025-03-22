package db

import (
	"fmt"
	"github.com/linxGnu/grocksdb"
)

// KeyValuePair represents a key-value pair with optional error
type KeyValuePair struct {
	Key   string
	Value []byte
	Err   error
}

type RocksDB struct {
	db *grocksdb.DB
	ro *grocksdb.ReadOptions
	wo *grocksdb.WriteOptions
}

func NewRocksDB(path string) (*RocksDB, error) {
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)

	db, err := grocksdb.OpenDb(opts, path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &RocksDB{
		db: db,
		ro: grocksdb.NewDefaultReadOptions(),
		wo: grocksdb.NewDefaultWriteOptions(),
	}, nil
}

func (r *RocksDB) Close() {
	r.ro.Destroy()
	r.wo.Destroy()
	r.db.Close()
}

func (r *RocksDB) Put(key string, value []byte) error {
	return r.db.Put(r.wo, []byte(key), value)
}

func (r *RocksDB) Get(key string) ([]byte, bool, error) {
	slice, err := r.db.Get(r.ro, []byte(key))
	if err != nil {
		return nil, false, fmt.Errorf("failed to get key: %w", err)
	}
	defer slice.Free()

	if slice.Size() == 0 {
		return nil, false, nil
	}

	value := make([]byte, slice.Size())
	copy(value, slice.Data())
	return value, true, nil
}

func (r *RocksDB) Delete(key string) error {
	return r.db.Delete(r.wo, []byte(key))
}
func (r *RocksDB) GetByPrefix(prefix string) chan KeyValuePair {
	ch := make(chan KeyValuePair)

	go func() {
		defer close(ch)

		it := r.db.NewIterator(r.ro)
		defer it.Close()

		prefixBytes := []byte(prefix)
		for it.Seek(prefixBytes); it.Valid(); it.Next() {
			key := it.Key()
			if !hasPrefix(key.Data(), prefixBytes) {
				key.Free()
				break
			}

			value := it.Value()
			keyStr := string(key.Data())
			valueData := make([]byte, value.Size())
			copy(valueData, value.Data())

			key.Free()
			value.Free()

			ch <- KeyValuePair{
				Key:   keyStr,
				Value: valueData,
			}
		}

		if err := it.Err(); err != nil {
			ch <- KeyValuePair{Err: fmt.Errorf("iterator error: %w", err)}
		}
	}()

	return ch
}

func (r *RocksDB) GetMultiple(keys []string) chan KeyValuePair {
	ch := make(chan KeyValuePair)

	go func() {
		defer close(ch)

		for _, key := range keys {
			value, exists, err := r.Get(key)
			if err != nil {
				ch <- KeyValuePair{
					Key: key,
					Err: err,
				}
				continue
			}

			if !exists {
				continue
			}

			ch <- KeyValuePair{
				Key:   key,
				Value: value,
			}
		}
	}()

	return ch
}

func hasPrefix(s, prefix []byte) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}
