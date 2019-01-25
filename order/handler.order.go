package order

import (
  "log"
  "encoding/json"
  "net/http"
  "strconv"
  
  "github.com/fannyhasbi/stall-pos/common"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {
  if len(r.FormValue("id_product")) == 0 || len(r.FormValue("id_employee")) == 0 {
    common.SendJSONError(w, r, "Bad Request", http.StatusBadRequest)
    return
  }

  id_employee, err := strconv.Atoi(r.FormValue("id_employee"))
  if err != nil {
    log.Println(err)
    return
  }

  var id_customer int
  if len(r.FormValue("id_customer")) == 0 {
    id_customer = 0
  } else {
    id_customer, err = strconv.Atoi(r.FormValue("id_customer"))
    if err != nil {
      log.Println(err)
      return
    }
  }

  ord := orderInsert{
    EmployeeID: id_employee,
    CustomerID: id_customer,
  }

  orderId, err := insertOrder(ord)

  if err != nil {
    log.Println(err)
    common.SendJSONError(w, r, err.Error(), http.StatusInternalServerError)
    return
  }

  id_product, err := strconv.Atoi(r.FormValue("id_product"))
  if err != nil {
    log.Println(err)
    return
  }
  
  ordDetails := orderDetailsInsert{
    OrderID: orderId,
    ProductID: id_product,
  }

  err = insertOrderDetails(ordDetails)
  if err != nil {
    log.Println(err)
    common.SendJSONError(w, r, err.Error(), http.StatusInternalServerError)

    return
  }

  response := commonResponse{
    Status: http.StatusOK,
    Data: orderId,
  }

  res, err := json.Marshal(response)
  if err != nil {
    log.Println(err)
    common.SendJSONError(w, r, err.Error(), http.StatusInternalServerError)
    return
  }

  common.SendJSON(w, r, res, http.StatusOK)
}