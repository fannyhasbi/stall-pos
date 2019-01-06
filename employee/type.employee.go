package employee

type Employee struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
}

type ResponseEmployee struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
