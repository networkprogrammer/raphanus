package raphanus

import "github.com/msoap/raphanus/common"

// GetDict - get dict value by key
func (db *DB) GetDict(key string) (value raphanuscommon.DictValue, err error) {
	if db.withLock {
		db.RLock()
		defer db.RUnlock()
	}

	rawVal, ok := db.data[key]
	if !ok {
		return value, raphanuscommon.ErrKeyNotExists
	}

	value, ok = rawVal.(raphanuscommon.DictValue)
	if !ok {
		return value, raphanuscommon.ErrKeyTypeMissmatch
	}

	return value, nil
}

// SetDict - create/update dict value by key
func (db *DB) SetDict(key string, value raphanuscommon.DictValue, ttl int) error {
	if db.withLock {
		db.Lock()
		defer db.Unlock()
	}

	if !isValidKey(key) {
		return raphanuscommon.ErrKeyIsNotValid
	}

	db.data[key] = value
	db.setTTL(key, ttl)

	return nil
}

// UpdateDict - update dict value by exists key
func (db *DB) UpdateDict(key string, value raphanuscommon.DictValue) (err error) {
	if db.withLock {
		db.Lock()
		defer db.Unlock()
	}

	if _, ok := db.data[key]; !ok {
		return raphanuscommon.ErrKeyNotExists
	}

	db.data[key] = value

	return nil
}

// GetDictItem - get item from dict value by exists key
func (db *DB) GetDictItem(key string, dictKey string) (string, error) {
	if db.withLock {
		db.RLock()
		defer db.RUnlock()
	}

	if err := db.validateDictParams(key, dictKey); err != nil {
		return "", err
	}

	return db.data[key].(raphanuscommon.DictValue)[dictKey], nil
}

// SetDictItem - set item on dict value by exists key
func (db *DB) SetDictItem(key, dictKey, dictValue string) error {
	if db.withLock {
		db.RLock()
		defer db.RUnlock()
	}

	if err := db.validateDictParams(key, dictKey); err != nil {
		return err
	}

	db.data[key].(raphanuscommon.DictValue)[dictKey] = dictValue

	return nil
}

// RemoveDictItem - remove item on dict value by exists key
func (db *DB) RemoveDictItem(key, dictKey string) error {
	if db.withLock {
		db.RLock()
		defer db.RUnlock()
	}

	if err := db.validateDictParams(key, dictKey); err != nil {
		return err
	}

	delete(db.data[key].(raphanuscommon.DictValue), dictKey)

	return nil
}

func (db *DB) validateDictParams(key, dictKey string) error {
	if len(key) == 0 || len(dictKey) == 0 {
		return raphanuscommon.ErrDictKeyIsEmpty
	}

	rawVal, ok := db.data[key]
	if !ok {
		return raphanuscommon.ErrKeyNotExists
	}

	value, ok := rawVal.(raphanuscommon.DictValue)
	if !ok {
		return raphanuscommon.ErrKeyTypeMissmatch
	}

	if _, ok := value[dictKey]; !ok {
		return raphanuscommon.ErrDictKeyNotExists
	}

	return nil
}
