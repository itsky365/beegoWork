package models

import "fmt"

type Users struct {
    Id     int64
    Name   string
    Emails []string
}

func (u Users) String() string {
    return fmt.Sprintf("Users<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
    Id       int64
    Title    string
    AuthorId int64
    Author   *Users
}

func (s Story) String() string {
    return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

type Zhuxh struct {
    Id int
    Name string
    Email  []string `pg:",array"`
}

func (z Zhuxh) String() string {
    return fmt.Sprintf("Zhuxh<%d %s %v>", z.Id, z.Name, z.Email)
}