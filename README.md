## About
A small library for interacting with shikimori, written in golang.
The library allows you to search the shikimori database.
Work with API occurs only through OAuth2.

## Install
```
go get -u github.com/vexilology/goshikimori
```

## Examples
* [Click her](https://github.com/vexilology/goshikimori/tree/main/examples)

## Available functions
```golang
SearchUser(string) // Find users. Check Users request.
SearchAnime(string, interface) // Find animes. Check Animes request.
SearchManga(string, interface) // Find mangas. Check Mangas request.
SearchClub(string, interface) // Find clubs. Check Clubs request.
FastIdAnime(string) // Anime id search.
FastIdManga(string) // Manga id search.
NekoSearch(string) // Search by anime name in achievements.
SearchAchievement(int) // Find achievements with anime id. Check Achievements request.
SearchAnimeScreenshots(int) // Find anime-screenshots with anime id. Check AnimeScreenshots request.
SearchSimilarAnime(int) // Find similar anime with anime id. Check Animes request.
SearchSimilarManga(int) // Find similar manga with manga id. Check Mangas request.
SearchRelatedAnime(int) // Find related anime with anime id. Check RelatedAnimes request.
SearchRelatedManga(int) // Find related manga with manga id. Check RelatedMangas request.
SearchAnimeVideos(int) // Find anime-videos with anime id. Check AnimeVideos request.
SearchAnimeRoles(int) // Find anime roles with anime id. Check AnimeRoles/MangaRoles request.
SearchMangaRoles(int) // Find manga roles with anime id. Check AnimeRoles/MangaRoles request.
SearchBans() // Find last bans. Check Bans request.
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
      <li>Full_years</li>
      <li>Last_online</li>
      <li>Website</li>
      <li>Location</li>
      <li>Banned</li>
      <li>About</li>
      <li>AboutHTML</li>
      <li>Common_Info</li>
      <li>Show_comments</li>
      <li>In_friends</li>
      <li>Is_ignored</li>
      <li>Style_id</li>
    </ul>
</details>
<details>
  <summary>Interface_Anime</summary>
    <ul>
      <li>Limit: 50 maximum</li>
      <li>Kind: tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48</li>
      <li>Status: anons, ongoing, released</li>
      <li>Season: summer_2017, 2016, 2014_2016, 199x</li>
      <li>Score: 9 maximum</li>
      <li>Rating: none, g, pg, pg_13, r, r_plus, rx</li>
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
  <summary>Interface_Manga</summary>
    <ul>
      <li>Limit: 50 maximum</li>
      <li>Kind: manga, manhwa, manhua, light_novel, novel, one_shot, doujin</li>
      <li>Status: anons, ongoing, released, paused, discontinued</li>
      <li>Season: summer_2017, "spring_2016,fall_2016", "2016,!winter_2016", 2016, 2014_2016, 199x</li>
      <li>Score: 9 maximum</li>
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
  <summary>Interface_Club</summary>
    <ul>
      <li>Limit: 30 maximum</li>
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
            <li>Url</li>
          </ul>
        </details>
      </li>
    </ul>
</details>
<details>
  <summary>Bans request</summary>
    <ul>
      <li>Id</li>
      <li>User_id</li>
      <li>
        <details>
          <summary>Comment</summary>
          <ul>
            <li>Id</li>
            <li>Commentable_id</li>
            <li>Commentable_type</li>
            <li>Body</li>
            <li>User_id</li>
            <li>Created_at</li>
            <li>Updated_at</li>
            <li>Is_offtopic</li>
          </ul>
        </details>
      </li>
      <li>Moderator_id</li>
      <li>Reason</li>
      <li>Created_at</li>
      <li>Duration_minutes</li>
      <li>
        <details>
          <summary>User</summary>
          <ul>
            <li>Id</li>
            <li>Nickname</li>
            <li>Avatar</li>
            <li>
              <details>
                <summary>Image</summary>
                <ul>
                  <li>X160</li>
                  <li>X148</li>
                  <li>X80</li>
                  <li>X64</li>
                  <li>X48</li>
                  <li>X32</li>
                  <li>X16</li>
                </ul>
              </details>
            </li>
          </ul>
        </details>
      </li>
      <li>Last_online_at</li>
    </ul>
</details>

## Shikimori documentation
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)
