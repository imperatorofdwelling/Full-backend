package config

import "os"

type Encryption struct {
	EncryptionKey string
}

func InitEncryption() *Encryption {
	return &Encryption{
		EncryptionKey: os.Getenv("ENCRYPTION_KEY"),
	}
}
