package session

import (
	"sync"
	"time"
)

type SessionObj struct {
	IsPrivate bool
	Key       string
	Value     string
	Expired   int64
}

type Session struct {
	sync.RWMutex
	sessions map[string]SessionObj
}

func New() *Session {
	return &Session{
		sessions: make(map[string]SessionObj),
	}
}

func (s *Session) Get(key string) (SessionObj, bool) {
	s.Lock()
	defer s.Unlock()
	sess, ok := s.sessions[key]
	return sess, ok
}

func (s *Session) Set(key, value string, isPrivate bool, exp int64) {
	s.Lock()
	defer s.Unlock()

	s.sessions[key] = SessionObj{
		Key:       key,
		Value:     value,
		IsPrivate: isPrivate,
		Expired:   exp,
	}
}

func (s *Session) IsKeyPrivate(key string) (SessionObj, bool) {
	s.Lock()
	defer s.Unlock()
	sess, ok := s.sessions[key]
	if !ok {
		return sess, false
	}

	return sess, sess.IsPrivate
}

func (s *Session) IsExpired(sess SessionObj) bool {
	return (sess.Expired - time.Now().Unix()) <= 0
}
