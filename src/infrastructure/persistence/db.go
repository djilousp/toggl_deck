package persistence

import (
	"fmt"

	"github.com/djilousp/toggl_deck/src/application/protocols"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



type Repositories struct {
	Deck protocols.DeckRepository
	db   *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		Deck: NewDeckRepository(db),
		db:   db,
	}, nil
}

//Close the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//This migrates all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&Deck{}, &Card{}, &DecksCards{}).Error
}
