package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/KanybekMomukeyev/ConcurrentDB/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/mock"
)

var dbMng *database.DbManager

func init() {
	dbMng = database.NewDbManager("dbname=fortest host=localhost sslmode=disable")
}

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	//assert.Nil(t, dbMng)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, dbMng) {

	}
}
/*
  Test objects
*/

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MyMockedObject struct{
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MyMockedObject) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestSomething2(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// setup expectations
	testObj.On("DoSomething", 123).Return(true, nil)

	// call the code we are testing
	testObj.DoSomething(126)

	// assert that the expectations were met
	testObj.AssertExpectations(t)
}