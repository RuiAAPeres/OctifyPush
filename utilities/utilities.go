package utilities

func PanicError(anError error) {
	if anError != nil {
		panic(anError)
	}
}