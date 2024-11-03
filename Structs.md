```go
package main

import "fmt"

// Define the User struct
type User struct {
    ID   int
    Name string
}

// Method that returns a greeting string
func (u *User) Greet() string {
    return "Hello, " + u.Name
}


// Normal function that takes a User and returns a greeting string
func greetUser(u User) string {
    return "Hello, " + u.Name
}

func main() {
    user := User{ID: 1, Name: "Alice"}
    fmt.Println(user.Greet()) // Output: Hello, Alice
    fmt.Println(greetUser(user)) // Output: Hello, Alice
}
```
In this example, the Greet method is associated with the User struct, and any instance of User can call this method to return a greeting.