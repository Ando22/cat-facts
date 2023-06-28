package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Ando22/user-message/backend/database"
	"github.com/Ando22/user-message/backend/models"
	"github.com/Ando22/user-message/backend/utils"
)

// GetFacts handles the retrieval of messages
func GetFacts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := database.GetDB()
	tx, err := db.Begin()

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch facts a")
		return
	}

	// Get cat fact data from open source api
	err = FetchFact(ctx, tx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch facts b")
		tx.Rollback()
		return
	}

	fact, err := tx.QueryContext(ctx, "SELECT * FROM facts ORDER BY id ASC")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch facts c")
		tx.Rollback()
		return
	}

	defer fact.Close()

	logString := fmt.Sprintf("%+v", fact)
	fmt.Println(logString)

	var messages []models.Fact
	for fact.Next() {
		var message models.Fact
		err := fact.Scan(&message.ID, &message.Fact, &message.Length)
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch facts d")
			tx.Rollback()
			return
		}
		messages = append(messages, message)
	}

	fmt.Println(messages)

	err = tx.Commit()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch facts c")
		return
	}

	jsonResponse(w, http.StatusOK, messages)
}

// FetchFact handles the retrieval of a cat fact
func FetchFact(ctx context.Context, tx *sql.Tx) error {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var fact models.Fact
	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO facts (fact, length) VALUES (?, ?)`, fact.Fact, fact.Length)
	return err
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
