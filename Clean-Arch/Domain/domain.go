package domain

// User represents the core user entity.
type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	UserType     string
	Phone        string
	CreatedAt    int64
	UpdatedAt    int64
	Token        string
	RefreshToken string
}

// Task represents the core task entity.
type Task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}
