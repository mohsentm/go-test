/*
 * BUG: Invalid user ID defined in the expected data.
 *      There was a user with ID 350 which doesn't exist. I replaced the ID with correct one (123)
 *      I added an extra test for covering the not found exception
 */
package test2

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func NewDatabase() *Database {
	return &Database{
		database: map[int]*DBPerson{
			621: {ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
			123: {ID: 123, Name: "FriendNo1", Friends: []int{621, 631}},
			251: {ID: 251, Name: "SecondBestFriend", Friends: []int{621}},
			631: {ID: 631, Name: "ThirdWh33l", Friends: []int{621, 123, 251}},
		},
	}
}

func (d *Database) GetUser(id int, out chan *DBPerson) error {
	p, ok := d.database[id]
	if !ok {
		return ErrPersonNotFound
	}
	time.Sleep(time.Millisecond * 300)
	out <- p
	return nil
}

type testCase2 struct {
	request int
	result  *PopulatedPerson
	err     error
}

func TestPopulate(t *testing.T) {
	testCases := []testCase2{
		{
			request: 621,
			result: &PopulatedPerson{
				ID:   621,
				Name: "XxDragonSlayerxX",
				Friends: []DBPerson{
					{ID: 123, Name: "FriendNo1", Friends: []int{621, 631}},
					{ID: 251, Name: "SecondBestFriend", Friends: []int{621}},
					{ID: 631, Name: "ThirdWh33l", Friends: []int{621, 123, 251}},
				},
			},
			err: nil,
		},
		{
			request: 350,
			result:  nil,
			err:     ErrPersonNotFound,
		},
		{
			request: 123,
			result: &PopulatedPerson{
				ID:   123,
				Name: "FriendNo1",
				Friends: []DBPerson{
					{ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
					{ID: 631, Name: "ThirdWh33l", Friends: []int{621, 123, 251}},
				},
			},
			err: nil,
		},
		{
			request: 251,
			result: &PopulatedPerson{
				ID:   251,
				Name: "SecondBestFriend",
				Friends: []DBPerson{
					{ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
				},
			},
			err: nil,
		},
		{
			request: 631,
			result: &PopulatedPerson{
				ID:   631,
				Name: "ThirdWh33l",
				Friends: []DBPerson{
					{ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
					{ID: 123, Name: "FriendNo1", Friends: []int{621, 631}},
					{ID: 251, Name: "SecondBestFriend", Friends: []int{621}},
				},
			},
			err: nil,
		},
	}
	for ind, test := range testCases {
		t.Run(fmt.Sprint(ind), func(t *testing.T) {
			var db = NewDatabase()
			res, err := populate(db, test.request)
			if !cmp.Equal(res, test.result) {
				t.Log("result is incorrect", cmp.Diff(res, test.result))
				t.Fail()
			}
			if !errors.Is(err, test.err) {
				t.Log("err is incorrect", cmp.Diff(err, test.err))
				t.Fail()
			}
		})
	}
}
