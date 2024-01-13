package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictonary := Dictonary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictonary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, got := dictonary.Search("unkown")

		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictonary := Dictonary{}
		word := "test"
		defination := "this is just a test"

		err := dictonary.Add(word, defination)

		assertError(t, err, nil)

		assertDefination(t, dictonary, word, defination)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictonary := Dictonary{word: defination}

		err := dictonary.Add(word, defination)

		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictonary := Dictonary{word: defination}
		newDefination := "new defination"

		dictonary.Update(word, newDefination)

		assertDefination(t, dictonary, word, newDefination)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictonary := Dictonary{}

		err := dictonary.Update(word, defination)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictonary := Dictonary{word: "test defination"}

	dictonary.Delete(word)

	_, err := dictonary.Search(word)

	if err != ErrorNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertDefination(t testing.TB, dictonary Dictonary, word, defination string) {
	t.Helper()

	got, err := dictonary.Search("test")
	if err != nil {
		t.Fatal("should find added word.", err)
	}

	assertStrings(t, got, defination)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
