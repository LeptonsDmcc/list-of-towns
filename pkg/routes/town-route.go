package routes

import (
	"listof/pkg/controllers"

	"github.com/gorilla/mux"
)

func SetupCityRoutes(router *mux.Router) {
	router.HandleFunc("/towns", controllers.GetTowns).Methods("GET")
	router.HandleFunc("/lgs/{lg_id}/towns", controllers.GetTownsByLGA).Methods("GET")
}
