package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Host string
	Port string
	User string
	Pass string
	DB   string
)

func loadDotEnv() error {
	var candidates []string

	// 1) Allow explicit override for local testing/CI.
	if envFile := os.Getenv("APP_ENV_FILE"); envFile != "" {
		candidates = append(candidates, envFile)
	}

	// 2) Works reliably with `bazel run`.
	if ws := os.Getenv("BUILD_WORKSPACE_DIRECTORY"); ws != "" {
		candidates = append(candidates, filepath.Join(ws, ".env"))
	}

	// 3) Bazel runfiles directory in Linux/macOS.
	if runfilesDir := os.Getenv("RUNFILES_DIR"); runfilesDir != "" {
		candidates = append(candidates,
			filepath.Join(runfilesDir, "_main", ".env"),
			filepath.Join(runfilesDir, "go_workspace", ".env"),
		)
	}

	// 4) Bazel runfiles next to executable.
	if exe, err := os.Executable(); err == nil {
		runfilesBase := exe + ".runfiles"
		candidates = append(candidates,
			filepath.Join(runfilesBase, "_main", ".env"),
			filepath.Join(runfilesBase, "go_workspace", ".env"),
		)
	}

	// 5) Works with `go run` from project root.
	candidates = append(candidates, ".env")

	// 6) Fallback to source-relative lookup.
	_, b, _, ok := runtime.Caller(0)
	if ok {
		basepath := filepath.Dir(b)
		candidates = append(candidates, filepath.Join(basepath, "../../../.env"))
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return godotenv.Load(path)
		}
	}

	return fmt.Errorf(".env not found in known locations")
}

func init() {
	err := loadDotEnv()
	if err != nil {
		log.Println("Aviso: não foi possível carregar o .env, usando variáveis de ambiente do sistema")
	}

	Host = os.Getenv("POSTGRES_HOST")
	Port = os.Getenv("POSTGRES_PORT")
	User = os.Getenv("POSTGRES_USER")
	Pass = os.Getenv("POSTGRES_PASSWORD")
	DB = os.Getenv("POSTGRES_DB")

	if Host == "" || Port == "" || User == "" || Pass == "" || DB == "" {
		log.Println("Aviso: algumas variáveis do .env não foram carregadas corretamente")
		log.Printf("POSTGRES_HOST = %s\nPOSTGRES_PORT = %s\nPOSTGRES_USER = %s\nPOSTGRES_PASSWORD = %s\nPOSTGRES_DB = %s\n",
			Host, Port, User, Pass, DB)
	}
}

func GetPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Pass, DB)
}

func ConnectDB() (*sql.DB, error) {
	dsn := GetPostgresDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir conexão: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar: %w", err)
	}

	fmt.Println("Conexão com o PostgreSQL realizada com sucesso!")
	return db, nil
}
