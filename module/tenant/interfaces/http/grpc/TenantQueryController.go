package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gomora/internal/errors"
	"gomora/module/tenant/application"
	grpcPB "gomora/module/tenant/interfaces/http/grpc/pb"
)

// TenantQueryController handles the grpc tenant query requests
type TenantQueryController struct {
	application.TenantQueryServiceInterface
}

// GetTenantByID retrieves the tenant id from the proto
func (controller *TenantQueryController) GetTenantByID(ctx context.Context, req *grpcPB.TenantRequest) (*grpcPB.TenantResponse, error) {
	res, err := controller.TenantQueryServiceInterface.GetTenantByID(context.TODO(), req.Id)
	if err != nil {
		var code codes.Code

		switch err.Error() {
		case errors.DatabaseError:
			code = codes.Internal
		case errors.MissingRecord:
			code = codes.NotFound
		default:
			code = codes.Unknown
		}

		st := status.New(code, err.Error())

		return nil, st.Err()
	}

	createProtoTime, _ := ptypes.TimestampProto(res.CreatedAt)
	updateProtoTime, _ := ptypes.TimestampProto(res.UpdatedAt)

	return &grpcPB.TenantResponse{
		Id:            res.ID,
		Name:          res.Name,
		Alias:         res.Alias,
		Email:         res.Email,
		Code:          res.Code,
		Address:       res.Address,
		ContactNumber: res.ContactNumber,
		Avatar:        res.Avatar,
		IsActive:      res.IsActive,
		CreatedAt:     createProtoTime,
		UpdatedAt:     updateProtoTime,
	}, nil
}
