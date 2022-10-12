package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

type Order struct {
	OrderID      string    `json:"orderId" example: "1"`
	CustomerName string    `json:"customerName" example: "Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example: "2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemID      string `json:"itemId" example:"A1B2C3"`
	Description string `json:"description" example:"A random description"`
	Quantity    string `json:"quantity" example:"1"`
}

var orders []Order
var prevOrderID = 0

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()

	//create
	router.HandleFunc("/orders", createOrder).Methods("POST")
	//get all
	router.HandleFunc("/orders", getOrders).Methods("GET")
	//get where
	router.HandleFunc("/order/{orderId}", getOrderWhere).Methods("GET")
	//delete where
	router.HandleFunc("/order/{orderId}", getOrderWhere).Methods("DELETE")
	// update
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	//swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//call to browser http://loalhost:8080/swagger/index.html

	log.Fatal(http.ListenAndServe(":8080", router))
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// GetOrder godoc
// @Summary Get details of all orders where orderId
// @Description Get details of all orders where orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 200 {object} Order
// @Router /order/{orderId} [get]
func getOrderWhere(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]
	for _, order := range orders {
		if order.OrderID == inputOrderId {
			json.NewEncoder(w).Encode(orders)
			return
		}
	}
}

// deleteOrder godoc
// @Summary Delete data order where orderId
// @Description Delete data order where orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 204 "No Content"
// @Router /order/{orderId} [delete]
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderId {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

// UpdateOrder godoc
// @Summary Update order identified by the given orderId
// @Description Update the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be updated"
// @Success 200 {object} Order
// @Router /orders/{orderId} [put]
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}

//Swagger:
//setiap ada perubahan harus jalankan command: swag init -g main.go
