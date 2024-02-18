package database

import (
	"context"

	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

// Create a new Store, which has a connection to a gorm.DB.
func New() *Store {
	db, err := SQLiteConn()
	if err != nil {
		panic(err)
	}

	return &Store{
		DB: db,
	}
}

// PINGSSSSSSSSSSSSSSSS
func (s *Store) Ping() error {
	ctx := context.Background()

	db, err := s.DB.DB()
	if err != nil {
		return err
	}

	return db.PingContext(ctx)
}

// CLOSSESSSS
func (s *Store) Close() error {
	db, err := s.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// Gets ALL the records from table
func (s *Store) List(i interface{}, t string, primary ...interface{}) error {
	result := s.DB.Table(t).Find(i, primary...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Gets a record from THE table
func (s *Store) Get(i interface{}, t string, primary ...interface{}) error {
	result := s.DB.Table(t).Find(i, primary...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Adds a record to THE table
func (s *Store) Add(i interface{}, t string) error {
	result := s.DB.Table(t).Create(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Updates a record to THE table
func (s *Store) Update(i interface{}, t string) error {
	result := s.DB.Table(t).Save(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Deletes a record FROM THE table
func (s *Store) Delete(i interface{}, t string, primary ...interface{}) error {
	result := s.DB.Table(t).Delete(i, primary...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Checks if there is a record
func (s *Store) Contains(i interface{}, t string, primary ...interface{}) error {
	result := s.DB.Table(t).Model(i).First(i, primary...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Automatically Migrate all of the given Models
func (s *Store) AutoMigrate(i ...interface{}) error {
	if len(i) > 0 {
		for _, m := range i {
			err := s.DB.Migrator().AutoMigrate(&m)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
