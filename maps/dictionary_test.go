package maps

import "testing"

func TestSearch(t *testing.T) {

	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q given, %q", got, want, "test")
		}
	}

	assertErrors := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q given, %q", got, want, "test")
		}
	}

	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertErrors(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	assertDef := func(t *testing.T, dict Dictionary, word, def string) {
		t.Helper()
		got, err := dict.Search(word)
		if err != nil {
			t.Fatal("should find added word: ", err)
		}
		if def != got {
			t.Errorf("got %q want %q", got, def)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if got == nil {
			if want == nil {
				return
			}
			t.Fatal("expected to get an error")
		}
	}

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word, def := "test", "this is just a test"
		err := dict.Add(word, def)
		assertError(t, err, nil)
		assertDef(t, dict, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word, def := "test", "this is just a test"
		dict := Dictionary{word: def}
		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExist)
		assertDef(t, dict, word, def)
	})
}

func TestUpdate(t *testing.T) {

	assertDef := func(t *testing.T, dict Dictionary, word, newDef string) {
		got, err := dict.Search(word)
		if err != nil {
			t.Fatal("should find updated word: ", err)
		}

		if got != newDef {
			t.Errorf("got %q want %q", got, newDef)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if got == nil {
			if want == nil {
				return
			}
			t.Fatal("expected to get an error")
		}
	}

	t.Run("Update existing", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		newDef := "new definition"

		err := dict.Update(word, newDef)
		assertError(t, err, nil)
		assertDef(t, dict, word, newDef)
	})

	t.Run("New word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}
	dict.Delete(word)
	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
