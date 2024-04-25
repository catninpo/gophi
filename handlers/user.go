package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/catninpo/gophi"
	view "github.com/catninpo/gophi/views/user"
)

type UserHandler struct {
	UserService gophi.UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserService.UserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if r.Header.Get("Accept") == "application/json" {
		b, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}

	err = view.Show(user).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
