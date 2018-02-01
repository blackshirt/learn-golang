///////////////////////
///// delivery.go /////
///////////////////////

package cleanarch

import (
	"encoding/json"
	"io/ioutil"
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

func (h *HTTPTrainHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	// todo: check start limit should be
	start, _ := strconv.Atoi(r.FormValue("start"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	res, err := h.ItemUCase.Fetch(start, limit)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *HTTPTrainHandler) Posts(w http.ResponseWriter, r *http.Request) {
	//res, err := h.ItemUCase.Posts()
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//json.NewEncoder(w).Encode(err.Error())
	//return
	//}
	//json.NewEncoder(w).Encode(body)
	w.Header().Set("Content-Type", "application/json")

	b, _ := ioutil.ReadAll(r.Body)

	w.Write(b)
}
