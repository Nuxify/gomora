package http

// CreateRecordRequest request struct for create record
type CreateRecordRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// RecordResponse response struct
type RecordResponse struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"createdAt"`
}
