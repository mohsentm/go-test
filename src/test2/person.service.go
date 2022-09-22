package test2

import (
	"errors"
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

	personModel, err := getUserData(db, id)
	if err != nil {
		return nil, err
	}
	populatedPerson := PopulatedPerson{
		ID:   personModel.ID,
		Name: personModel.Name,
	}

	personModel.appendFriends(db, &populatedPerson)

	return &populatedPerson, nil
}

func (person *DBPerson) appendFriends(db *Database, populatedPerson *PopulatedPerson) {
	for _, personID := range person.Friends {
		personModel, _ := getUserData(db, personID)
		populatedPerson.Friends = append(populatedPerson.Friends, *personModel)
	}
}

func getUserData(db *Database, id int) (*DBPerson, error) {
	errChan := make(chan error)
	dataChan := make(chan *DBPerson)

	go func(id int) {
		defer close(dataChan)
		defer close(errChan)

		err := db.GetUser(id, dataChan)
		if err != nil {
			errChan <- err
		}
	}(id)

	select {
	case personData := <-dataChan:
		return personData, nil
	case err := <-errChan:
		return nil, err
	}
}
