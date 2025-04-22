package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	"gomora/interfaces/http/rest/viewmodels"
	"gomora/internal/errors"
	apiError "gomora/internal/errors"
	"gomora/module/record/application"
	serviceTypes "gomora/module/record/infrastructure/service/types"
	types "gomora/module/record/interfaces/http"
)

// RecordCommandController request controller for record command
type RecordCommandController struct {
	application.RecordCommandServiceInterface
}

// CreateRecord request handler to create record
func (controller *RecordCommandController) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var request types.CreateRecordRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// validate request
	err := types.Validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   types.ValidationErrors[errors[0].StructNamespace()],
				ErrorCode: apiError.InvalidPayload,
			}

			response.JSON(w)
			return
		}

		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	record := serviceTypes.CreateRecord{
		ID:   request.ID,
		Data: request.Data,
	}

	res, err := controller.RecordCommandServiceInterface.CreateRecord(context.TODO(), record)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving record."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "Record ID already exist."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully created record.",
		Data: &types.CreateRecordResponse{
			ID:        res.ID,
			Data:      res.Data,
			CreatedAt: time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// GenerateToken request handler to generate token
func (controller *RecordCommandController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	res, err := controller.RecordCommandServiceInterface.GenerateToken(context.TODO())
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while generating token."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	// set httponly cookie
	controller.setJWTCookie(w, res.AccessToken, res.ExpiresAt)

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully generated token.",
		Data: &types.GenerateTokenResponse{
			AccessToken: res.AccessToken,
		},
	}

	response.JSON(w)
}

func (controller *RecordCommandController) setJWTCookie(w http.ResponseWriter, token string, expiresAt time.Time) {
	cookie := &http.Cookie{
		Name:     "jwt", // required by jwtauth.Verifier
		Value:    token,
		Path:     "/",
		Expires:  expiresAt,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	// enforce domain for production
	if os.Getenv("API_ENV") == "production" {
		cookie.Domain = ".nuxify.tech" // allow access to all subdomains only
	}

	http.SetCookie(w, cookie)
}

func (controller *RecordCommandController) clearJWTCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	// enforce domain for production
	if os.Getenv("API_ENV") == "production" {
		cookie.Domain = ".nuxify.tech" // allow access to all subdomains only
	}

	http.SetCookie(w, cookie)
}
