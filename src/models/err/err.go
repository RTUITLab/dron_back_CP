package err

type Error struct {
	Error		string		`json:"error"`
}

func FromError(err error) Error {
	return Error{
		Error: err.Error(),
	}
}

func FromString(err string) Error {
	return Error{
		Error: err,
	}
}