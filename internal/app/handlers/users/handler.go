package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test_go/internal/app/handlers"
	"test_go/internal/app/model"
	"test_go/internal/app/store"
	"test_go/pkg/auth"

	"log"

	"github.com/julienschmidt/httprouter"
)

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
	singUrl  = "/singIn"
)

type handler struct {
	store store.Store
}

func New(s store.Store) handlers.Handler {
	return &handler{s}
}

// /Прописать коды ответов
func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersUrl, h.GetList)
	router.POST(usersUrl, h.CreateUser)
	router.GET(userUrl, h.GetUser)
	router.PUT(userUrl, h.UpdateUser)
	router.PATCH(userUrl, h.PartiallyUpdateUser)
	router.DELETE(userUrl, h.DeleteUser)
	router.POST(singUrl, h.SingIn)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users, err := h.store.Repository().GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		w.Write([]byte(fmt.Sprintf("name: %s, status: %s, password: %s \n\n", u.Name, u.Status, u.Password)))
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	type request struct {
		Name     string `json:"name"`
		Status   string `json:"status"`
		Password string `json:"password"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Fatal(err)
	}

	e := &model.Employee{
		Name:     req.Name,
		Status:   req.Status,
		Password: auth.GeneratePasswordHash(req.Password),
	}
	if err := h.store.Repository().Create(e); err != nil {
		log.Fatal(err)
	}

	log.Printf("name: %s, status: %s, password: %s", e.Name, e.Status, e.Password)
}

func (h *handler) SingIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Fatal(err)
	}

	e := &model.Employee{
		Name:     req.Name,
		Password: auth.GeneratePasswordHash(req.Password),
	}
	user, err := h.store.Repository().FindUser(e)
	if err != nil {
		log.Fatal(err)
	}
	id := user.ID

	token, err := auth.GenerateTokenJWT(id)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("Ваш токен: '%s'", token)))

}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("Get user")))
	w.WriteHeader(200)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("Update user")))
	w.WriteHeader(204)
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("Part update user")))
	w.WriteHeader(204)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("Delete user")))
	w.WriteHeader(204)
}
