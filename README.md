## About
A small library for interacting with shikimori, written in golang.
The library allows you to search the shikimori database.
Work with API occurs only through OAuth2.

## Install
```
go get github.com/vexilology/goshikimori
```

## Examples
``` golang
package main

import (
  "fmt"

  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

// Found User
func main() {
  c := conf()
  s := c.SearchAnime("Initial D")
  fmt.Println(s.Name, s.Status, s.Score)
}
```
``` golang
package main

import (
  "fmt"

  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

// Found Achievements
func main() {
  c := conf()
  u := c.SearchUser("incarnati0n")
  r := c.SearchAchievement(u.Id)
  for _, v := range r {
    if v.Neko_id == g.NekoSearch("Initial D") {
      fmt.Printf("level: %d - progress %d\n", v.Level, v.Progress)
      fmt.Printf("created: %v - updated: %v\n", v.Created_at, v.Updated_at)
    }
  }
}
```

## Available functions
```golang
SearchUser(string) // Find users. Check User request.
SearchAnime(string) // Find animes. Check Anime request.
SearchManga(string) // Find mangas. Check Manga request.
SearchRanobe(string) // Find ranobes. Check Manga request.
SearchClub(string) // Find clubs. Check Club request.
SearchAchievement(int) // Find achievements with anime ID. Check Achievements request.
NekoSearch(string) // Search by anime name in achievements.
SearchSimilarAnime(int) // Find similar anime with anime ID. Check Anime request.
SearchSimilarManga(int) // Find similar manga with manga ID. Check Manga request.
SearchSimilarRanobe(int) // Find similar ranobe with ranobe ID. Check Manga request.
```

## Available API
<details>
  <summary>User request</summary>
    <ul>
      <li>Id</li>
      <li>Nickname</li>
      <li>Avatar</li>
      <li>
        <details>
          <summary>Image</summary>
            <ul>
              <li>Image.X160</li>
              <li>Image.X148</li>
              <li>Image.X80</li>
              <li>Image.X64</li>
              <li>Image.X48</li>
              <li>Image.X32</li>
              <li>Image.X16</li>
            </ul>
        </details>
      </li>
      <li>Online</li>
      <li>Name</li>
      <li>Sex</li>
      <li>Full_Years</li>
      <li>Last_Online</li>
      <li>Website</li>
      <li>Location</li>
      <li>Banned</li>
      <li>About</li>
      <li>AboutHTML</li>
      <li>Common_Info</li>
      <li>Show_Comments</li>
      <li>In_Friends</li>
      <li>Is_Ignored</li>
      <li>Style_Id</li>
    </ul>
</details>

<details>
  <summary>Anime request</summary>
    <ul>
      <li>Id</li>
      <li>Name</li>
      <li>Russian</li>
      <li>
        <details>
          <summary>Image</summary>
            <ul>
              <li>Image.Original</li>
              <li>Image.Preview</li>
              <li>Image.X96</li>
              <li>Image.X48</li>
            </ul>
        </details>
      </li>
      <li>Url</li>
      <li>Kind</li>
      <li>Score</li>
      <li>Status</li>
      <li>Episodes</li>
      <li>Episodes_aired</li>
      <li>Aired_on</li>
      <li>Released_on</li>
    </ul>
</details>

<details>
  <summary>Manga/Ranobe request</summary>
    <ul>
      <li>Id</li>
      <li>Name</li>
      <li>Russian</li>
      <li>
        <details>
          <summary>Image</summary>
            <ul>
              <li>Image.Original</li>
              <li>Image.Preview</li>
              <li>Image.X96</li>
              <li>Image.X48</li>
            </ul>
        </details>
      </li>
      <li>Url</li>
      <li>Kind</li>
      <li>Score</li>
      <li>Status</li>
      <li>Volumes</li>
      <li>Chapters</li>
      <li>Aired_on</li>
      <li>Released_on</li>
    </ul>
</details>

<details>
  <summary>Club request</summary>
    <ul>
      <li>Id</li>
      <li>Name</li>
      <li>
        <details>
          <summary>Logo</summary>
            <ul>
              <li>Logo.Original</li>
              <li>Logo.Main</li>
              <li>Logo.X96</li>
              <li>Logo.X73</li>
              <li>Logo.X48</li>
            </ul>
        </details>
      </li>
      <li>Is_censored</li>
      <li>Join_policy</li>
      <li>Comment_policy</li>
    </ul>
</details>

<details>
  <summary>Achievement request</summary>
    <ul>
      <li>Id</li>
      <li>Neko_id</li>
      <li>Level</li>
      <li>Progress</li>
      <li>User_id</li>
      <li>Created_at</li>
      <li>Updated_at</li>
    </ul>
</details>

## Shikimori documentation
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)
