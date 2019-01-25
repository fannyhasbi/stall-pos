package employee

import (
	"encoding/json"
	"database/sql"
	"log"
	"net/http"

	"github.com/fannyhasbi/stall-pos/common"
)

func HandleEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		common.SendJSONError(w, r, "Bad Request", http.StatusBadRequest)
		return
	}

	db, err := common.Connect()
	if err != nil {
		log.Println(err)

		common.SendJSONError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	row := db.QueryRow("SELECT id, name, username FROM employee WHERE username = ? AND password = MD5(?)", username, password)
	err = row.Scan(&employee.ID, &employee.Name, &employee.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			response := ResponseEmployee{
				Status: http.StatusNotFound,
				Data:   nil,
			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				common.SendJSONError(w, r, err.Error(), http.StatusInternalServerError)
				return
			}

			common.SendJSON(w, r, res, http.StatusOK)
			return
		} else {
			log.Println(err)
			return
		}
	}

	response := ResponseEmployee{
		Status: http.StatusOK,
		Data:   employee,
	}

	res, err := json.Marshal(response)
	if err != nil {
		common.SendJSONError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SendJSON(w, r, res, http.StatusOK)
}
