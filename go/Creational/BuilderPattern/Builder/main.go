package main

import "fmt"

type User struct {
    FirstName, LastName, Address, Phone string
}
type UserBuilder struct{
    User *User
}
func NewUserBuilder(firstname, lastname string) *UserBuilder {
    return &UserBuilder{
        &User{
            FirstName: firstname,
            LastName:  lastname,
        },
    }
}
func (ub *UserBuilder) StaysAt(address string) *UserBuilder {
    ub.User.Address = address
    return ub
}
func (ub *UserBuilder) ConnectOn(phone string) *UserBuilder {
    ub.User.Phone = phone
    return ub
}
func (ub *UserBuilder) Build() *User {
    return ub.User
}
func main() {
    ub := NewUserBuilder("Rajiv", "Birla")
    user := ub.Build()
    fmt.Println("User: ", *user)

    user = ub.StaysAt("Golden City").ConnectOn("9090909090").Build()
    fmt.Println("User Info: ", *user)
}