package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"gomora/interfaces/http/rest/viewmodels"
	"gomora/internal/errors"
	"gomora/module/tenant/application"
)

type TenantQueryController struct {
	application.TenantQueryServiceInterface
}

// GetTenantByID retrieves the tenant id from the rest request
func (controller *TenantQueryController) GetTenantByID(w http.ResponseWriter, r *http.Request) {
	tenantID := chi.URLParam(r, "id")

	if len(tenantID) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid tenant ID",
			ErrorCode: errors.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	res, err := controller.TenantQueryServiceInterface.GetTenantByID(context.TODO(), tenantID)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Encountered persistence problem."
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No records found."
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
		Message: "Tenant successfully fetched.",
		Data:    res,
	}

	response.JSON(w)
	return
}
