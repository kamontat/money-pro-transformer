package profile

import (
	transaction "moneypro.kamontat.net/models-transaction"
)

// Loader will get mapper format and load to profile object
func Loader(profile *Profile, mapper []map[string]string) (*Profile, error) {
	for _, data := range mapper {
		t, err := transaction.Builder(data)
		if err != nil {
			return profile, err
		}
		profile.AddTransaction(t)
	}

	return profile, nil
}
