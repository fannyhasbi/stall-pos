package product

import (
  "log"
  "encoding/json"
  "net/http"

  "github.com/fannyhasbi/stall-pos/common"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
  var product Product
  var arrProducts []Product

  db := common.Connect()
  defer db.Close()

  rows, err := db.Query("SELECT * FROM product")
  if err != nil {
    log.Println(err)
  }

  for rows.Next() {
    if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
      log.Println(err)
      return
    }

    arrProducts = append(arrProducts, product)
  }

  response := ResponseProduct{
    Status: http.StatusOK,
    Data: arrProducts,
  }  

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}