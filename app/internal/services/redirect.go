// Redirect service
package services

import (
	"context"
	"math/rand"
	"short_url/app/internal/db"
	"short_url/app/internal/models"
	"time"
)

// Redirect Service
type RedirService struct {
	DB *db.DB
}

func (s *RedirService) Create(ctx context.Context, redirect models.CreateRedirectForm, coll string) (models.Redirect, error) {
	result, err := s.DB.Create(ctx, redirect, coll)
	if err != nil {
		return models.Redirect{}, err
	}
	response := models.Redirect{}
	id := result.InsertedID.(db.ObjectID)
	response.ID = id
	response.Url = redirect.Url
	response.Redirect = redirect.Redirect
	response.Count = redirect.Count
	response.Deleted = redirect.Deleted
	response.CreatedAt = redirect.CreatedAt
	response.UpdatedAt = redirect.UpdatedAt
	return response, nil
}

func (s *RedirService) FindByRedirect(ctx context.Context, url string, coll string) (models.Redirect, error) {
	redirect := models.Redirect{}
	filter := db.MType{
		"redirect": url,
		"deleted":  false,
	}
	result := s.DB.FindOneBy(ctx, filter, coll)
	err := result.Decode(&redirect)
	return redirect, err
}

func (s *RedirService) FindByUrl(ctx context.Context, url string, coll string) (models.Redirect, error) {
	redirect := models.Redirect{}
	filter := db.MType{
		"url":     url,
		"deleted": false,
	}
	result := s.DB.FindOneBy(ctx, filter, coll)
	err := result.Decode(&redirect)
	return redirect, err
}

func (s *RedirService) Update(ctx context.Context, id db.ObjectID, redirect interface{}, coll string) (*db.UpdateResult, error) {
	return s.DB.Update(ctx, id, redirect, coll)
}

func (s *RedirService) RedirectGenerator() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	str := string(b)
	return str
}
