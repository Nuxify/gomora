package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid payload sent.",
			ErrorCode: apiError.InvalidPayload,
		}

		response.JSON(w)
		return
	}

	// verify content must not empty
	if len(request.Data) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Data input cannot be empty.",
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
			httpCode = http.StatusUnprocessableEntity
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
		Data: &types.RecordResponse{
			ID:        res.ID,
			Data:      res.Data,
			CreatedAt: time.Now().Unix(),
		},
	}

	response.JSON(w)
}
