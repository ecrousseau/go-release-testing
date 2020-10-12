package main
import "fmt"
import "rsc.io/quote"
func myQuote() string {
	return quote.Hello()
}
func main() {
    fmt.Println(myQuote())
}