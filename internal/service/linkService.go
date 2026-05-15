package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/yihune21/link-shortner/internal/database"
)

type LinkService struct {
    q *database.Queries
}

func NewLinkService(q *database.Queries) *LinkService  {
	return &LinkService{q : q}
} 

type Link struct {
	ID       uuid.UUID   `json:"id"`
	ShortLink    	string `json:"short_link"`
	OriginalLink  		string `json:"original_link"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func mapLink(dbLink database.Link) Link{
     return Link{
        ID: dbLink.ID,
		ShortLink: dbLink.ShortLink,
		OriginalLink: dbLink.OriginalLink,
		UserId: dbLink.UserID,
		CreatedAt: dbLink.CreatedAt,
	 }
}
func mapLinks(dbLinks []database.Link)[]Link{
    Links := []Link{}
	for _ , Link := range dbLinks {
		Links = append(Links,mapLink(Link))
	}
    return Links
}

func generateRandomKey(length int) (string, error) {
    const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    result := make([]byte, length)
    for i := 0; i < length; i++ {
        num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
        if err != nil {
            return "", err
        }
        result[i] = alphabet[num.Int64()]
    }
    return string(result), nil
}
func isUniqueConstraintError(err error) bool {
    if pqErr, ok := err.(*pq.Error); ok {
        return pqErr.Code == "23505"
    }
    return false
}

func (s *LinkService) CreateLink(ctx context.Context, original_link string, dbUser *database.User) (Link, error) {
    const maxRetries = 3
    var dbLink database.Link
    var err error

    for i := 0; i < maxRetries; i++ {
        shortKey, genErr := generateRandomKey(6)
        if genErr != nil {
            return Link{}, errors.New("failed to generate short key")
        }
        dbLink, err = s.q.CreateLink(ctx, database.CreateLinkParams{
            ID:           uuid.New(),
            ShortLink:    shortKey, 
            OriginalLink: original_link,
            UserID:       dbUser.ID,
            CreatedAt:    time.Now().UTC(),
        })

        if err == nil {
            return mapLink(dbLink), nil
        }

        if !isUniqueConstraintError(err) {
            return Link{}, fmt.Errorf("database error: %w", err)
        }

    }

    return Link{}, errors.New("failed to generate a unique short link after multiple attempts")
}

func (s *LinkService)ListLink(ctx context.Context)( []Link , error)  {
	dbLinks , err := s.q.ListLinks(ctx)
	if err != nil{
		return []Link{},err
	}

	return mapLinks(dbLinks),nil
}
func (s *LinkService)GetLinkById(ctx context.Context , id uuid.UUID) (Link , error)  {
	dbLink , err  := s.q.GetLinkById(ctx , id)
	if err != nil {
		return  Link{},ErrNotFound
	}
	return mapLink(dbLink) , nil
}
func (s *LinkService)GetLinksByShortLink(ctx context.Context , short_link string) (Link , error)  {
	dbLink , err  := s.q.GetLinksByShortLink(ctx , short_link)
	if err != nil {
		return  Link{},err
	}
	return mapLink(dbLink) , nil
}
func (s *LinkService)GetLinksByUserId(ctx context.Context , dbUser *database.User) ([]Link , error)  {
	id := dbUser.ID
	dbLink , err  := s.q.GetLinksByUserId(ctx , id)
	if err != nil {
		return  []Link{},ErrNotFound
	}
	return mapLinks(dbLink) , nil
}
func (s *LinkService)DeleteLink(ctx context.Context , id uuid.UUID) error {
	err  := s.q.DeleteLink(ctx , id)
	if err != nil {
		return  ErrNotFound
	}
	return nil
}
