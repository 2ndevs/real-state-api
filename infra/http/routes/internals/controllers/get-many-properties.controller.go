package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"main/utils/libs"
	"net/http"
)

func GetManyProperties(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	propertiesService := application.GetManyPropertiesService{Database: database}

	searchFilter := request.URL.Query().Get("search")
	latitudeFilter := request.URL.Query().Get("latitude")
	longitudeFilter := request.URL.Query().Get("longitude")
	filters := application.GetManyPropertiesFilters{}

	if searchFilter != "" {
		filters.Search = &searchFilter
	}

	latitude := libs.ValidateAndConvertCoordinate(latitudeFilter, -90, 90)
	longitude := libs.ValidateAndConvertCoordinate(longitudeFilter, -180, 180)
	if latitude != nil && longitude != nil {
		filters.Latitude = latitude
		filters.Longitude = longitude
	}

	properties, getPropertiesErr := propertiesService.Execute(filters)
	if getPropertiesErr != nil {
		http.Error(write, getPropertiesErr.Error(), http.StatusInternalServerError)
		return
	}

	var response []presenters.PropertyToHTTP

	for _, property := range *properties {
		response = append(response, httpPresenter.ToHTTP(property))
	}

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}
