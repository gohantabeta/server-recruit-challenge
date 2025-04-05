package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gohantabeta/server-recruit-challenge-sample/model"
	"github.com/gohantabeta/server-recruit-challenge-sample/service"
)

type albumController struct {
	albumService  service.AlbumService
	singerService service.SingerService
}

func NewAlbumController(as service.AlbumService, ss service.SingerService) *albumController {
	return &albumController{
		albumService:  as,
		singerService: ss,
	}
}

// GET /albums のハンドラー
func (c *albumController) GetAlbumListHandler(w http.ResponseWriter, r *http.Request) {
	albums, err := c.albumService.GetAlbumListService(r.Context())
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	// レスポンス用の構造体
	type AlbumResponse struct {
		ID     model.AlbumID `json:"id"`
		Title  string        `json:"title"`
		Singer *model.Singer `json:"singer"`
	}

	// レスポンスの配列を作成
	response := make([]*AlbumResponse, 0, len(albums))
	for _, album := range albums {
		// 各アルバムの歌手情報を取得
		singer, err := c.singerService.GetSingerService(r.Context(), album.SingerID)
		if err != nil {
			errorHandler(w, r, 500, err.Error())
			return
		}

		response = append(response, &AlbumResponse{
			ID:     album.ID,
			Title:  album.Title,
			Singer: singer,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

// GET /albums/{id} のハンドラー
func (c *albumController) GetAlbumDetailHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	albumID, err := strconv.Atoi(idString)
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	album, err := c.albumService.GetAlbumService(r.Context(), model.AlbumID(albumID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	// レスポンス用の構造体
	type AlbumResponse struct {
		ID     model.AlbumID `json:"id"`
		Title  string        `json:"title"`
		Singer *model.Singer `json:"singer"`
	}

	// 歌手情報を取得
	singer, err := c.singerService.GetSingerService(r.Context(), album.SingerID)
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	response := &AlbumResponse{
		ID:     album.ID,
		Title:  album.Title,
		Singer: singer,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

// POST /albums のハンドラー
func (c *albumController) PostAlbumHandler(w http.ResponseWriter, r *http.Request) {
	var album *model.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		err = fmt.Errorf("invalid body param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.albumService.PostAlbumService(r.Context(), album); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// DELETE /albums/{id} のハンドラー
func (c *albumController) DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	albumID, err := strconv.Atoi(idString)
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.albumService.DeleteAlbumService(r.Context(), model.AlbumID(albumID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.WriteHeader(204)
}
