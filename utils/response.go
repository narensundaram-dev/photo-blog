package utils

// Error struct
type Error struct {
	Message     string
	Description string
}

// Data struct
type Data struct {
	Message string
	Body    interface{}
}

// Response struct
type Response struct {
	Data
	Error
}

func (err *Error) Error() string {
	return err.Message
}
