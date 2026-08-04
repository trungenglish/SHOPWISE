package accessories

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrSyncInProgress = errors.New("Phong Vu sync is already running")

type CatalogSource struct {
	URL      string
	Category string
}

type SyncedProduct struct {
	RetailerProductID string
	Name              string
	Brand             string
	Category          string
	ImageURL          string
	SourceURL         string
	Price             int64
	OriginalPrice     int64
	InStock           bool
}

type CatalogFetcher interface {
	Fetch(context.Context, CatalogSource) ([]SyncedProduct, error)
}

type SyncRepository interface {
	UpsertPhongVuCatalog(context.Context, []SyncedProduct, time.Time) error
	MarkPhongVuSyncFailure(context.Context, time.Time, error) error
}

type Syncer struct {
	fetcher    CatalogFetcher
	repository SyncRepository
	sources    []CatalogSource
	now        func() time.Time
	mutex      sync.Mutex
	running    bool
}

func NewSyncer(fetcher CatalogFetcher, repository SyncRepository, sources []CatalogSource, now func() time.Time) *Syncer {
	return &Syncer{fetcher: fetcher, repository: repository, sources: sources, now: now}
}

func (syncer *Syncer) Sync(ctx context.Context) error {
	// ponytail: process-local lock; use a Redis lock if multiple worker replicas are deployed.
	syncer.mutex.Lock()
	if syncer.running {
		syncer.mutex.Unlock()
		return ErrSyncInProgress
	}
	syncer.running = true
	syncer.mutex.Unlock()
	defer func() {
		syncer.mutex.Lock()
		syncer.running = false
		syncer.mutex.Unlock()
	}()

	runAt := syncer.now().UTC()
	products := make([]SyncedProduct, 0)
	for _, source := range syncer.sources {
		items, err := syncer.fetcher.Fetch(ctx, source)
		if err != nil {
			wrapped := fmt.Errorf("sync %s: %w", source.URL, err)
			_ = syncer.repository.MarkPhongVuSyncFailure(ctx, runAt, wrapped)
			return wrapped
		}
		products = append(products, items...)
	}
	if err := syncer.repository.UpsertPhongVuCatalog(ctx, products, runAt); err != nil {
		_ = syncer.repository.MarkPhongVuSyncFailure(ctx, runAt, err)
		return fmt.Errorf("upsert Phong Vu catalog: %w", err)
	}
	return nil
}
