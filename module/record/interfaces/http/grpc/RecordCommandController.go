package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gomora/internal/errors"
	"gomora/module/record/application"
	serviceTypes "gomora/module/record/infrastructure/service/types"
	grpcPB "gomora/module/record/interfaces/http/grpc/pb"
)

// RecordCommandController handles the grpc record command requests
type RecordCommandController struct {
	application.RecordCommandServiceInterface
	grpcPB.UnimplementedRecordCommandServiceServer
}

// CreateRecord creates a new record
func (controller *RecordCommandController) CreateRecord(ctx context.Context, req *grpcPB.CreateRecordRequest) (*grpcPB.RecordResponse, error) {
	record := serviceTypes.CreateRecord{
		ID:   req.Id,
		Data: req.Data,
	}

	res, err := controller.RecordCommandServiceInterface.CreateRecord(context.TODO(), record)
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

		st := status.New(code, fmt.Sprintf("[RECORD] %s", err.Error()))

		return nil, st.Err()
	}

	createProtoTime := timestamppb.New(time.Now())

	return &grpcPB.RecordResponse{
		Id:        res.ID,
		Data:      res.Data,
		CreatedAt: createProtoTime,
	}, nil
}
