package domain

type CategoryEnum string

const (
    Auditorium CategoryEnum = "auditorium"
    Bistro CategoryEnum = "bistro"
    Cinema CategoryEnum = "cinema"
    Lounge CategoryEnum = "lounge"
)

type Category struct {
    Id string `json:"id"`
    Name CategoryEnum `json:"name"`
}
