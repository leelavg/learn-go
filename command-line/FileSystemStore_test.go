package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	// t.Run("/league from a reader", func(t *testing.T) {

	// 	database, cleanDatabase := createTempFile(t, `[
	// 		{"Name": "Leela", "Wins": 10},
	// 		{"Name": "Neela", "Wins": 33}]`)
	// 	defer cleanDatabase()

	// 	store, err := NewFileSystemPlayerStore(database)
	// 	AssertNoError(t, err)

	// 	got := store.GetLeague()

	// 	want := League{
	// 		{"Leela", 10},
	// 		{"Neela", 33},
	// 	}

	// 	AssertLeague(t, got, want)

	// 	got = store.GetLeague()
	// 	AssertLeague(t, got, want)

	// })

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Leela", "Wins": 10},
			{"Name": "Neela", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Leela")
		AssertScoreEquals(t, got, 10)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Leela", "Wins": 10},
			{"Name": "Neela", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Leela")
		got := store.GetPlayerScore("Leela")
		AssertScoreEquals(t, got, 11)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Leela", "Wins": 10},
			{"Name": "Neela", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Venkat")
		got := store.GetPlayerScore("Venkat")
		AssertScoreEquals(t, got, 1)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Leela", "Wins": 10},
			{"Name": "Neela", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Neela", 33},
			{"Leela", 10},
		}

		AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}

func AssertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
