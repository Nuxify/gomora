package service

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/segmentio/ksuid"

	"gomora/module/record/domain/entity"
	"gomora/module/record/domain/repository"
	repositoryTypes "gomora/module/record/infrastructure/repository/types"
	"gomora/module/record/infrastructure/service/types"
)

// RecordCommandService handles the record command service logic
type RecordCommandService struct {
	repository.RecordCommandRepositoryInterface
}

// CreateRecord create a record
func (service *RecordCommandService) CreateRecord(ctx context.Context, data types.CreateRecord) (entity.Record, error) {
	record := repositoryTypes.CreateRecord{
		ID:   data.ID,
		Data: data.Data,
	}

	// check id if empty create new unique id
	if len(record.ID) == 0 {
		record.ID = generateID()
	}

	res, err := service.RecordCommandRepositoryInterface.InsertRecord(record)
	if err != nil {
		return entity.Record{}, err
	}

	return res, nil
}

// GenerateToken generates a jwt token
func (service *RecordCommandService) GenerateToken(ctx context.Context) (types.JWTResponse, error) {
	// create access token
	expiresAt := time.Now().Add(time.Minute * 10080) // TODO: currently 7 days
	accessTokenClaims := jwt.MapClaims{
		"iss": "gomora",
		"iat": time.Now().Unix(),
		"exp": expiresAt.Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return types.JWTResponse{}, err
	}

	return types.JWTResponse{
		AccessToken: token,
		ExpiresAt:   expiresAt,
	}, nil
}

// generateID generates unique id
func generateID() string {
	return ksuid.New().String()
}
