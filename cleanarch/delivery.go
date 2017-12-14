package cleanarch

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HTTPTrainHandler struct {
	ItemUCase TrainUCase
}

func NewHTTPTrainHandler(tr TrainUCase) *HTTPTrainHandler {
	return &HTTPTrainHandler{ItemUCase: tr}
}

// serve single /path/{id}
func (h *HTTPTrainHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	res, err := h.ItemUCase.GetById(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(res)
}
