package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test_go/internal/app/handlers"
	"test_go/internal/app/model"
	"test_go/internal/app/store"
	"test_go/pkg/auth"
	"time"

	"log"

	"github.com/julienschmidt/httprouter"
)

const (
	usersUrl = "/users"
	userUrl  = "/profile"
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
	router.GET(userUrl, authorizeMiddleware(http.HandlerFunc(h.GetUser)))
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
		Password: req.Password,
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

	log.Println(req)

	e := &model.Employee{
		Name:     req.Name,
		Password: req.Password,
	}
	user, err := h.store.Repository().FindUser(e)
	if err != nil {
		log.Fatal(err)
	}
	id := user.ID
	status := user.Status

	jwt, err := auth.GenerateTokenJWT(id, status)
	if err != nil {
		log.Fatal(err)
	}

	refresh, err := auth.GenerateTokenRefresh()
	if err != nil {
		log.Fatal(err)
	}

	timeClose := int(time.Now().Add(auth.TokenTTL).Unix())
	session := &model.Session{
		UserId:       id,
		Status:       status,
		RefreshToken: refresh,
		TimeClose:    timeClose,
	}

	if err := h.store.Sessions().Create(session); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Токен сессии: '%s'", jwt)))

}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Get user...")
	idCook, statusCook := r.Cookies()[0], r.Cookies()[1]
	w.Write([]byte(fmt.Sprintf("User: %s, %s", idCook.Value, statusCook.Value)))

	//var user *auth.ResultClaims

	//mapstructure.Decode(claims.(jwt.MapClaims), &user)
	//fmt.Println(user)
	//w.Write([]byte(fmt.Sprintf("User: %d, %s", user.UserId, user.Status)))
}
