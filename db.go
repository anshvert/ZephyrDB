package ZephyrDB

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// HUhaha

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
	file, err := os.ReadFile(db.file)
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

func (db *Database) Save() {
	file, err := json.MarshalIndent(db.data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to encode database: %v", err)
	}
	err = os.WriteFile(db.file, file, 0644)
	if err != nil {
		log.Fatalf("Failed to write database: %v", err)
	}
}

func (db *Database) Set(key string, value interface{}) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
	db.Save()
}

func (db *Database) Get(key string) (interface{}, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, ok := db.data[key]
	return value, ok
}

//func main() {
//	db := NewDatabase("./db.json")
//	db.Set("foo", "bar")
//	fmt.Println(db.Get("foo"))
//}
