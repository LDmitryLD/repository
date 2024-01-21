package storages

import (
	"projects/LDmitryLD/repository/app/internal/db/adapter"
	"projects/LDmitryLD/repository/app/internal/modules/user/storage"
)

type Storages struct {
	User storage.UserRepository
}

func NewStorages(sqlAdapter *adapter.SQLAdapter) *Storages {
	return &Storages{
		User: storage.NewUserStorage(sqlAdapter),
	}
}
