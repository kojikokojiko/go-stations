package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		cdr := model.CreateTODORequest{}

		err := json.NewDecoder(r.Body).Decode(&cdr)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer r.Body.Close()
		if cdr.Subject == "" {

			// err.Errorを消して、エラー分をべた書きにしたら治ったなぜ？cdr

			http.Error(w, "400　BadRequest", http.StatusBadRequest)
			log.Println("sdfsdfdsfsdfds")
			return
		} else {

			ctx := r.Context()

			todo, err := h.svc.CreateTODO(ctx, cdr.Subject, cdr.Description)

			if err != nil {
				return
			}

			response := &model.CreateTODOResponse{TODO: todo}

			err = json.NewEncoder(w).Encode(response)

			if err != nil {
				fmt.Println(err)
			}
			return
		}
	default:
		log.Println("requiest is not post")
		http.Error(w, "requiest is not post", http.StatusBadRequest)
	}

}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
