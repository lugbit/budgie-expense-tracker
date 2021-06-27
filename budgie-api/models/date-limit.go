package models

// model for date range and pagination offset and limit
type DateRangeLimit struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}
