package lists

//import the base interface
import (
	"github.com/Golang_play/ds/baseints"
)

// Now create a interface so that the single linked list and double linked list
// can inhirit the list functionality, some of the functionality are
// Insert, Add, Remove, Get, etc ...

type List interface {
	Insert() bool
	InsertAfter(index int, value interface{})
	Remove(index int) bool
	// As we don't know the return value lets make the return val as interface
	Get(index int) (interface{}, bool)
	Contains(value interface{}) bool

	// Now we need to inherit the base interface functions like size, empty
	// from the baseint package

	baseints.Baseint
}
