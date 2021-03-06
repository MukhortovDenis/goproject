package pkg

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
)

// Путь к статике для рендеринга html со стороны сервера
var dirWithHTML string = "./ui/"

// Создание структуры, в которой подбираются данные из окружения
var configEnv = init_env()

// URI к бд
var DBConn string = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=require", configEnv.Dialect, configEnv.DataUser, configEnv.DataPass, configEnv.DataHost, configEnv.DataPort, configEnv.DataName)

type Handler struct {
	Store   *sessions.CookieStore
	Storage *sql.DB
}

func (h *Handler) MainHandle() *chi.Mux {
	router := NewRouter()
	fileServer(router)

	router.Get("/signin", h.signin)
	router.Get("/signup", h.signup)
	router.Get("/", h.index)
	router.Get("/settings", h.settings)
	router.Get("/settings-appearance", h.settingsAppearance)
	router.Get("/cabinet", h.cabinet)
	router.Get("/cabinet-info", h.cabinetInfo)
	router.Get("/cabinet-password", h.cabinetPassword)
	router.Get("/quit", h.quit)
	router.Get("/chests", h.chests)
	router.Get("/chest", h.giveChests)
	router.Get("/open-chest", h.openChest)
	router.Get("/inventory", h.inventory)
	router.Post("/save_user", h.save)
	router.Post("/check_user", h.check)

	router.Patch("/cabinet-info-change", h.changeCabinetInfo)
	router.Patch("/cabinet-password-change", h.changeCabinetPassword)
	return router
}
