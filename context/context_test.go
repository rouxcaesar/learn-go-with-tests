package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1) // the second int argument represents the bufferlength of the channel

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("Spy store got cancelled")
				return
			default:
				time.Sleep(100 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestHandler(t *testing.T) {
	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		//store.assertWasNotCancelled()
	})

	//	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
	//		store := &SpyStore{response: data, t: t}
	//		svr := Server(store)
	//
	//		request := httptest.NewRequest(http.MethodGet, "/", nil)
	//
	//		cancellingCtx, cancel := context.WithCancel(request.Context())
	//		time.AfterFunc(5*time.Millisecond, cancel)
	//		request = request.WithContext(cancellingCtx)
	//
	//		response := httptest.NewRecorder()
	//
	//		svr.ServeHTTP(response, request)
	//
	//		store.assertWasCancelled()
	//	})
}

//func (s *SpyStore) assertWasCancelled() {
//	s.t.Helper()
//	if !s.cancelled {
//		s.t.Errorf("store was not told to cancel")
//	}
//}
//
//func (s *SpyStore) assertWasNotCancelled() {
//	s.t.Helper()
//	if s.cancelled {
//		s.t.Errorf("store was told to cancel")
//	}
//}
