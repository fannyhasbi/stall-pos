package order

import (
  "log"
  "net/http"
  "strconv"
  
  "github.com/fannyhasbi/stall-pos/common"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {
  if len(r.FormValue("id_product")) == 0 || len(r.FormValue("id_employee")) == 0 {
    response := commonResponse{
      Status: http.StatusBadRequest,
      Data: nil,
    }
    
    common.SendJSON(w, r, response)
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
    response := commonResponse{
      Status: http.StatusInternalServerError,
      Data: nil,
    }
    common.SendJSON(w, r, response)
    log.Println(err)
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
    response := commonResponse{
      Status: http.StatusInternalServerError,
      Data: nil,
    }
    common.SendJSON(w, r, response)
    log.Println(err)

    return
  }

  response := commonResponse{
    Status: http.StatusOK,
    Data: orderId,
  }

  common.SendJSON(w, r, response)
}