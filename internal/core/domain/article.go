package domain

import "time"

type Article struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Summary string `json:"summary"`
    Content string `json:"content"`// Markdown 
    Author string `json:"author"`
    Category *Category  `json:"category"`
    Keywords []*Keyword `json:"keywords"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

