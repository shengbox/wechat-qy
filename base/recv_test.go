package base

import (
	"sync"
	"testing"
)

func TestGenerateNonce_Uniqueness(t *testing.T) {
	const count = 1000
	nonces := make(map[string]bool)
	var mu sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nonce := GenerateNonce()

			mu.Lock()
			if _, exists := nonces[nonce]; exists {
				t.Errorf("Duplicate nonce detected: %s", nonce)
			}
			nonces[nonce] = true
			mu.Unlock()
		}()
	}
	wg.Wait()
}
