package cmd

func SetFunc(storage map[string]string, k string, v string) map[string]string {
	storage[k] = v

	return storage
}

func GetFunc(storage map[string]string, k string) *string {
	_, ok := storage[k]
	if ok {
		v := storage[k]
		return &v
	}

	return nil
}