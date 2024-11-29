package repositories

import (
	"log"
	"sync"
)

func check(err error) {
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

type Repository interface {
	CreateRepository() *Repository
}

var instance *SqliteDb
var once sync.Once

func Connect() *SqliteDb {
	once.Do(func() {
		db, err := newSqliteDb("../st.db")
		if err != nil {
			log.Fatalln("Unable to connect to db. Quitting instead.")
		}

		instance = db
	})

	return instance
}

func CloseDb() {
	instance.close()
}

func RegisterRepositories() {
	Connect()

	// each of the repositories should be invoked here just so that
	// we ensure all of the repositories are created before we start connecting
	// services to them
	CreateConfigRepository()
}
