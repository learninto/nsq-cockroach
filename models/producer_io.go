package models

type ProducerPostIo struct {
	Body        string `json:"body"`
	StorageType string `json:"storage_type"`
	StorageName string `json:"storage_name"`
}
