package postgres

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresIntegrational(t *testing.T) {
	t.Setenv("PG_URL", "postgres://postgres:postgres@localhost/links?sslmode=disable")
	type test struct {
		input string
		err   error
	}

	tests := []test{
		{input: "", err: errors.New("no PG_URL was provided")},
		{input: "postgres://dummy:dummy@localhost/links?sslmode=disable", err: errors.New("Expecting error")},
		{input: "postgressssss://dummy:dummy@localhost/links?sslmode=disable", err: errors.New("Bad chema")},
		{input: os.Getenv("PG_URL"), err: nil},
	}

	for _, tc := range tests {
		_, err := NewPgDb(tc.input)

		//we expection error return
		if tc.err != nil {
			assert.Error(t, err)
		}
	}
}
