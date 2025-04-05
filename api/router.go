package api

import (
	"net/http"

	"github.com/gohantabeta/server-recruit-challenge-sample/api/middleware"
	"github.com/gohantabeta/server-recruit-challenge-sample/controller"
	"github.com/gohantabeta/server-recruit-challenge-sample/infra/mysqldb"
	"github.com/gohantabeta/server-recruit-challenge-sample/service"
)

func NewRouter(
	dbUser, dbPass, dbHost, dbName string,
) (http.Handler, error) {
	dbClient, err := mysqldb.Initialize(dbUser, dbPass, dbHost, dbName)
	if err != nil {
		return nil, err
	}
	// 接続確認
	if err := dbClient.Ping(); err != nil {
		return nil, err
	}

	singerRepo := mysqldb.NewSingerRepository(dbClient)
	singerService := service.NewSingerService(singerRepo)
	singerController := controller.NewSingerController(singerService)
	albumRepo := mysqldb.NewAlbumRepository(dbClient)
	albumService := service.NewAlbumService(albumRepo)
	albumController := controller.NewAlbumController(albumService, singerService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /singers", singerController.GetSingerListHandler)
	mux.HandleFunc("GET /singers/{id}", singerController.GetSingerDetailHandler)
	mux.HandleFunc("POST /singers", singerController.PostSingerHandler)
	mux.HandleFunc("DELETE /singers/{id}", singerController.DeleteSingerHandler)
	mux.HandleFunc("GET /albums", albumController.GetAlbumListHandler)
	mux.HandleFunc("GET /albums/{id}", albumController.GetAlbumDetailHandler)
	mux.HandleFunc("POST /albums", albumController.PostAlbumHandler)
	mux.HandleFunc("DELETE /albums/{id}", albumController.DeleteAlbumHandler)

	wrappedMux := middleware.LoggingMiddleware(mux)

	return wrappedMux, nil
}
