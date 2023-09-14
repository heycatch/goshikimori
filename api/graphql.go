package api

type GraphQL struct {
	Data Data `json:"data"`
	Errors []MessageGraphQL `json:"errors"`
}

type Data struct {
	Animes []AnimesGraphQL  `json:"animes"`
}

type MessageGraphQL struct {
	Message string `json:"message"`
}

/*
# Removed:
	Season: nil;
	NextEpisodeAt: nil;
	ExternalLinks: nil;
*/
type AnimesGraphQL struct {
	Id 							string 									`json:"id"`
	MalId 					string 									`json:"malId"`
	Name 						string 									`json:"name"`
	Russian 				string 									`json:"russian"`
	LicenseNameRu 	string 									`json:"licenseNameRu"`
	English 				string 									`json:"english"`
	Japanese 				string 									`json:"japanese"`
	Synonyms 				[]string 								`json:"synonyms"`
	Kind 						string 									`json:"kind"`
	Rating 					string 									`json:"rating"`
	Score 					float32 								`json:"score"`
	Status 					string 									`json:"status"`
	Episodes 				int 										`json:"episodes"`
	EpisodesAired 	int 										`json:"episodesAired"`
	Duration 				int 										`json:"duration"`
	AiredOn 				AiredOnGraphQL 					`json:"airedOn"`
	ReleasedOn 			AiredOnGraphQL 					`json:"releasedOn"`
	Url 						string 									`json:"url"`
	Poster 					PosterGraphQL 					`json:"poster"`
	Fansubbers 			[]string 								`json:"fansubbers"`
	Fandubbers 		  []string 								`json:"fandubbers"`
	Licensors 			[]string 								`json:"licensors"`
	CreatedAt 			string 									`json:"createdAt"`
	UpdatedAt 			string 									`json:"updatedAt"`
	IsCensored 			bool 										`json:"isCensored"`
	Genres 				  []GenresGraphQL 				`json:"genres"`
	Studios 			  []StudiosGraphQL 				`json:"studios"`
	PersonRoles     []PersonRolesGraphQL 		`json:"personRoles"`
	CharacterRoles  []CharacterRolesGraphQL `json:"characterRoles"`
	Related 			  []RelatedGraphQL 				`json:"related"`
	Videos 				  []VideosGraphQL 				`json:"videos"`
	Screenshots 	  []ScreenshotsGraphQL 		`json:"screenshots"`
	ScoresStats 	  []ScoresStatusGraphQL 	`json:"scoresStats"`
	StatusesStats   []StatusesStatsGraphQL  `json:"statusesStats"`
	Description     string 									`json:"description"`
	DescriptionHTML string 									`json:"descriptionHtml"`
}

type AiredOnGraphQL struct {
	Year  int 	 `json:"year"`
	Month int 	 `json:"month"`
	Day   int 	 `json:"day"`
	Date  string `json:"date"`
}

type PosterGraphQL struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	MainUrl     string `json:"mainUrl"`
}

type GenresGraphQL struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Kind    string `json:"kind"`
}

type StudiosGraphQL struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

type PersonRolesGraphQL struct {
	Id      string  	`json:"id"`
	RolesRu []string 	`json:"rolesRu"`
	RolesEn []string 	`json:"rolesEn"`
	Person  Persons 	`json:"person"`
}

type Persons struct {
	Id 		 string  `json:"id"`
	Name 	 string  `json:"name"`
	Poster Posters `json:"poster"`
}

type Posters struct {
	Id string `json:"id"`
}

type CharacterRolesGraphQL struct {
	Id        string   `json:"id"`
	RolesRu   []string `json:"rolesRu"`
	RolesEn   []string `json:"rolesEn"`
	Character Persons  `json:"character"`
}

// If the Anime or Manga is not found, nil is returned.
type RelatedGraphQL struct {
	Id 				 string 	  `json:"id"`
	Anime 		 ObjectInfo `json:"anime"`
	Manga 		 ObjectInfo `json:"manga"`
	RelationRu string 		`json:"relationRu"`
	RelationEn string 		`json:"relationEn"`
}

type ObjectInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type VideosGraphQL struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type ScreenshotsGraphQL struct {
	Id 					string `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	X166Url     string `json:"x166Url"`
	X332Url 		string `json:"x332Url"`
}

type ScoresStatusGraphQL struct {
	Score int `json:"score"`
	Count int `json:"count"`
}

type StatusesStatsGraphQL struct {
	Status string `json:"status"`
	Count  int		`json:"count"`
}
