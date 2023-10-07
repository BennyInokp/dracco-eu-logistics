package logs

import "time"

type AuditLog struct {
	LogID       int       `json:"log_id"`
	TableName   string    `json:"table_name"`
	Action      string    `json:"action"`
	ChangeDate  time.Time `json:"change_date"`
	UserID      int       `json:"user_id"`
	ExecutedSQL string    `json:"executed_sql"`
}
