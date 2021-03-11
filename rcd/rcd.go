package rcd

// status codes
const (
	FilenameExist = 1000
)

var text = map[int]string{
	FilenameExist: "Filename is already used",
}

// Stats -- return string per code
func Stats(code int) string {
	return text[code]
}
