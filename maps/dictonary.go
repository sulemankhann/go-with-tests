package main

type Dictonary map[string]string

const (
	ErrorNotFound       = DictonaryErr("could not find the world you are looking for")
	ErrWordExists       = DictonaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictonaryErr("cannot update word because it does not exists")
)

type DictonaryErr string

func (e DictonaryErr) Error() string {
	return string(e)
}

func (d Dictonary) Search(word string) (string, error) {
	defination, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return defination, nil
}

func (d Dictonary) Add(word, defination string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = defination
	case nil:
		return ErrWordExists
	default:
		return err

	}

	return nil
}

func (d Dictonary) Update(word, defination string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = defination
	default:
		return err

	}

	return nil
}

func (d Dictonary) Delete(word string) {
	delete(d, word)
}
