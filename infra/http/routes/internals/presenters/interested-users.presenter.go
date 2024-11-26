package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type InterestedUsersPresenter struct{}

type InterestedUsersFromHTTP struct {
	FirstName string  `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string  `json:"last_name" validate:"required,min=2,max=50"`
	Phone     string  `json:"phone" validate:"required,min=11,max=11"`
	Email     *string `json:"email" validate:"email,omitempty"`
	Answered  *bool   `json:"answered"`

	PropertyID uint  `json:"property_id" validate:"required,min=1"`
	StatusID   *uint `json:"status_id" validate:"gte=1,lte=2"`
}

type InterestedUsersToHTTP struct {
	ID        uint    `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Phone     string  `json:"phone"`
	Email     *string `json:"email"`
	Answered  *bool   `json:"answered"`

	StatusID   *uint              `json:"status_id"`
	Status     *entities.Status   `json:"status"`
	PropertyID uint               `json:"property_id"`
	Property   *entities.Property `json:"property"`
}

func (InterestedUsersPresenter) FromHTTP(request *http.Request) (*InterestedUsersFromHTTP, error) {
	var interestedUserRequest InterestedUsersFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&interestedUserRequest)
	if err != nil {
		return nil, err
	}

	return &interestedUserRequest, nil
}

func (InterestedUsersPresenter) ToHTTP(interestedUser entities.InterestedUser) InterestedUsersToHTTP {
	return InterestedUsersToHTTP{
		ID:        interestedUser.ID,
		FirstName: interestedUser.FirstName,
		LastName:  interestedUser.LastName,
		Phone:     interestedUser.Phone,
		Email:     interestedUser.Email,
		Answered:  interestedUser.Answered,

		StatusID: interestedUser.StatusID,
		Status:   interestedUser.Status,

		PropertyID: interestedUser.PropertyID,
		Property:   interestedUser.Property,
	}
}
