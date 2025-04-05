package model

type AlbumID int

type Album struct {
	ID      AlbumID `json:"id"`
	Title   string  `json:"title"`
	AlbumID AlbumID `json:"album_id"` // モデル Album の ID と紐づきます
}

func (a *Album) Validate() error {
	if a.Title == "" {
		return ErrInvalidParam
	}
	if len(a.Title) > 255 {
		return ErrInvalidParam
	}
	return nil
}
