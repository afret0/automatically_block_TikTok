package script

var code *Code

type Code struct {
	Failed int
}

func init() {
	code = new(Code)
	code.Failed = 100201
}

func GetCode() *Code {
	return code
}
