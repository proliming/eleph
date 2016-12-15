package list

import "eleph"

// An ordered collection.
// The user can access elements by their integer index.
// List typically allow duplicate elements.
type List interface {
    // Inserts the specified element at the specified position in this list.
    Add(index int, e interface{})

    //  Returns the element at the specified position in this list.
    Get(index int) (value interface{}, err error)

    //Replaces the element at the specified position in this list with the
    //specified element (optional operation).
    Set(index int, e interface{}) (oldValue interface{}, err error)

    // Removes the element at the specified position in this list
    Remove(index int) error

    // Returns the index of the first occurrence of the specified element
    // in this list, or -1 if this list does not contain the element.
    IndexOf(e interface{}) int

    // Returns a view of the portion of this list between the specified
    // from, inclusive, and to, exclusive.
    Sub(from, to int) (subList List, err error)

    collection.Collection
}
