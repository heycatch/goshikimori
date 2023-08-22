package main

import (
  "fmt"
  g "github.com/heycatch/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

	// Search topics
  o := &g.Options{Page: "1", Limit: "1", Forum: "", Linked_id: "", Linked_type: ""}
  t, status, err := c.SearchTopics(o)
  if err != nil {
    fmt.Println(err)
    return
  }
	if status == 200 {
		for _, v := range t {
			fmt.Println(v.Body, v.Comments_count, v.Created_at, v.HTMLBody)
		}
	}

	// Search topics updates
	ou := &g.Options{Page: "1", Limit: "5"}
	tu, status_updates, err := c.SearchTopicsUpdates(ou)
	if err != nil {
		fmt.Println(err)
		return
	}
	if status_updates == 200 {
		for _, v := range tu {
			fmt.Println(v.Linked.Id, v.Linked.Name, v.Linked.Russian, v.Linked.Url)
		}
	}

	// Search topics hot
	oh := &g.Options{Limit: "5"}
	th, status_hot, err := c.SearchTopicsHot(oh)
	if err != nil {
		fmt.Println(err)
		return
	}
	if status_hot == 200 {
		for _, v := range th {
			fmt.Println(v.Id, v.Topic_title, v.Body, v.Created_at, v.Comments_count)
		}
	}

	// Search topic id
	ti, status_id, err := c.SearchTopicsId(368370)
	if err != nil {
		fmt.Println(err)
		return
	}
	if status_id == 200 && ti.Id != 0 {
		fmt.Println(ti.Id, ti.Created_at, ti.Comments_count, ti.HTMLBody)
	}

	// Ignore/Unignore topic
	ignore, status_ignore, err := c.AddIgnoreTopic(368370)
	//ignore, status_ignore, err := c.RemoveIgnoreTopic(368370)
	if err != nil {
		fmt.Println(err)
		return
	}
	if status_ignore == 200 { fmt.Println(ignore.Is_ignored, ignore.Topic_id) }
}
