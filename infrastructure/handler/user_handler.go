package handler

import (
	"encoding/json"
	"net/http"

	"nofu-be/domain" // interface UserUseCase
	"nofu-be/domain/user"
)

// UserHandler menangani HTTP request untuk endpoint /api/v1/users
type UserHandler struct {
	usecase domain.UserUseCase
}

// NewUserHandler membuat instance UserHandler dengan dependency injection
func NewUserHandler(uc domain.UserUseCase) *UserHandler {
	return &UserHandler{usecase: uc}
}

// GetAllUsers mengembalikan daftar semua user dalam format JSON
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		http.Error(w, "Gagal mengambil data user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users) // ignore error sederhana untuk contoh
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.User

	// Parse body JSON
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validasi dasar
	if user.CUsername == "" || user.CPassword == "" || user.CFullname == "" {
		http.Error(w, "Username, Password, dan Fullname wajib diisi", http.StatusBadRequest)
		return
	}

	// Panggil usecase untuk menyimpan
	if err := h.usecase.CreateUser(user); err != nil {
		http.Error(w, "Gagal menyimpan user", http.StatusInternalServerError)
		return
	}

	// Kembalikan response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User berhasil ditambahkan",
	})
}
