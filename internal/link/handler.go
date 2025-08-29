package link

import (
	"fmt"
	"go/adv-dev/configs"
	"go/adv-dev/pkg/middleware"
	"go/adv-dev/pkg/req"
	"go/adv-dev/pkg/res"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type linkHandler struct {
	LinkRepository *LinkRepository
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &linkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link/create", handler.create())
	router.HandleFunc("GET /{hash}", handler.goTo())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.delete())
}

func (handler *linkHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, request)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JsonResponse(w, createdLink, http.StatusCreated)
	}
}

func (handler *linkHandler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		hash := request.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, request, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *linkHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](&w, request)
		if err != nil {
			return
		}
		idString := request.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		email, ok := request.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}
		res.JsonResponse(w, link, 201)
	}
}

func (handler *linkHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		idString := request.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JsonResponse(w, nil, 200)
	}
}
