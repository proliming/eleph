package maps

// An object that maps keys to values.  A map cannot contain duplicate keys;
// each key can map to at most one value.
type Map interface {
    // Returns the number of key-value mappings in this map.
    // max number of key-value mappings can this map contains depends on platform math.MaxInt64
    Size() int

    // Returns true if this map is empty
    Empty() bool

    // Returns <tt>true</tt> if this map contains a mapping for the specified key.
    ContainsKey(k interface{}) bool

    //Returns <tt>true</tt> if this map maps one or more keys to the
    //specified value.
    ContainsValue(v interface{}) bool

    // Put a key-value mapping into this map.
    Put(k interface{}, v interface{}) error

    //Returns the value to which the specified key is mapped, or nil if this map contains no mapping for the key.
    Get(k interface{}) (value interface{}, err error)

    // Removes the mapping for a key from this map if it is present
    Remove(k interface{}) (value interface{}, err error)

    // Returns a Set view of the keys contained in this map.
    KeySet() []interface{}

    // Removes all of the mappings from this map (optional operation)
    Clear()
}