package test2

type Database struct {
	database map[int]*DBPerson
	repository
}

type repository interface {
	GetUser(id int, out chan *DBPerson) error
}
