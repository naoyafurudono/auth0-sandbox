package handler

import (
	"encoding/json"
	"net/http"

	"github.com/naoyafurudono/auth0-sandbox/backend/internal/middleware"
	"github.com/naoyafurudono/auth0-sandbox/backend/internal/model"
)

type UserHandler struct {
	store *model.Store
}

func NewUserHandler(store *model.Store) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	auth0ID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.store.GetOrCreateUser(auth0ID, "", "")
	if err != nil {
		writeError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, user, http.StatusOK)
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	auth0ID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.store.GetOrCreateUser(auth0ID, "", "")
	if err != nil {
		writeError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	profile, err := h.store.GetUserProfile(user.ID)
	if err == model.ErrNotFound {
		writeError(w, "Profile not found", http.StatusNotFound)
		return
	}
	if err != nil {
		writeError(w, "Failed to get profile", http.StatusInternalServerError)
		return
	}

	writeJSON(w, profile, http.StatusOK)
}

func (h *UserHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	auth0ID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.store.GetOrCreateUser(auth0ID, "", "")
	if err != nil {
		writeError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	var update model.UserProfileUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	profile, err := h.store.CreateOrUpdateUserProfile(user.ID, &update)
	if err != nil {
		writeError(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	writeJSON(w, profile, http.StatusOK)
}

func (h *UserHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	auth0ID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.store.GetOrCreateUser(auth0ID, "", "")
	if err != nil {
		writeError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	data, err := h.store.GetUserData(user.ID)
	if err != nil {
		writeError(w, "Failed to get user data", http.StatusInternalServerError)
		return
	}

	writeJSON(w, data, http.StatusOK)
}

func (h *UserHandler) CreateUserData(w http.ResponseWriter, r *http.Request) {
	auth0ID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.store.GetOrCreateUser(auth0ID, "", "")
	if err != nil {
		writeError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	var create model.UserDataCreate
	if err := json.NewDecoder(r.Body).Decode(&create); err != nil {
		writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if create.Content == "" {
		writeError(w, "Content is required", http.StatusBadRequest)
		return
	}

	data, err := h.store.CreateUserData(user.ID, &create)
	if err != nil {
		writeError(w, "Failed to create user data", http.StatusInternalServerError)
		return
	}

	writeJSON(w, data, http.StatusCreated)
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, message string, status int) {
	writeJSON(w, model.ErrorResponse{Message: message}, status)
}
