package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"main/utils"
	"main/utils/libs"
	"net/http"
)

func GetManyProperties(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	propertiesService := application.GetManyPropertiesService{Database: database}

	searchFilter := request.URL.Query().Get("search")
	latitudeFilter := request.URL.Query().Get("latitude")
	longitudeFilter := request.URL.Query().Get("longitude")
	isNewFilter := request.URL.Query().Get("is_new")
	withDiscountFilter := request.URL.Query().Get("with_discount")
	recentlySold := request.URL.Query().Get("recently_sold")
	recentlyBuilt := request.URL.Query().Get("recently_built")
	isSpecial := request.URL.Query().Get("is_special")
	isApartment := request.URL.Query().Get("is_apartment")
	allowFinancing := request.URL.Query().Get("allow_financing")

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

	filters.IsNew = utils.ParseParamToBool(isNewFilter)
	filters.WithDiscount = utils.ParseParamToBool(withDiscountFilter)
	filters.RecentlySold = utils.ParseParamToBool(recentlySold)
	filters.RecentlyBuilt = utils.ParseParamToBool(recentlyBuilt)
	filters.IsSpecial = utils.ParseParamToBool(isSpecial)
	filters.IsApartment = utils.ParseParamToBool(isApartment)
	filters.AllowFinancing = utils.ParseParamToBool(allowFinancing)

	properties, getPropertiesErr := propertiesService.Execute(filters)
	if getPropertiesErr != nil {
		core.HandleHTTPStatus(write, getPropertiesErr)
		return
	}

	var response []presenters.PropertyToHTTP

	for _, property := range *properties {
		response = append(response, httpPresenter.ToHTTP(property))
	}

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}
