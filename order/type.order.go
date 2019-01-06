package order

type OrderDetails struct {
	ID        int     `json:"id" db:"id"`
	ProductID int     `json:"product_id" db:"product_id"`
	OrderID   int     `json:"order_id" db:"order_id"`
	Amount    float64 `json:"amount"`
}

type orderInsert struct {
  EmployeeID int
  CustomerID int
}

type orderDetailsInsert struct {
  OrderID int
  ProductID int
}

type commonResponse struct {
	Status int `json:"status"`
	Data   interface{} `json:"data"`
}
