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

  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "Api Test",
    "9z4yB8Gi_zsfwPAws2SvHj-zeojWOvhxYMauEQYvEqM",
  )
}

// To search for achievements you need ONLY ID
func foundId() int {
  c := conf()
  r := c.SearchUser("incarnati0n")
  return r.Id
}

// Found Achievements
func main() {
  c := conf()
  r := c.SearchAchievement(foundId())
  for _, v := range r {
    if v.Neko_id == g.ConvertNeko("Initial D") {
      fmt.Printf("level: %d - progress %d\n", v.Level, v.Progress)
      fmt.Printf("created: %v - updated: %v\n", v.Created_at, v.Updated_at)
    }
  }
}
```

## Available functions
```golang
SearchUser(string) // found users
SearchAnime(string) // found animes
SearchManga(string) // found mangas
SearchRanobe(string) // found ranobes
SearchClub(string) // found clubs
SearchAchievement(int) // found achievements
ConvertNeko(string) // search by anime name in achievements
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
