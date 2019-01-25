package order

import (
  "fmt"
  "github.com/fannyhasbi/stall-pos/common"
)

func insertOrder(ord orderInsert) (int, error) {
  db, err := common.Connect()
  if err != nil {
    return 0, err
  }

  defer db.Close()

  var query string

  if ord.CustomerID != 0 {
    query = fmt.Sprintf("INSERT INTO orders (cust_type, emp_id, cust_id) VALUES (2, %d, %d)", ord.EmployeeID, ord.CustomerID)
  } else {
    query = fmt.Sprintf("INSERT INTO orders (emp_id) VALUES (%d)", ord.EmployeeID)
  }

  _, err = db.Exec(query)
  if err != nil {
    return 0, err
  }

  var lastId int

  row := db.QueryRow("SELECT LAST_INSERT_ID() as id")
  err = row.Scan(&lastId)

  if err != nil {
    return 0, err
  }

  return lastId, nil
}

func insertOrderDetails(ord orderDetailsInsert) error {
  db, err := common.Connect()
  if err != nil {
    return err
  }
  defer db.Close()

  query := fmt.Sprintf("INSERT INTO order_details (product_id, order_id) VALUES (%d, %d)", ord.ProductID, ord.OrderID)

  _, err = db.Exec(query)
  if err != nil {
    return err
  }

  return nil
}