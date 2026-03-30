package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	// Las variables de entorno ya fueron cargadas por config.Load()
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// 2. Crear configuración del Pool
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("No se pudo parsear la config de la DB: %v", err)
	}

	// Optimizaciones para el plan gratuito de Supabase
	config.MaxConns = 10 // No necesitamos las 60 permitidas para empezar
	config.MinConns = 2

	// 3. Conectar
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("No se pudo conectar a Supabase: %v", err)
	}

	// 4. Verificar conexión (Ping)
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("La base de datos no responde al ping: %v", err)
	}

	fmt.Println("Conexión exitosa a Supabase!")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}
