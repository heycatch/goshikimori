package main

import (
	"fmt"

	g "github.com/heycatch/goshikimori"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func main() {
	c := config()

	// Search anime topic.
	fast_anime, status, err := c.FastIdAnime("initial d first stage")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	a := &g.Options{Page: 1, Limit: 10}
	topic_anime, status, err := fast_anime.SearchTopicsAnime(a)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range topic_anime {
		fmt.Println(v.Id, v.HTMLBody, v.Comments_count, v.Last_comment_viewed)
	}

	// Search manga topic.
	fast_manga, status, err := c.FastIdManga("naruto")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	m := &g.Options{Page: 1, Limit: 10}
	topic_manga, status, err := fast_manga.SearchTopicsManga(m)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range topic_manga {
		fmt.Println(v.Id, v.HTMLBody, v.Comments_count, v.Last_comment_viewed)
	}

	// Search ranobe topic.
	fast_ranobe, status, err := c.FastIdRanobe("sword art")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	r := &g.Options{Page: 1, Limit: 10}
	topic_ranobe, status, err := fast_ranobe.SearchTopicsRanobe(r)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range topic_ranobe {
		fmt.Println(v.Id, v.HTMLBody, v.Comments_count, v.Last_comment_viewed)
	}

	// Search topics.
	o := &g.Options{Page: 1, Limit: 1, Forum: g.TOPIC_FORUM_ANIMANGA}
	t, status, err := c.SearchTopics(o)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range t {
		fmt.Println(v.Body, v.Comments_count, v.Created_at, v.HTMLBody)
	}

	// Search topics updates.
	ou := &g.Options{Page: 1, Limit: 5}
	tu, status, err := c.SearchTopicsUpdates(ou)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range tu {
		fmt.Println(v.Linked.Id, v.Linked.Name, v.Linked.Russian, v.Linked.Url)
	}

	// Search topics hot.
	oh := &g.Options{Limit: 5}
	th, status, err := c.SearchTopicsHot(oh)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range th {
		fmt.Println(v.Id, v.Topic_title, v.Body, v.Created_at, v.Comments_count)
	}

	// Search topic id.
	ti, status, err := c.SearchTopicsId(368370)
	if err != nil {
		fmt.Println(status, err)
		return
	}
	if status == 200 && ti.Id != 0 {
		fmt.Println(ti.Id, ti.Created_at, ti.Comments_count, ti.HTMLBody)
	}

	// Ignore/Unignore topic.
	ignore, status, err := c.AddIgnoreTopic(368370)
	//ignore, status, err := c.RemoveIgnoreTopic(368370)
	if err != nil {
		fmt.Println(err)
		return
	}
	if status == 200 {
		fmt.Println(ignore.Is_ignored, ignore.Topic_id)
	}
}
