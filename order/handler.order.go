package order

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fannyhasbi/stall-pos/common"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {
	if len(r.FormValue("id_product")) == 0 || len(r.FormValue("id_employee")) == 0 {
		common.CommonResponse(w, r, http.StatusBadRequest, nil)
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
		common.CommonResponse(w, r, http.StatusInternalServerError, nil)

		log.Println(err)
		return
	}

	id_product, err := strconv.Atoi(r.FormValue("id_product"))
	if err != nil {
		log.Println(err)
		return
	}

	ordDetails := orderDetailsInsert{
		OrderID:   orderId,
		ProductID: id_product,
	}

	err = insertOrderDetails(ordDetails)
	if err != nil {
		common.CommonResponse(w, r, http.StatusInternalServerError, nil)
		log.Println(err)

		return
	}

	common.CommonResponse(w, r, http.StatusOK, orderId)
}
