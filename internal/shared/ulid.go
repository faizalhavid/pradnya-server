package shared

import (
	"crypto/rand"

	"github.com/oklog/ulid/v2"
)

func New() (string, error) {
	id, err := ulid.New(ulid.Now(), rand.Reader)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
