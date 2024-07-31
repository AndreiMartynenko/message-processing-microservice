package postgres

const (
	InsertMessage = "INSERT INTO messages (content, status) VALUES ($1, 'pending')"
	GetStatistics = "SELECT status, COUNT(*) FROM messages GROUP BY status"
)
