// The root interface in collection hierarchy
// A collection represents a group of objects, known as its <i>elements</i>
package collection

type Collection interface {
    // Returns the number of elements in this collection.
    // max number of elements can this collection contains depends on platform math.MaxInt64
    Size() int

    // Returns true if this collection is empty
    Empty() bool

    // Returns <tt>true</tt> if this collection contains the specified element.
    Contains(e interface{}) (bool, error)

    // Returns an slice containing all of the elements in this collection
    ToSlice() [] interface{}

    // Add an element into this collection. Returns <tt>true</tt> if this collection changed as a
    // result of the call.  (Returns <tt>false</tt> if this collection does
    // not permit duplicates and already contains the specified element.)
    Add(e interface{}) (bool, error)

    // Removes a single instance of the specified element from this collection.
    // If this collection contains one or more such elements, returns true.
    Remove(e interface{}) (bool, error)


    // Removes all of the elements from this collection
    Clear()
}

