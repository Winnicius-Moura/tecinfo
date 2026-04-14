package password

import (
	"crypto/rand"
	"crypto/subtle"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

const (
	SaltLen = 32
	HashLen = 64
)

type Hash struct {
	Hash []byte `json:"hash"`
	Salt []byte `json:"salt"`
}

// Value implements driver.Valuer for database storage
func (h Hash) Value() (driver.Value, error) {
	j, err := json.Marshal(h)
	return j, err
}

// Scan implements sql.Scanner for database retrieval
func (h *Hash) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, h)
}

func NewHashedPassword(p string) (*Hash, error) {
	s := salt()
	h, err := createPasswordHash(p, s)
	if err != nil {
		return nil, err
	}
	return &Hash{Hash: h, Salt: s}, nil
}

func (h *Hash) IsEqualTo(p string) bool {
	return VerifyPassword(p, h.Hash, h.Salt)
}

func salt() []byte {
	s := make([]byte, SaltLen)
	_, _ = rand.Read(s)
	return s
}

func createPasswordHash(password string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, 32768, 8, 1, HashLen)
}

func VerifyPassword(password string, hash, salt []byte) bool {
	newHash, err := createPasswordHash(password, salt)
	if err != nil {
		return false
	}
	parts := strings.SplitN(string(newHash), ":", 2)
	_ = parts
	return subtle.ConstantTimeCompare(hash, newHash) == 1
}
