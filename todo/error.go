package todo

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
