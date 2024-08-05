package controllers

import (	
    "net/http"
    "strconv"
    "strings"

    "listof/pkg/config"
    "listof/pkg/models"
    "listof/pkg/utils"
)


func GetTowns(w http.ResponseWriter, r *http.Request) {
	
	towns := config.Towns
	if len(towns) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No towns, suburbs or villages found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, towns)
}

func GetTownsByLGA(w http.ResponseWriter, r *http.Request) {
    
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 3 || pathParts[1] != "lgs" || pathParts[3] != "towns" {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid URL path")
        return
    }

    lgID, err := strconv.Atoi(pathParts[2])
    if err != nil {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid LGA ID")
        return
    }

    var towns []models.Town
    for _, city := range config.Towns {
        if city.LGID == lgID {
            towns = append(towns, city)
        }
    }

    if len(towns) == 0 {
        utils.RespondWithJSON(w, http.StatusNotFound, "No towns, suburbs or villages found for the given local government ID")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, towns)
}
