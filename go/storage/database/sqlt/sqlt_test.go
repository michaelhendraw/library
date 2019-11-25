package sqlt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/michaelhendraw/library/go/storage/database/sqlt"
)

func TestNewFailed(t *testing.T) {

	db := New(Config{
		Driver: "mysql",
		Master: "master",
		Slave:  []string{"slave"},
	})

	assert.Nil(t, db)
}
