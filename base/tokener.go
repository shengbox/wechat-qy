package base

import (
	"sync"
	"time"
)

// TokenFetcher 包含向 API 服务器获取令牌信息的操作
type TokenFetcher interface {
	FetchToken() (token string, expiresIn int64, err error)
}

// Tokener 用于管理应用套件或企业号的令牌信息
type Tokener struct {
	mu           sync.RWMutex
	token        string
	expiresIn    int64
	tokenFetcher TokenFetcher
}

// NewTokener 方法用于创建 Tokener 实例
func NewTokener(tokenFetcher TokenFetcher) *Tokener {
	return &Tokener{tokenFetcher: tokenFetcher}
}

// Token 方法用于获取应用套件令牌
func (t *Tokener) Token() (token string, err error) {
	t.mu.RLock()
	if t.isValidToken() {
		tok := t.token
		t.mu.RUnlock()
		return tok, nil
	}
	t.mu.RUnlock()

	t.mu.Lock()
	defer t.mu.Unlock()

	// Double-check after acquiring write lock
	if t.isValidToken() {
		return t.token, nil
	}

	if err = t.refreshTokenLocked(); err != nil {
		return "", err
	}

	return t.token, nil
}

// RefreshToken 方法用于刷新令牌信息
func (t *Tokener) RefreshToken() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.refreshTokenLocked()
}

func (t *Tokener) refreshTokenLocked() error {
	token, expiresIn, err := t.tokenFetcher.FetchToken()
	if err != nil {
		return err
	}

	expiresIn = time.Now().Add(time.Second * time.Duration(expiresIn)).Unix()

	t.token = token
	t.expiresIn = expiresIn

	return nil
}

func (t *Tokener) isValidToken() bool {
	now := time.Now().Unix()

	if now >= t.expiresIn || t.token == "" {
		return false
	}

	return true
}

type Ticket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

