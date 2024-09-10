package goshikimori

import "time"

const (
  MAX_EXPECTATION time.Duration = 10
  CUSTOM_MAX_EXPECTATION time.Duration = 40

  SITE string = "https://shikimori.one/api/"

  TARGET_TYPE_ANIME string = "Anime"
  TARGET_TYPE_MANGA string = "Manga"

  SEASON_198x string = "198x"
  SEASON_199x string = "199x"
  SEASON_2000_2010 string = "2000_2010"
  SEASON_2010_2014 string = "2010_2014"
  SEASON_2015_2019 string = "2015_2019"
  SEASON_2020_2021 string = "2020_2021"
  SEASON_2022 string = "2022"
  SEASON_2023 string = "2023"
  SEASON_198x_NOT_EQUAL string = "!198x"
  SEASON_199x_NOT_EQUAL string = "!199x"
  SEASON_2000_2010_NOT_EQUAL string = "!2000_2010"
  SEASON_2010_2014_NOT_EQUAL string = "!2010_2014"
  SEASON_2015_2019_NOT_EQUAL string = "!2015_2019"
  SEASON_2020_2021_NOT_EQUAL string = "!2020_2021"
  SEASON_2022_NOT_EQUAL string = "!2022"
  SEASON_2023_NOT_EQUAL string = "!2023"

  MY_LIST_PLANNED string = "planned"
  MY_LIST_WATCHING string = "watching"
  MY_LIST_REWATCHING string = "rewatching"
  MY_LIST_COMPLETED string = "completed"
  MY_LIST_ON_HOLD string = "on_hold"
  MY_LIST_DROPPED string = "dropped"

  TOPIC_FORUM_ALL string = "all"
  TOPIC_FORUM_COSPLAY string = "cosplay"
  TOPIC_FORUM_ANIMANGA string = "animanga"
  TOPIC_FORUM_SITE string = "site"
  TOPIC_FORUM_GAMES string = "games"
  TOPIC_FORUM_VN string = "vn"
  TOPIC_FORUM_CONTEST string = "contest"
  TOPIC_FORUM_OFFTOPIC string = "offtopic"
  TOPIC_FORUM_CLUBS string = "clubs"
  TOPIC_FORUM_MYCLUBS string = "my_clubs"
  TOPIC_FORUM_CRITIQUES string = "critiques"
  TOPIC_FORUM_NEWS string = "news"
  TOPIC_FORUM_COLLECTIONS string = "collections"
  TOPIC_FORUM_ARTICLES string = "articles"

  TOPIC_LINKED_TYPE_ANIME string = "Anime"
  TOPIC_LINKED_TYPE_MANGA string = "Manga"
  TOPIC_LINKED_TYPE_RANOBE string = "Ranobe"
  TOPIC_LINKED_TYPE_CHARACTER string = "Character"
  TOPIC_LINKED_TYPE_PERSON string = "Person"
  TOPIC_LINKED_TYPE_CLUB string = "Club"
  TOPIC_LINKED_TYPE_CLUBPAGE string = "ClubPage"
  TOPIC_LINKED_TYPE_CRITIQUE string = "Critique"
  TOPIC_LINKED_TYPE_REVIEW string = "Review"
  TOPIC_LINKED_TYPE_CONTEST string = "Contest"
  TOPIC_LINKED_TYPE_COSPLAYGALLYRY string = "CosplayGallery"
  TOPIC_LINKED_TYPE_COLLECTION string = "Collection"
  TOPIC_LINKED_TYPE_ARTICLE string = "Article"

  MESSAGE_TYPE_INBOX string = "inbox"
  MESSAGE_TYPE_PRIVATE string = "private"
  MESSAGE_TYPE_SENT string = "sent"
  MESSAGE_TYPE_NEWS string = "news"
  MESSAGE_TYPE_NOTIFICATIONS string = "notifications"

  ANIME_ORDER_ID string = "id"
  ANIME_ORDER_RANKED string = "ranked"
  ANIME_ORDER_KIND string = "kind"
  ANIME_ORDER_POPULARITY string = "popularity"
  ANIME_ORDER_NAME string = "name"
  ANIME_ORDER_AIRED_ON string = "aired_on"
  ANIME_ORDER_EPISODES string = "episodes"
  ANIME_ORDER_STATUS string = "status"

  ANIME_KIND_MOVIE string = "movie"
  ANIME_KIND_MOVIE_NOT_EQUAL string = "!movie"
  ANIME_KIND_OVA string = "ova"
  ANIME_KIND_ONA string = "ona"
  ANIME_KIND_OVA_NOT_EQUAL string = "!ova"
  ANIME_KIND_ONA_NOT_EQUAL string = "!ona"
  ANIME_KIND_SPECIAL string = "special"
  ANIME_KIND_SPECIAL_NOT_EQUAL string = "!special"
  ANIME_KIND_MUSIC string = "music"
  ANIME_KIND_MUSIC_NOT_EQUAL string = "!music"
  ANIME_KIND_TV string = "tv"
  ANIME_KIND_TV_13 string = "tv_13"
  ANIME_KIND_TV_24 string = "tv_24"
  ANIME_KIND_TV_48 string = "tv_48"
  ANIME_KIND_TV_NOT_EQUAL string = "!tv"
  ANIME_KIND_TV_13_NOT_EQUAL string = "!tv_13"
  ANIME_KIND_TV_24_NOT_EQUAL string = "!tv_24"
  ANIME_KIND_TV_48_NOT_EQUAL string = "!tv_48"

  ANIME_STATUS_ANONS string = "anons"
  ANIME_STATUS_ANONS_NOT_EQUAL string = "!anons"
  ANIME_STATUS_ONGOING string = "ongoing"
  ANIME_STATUS_ONGOING_NOT_EQUAL string = "!ongoing"
  ANIME_STATUS_RELEASED string = "released"
  ANIME_STATUS_RELEASED_NOT_EQUAL string = "!released"

  ANIME_RATING_NONE string = "none"
  ANIME_RATING_G string = "g"
  ANIME_RATING_PG string = "pg"
  ANIME_RATING_PG_13 string = "pg_13"
  ANIME_RATING_R string = "r"
  ANIME_RATING_R_PLUS string = "r_plus"
  ANIME_RATING_RX string = "rx"
  ANIME_RATING_G_NOT_EQUAL string = "!g"
  ANIME_RATING_PG_NOT_EQUAL string = "!pg"
  ANIME_RATING_PG_13_NOT_EQUAL string = "!pg_13"
  ANIME_RATING_R_NOT_EQUAL string = "!r"
  ANIME_RATING_R_PLUS_NOT_EQUAL string = "!r_plus"
  ANIME_RATING_RX_NOT_EQUAL string = "!rx"

  ANIME_DURATION_S string = "S"
  ANIME_DURATION_D string = "D"
  ANIME_DURATION_F string = "F"
  ANIME_DURATION_S_NOT_EQUAL string = "!S"
  ANIME_DURATION_D_NOT_EQUAL string = "!D"
  ANIME_DURATION_F_NOT_EQUAL string = "!F"

  MANGA_ORDER_ID string = "id"
  MANGA_ORDER_RANKED string = "ranked"
  MANGA_ORDER_KIND string = "kind"
  MANGA_ORDER_POPULARITY string = "popularity"
  MANGA_ORDER_NAME string = "name"
  MANGA_ORDER_AIRED_ON string = "aired_on"
  MANGA_ORDER_VOLUMES string = "volumes"
  MANGA_ORDER_CHAPTERS string = "chapters"
  MANGA_ORDER_STATUS string = "status"

  MANGA_KIND_MANGA string = "manga"
  MANGA_KIND_MANHWA string = "manhwa"
  MANGA_KIND_MANHUA string = "manhua"
  MANGA_KIND_NOVEL string = "novel"
  MANGA_KIND_LIGHT_NOVEL string = "light_novel"
  MANGA_KIND_ONE_SHOT string = "one_shot"
  MANGA_KIND_DOUJIN string = "doujin"
  MANGA_KIND_MANGA_NOT_EQUAL string = "!manga"
  MANGA_KIND_MANHWA_NOT_EQUAL string = "!manhwa"
  MANGA_KIND_MANHUA_NOT_EQUAL string = "!manhua"
  MANGA_KIND_NOVEL_NOT_EQUAL string = "!novel"
  MANGA_KIND_LIGHT_NOVEL_NOT_EQUAL string = "!light_novel"
  MANGA_KIND_ONE_SHOT_NOT_EQUAL string = "!one_shot"
  MANGA_KIND_DOUJIN_NOT_EQUAL string = "!doujin"

  MANGA_STATUS_ANONS string = "anons"
  MANGA_STATUS_ONGOING string = "ongoing"
  MANGA_STATUS_RELEASED string = "released"
  MANGA_STATUS_PAUSED string = "paused"
  MANGA_STATUS_DISCONTINUED string = "discontinued"
  MANGA_STATUS_ANONS_NOT_EQUAL string = "!anons"
  MANGA_STATUS_ONGOING_NOT_EQUAL string = "!ongoing"
  MANGA_STATUS_RELEASED_NOT_EQUAL string = "!released"
  MANGA_STATUS_PAUSED_NOT_EQUAL string = "!paused"
  MANGA_STATUS_DISCONTINUED_NOT_EQUAL string = "!discontinued"

  PEOPLE_KIND_SEYU string = "seyu"
  PEOPLE_KIND_MANGAKA string = "mangaka"
  PEOPLE_KIND_PRODUCER string = "producer"


  FAVORITES_LINKED_TYPE_ANIME string = "Anime"
  FAVORITES_LINKED_TYPE_MANGA string = "Manga"
  FAVORITES_LINKED_TYPE_RANOBE string = "Ranobe"
  FAVORITES_LINKED_TYPE_PERSON string = "Person"
  FAVORITES_LINKED_TYPE_CHARACTER string = "Character"

  FAVORITES_KIND_COMMON string = "common"
  FAVORITES_KIND_SEYU string = "seyu"
  FAVORITES_KIND_MANGAKA string = "mangaka"
  FAVORITES_KIND_PRODUCER string = "producer"
  FAVORITES_KIND_PERSON string = "person"

  GENRES_ANIME string = "anime"
  GENRES_MANGA string = "manga"

  GENERATE_GENRES_ANIME string = "Anime"
  GENERATE_GENRES_MANGA string = "Manga"

  UNREAD_MESSAGES_IDS_NEWS string = "news"
  UNREAD_MESSAGES_IDS_MESSAGES string = "messages"
  UNREAD_MESSAGES_IDS_NOTIFICATIONS string = "notifications"

  GRAPHQL_ORDER_FIELD_ID string = "id"
  GRAPHQL_ORDER_FIELD_UPDATED_AT string = "updated_at"
  GRAPHQL_ORDER_ORDER_ASC string = "asc"
  GRAPHQL_ORDER_ORDER_DESC string = "desc"
)