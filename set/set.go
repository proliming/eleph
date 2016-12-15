package set

import "eleph"

// A collection that contains no duplicate elements.
type Set interface {
    // Adds all of the elements in the specified collection to this set
    // if they're not already present (optional operation).
    AddAll(c collection.Collection) (bool, error)

    //Removes from this set all of its elements that are contained in
    // the specified collection (optional operation)
    RemoveAll(c collection.Collection) (bool, error)

    // Returns <tt>true</tt> if this set contains all of the elements
    // of the specified collection.
    ContainsAll(c collection.Collection) (bool, error)

    collection.Collection
}
