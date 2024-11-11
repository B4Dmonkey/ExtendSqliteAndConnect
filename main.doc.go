/* 
This package handles the boilerplate for connecting to an extended sqlite3 database.

	To connect to the database, use the `ConnectToExtendedSqliteDatabase` function.

	Upon connecting to the database, the sqlite3 driver is extended using an init with a sync.Once to ensure the extension is only done once.

	The Database connection is extended with the following functions:
		- `uuid` - returns a new UUID using the `github.com/google/uuid` package
		- `current_timestamp` - returns the current timestamp in the format "2006-01-02 15:04:05"
*/
package database_connect