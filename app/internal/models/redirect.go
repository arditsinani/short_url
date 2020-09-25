package models

import (
	"short_url/app/internal/db"
	"time"
)

// Redirect model
type Redirect struct {
	ID        db.ObjectID `updatable:"false" unique:"true" db:"_id" json:"id" bson:"_id"`
	Url       string      `updatable:"false" unique:"true" db:"url" json:"url" bson:"url"`
	Redirect  string      `updatable:"true" unique:"false" db:"redirect" json:"redirect" bson:"redirect"`
	Count     int         `updatable:"true" unique:"false" db:"count" json:"count" bson:"count"`
	Deleted   bool        `updatable:"true" unique:"false" db:"deleted" json:"deleted" bson:"deleted" default:"false"`
	CreatedAt time.Time   `updatable:"false" unique:"false" db:"created_at" json:"created_at" bson:"created_at"`
	UpdatedAt time.Time   `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at"`
}

type CreateRedirectForm struct {
	Url       string    `updatable:"false" unique:"true" db:"url" json:"url" bson:"url"`
	Redirect  string    `updatable:"false" unique:"false" db:"redirect" json:"redirect" bson:"redirect"`
	Count     int       `updatable:"true" unique:"false" db:"count" json:"count" bson:"count"`
	Deleted   bool      `updatable:"true" unique:"false" db:"deleted" json:"deleted" bson:"deleted" default:"false"`
	CreatedAt time.Time `updatable:"false" unique:"false" db:"created_at" json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at"`
}

type UpdateRedirectForm struct {
	Count     int       `updatable:"true" unique:"false" db:"count" json:"count" bson:"count"`
	UpdatedAt time.Time `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at"`
}

type DeleteRedirectForm struct {
	Deleted   bool      `updatable:"true" unique:"false" db:"deleted" json:"deleted" bson:"deleted" default:"false"`
	UpdatedAt time.Time `updatable:"true" unique:"false" db:"updated_at" json:"updated_at" bson:"updated_at"`
}

type RedirectsCollection struct {
	Collection []Redirect
}

const (
	REDIRECTSCOLLECTION = "redirects"
)
