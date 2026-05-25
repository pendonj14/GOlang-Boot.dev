package main

func reformat(message string, formatter func(string) string) string {
	firstResult := formatter(message)
	secondResult := formatter(firstResult)
	thirdResult := formatter(secondResult)
	finalResult := "TEXTIO: " + thirdResult
	return finalResult
}
