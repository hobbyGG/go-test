package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 我们在测试中的存储方式，使用map存储
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string // 用于spy函数调用
}

// 为存储方式创建指定方法
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordinWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := &PlayerServer{&store}

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}
