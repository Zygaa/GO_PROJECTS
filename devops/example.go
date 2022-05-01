package main
import "fmt"
func main()  {
  var challenge string = "#90DaysOfDevOps"
  const daysTotal uint = 90
  var daysComplete uint = 11
  var name string
  fmt.Println("Enter your name:")
  fmt.Scanln(&name)
  fmt.Println("Welcome to", challenge, name)
  fmt.Println("This is a", daysTotal, "challenge and you have completed", daysComplete, "days")
  fmt.Println("Great work!")
}
