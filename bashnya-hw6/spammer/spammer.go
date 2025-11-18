package main

import (
	"fmt"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	chans := make([]chan any, 0, len(cmds)+1)
	for range cap(chans) {
		chans = append(chans, make(chan any))
	}
	var wg sync.WaitGroup
	for i, command := range cmds {
		wg.Go(func() {
			defer close(chans[i+1])
			command(chans[i], chans[i+1])
		})
	}
	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	var emails sync.Map
	var wg sync.WaitGroup
	for email := range in {
		wg.Go(func() {
			user := GetUser(email.(string))
			if _, ok := emails.LoadOrStore(user.Email, struct{}{}); !ok {
				out <- user
			}
		})
	}
	wg.Wait()
}

func findMessages(wg *sync.WaitGroup, out chan any, users []User) {
	wg.Go(func() {
		messages, _ := GetMessages(users...)
		for _, msg := range messages {
			out <- msg
		}
	})
}

func SelectMessages(in, out chan interface{}) {
	batch := make([]User, 0, GetMessagesMaxUsersBatch)
	var wg sync.WaitGroup
	for user := range in {
		batch = append(batch, user.(User))
		if len(batch) == GetMessagesMaxUsersBatch {
			findMessages(&wg, out, batch)
			batch = make([]User, 0, GetMessagesMaxUsersBatch)
		}
	}
	findMessages(&wg, out, batch)
	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, HasSpamMaxAsyncRequests)
	for msgId := range in {
		sem <- struct{}{}
		wg.Go(func() {
			defer func() { <-sem }()
			spam, _ := HasSpam(msgId.(MsgID))
			out <- MsgData{ID: msgId.(MsgID), HasSpam: spam}
		})
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	messages := []MsgData{}
	for msgData := range in {
		messages = append(messages, msgData.(MsgData))
	}
	sort.Slice(messages, func(i, j int) bool {
		a, b := messages[i], messages[j]
		if a.HasSpam == b.HasSpam {
			return a.ID < b.ID
		}
		return a.HasSpam
	})
	for _, msgData := range messages {
		out <- fmt.Sprintf("%v %d", msgData.HasSpam, msgData.ID)
	}
}
