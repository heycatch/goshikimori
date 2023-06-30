package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"sync"

	//g "github.com/vexilology/goshikimori"
)

type Session struct {
	App, Token string
}

type Cache struct {
	items map[string]*item
	mu    sync.Mutex
}

type item struct {
	value 	string
	expires int64
}

func New() *Cache {
	c := &Cache{items: make(map[string]*item)}

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				c.mu.Lock()
				for k, v := range c.items {
					if v.Expired(time.Now().UnixNano()) {
						log.Printf("%v has expires at %d", c.items, time.Now().UnixNano())
						delete(c.items, k)
					}
				}
				c.mu.Unlock()
			}
		}
	}()

	return c
}

func (i *item) Expired(time int64) bool {
	if i.expires == 0 { return true }
	return time > i.expires
}

func (c *Cache) Get(key string) string {
	c.mu.Lock()
	var s string
	if v, ok := c.items[key]; ok {
		s = v.value
	}
	c.mu.Unlock()

	return s
}

func (c *Cache) Put(key, value string, expires int64) {
	c.mu.Lock()
	if _, ok := c.items[key]; !ok {
		c.items[key] = &item{
			value: value,
			expires: expires,
		}
	}
	c.mu.Unlock()
}

func auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/fail", http.StatusBadRequest)
		return
	}

	var s Session

  data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  if err := json.Unmarshal(data, &s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
  }

	// key(app) + value(token) -> 24h life in memory -> afterwards is deleted
	New().Put(s.App, s.Token, time.Now().Add(time.Second*5).UnixNano())
}

func animes(w http.ResponseWriter, r *http.Request) {
	var c Cache
	log.Println(c.Get("test api"))
}

func mangas(w http.ResponseWriter, r *http.Request) {
	var c Cache
	log.Println(c.Get("test api"))
}

func main() {
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {/* empty function */})

	http.HandleFunc("/api/animes", animes)
	http.HandleFunc("/api/mangas", mangas)

	log.Println("server started...")
	http.ListenAndServe(":1337", nil)
}
