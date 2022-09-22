package test2

import (
	"errors"
	"sync"
)

var (
	//ErrNotImplemented = errors.New("not_implemented")
	ErrPersonNotFound = errors.New("account_not_found")
)

type PopulatedPerson struct {
	ID      int
	Name    string
	Friends []DBPerson
}

func populate(db *Database, id int) (*PopulatedPerson, error) {
	dbOut := make(chan *DBPerson)

	var err error
	go func(id int, err *error) {
		defer close(dbOut)
		Error := db.GetUser(id, dbOut)
		err = &Error
		if Error != nil {
			close(dbOut)
		}
	}(id, &err)

	personModel := <-dbOut
	if err != nil {
		return nil, err
	}

	populatedPerson := PopulatedPerson{
		ID:   personModel.ID,
		Name: personModel.Name,
	}

	_ = personModel.getFriends(db, &populatedPerson)

	return &populatedPerson, nil
}

func (person *DBPerson) getFriends(db *Database, populatedPerson *PopulatedPerson) error {
	dbOut := make(chan *DBPerson, len(person.Friends))
	defer close(dbOut)
	var wg sync.WaitGroup

	for _, personID := range person.Friends {
		wg.Add(1)
		go func(personID int) {
			defer wg.Done()
			_ = db.GetUser(personID, dbOut)
		}(personID)

		personModel := <-dbOut
		populatedPerson.Friends = append(populatedPerson.Friends, *personModel)

	}

	wg.Wait()
	return nil
}
