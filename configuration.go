package goshikimori

import (
	"net/url"
	"strconv"

	"github.com/heycatch/goshikimori/concat"
)

type Configuration struct {
	Application, AccessToken string
}

type FastId struct {
	Id   int
	Conf Configuration
	Err  error
}

// Getting an id(anime, manga, ranobe, user, person, group).
//
// More information can be found in the [example1] and [example2].
//
// [example1]: https://github.com/heycatch/goshikimori/blob/master/examples/custom_fastid
// [example2]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func (f *FastId) GetFastId() int { return f.Id }

// To create a custom id(anime, manga, ranobe, user, person, group).
//
// More information can be found in the [example1] and [example2].
//
// [example1]: https://github.com/heycatch/goshikimori/blob/master/examples/custom_fastid
// [example2]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func (c *Configuration) SetFastId(id int) *FastId {
	return &FastId{Id: id, Conf: *c, Err: nil}
}

// Getting the configuration.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func (c *Configuration) GetConfiguration() (string, string) {
	return c.Application, c.AccessToken
}

// To register the application, follow the link from [OAuth].
//
// More information can be found in the [example].
//
// [OAuth]: https://github.com/heycatch/goshikimori/blob/master/examples/first_steps
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func SetConfiguration(appname, token string) *Configuration {
	return &Configuration{Application: appname, AccessToken: token}
}

type Options struct {
	Order, Kind, Status, Season, Rating,
	Type, Target_type, Duration, Mylist,
	Forum, Linked_type string
	Page, Limit, Score, Linked_id, Target_id int
	Censored                                 bool
	Genre_v2                                 []int
}

type Result interface {
	OptionsOnlyPageLimit(int, int) string
	OptionsAnime() string
	OptionsManga() string
	OptionsRanobe() string
	OptionsCalendar() string
	OptionsAnimeRates() string
	OptionsMangaRates() string
	OptionsUserHistory() string
	OptionsMessages() string
	OptionsPeople() string
	OptionsTopics() string
	OptionsTopicsHot() string
	OptionsRandomAnime() string
	OptionsRandomManga() string
	OptionsRandomRanobe() string
}

func (o *Options) OptionsOnlyPageLimit(page, limit int) string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= page {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= limit {
		v.Add("limit", strconv.Itoa(o.Limit))
	}

	return v.Encode()
}

func (o *Options) OptionsTopics() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 30 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Forum != "" {
		v.Add("forum", o.Forum)
	}
	// linked_id and linked_type are only used together.
	if o.Linked_id >= 1 && o.Linked_type != "" {
		v.Add("linked_id", strconv.Itoa(o.Linked_id))
		v.Add("linked_type", o.Linked_type)
	}

	return v.Encode()
}

func (o *Options) OptionsMessages() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 100 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	// The type is required.
	if o.Type == "" {
		v.Add("type", MESSAGE_TYPE_NEWS)
	} else {
		v.Add("type", o.Type)
	}

	return v.Encode()
}

func (o *Options) OptionsUserHistory() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 100 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Target_type != "" {
		v.Add("target_type", o.Target_type)
	}
	if o.Target_id > 0 {
		v.Add("target_id", strconv.Itoa(o.Target_id))
	}

	return v.Encode()
}

func (o *Options) OptionsAnime() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Order != "" {
		v.Add("order", o.Order)
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Rating != "" {
		v.Add("rating", o.Rating)
	}
	if o.Duration != "" {
		v.Add("duration", o.Duration)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := concat.MapGenresAnime(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsManga() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Order != "" {
		v.Add("order", o.Order)
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := concat.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsRanobe() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Order != "" {
		v.Add("order", o.Order)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := concat.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsCalendar() string {
	v := url.Values{}

	v.Add("censored", strconv.FormatBool(o.Censored))

	return v.Encode()
}

func (o *Options) OptionsAnimeRates() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 5000 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	return v.Encode()
}

// FIXME: The manga has no status, ranobe is missing.
func (o *Options) OptionsMangaRates() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 5000 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	return v.Encode()
}

func (o *Options) OptionsPeople() string {
	v := url.Values{}

	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}

	return v.Encode()
}

func (o *Options) OptionsTopicsHot() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 10 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}

	return v.Encode()
}

func (o *Options) OptionsRandomAnime() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Rating != "" {
		v.Add("rating", o.Rating)
	}
	if o.Duration != "" {
		v.Add("duration", o.Duration)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := concat.MapGenresAnime(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsRandomManga() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := concat.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsRandomRanobe() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := concat.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}
