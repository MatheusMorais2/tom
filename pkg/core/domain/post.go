package domain

import "time"

type Post struct {
    Id string
    Title string
    Summary string
    Content string // Markdown
    Category string
    Keywords []string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Category string
const (
    Auditorium Category = "auditorium"
    Bistro Category = "bistro"
    Cinema Category = "cinema"
    Lounge Category = "lounge"
)
