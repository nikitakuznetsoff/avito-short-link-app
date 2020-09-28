package database

import (
	"database/sql"
	"shortlinkapp/pkg/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) Get(key string) (*models.ShortLink, error) {
	shortLink := &models.ShortLink{}
	err := repo.DB.QueryRow("SELECT id, link FROM links WHERE id = ?", key).
				Scan(&shortLink.ID, &shortLink.Link)
	if err != nil {
		return nil, err
	}
	return shortLink, nil
}

func (repo *Repository) Set(sl *models.ShortLink) (int64, error) {
	result, err := repo.DB.Exec("INSERT INTO links (`id`, `link`) VALUES (?, ?)",
								sl.ID, sl.Link)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}