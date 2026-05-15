package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
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

func (s *LinkService)CreateLink(ctx context.Context ,original_link string ,dbUser *database.User)( Link , error)  {
	
	//short link computation todo
	shorten_url :=  original_link


	dbLink, err := s.q.CreateLink(ctx ,database.CreateLinkParams{
		ID: uuid.New(),
		ShortLink:shorten_url ,
		OriginalLink: original_link,
	    UserID:dbUser.ID,
		CreatedAt: time.Now().UTC(),
	 } )

	if err != nil {
		return Link{} , errors.New(err.Error())
	}
	return mapLink(dbLink) , nil
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

func (s *LinkService)GetLinksByUserId(ctx context.Context , id uuid.UUID) ([]Link , error)  {
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