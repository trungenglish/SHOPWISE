package accessories

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type catalogFetcherStub struct {
	products []SyncedProduct
	err      error
	calls    atomic.Int32
	release  chan struct{}
}

func (stub *catalogFetcherStub) Fetch(context.Context, CatalogSource) ([]SyncedProduct, error) {
	stub.calls.Add(1)
	if stub.release != nil {
		<-stub.release
	}
	return stub.products, stub.err
}

type syncRepositoryStub struct {
	mu       sync.Mutex
	upserts  int
	failures int
}

func (stub *syncRepositoryStub) UpsertPhongVuCatalog(context.Context, []SyncedProduct, time.Time) error {
	stub.mu.Lock()
	defer stub.mu.Unlock()
	stub.upserts++
	return nil
}

func (stub *syncRepositoryStub) MarkPhongVuSyncFailure(context.Context, time.Time, error) error {
	stub.mu.Lock()
	defer stub.mu.Unlock()
	stub.failures++
	return nil
}

func TestSyncFailureKeepsExistingSnapshot(t *testing.T) {
	t.Parallel()

	repository := &syncRepositoryStub{}
	syncer := NewSyncer(&catalogFetcherStub{err: errors.New("source unavailable")}, repository, []CatalogSource{{URL: "https://phongvu.vn/c/chuot-co-day", Category: "mouse"}}, time.Now)
	if err := syncer.Sync(context.Background()); err == nil {
		t.Fatal("Sync() error = nil, want source error")
	}
	if repository.upserts != 0 || repository.failures != 1 {
		t.Fatalf("upserts/failures = %d/%d, want 0/1", repository.upserts, repository.failures)
	}
}

func TestSyncPreventsOverlappingRuns(t *testing.T) {
	t.Parallel()

	release := make(chan struct{})
	fetcher := &catalogFetcherStub{products: []SyncedProduct{{RetailerProductID: "PV-1"}}, release: release}
	repository := &syncRepositoryStub{}
	syncer := NewSyncer(fetcher, repository, []CatalogSource{{URL: "https://phongvu.vn/c/chuot-co-day", Category: "mouse"}}, time.Now)
	firstDone := make(chan error, 1)
	go func() { firstDone <- syncer.Sync(context.Background()) }()
	for fetcher.calls.Load() == 0 {
		time.Sleep(time.Millisecond)
	}
	if err := syncer.Sync(context.Background()); !errors.Is(err, ErrSyncInProgress) {
		t.Fatalf("overlapping Sync() error = %v", err)
	}
	close(release)
	if err := <-firstDone; err != nil {
		t.Fatalf("first Sync() error = %v", err)
	}
}
