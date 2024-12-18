package MyDB

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Database struct {
	data map[string]interface{}
	mu   sync.RWMutex
	file string
}

func NewDatabase(fileName string) *Database {
	db := &Database{
		data: make(map[string]interface{}),
		file: fileName,
	}
	db.load()
	return db
}

func (db *Database) load() {
	file, err := ioutil.ReadFile(db.file)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Failed to read database file: %v", err)
	}
	err = json.Unmarshal(file, &db.data)
	if err != nil {
		log.Fatalf("Failed to parse database file: %v", err)
	}
}

func main() {
	db := NewDatabase("./db.json")
}
