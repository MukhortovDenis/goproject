package pkg

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// Путь к статике для рендеринга html со стороны сервера
var dirWithHTML string = "./ui/"

// Создание структуры, в которой подбираются данные из окружения
var configEnv = init_env()

// URI к бд
var dbConn string = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=require", configEnv.Dialect, configEnv.DataUser, configEnv.DataPass, configEnv.DataHost, configEnv.DataPort, configEnv.DataName)

//Создание хранилища куки с рандомным ключом
var store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

type Handler struct {
}

func (h *Handler) MainHandle() *chi.Mux {
	router := NewRouter()
	fileServer(router)
	router.Get("/signin", h.signin)
	router.Get("/signup", h.signup)
	router.Get("/", h.index)
	router.Get("/cabinet", h.cabinet)
	router.Get("/save_user", h.save)
	router.Get("/cabinet-info", h.cabinetInfo)
	router.Get("/quit", h.quit)
	router.Get("/check_user", h.check)
	router.Get("/cabinet-password", h.cabinetPassword)
	return router
}
