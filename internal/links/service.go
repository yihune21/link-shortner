package links

import (
	"context"

	"github.com/google/uuid"
	"github.com/yihune21/link-shortner/internal/database"
)


type Service interface {
	ListLinks(ctx context.Context, id uuid.UUID) ([]database.Link, error)
	CreateLink(cxt context.Context , params database.CreateLinkParams) (database.Link , error)
}


type svc struct {
	db *database.Queries
}

func NewService(repo *database.Queries) Service {
	return &svc{
		db: repo,
	}
}


func (s *svc) ListLinks(ctx context.Context, id uuid.UUID) ([]database.Link, error) {
	
	links, err := s.db.ListLinksById(ctx, id)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (s *svc)CreateLink(ctx context.Context , params database.CreateLinkParams) (database.Link, error) {
	link , err := s.db.CreateLink(ctx , params)
	if err != nil {
		return link, err
	}

	return link, nil
}