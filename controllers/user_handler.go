package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tutorials/dynamic_db_conn_project/constants"
	"tutorials/dynamic_db_conn_project/dal"
	response "tutorials/dynamic_db_conn_project/json_response"
	"tutorials/dynamic_db_conn_project/models"

	"github.com/go-chi/chi/v5"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	dbConn := r.Context().Value(constants.DB_CONN_CTX_KEY).(*sql.DB)

	var err error
	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.RespondWithJSONWithPayload(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
		return
	}

	err = dal.RegisterUser(dbConn, user)
	if err != nil {
		response.RespondWithJSONWithPayload(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to register user",
		})
		return
	}

	response.RespondWithJSONWithPayload(w, http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	dbConn := r.Context().Value(constants.DB_CONN_CTX_KEY).(*sql.DB)

	users, err := dal.GetUsers(dbConn)
	if err != nil {
		response.RespondWithJSONWithPayload(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get users",
		})
		return
	}

	response.RespondWithJSONWithPayload(w, http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	dbConn := r.Context().Value(constants.DB_CONN_CTX_KEY).(*sql.DB)
	userID := chi.URLParam(r, "id")

	var err error
	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.RespondWithJSONWithPayload(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
		return
	}

	err = dal.UpdateUser(dbConn, userID, user)
	if err != nil {
		response.RespondWithJSONWithPayload(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update user",
		})
		return
	}

	response.RespondWithJSONWithPayload(w, http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	dbConn := r.Context().Value(constants.DB_CONN_CTX_KEY).(*sql.DB)
	userID := chi.URLParam(r, "id")

	err := dal.DeleteUser(dbConn, userID)
	if err != nil {
		response.RespondWithJSONWithPayload(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete user",
		})
		return
	}

	response.RespondWithJSONWithPayload(w, http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}