package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

type ResPack struct {
	r   *http.Response
	err error
}

func work(ctx context.Context) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	defer wg.Done()
	c := make(chan ResPack, 1)

	req, _ := http.NewRequest("GET", "http://localhost:9200", nil)
	go func() {
		resp, err := client.Do(req)
		pack := ResPack{r: resp, err: err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		log.Println("Timeout")
	case res := <-c:
		if res.err != nil {
			log.Println("res.err=", res.err)
			return
		}
		defer res.r.Body.Close()

		out, _ := ioutil.ReadAll(res.r.Body)
		log.Println("Server Response:", string(out))
	}
	return
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg.Add(1)
	go work(ctx)
	wg.Wait()
	log.Println("Finished")
}
