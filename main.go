package main

import (
	"log"
	"net/http"
	"strconv"

	"nofu-be/application/usecase"
	"nofu-be/config"
	"nofu-be/infrastructure/database"
	"nofu-be/infrastructure/handler"
	"nofu-be/infrastructure/repository"
)

func main() {
	// 1️⃣ Load konfigurasi
	cfg := config.LoadConfig()

	// 2️⃣ Hubungkan ke PostgreSQL
	database.ConnectDB(cfg)

	// 3️⃣ Inisialisasi repository, use‑case, dan handler
	userRepo := repository.NewUserRepository()
	userUC := usecase.NewUserUseCase(userRepo) // mengembalikan domain.UserUseCase
	userHandler := handler.NewUserHandler(userUC)

	// 4️⃣ Setup router
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/users", userHandler.GetAllUsers)
	mux.HandleFunc("POST /api/v1/user/create", userHandler.CreateUser)

	// 5️⃣ Jalankan server
	addr := ":" + strconv.Itoa(cfg.Port)
	log.Printf("Server berjalan di http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
