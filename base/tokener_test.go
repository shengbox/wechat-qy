package base

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type mockTokenFetcher struct {
	count int32
}

func (f *mockTokenFetcher) FetchToken() (string, int64, error) {
	atomic.AddInt32(&f.count, 1)
	time.Sleep(10 * time.Millisecond) // Simulate network delay
	return "mock-token", 2, nil      // Expires in 2 seconds
}

func TestTokener_Concurrency(t *testing.T) {
	fetcher := &mockTokenFetcher{}
	tokener := NewTokener(fetcher)

	var wg sync.WaitGroup
	const concurrentRequests = 50

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			token, err := tokener.Token()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if token != "mock-token" {
				t.Errorf("Unexpected token value: %s", token)
			}
		}()
	}

	wg.Wait()

	// Under high concurrency, due to double-checked locking,
	// FetchToken should only be called once.
	calls := atomic.LoadInt32(&fetcher.count)
	if calls != 1 {
		t.Errorf("Expected FetchToken to be called exactly 1 time, got %d", calls)
	}
}

type errTokenFetcher struct{}

func (f *errTokenFetcher) FetchToken() (string, int64, error) {
	return "", 0, errors.New("network error")
}

func TestTokener_FetchError(t *testing.T) {
	tokener := NewTokener(&errTokenFetcher{})
	_, err := tokener.Token()
	if err == nil {
		t.Error("Expected error from Tokener.Token(), got nil")
	}
}
