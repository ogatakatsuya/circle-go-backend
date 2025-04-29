package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func NewClient() *supabase.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding without it")
	}

	key := os.Getenv("SUPABASE_KEY")
	url := os.Getenv("SUPABASE_URL")
	if key == "" || url == "" {
		log.Fatal("SUPABASE_KEY and SUPABASE_URL must be set")
	}

	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatalf("failed to create Supabase client: %v", err)
	}

	return client
}
