package maps

//Dictionary type is custom type on top of map
//with keys and values defined as string
type Dictionary map[string]string

const (
	ErrWordNotFound      = DictionaryErr("could not find the word you are looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exist")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exists")
)

//DictionaryErr is wrapper for all errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

//Search for the definition of the given word
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

//Add method will add new word and the definition to the dictonary
//Note: as Dictionary type is created on top of maps we don't have to use pointer
//because map is reference type.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

//Update definition of the existing word
//return error if word does not exists
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

//Delete remove word and definition from Dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
