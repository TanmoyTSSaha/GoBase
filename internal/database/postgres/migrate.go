package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func PerformMigration(db *sql.DB, migrationDir string) error {
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			filePath := filepath.Join(migrationDir, file.Name())
			if err := executeSQLFile(db, filePath); err != nil {
				return fmt.Errorf("failed to execute migration %s: %v", filePath, err)
			}

			log.Printf("Successfully applied migrations: %s", file.Name())
		}
	}

	return nil
}

func executeSQLFile(db *sql.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file %s: %v", filePath, err)
	}

	if _, err := db.Exec(string(content)); err != nil {
		return fmt.Errorf("failed to execute SQL file %s: %v", filePath, err)
	}

	return nil
}
