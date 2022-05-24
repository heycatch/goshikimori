package api

import "time"

type UserSize struct {
  X160 string `json:"x160"`
  X148 string `json:"x148"`
  X80  string `json:"x80"`
  X64  string `json:"x64"`
  X48  string `json:"x48"`
  X32  string `json:"x32"`
  X16  string `json:"x16"`
}

type Users struct {
  Id            int      `json:"id"`
  Nickname      string   `json:"nickname"`
  Avatar        string   `json:"avatar"`
  Image         UserSize `json:"image"`
  Online        string   `json:"last_online_at"`
  Name          string   `json:"name"`
  Sex           string   `json:"sex"`
  Full_years    int      `json:"full_years"`
  Last_online   string   `json:"last_online"`
  Website       string   `json:"website"`
  Location      string   `json:"location"`
  Banned        bool     `json:"banned"`
  About         string   `json:"about"`
  AboutHTML     string   `json:"about_html"`
  Common_info   []string `json:"common_info"`
  Show_comments bool     `json:"show_comments"`
  In_friends    bool     `json:"in_friends"`
  Is_ignored    bool     `json:"is_ignored"`
  Style_id      int      `json:"style_id"`
}

type AnimeSize struct {
  Original string `json:"original"`
  Preview  string `json:"preview"`
  X96      string `json:"x96"`
  X48      string `json:"x48"`
}

type Animes struct {
  Id             int       `json:"id"`
  Name           string    `json:"name"`
  Russian        string    `json:"russian"`
  Image          AnimeSize `json:"image"`
  Url            string    `json:"url"`
  Kind           string    `json:"kind"`
  Score          string    `json:"score"`
  Status         string    `json:"status"`
  Episodes       int       `json:"episodes"`
  Episodes_aired int       `json:"episodes_aired"`
  Aired_on       string    `json:"aired_on"`
  Released_on    string    `json:"released_on"`
}

type Mangas struct {
  Id          int       `json:"id"`
  Name        string    `json:"name"`
  Russian     string    `json:"russian"`
  Image       AnimeSize `json:"image"`
  Url         string    `json:"url"`
  Kind        string    `json:"kind"`
  Score       string    `json:"score"`
  Status      string    `json:"status"`
  Volumes     int       `json:"volumes"`
  Chapters    int       `json:"chapters"`
  Aired_on    string    `json:"aired_on"`
  Released_on string    `json:"released_on"`
}

type ClubSize struct {
  Original string `json:"original"`
  Main     string `json:"main"`
  X96      string `json:"x96"`
  X73      string `json:"x73"`
  X48      string `json:"x48"`
}

type Clubs struct {
  Id             int      `json:"id"`
  Name           string   `json:"name"`
  Logo           ClubSize `json:"logo"`
  Is_censored    bool     `json:"is_censored"`
  Join_policy    string   `json:"join_policy"`
  Comment_policy string   `json:"comment_policy"`
}

type Achievements struct {
  Id         int       `json:"id"`
  Neko_id    string    `json:"neko_id"`
  Level      int       `json:"level"`
  Progress   int       `json:"progress"`
  User_id    int       `json:"user_id"`
  Created_at time.Time `json:"created_at"`
  Updated_at time.Time `json:"updated_at"`
}

// NOTES: removed Manga -> bool
// json: cannot unmarshal manga now
type RelatedAnimes struct {
  Relation         string `json:"relation"`
  Relation_Russian string `json:"relation_russian"`
  Anime            Animes `json:"anime"`
}

// NOTES: removed Anime -> bool
type RelatedMangas struct {
  Relation         string `json:"relation"`
  Relation_Russian string `json:"relation_russian"`
  Manga            Mangas `json:"manga"`
}

type AnimeScreenshots struct {
  Original string `json:"original"`
  Preview  string `json:"preview"`
}

type AnimeVideos struct {
  Id         int    `json:"id"`
  Url        string `json:"url"`
  Image_url  string `json:"image_url"`
  Player_url string `json:"player_url"`
  Name       string `json:"name"`
  Kind       string `json:"kind"`
  Hosting    string `json:"hosting"`
}

type CharacterInfo struct {
  Id      int       `json:"id"`
  Name    string    `json:"name"`
  Russian string    `json:"russian"`
  Image   AnimeSize `json:"image"`
  Url     string    `json:"url"`
}

// NOTES: removed Person -> nil
type Roles struct {
  Roles         []string      `json:"roles"`
  Roles_Russian []string      `json:"roles_russian"`
  Character     CharacterInfo `json:"character"`
}

type CommentInfo struct {
  Id               int       `json:"id"`
  Commentable_id   int       `json:"commentable_id"`
  Commentable_type string    `json:"commentable_type"`
  Body             string    `json:"body"`
  User_id          int       `json:"user_id"`
  Created_at       time.Time `json:"created_at"`
  Updated_at       time.Time `json:"updated_at"`
  Is_offtopic      bool      `json:"is_offtopic"`
}

type UserInfo struct {
  Id       int      `json:"id"`
  Nickname string   `json:"nickname"`
  Avatar   string   `json:"avatar"`
  Image    UserSize `json:"image"`
}

type Bans struct {
  Id               int         `json:"id"`
  User_id          int         `json:"user_id"`
  Comment          CommentInfo `json:"comment"`
  Moderator_id     int         `json:"moderator_id"`
  Reason           string      `json:"reason"`
  Created_at       time.Time   `json:"created_at"`
  Duration_minutes int         `json:"duration_minutes"`
  User             UserInfo    `json:"user"`
  Last_online_at   time.Time   `json:"last_online_at"`
}

type Calendar struct {
  Next_episode    int       `json:"next_episode"`
  Next_episode_at time.Time `json:"next_episode_at"`
  Duration        int       `json:"duration"`
  Anime           Animes    `json:"anime"`
}
