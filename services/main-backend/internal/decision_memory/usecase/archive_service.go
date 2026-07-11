package usecase

import (
	"context"
	"log"
	"time"
)

type ArchiveService struct {
	repo Repository
}

func NewArchiveService(repo Repository) *ArchiveService {
	return &ArchiveService{repo: repo}
}

// RunArchiveJob should be triggered by a scheduler (e.g. cron).
// It archives all sessions that haven't been updated in 90 days.
func (a *ArchiveService) RunArchiveJob(ctx context.Context) error {
	ninetyDaysAgo := time.Now().UTC().AddDate(0, 0, -90)
	
	count, err := a.repo.ArchiveInactiveSessions(ctx, ninetyDaysAgo)
	if err != nil {
		log.Printf("ArchiveJob error: %v", err)
		return err
	}
	
	log.Printf("ArchiveJob completed successfully. Archived %d sessions.", count)
	return nil
}
