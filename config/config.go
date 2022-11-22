package config

type Flags struct {
	URL         string
	Wordlist    string
	RequestType string
	StatusShow  int
	StatusHide  int
}

var (
	DefaultRequestType string = "GET"
	DefaultStatusShow  int    = 0
	DefaultStatusHide  int    = 404

	AppFlag Flags = Flags{}
)
