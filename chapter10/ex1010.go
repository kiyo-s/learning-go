package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PressureGuage struct {
	ch chan struct{}
}

func New(limit int) *PressureGuage {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGuage{
		ch: ch,
	}
}

func (pg *PressureGuage) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("キャパシティに余裕がありません")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	fmt.Println("Open: 'http://localhost:8080/request'")
	http.ListenAndServe(":8080", nil)
}
