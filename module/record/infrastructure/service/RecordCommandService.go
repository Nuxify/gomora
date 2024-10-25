package service

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
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

// CreateRecord create a record also generates a jwt token
func (service *RecordCommandService) CreateRecord(ctx context.Context, data types.CreateRecord) (entity.Record, string, error) {
	record := repositoryTypes.CreateRecord{
		ID:   data.ID,
		Data: data.Data,
	}

	// check id if empty create new unique id
	if len(record.ID) == 0 {
		record.ID = generateID()
	}

	// generate jwt token
	accessToken, err := service.GenerateJWT(record.ID)
	if err != nil {
		return entity.Record{}, "", err
	}

	res, err := service.RecordCommandRepositoryInterface.InsertRecord(record)
	if err != nil {
		return entity.Record{}, "", err
	}

	return res, accessToken, nil
}

// GenerateJWT generates a jwt token
func (service *RecordCommandService) GenerateJWT(id string) (string, error) {
	// create access token
	accessTokenClaims := jwt.MapClaims{
		"id":  id,
		"iss": "gomora",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * 30).Unix(), // 30 seconds for testing purpose
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateID generates unique id
func generateID() string {
	return ksuid.New().String()
}
