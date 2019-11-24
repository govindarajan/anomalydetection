package store

import (
	"database/sql"
	"time"

	"github.com/govindarajan/anomalydetection/log"
	_ "github.com/mattn/go-sqlite3"
)

const NO_EXPIRE int64 = -1

var dbConn KVStore

// KVStore is an interface which is used to store/fetch the values
type KVStore interface {
	Set(key string, val []byte) error
	Get(key string) ([]byte, error)

	ExpirableSet(key string, val []byte, expireAt time.Time) error
	ExpirableGet(key string) ([]byte, error)
	// will call the corresponding method to expire the record if store is not
	// supporting expiry feature. Otherwise Noop.
	Cleanup() error
}

// InitStore to initialize store.
func InitStore(s KVStore) {
	dbConn = s
	scheduleCleanup(s)
}

// GetStore to get the store which will be used for operations.
func GetStore() KVStore {
	return dbConn
}

func scheduleCleanup(s KVStore) {
	// Lets schedule cleanup for every one hour.
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case t := <-ticker.C:
				log.Debug("Running Cleanup at ", t)
				s.Cleanup()
			}
		}
	}()

}

// SQLite data store config
type SQLite struct {
	FilePath string `json:"file_path"`
	db       *sql.DB
}

// NewSQLite is used to get new instance of SQLite db
func NewSQLite(file string) (*SQLite, error) {
	// TODO: Validate if file exists
	s := &SQLite{FilePath: file}
	var err error
	s.db, err = sql.Open("sqlite3", "file:"+file+"?mode=rwc&cache=shared&_synchronous=0")
	//s.db, err = sql.Open("sqlite3", "file::memory:?mode=rwc&cache=shared")
	if err != nil {
		return nil, err
	}
	err = s.createTable()
	return s, err
}

// In SQLite, we have to create a table it seems.
func (s *SQLite) createTable() error {
	sql := "CREATE TABLE IF NOT EXISTS key_value (key VARCHAR NOT NULL PRIMARY KEY, value VARCHAR);"
	_, err := s.db.Exec(sql)
	if err != nil {
		return err
	}

	sql = "CREATE TABLE IF NOT EXISTS key_value_expire (key VARCHAR NOT NULL, expireAt DATETIME DEFAULT CURRENT_TIMESTAMP, value VARCHAR, PRIMARY KEY (expireAt, key));"
	_, err = s.db.Exec(sql)
	return err
}

// Set used to store the value
func (s *SQLite) Set(key string, val []byte) error {
	// Override the old value with new one.
	stmt, err := s.db.Prepare("REPLACE INTO key_value (key, value) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(key, val)
	return err
}

// Get - to get the value from store.
func (s *SQLite) Get(key string) ([]byte, error) {
	rows, err := s.db.Query("SELECT value FROM key_value WHERE key = ?", key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var val []byte
	for rows.Next() {
		rows.Scan(&val)
	}
	return val, nil
}

// ExpirableSet used to store the value with expiry time
func (s *SQLite) ExpirableSet(key string, val []byte, t time.Time) error {
	// Override the old value with new one.
	stmt, err := s.db.Prepare("REPLACE INTO key_value_expire (key, expireAt, value) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(key, t, val)

	return err
}

// ExpirableGet - to get the value from expirable store
func (s *SQLite) ExpirableGet(key string) ([]byte, error) {
	rows, err := s.db.Query("SELECT value FROM key_value_expire WHERE key = ?", key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var val []byte
	for rows.Next() {
		rows.Scan(&val)
	}
	return val, nil
}

// Cleanup is to delete all the values which are expired from store.
// Implement only if necessary
func (s *SQLite) Cleanup() error {
	now := time.Now()
	stmt, err := s.db.Prepare("DELETE from key_value_expire WHERE expireAt < ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(now)

	return err
}
