package employee

import (
  "log"
  "net/http"
  "database/sql"

  "github.com/fannyhasbi/stall-pos/common"
)

func HandleEmployee(w http.ResponseWriter, r *http.Request) {
  var employee Employee

  username := r.FormValue("username")
  password := r.FormValue("password")

  if len(username) == 0 || len(password) == 0 {
    response := ResponseEmployee{
      Status: http.StatusBadRequest,
      Data: nil,
    }

    common.SendJSON(w, r, response)
    return
  }

  db := common.Connect()
  defer db.Close()

  row := db.QueryRow("SELECT id, name, username FROM employee WHERE username = ? AND password = MD5(?)", username, password)
  err := row.Scan(&employee.ID, &employee.Name, &employee.Username)

  if err != nil {
    if err == sql.ErrNoRows {
      response := ResponseEmployee{
        Status: http.StatusNotFound,
        Data: nil,
      }

      common.SendJSON(w, r, response)
      return
    } else {
      log.Println(err)
      return
    }
  }

  response := ResponseEmployee{
    Status: http.StatusOK,
    Data: employee,
  }

  common.SendJSON(w, r, response)
}