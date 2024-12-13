package models

import "time"

type Food struct {
	ID        *int64    `db:"id" json:"id"`
	Name      *string   `db:"name" json:"name"`
	Calories  *float64  `db:"calories" json:"calories"`
	Photo     *string   `db:"photo" json:"photo"`
	Notes     *string   `db:"notes" json:"notes"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	IsVisible *bool `db:"is_visible" json:"is_visible"`
}