package structs

type User struct {
    ID       int
    Username string
    Email    string
    Active   bool
}

type Post struct {
    ID      int
    Title   string
    Content string
    Author  User
}
