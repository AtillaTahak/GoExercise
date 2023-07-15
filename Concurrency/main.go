package main

import (
	"fmt"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}
func main() {
	start := time.Now()
	userProfile,err := handleGetUserProfile(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(userProfile)
	fmt.Println(time.Since(start))

}

type Response struct {
	data any
	err error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	respch := make(chan Response, 3)
	wg := &sync.WaitGroup{}
	go handleGetUserComments(id, respch,wg)
	go handleGetUserLikes(id, respch,wg)
	go handleGetUserFriends(id, respch,wg)
	wg.Add(3)
	wg.Wait()
	close(respch)
	userProfile := &UserProfile{}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch resp.data.(type) {
		case []string:
			userProfile.Comments = resp.data.([]string)
		case int:
			userProfile.Likes = resp.data.(int)
		case []int:
			userProfile.Friends = resp.data.([]int)
		}	
	}
	return userProfile, nil
}

func handleGetUserComments(id int, respch chan Response,wg *sync.WaitGroup) {
	time.Sleep(200 * time.Millisecond)
	comments := []string{"comment1", "comment2", "comment3"}
	respch <- Response{data: comments, err: nil}
	wg.Done()

}

func handleGetUserLikes(id int,respch chan Response,wg *sync.WaitGroup) {
	time.Sleep(200 * time.Millisecond)

	respch <- Response{data: 33, err: nil}
	wg.Done()
}

func handleGetUserFriends(id int,respch chan Response,wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	respch <- Response{data: []int{11,34,854,455}, err: nil}
	wg.Done()

}
