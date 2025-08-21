package link

import (
	"fmt"
	"go/adv-dev/pkg/req"
	"go/adv-dev/pkg/res"
	"net/http"
)

type linkHandler struct {
	LinkRepository *LinkRepository
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &linkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link/create", handler.create())
	router.HandleFunc("GET /{hash}", handler.goTo())
	router.HandleFunc("PATCH /link/{id}", handler.update())
	router.HandleFunc("DELETE /link/{id}", handler.delete())
}

func (handler *linkHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, request)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
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

	}
}

func (handler *linkHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		id := request.PathValue("id")
		fmt.Println(id)
	}
}
