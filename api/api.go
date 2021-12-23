package api

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
  Full_Years    int      `json:"full_years"`
  Last_Online   string   `json:"last_online"`
  Website       string   `json:"website"`
  Location      string   `json:"location"`
  Banned        bool     `json:"banned"`
  About         string   `json:"about"`
  AboutHTML     string   `json:"about_html"`
  Common_Info   []string `json:"common_info"`
  Show_Comments bool     `json:"show_comments"`
  In_Friends    bool     `json:"in_friends"`
  Is_Ignored    bool     `json:"is_ignored"`
  Style_Id      int      `json:"style_id"`
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
