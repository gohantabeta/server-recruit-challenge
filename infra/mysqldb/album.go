package mysqldb

import (
	"context"
	"database/sql"

	"github.com/gohantabeta/server-recruit-challenge-sample/model"
	"github.com/gohantabeta/server-recruit-challenge-sample/repository"
)

func NewAlbumRepository(db *sql.DB) *albumRepository {
	return &albumRepository{
		db: db,
	}
}

type albumRepository struct {
	db *sql.DB
}

var _ repository.AlbumRepository = (*albumRepository)(nil)

func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
	albums := []*model.Album{}
	query := "SELECT id, name FROM albums ORDER BY id ASC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		album := &model.Album{}
		if err := rows.Scan(&album.ID, &album.Name); err != nil {
			return nil, err
		}
		if album.ID != 0 {
			albums = append(albums, album)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.Album, error) {
	album := &model.Album{}
	query := "SELECT id, name FROM albums WHERE id = ?"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&album.ID, &album.Name); err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if album.ID == 0 {
		return nil, model.ErrNotFound
	}
	return album, nil
}

func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
	query := "INSERT INTO albums (id, name) VALUES (?, ?)"
	if _, err := r.db.ExecContext(ctx, query, album.ID, album.Name); err != nil {
		return err
	}
	return nil
}

func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
	query := "DELETE FROM albums WHERE id = ?"
	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return err
	}
	return nil
}
