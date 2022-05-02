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

func main() {
  c := conf()
  s := c.SearchAnime("Initial D")
  fmt.Println(s.Name, s.Status, s.Score)
}
```
``` golang
package main

import (
  "os"
  "fmt"

  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  u := c.SearchUser("incarnati0n")
  r := c.SearchAchievement(u.Id)
  for _, v := range r {
    if v.Neko_id == g.NekoSearch("Initial D") {
      fmt.Printf("level: %d - progress %d\n", v.Level, v.Progress)
      fmt.Printf("created: %v - updated: %v\n", v.Created_at, v.Updated_at)
    } else {
      fmt.Println("Achievement not found")
      os.Exit(1)
    }
  }
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

func main() {
  c := conf()
  a := c.SearchAnime("Initial D")
  r := c.SearchRelatedAnime(a.Id)
  fmt.Println(r.Relation, r.Relation_Russian, r.Anime.Score)
}
```

## Available functions
```golang
SearchUser(string) // Find users. Check Users request.
SearchAnime(string) // Find animes. Check Animes request.
SearchManga(string) // Find mangas. Check Mangas request.
SearchClub(string) // Find clubs. Check Clubs request.
SearchAchievement(int) // Find achievements with anime ID. Check Achievements request.
NekoSearch(string) // Search by anime name in achievements.
SearchSimilarAnime(int) // Find similar anime with anime ID. Check Animes request.
SearchSimilarManga(int) // Find similar manga with manga ID. Check Mangas request.
SearchRelatedAnime(int) // Find related anime with anime ID. Check RelatedAnimes request.
SearchRelatedManga(int) // Find related manga with manga ID. Check RelatedMangas request.
SearchAnimeScreenshots(int) // Find anime-screenshots with anime ID. Check AnimeScreenshots request.
SearchAnimeVideos(int) // Find anime-videos with anime ID. Check AnimeVideos request.
SearchAnimeRoles(int) // Find anime roles with anime ID. Check AnimeRoles/MangaRoles request.
SearchMangaRoles(int) // Find manga roles with anime ID. Check AnimeRoles/MangaRoles request.
```

## Available API
<details>
  <summary>Users request</summary>
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
  <summary>Animes request</summary>
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
  <summary>Mangas request</summary>
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
  <summary>Clubs request</summary>
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
  <summary>Achievements request</summary>
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

<details>
  <summary>RelatedAnimes request</summary>
    <ul>
      <li>Relation</li>
      <li>Relation_Russian</li>
      <li>
        <details>
          <summary>Anime</summary>
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
      </li>
      <li>Manga</li>
    </ul>
</details>

<details>
  <summary>RelatedMangas request</summary>
    <ul>
      <li>Relation</li>
      <li>Relation_Russian</li>
      <li>
        <details>
          <summary>Manga</summary>
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
      </li>
      <li>Anime</li>
    </ul>
</details>

<details>
  <summary>AnimeScreenshots request</summary>
    <ul>
      <li>Original</li>
      <li>Preview</li>
    </ul>
</details>

<details>
  <summary>AnimeVideos request</summary>
    <ul>
      <li>Id</li>
      <li>Url</li>
      <li>Image_url</li>
      <li>Player_url</li>
      <li>Name</li>
      <li>Kind</li>
      <li>Hosting</li>
    </ul>
</details>

<details>
  <summary>AnimeRoles/MangaRoles request</summary>
    <ul>
      <li>Roles</li>
      <li>Roles_Russian</li>
      <li>
        <details>
          <summary>Character</summary>
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
          </ul>
          <li>Url</li>
        </details>
      </li>
    </ul>
</details>

## Shikimori documentation
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)
