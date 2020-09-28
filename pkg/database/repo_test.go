package database_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"shortlinkapp/pkg/database"
	"shortlinkapp/pkg/models"
	"testing"
)

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer db.Close()

	sl := &models.ShortLink{ID: "my-new-link", Link: "https://www.vk.com"}
	repo := &database.Repository{DB: db}
	row := sqlmock.NewRows([]string{"id", "link"}).AddRow(sl.ID, sl.Link)
	// Basic test
	mock.
		ExpectQuery("SELECT id, link FROM links WHERE").
		WithArgs(sl.ID, sl.Link).
		WillReturnRows(row)

	link, err := repo.Get(sl.ID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectation error: %s", err)
		return
	}
	if !reflect.DeepEqual(link, sl.Link) {
		t.Errorf("results not match, want %v, have %v", sl.Link, link)
		return
	}

	// Test with error
	mock.
		ExpectQuery("SELECT id, link FROM links WHERE").
		WithArgs(sl.ID, sl.Link).
		WillReturnError(fmt.Errorf("db_error"))

	link, err = repo.Get(sl.ID)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectation error: %s", err)
		return
	}
	if !reflect.DeepEqual(link, sl.Link) {
		t.Errorf("results not match, want %v, have %v", sl.Link, link)
		return
	}
}

func TestSet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer db.Close()

	sl := &models.ShortLink{ID: "my-new-link", Link: "https://www.vk.com"}
	repo := &database.Repository{DB: db}
	// Basic test
	mock.
		ExpectExec("INSERT INTO links (`id`, `link`) VALUES").
		WithArgs(sl.ID, sl.Link).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.Set(sl)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if id != 1 {
		t.Errorf("bad id: want %v, got %v", 1, id)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectation error: %s", err)
		return
	}

	// Test with err
	mock.
		ExpectExec("INSERT INTO links (`id`, `link`) VALUES").
		WithArgs(sl.ID, sl.Link).
		WillReturnError(fmt.Errorf("bad query"))

	id, err = repo.Set(sl)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectation error: %s", err)
		return
	}
}