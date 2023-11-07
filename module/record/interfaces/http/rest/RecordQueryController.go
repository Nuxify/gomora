package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"gomora/interfaces/http/rest/viewmodels"
	"gomora/internal/errors"
	"gomora/module/record/application"
	types "gomora/module/record/interfaces/http"
)

// RecordQueryController request controller for record query
type RecordQueryController struct {
	application.RecordQueryServiceInterface
}

// GetRecordByID retrieves the tenant id from the rest request
func (controller *RecordQueryController) GetRecordByID(w http.ResponseWriter, r *http.Request) {
	recordID := chi.URLParam(r, "id")

	if len(recordID) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid record ID",
			ErrorCode: errors.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	res, err := controller.RecordQueryServiceInterface.GetRecordByID(context.TODO(), recordID)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error while fetching record."
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No record found."
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
		Message: "Record successfully fetched.",
		Data: &types.RecordResponse{
			ID:        res.ID,
			Data:      res.Data,
			CreatedAt: res.CreatedAt.Unix(),
		},
	}

	response.JSON(w)
}
