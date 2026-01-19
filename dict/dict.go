package dict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var errNotFound = errors.New("There is no")
var errWordExists = errors.New("Already Exists")

// Search For a word
func (d Dictionary) Search(word string) (string, error) {
	val, exist := d[word]
	if exist {
		return val, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update Word to the dictionary
func (d Dictionary) Update(word string, def string) error {
	_, nonExists := d.Search(word)
	if nonExists == nil {
		d[word] = def
	} else {
		return nonExists
	}
	return nil
}

// Delete word to the Dictionary
func (d Dictionary) Delete(word string) error {
	_, nonExistst := d.Search(word)
	switch nonExistst {
	case nil:
		delete(d, "word")
	case errNotFound:
		return errors.New("There is no Word")
	}
	return nil
}
