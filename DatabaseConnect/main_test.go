package database_connect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToDatabase(t *testing.T) {
	db_conn, err := ConnectToExtendedSqliteDatabase("file::memory:?cache=shared")
	assert.NoError(t, err)
	assert.NotNil(t, db_conn)
}
