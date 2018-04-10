// package baseint is a base interface which has all the
// base function that all the data
// structures should implement.
// few them are IsEmpty, Serialize, Deserialize, Size, Clear
// interface is OO concept which helps in implementing polymorphism
// without the inheritance functioanlity
package baseints

type Baseint interface {
	IsEmpty()
	Size()
	Serialize()
	Deserialize()
}
