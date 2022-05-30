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
SearchUsers(name string, ExtraLimit interface) // Find users. Check Users request. Many results, search is not case sensitive. Check Users request.
SearchUser(name string) // Find user. Check User request. Single result, search is case sensitive. Check User request.
SearchUserFriends(id int) // Find user friends with user id. Check UserFriends request.
SearchUserClubs(id int) // Find user clubs with user id. Check UserClubs request.
SearchUserAnimeRates(id int, ExtraAnimeRates interface) // Find user anime rates with user id. Check UserAnimeRates request.
SearchUserMangaRates(id int, ExtraMangaRates interface) // Find user manga rates with user id. Check UserMangaRates request.
SearchUserFavourites(id int) // Find user favourites: anime, manga, characters, people,mangakas, seyu and producers with user id. Check UserFavourites request.
SearchUserHistory(id int, ExtraTargetType interface) // Find user history(anime, manga) with user id. Check UserHistory request.
SearchUserBans(id int) // Find user bans with user id. Check Bans request.
WhoAmi() // Verify who am i. Check WhoAmi request.
SearchAnime(name string, Extra interface) // Find animes. Check Animes request.
SearchManga(name string, Extra interface) // Find mangas. Check Mangas request.
SearchClub(name string, ExtraLimit interface) // Find clubs. Check Clubs request.
FastIdAnime(name string) // Anime id search.
FastIdManga(name string) // Manga id search.
NekoSearch(name string) // Search by anime name in achievements.
SearchAchievement(id int) // Find achievements with anime id. Check Achievements request.
SearchAnimeScreenshots(id int) // Find anime-screenshots with anime id. Check AnimeScreenshots request.
SearchSimilarAnime(id int) // Find similar anime with anime id. Check Animes request.
SearchSimilarManga(id int) // Find similar manga with manga id. Check Mangas request.
SearchRelatedAnime(id int) // Find related anime with anime id. Check RelatedAnimes request.
SearchRelatedManga(id int) // Find related manga with manga id. Check RelatedMangas request.
SearchAnimeVideos(id int) // Find anime-videos with anime id. Check AnimeVideos request.
SearchAnimeRoles(id int) // Find anime roles with anime id. Check AnimeRoles/MangaRoles request.
SearchMangaRoles(id int) // Find manga roles with anime id. Check AnimeRoles/MangaRoles request.
SearchBans() // Find last bans. Check Bans request.
SearchCalendar(ExtraCensored interface) // Find calendar. Check Calendar request.
```

## Available API
<details>
  <summary>Interface_Users</summary>
    <ul>
      <li>Limit: 100 maximum</li>
    </ul>
</details>
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
      <li>Last_online_at</li>
      <li>Name</li>
      <li>Sex</li>
      <li>Full_years</li>
      <li>Last_online</li>
      <li>Website</li>
      <li>Location</li>
      <li>Banned</li>
      <li>About</li>
      <li>AboutHTML</li>
      <li>[]Common_Info</li>
      <li>Show_comments</li>
      <li>In_friends</li>
      <li>Is_ignored</li>
      <li>Style_id</li>
    </ul>
</details>
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
      <li>Last_online_at</li>
      <li>Name</li>
      <li>Sex</li>
      <li>Full_years</li>
      <li>Last_online</li>
      <li>Website</li>
      <li>Location</li>
      <li>Banned</li>
      <li>About</li>
      <li>AboutHTML</li>
      <li>[]Common_Info</li>
      <li>Show_comments</li>
      <li>In_friends</li>
      <li>Is_ignored</li>
      <li>
        <details>
          <summary>Stats</summary>
            <ul>
              <details>
                <summary>Statuses</summary>
                  <ul>
                    <details>
                      <summary>[]Anime</summary>
                        <ul>
                          <li>Id</li>
                          <li>Grouped_id</li>
                          <li>Name</li>
                          <li>Size</li>
                          <li>Type</li>
                        </ul>
                    </details>
                    <details>
                      <summary>[]Manga</summary>
                        <ul>
                          <li>Id</li>
                          <li>Grouped_id</li>
                          <li>Name</li>
                          <li>Size</li>
                          <li>Type</li>
                        </ul>
                    </details>
                  </ul>
              </details>
            </ul>
        </details>
      </li>
      <li>Style_id</li>
    </ul>
</details>
<details>
  <summary>UserFriends request</summary>
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
      <li>Last_online_at</li>
    </ul>
</details>
<details>
  <summary>UserClubs request</summary>
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
  <summary>Interface_UserAnimeRates</summary>
    <ul>
      <li>Limit: 5000 maximum</li>
      <li>Status: planned, watching, rewatching, completed, on_hold, dropped</li>
      <li>Censored: true, false</li>
    </ul>
</details>
<details>
  <summary>UserAnimeRates request</summary>
    <ul>
      <li>Id</li>
      <li>Score</li>
      <li>Status</li>
      <li>Text</li>
      <li>Episodes</li>
      <li>Text_html</li>
      <li>Rewatches</li>
      <li>Created_at</li>
      <li>Updated_at</li>
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
            </ul>
        </details>
      </li>
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
  <summary>Interface_UserMangaRates</summary>
    <ul>
      <li>Limit: 5000 maximum</li>
      <li>Censored: true, false</li>
    </ul>
</details>
<details>
  <summary>UserMangaRates request</summary>
    <ul>
      <li>Id</li>
      <li>Score</li>
      <li>Status</li>
      <li>Text</li>
      <li>Chapters</li>
      <li>Volumes</li>
      <li>Text_html</li>
      <li>Rewatches</li>
      <li>Created_at</li>
      <li>Updated_at</li>
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
            </ul>
        </details>
      </li>
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
  <summary>UserFavourites request</summary>
    <ul>
      <li>
        <details>
          <summary>[]Animes</summary>
            <ul>
              <li>Id</li>
              <li>Name</li>
              <li>Russian</li>
              <li>Image</li>
            </ul>
        </details>
      </li>
      <li>
        <details>
          <summary>[]Mangas</summary>
            <ul>
              <li>Id</li>
              <li>Name</li>
              <li>Russian</li>
              <li>Image</li>
            </ul>
        </details>
      </li>
      <li>
        <details>
          <summary>[]Characters</summary>
            <ul>
              <li>Id</li>
              <li>Name</li>
              <li>Russian</li>
              <li>Image</li>
            </ul>
        </details>
      </li>
      <li>
        <details>
          <summary>[]People</summary>
            <ul>
              <li>Id</li>
              <li>Name</li>
              <li>Russian</li>
              <li>Image</li>
            </ul>
        </details>
      </li>
      <li>
        <details>
          <summary>[]Mangakas</summary>
            <ul>
              <li>Id</li>
              <li>Name</li>
              <li>Russian</li>
              <li>Image</li>
            </ul>
        </details>
      </li>
      <li>
        <details>
          <summary>[]Seyu</summary>
            <ul>
              <li>Id</li>
              <li>Name</li>
              <li>Russian</li>
              <li>Image</li>
            </ul>
        </details>
      </li>
      <li>
        <details>
          <summary>[]Producers</summary>
            <ul>
              <li>Id</li>
              <li>Name</li>
              <li>Russian</li>
              <li>Image</li>
            </ul>
        </details>
      </li>
    </ul>
</details>
<details>
  <summary>Interface_UserHistory</summary>
    <ul>
      <li>Limit: 100 maximum</li>
      <li>Target_type: Anime, Manga</li>
    </ul>
</details>
<details>
  <summary>UserHistory request</summary>
    <ul>
      <li>Id</li>
      <li>Created_at</li>
      <li>Description</li>
      <li>
        <details>
          <summary>Target</summary>
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
  <summary>WhoAmi request</summary>
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
      <li>Last_online_at</li>
      <li>Name</li>
      <li>Sex</li>
      <li>Website</li>
      <li>Birth_on</li>
      <li>Locale</li>
    </ul>
</details>
<details>
  <summary>Interface_Animes</summary>
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
  <summary>Interface_Mangas</summary>
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
  <summary>Interface_Clubs</summary>
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
      <li>[]Roles</li>
      <li>[]Roles_Russian</li>
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
            <li>Is_summary</li>
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
            <li>Last_online_at</li>
          </ul>
        </details>
      </li>
      <li>
        <details>
          <summary>Moderator</summary>
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
            <li>Last_online_at</li>
          </ul>
        </details>
      </li>
    </ul>
</details>
<details>
  <summary>Interface_Calendar</summary>
    <ul>
      <li>Censored: true, false</li>
    </ul>
</details>
<details>
  <summary>Calendar request</summary>
    <ul>
      <li>Next_episode</li>
      <li>Next_episode_at</li>
      <li>Duration</li>
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

## Shikimori documentation
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)
