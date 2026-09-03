package grpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gomora/internal/errors"
	"gomora/module/record/application"
	grpcPB "gomora/module/record/interfaces/http/grpc/pb"
)

// RecordQueryController handles the grpc record query requests
type RecordQueryController struct {
	application.RecordQueryServiceInterface
}

// GetRecordByID retrieves the record id from the proto
func (controller *RecordQueryController) GetRecordByID(ctx context.Context, req *grpcPB.GetRecordRequest) (*grpcPB.RecordResponse, error) {
	res, err := controller.RecordQueryServiceInterface.GetRecordByID(context.TODO(), req.Id)
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

	createProtoTime, _ := ptypes.TimestampProto(res.CreatedAt)

	return &grpcPB.RecordResponse{
		Id:        res.ID,
		Data:      res.Data,
		CreatedAt: createProtoTime,
	}, nil
}
