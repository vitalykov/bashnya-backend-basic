package main

import (
	"fmt"
	"slices"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	chans := make([]chan any, 0, len(cmds)+1)
	for range cap(chans) {
		chans = append(chans, make(chan any))
	}
	var wg sync.WaitGroup
	for i, c := range cmds {
		wg.Go(func() {
			defer close(chans[i+1])
			c(chans[i], chans[i+1])
		})
	}
	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	var emails sync.Map
	var wg sync.WaitGroup
	for email := range in {
		wg.Go(func() {
			u := GetUser(email.(string))
			if _, ok := emails.LoadOrStore(u.Email, struct{}{}); !ok {
				out <- u
			}
		})
	}
	wg.Wait()
}

func findMessages(wg *sync.WaitGroup, out chan any, users []User) {
	wg.Go(func() {
		msgs, _ := GetMessages(users...)
		for _, m := range msgs {
			out <- m
		}
	})
}

func SelectMessages(in, out chan interface{}) {
	batch := make([]User, GetMessagesMaxUsersBatch)
	i := 0
	var wg sync.WaitGroup
	for u := range in {
		batch[i] = u.(User)
		i = (i + 1) % GetMessagesMaxUsersBatch
		if i == 0 {
			batchCopy := make([]User, GetMessagesMaxUsersBatch)
			copy(batchCopy, batch)
			findMessages(&wg, out, batchCopy)
		}
	}
	if i > 0 {
		batchCopy := make([]User, GetMessagesMaxUsersBatch)
		copy(batchCopy, batch)
		findMessages(&wg, out, batchCopy[:i])
	}
	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, HasSpamMaxAsyncRequests)
	for mId := range in {
		sem <- struct{}{}
		wg.Go(func() {
			defer func() { <-sem }()
			spam, _ := HasSpam(mId.(MsgID))
			out <- MsgData{ID: mId.(MsgID), HasSpam: spam}
		})
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	msgs := []MsgData{}
	for mData := range in {
		msgs = append(msgs, mData.(MsgData))
	}
	slices.SortFunc(msgs, func(a, b MsgData) int {
		if a.HasSpam == b.HasSpam {
			if a.ID < b.ID {
				return -1
			}
			return 1
		}
		if a.HasSpam && !b.HasSpam {
			return -1
		}
		return 1
	})
	for _, mData := range msgs {
		out <- fmt.Sprintf("%v %d", mData.HasSpam, mData.ID)
	}
}
