package goshikimori

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/heycatch/goshikimori/concat"
	"github.com/heycatch/goshikimori/consts"
)

type Configuration struct {
	Application string
	AccessToken string
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
	Order       string
	Kind        string
	Status      string
	Season      string
	Rating      string
	Type        string
	Target_type string
	Duration    string
	Mylist      string
	Forum       string
	Linked_type string
	Page        int
	Limit       int
	Score       int
	Linked_id   int
	Target_id   int
	Genre_v2    []int
	Censored    bool
}

type Result interface {
	OptionsOnlyPageLimit(int, int) string
	OptionsAnime()                 string
	OptionsManga()                 string
	OptionsRanobe()                string
	OptionsCalendar()              string
	OptionsAnimeRates()            string
	OptionsMangaRates()            string
	OptionsUserHistory()           string
	OptionsMessages()              string
	OptionsPeople()                string
	OptionsTopics()                string
	OptionsTopicsHot()             string
	OptionsRandomAnime()           string
	OptionsRandomManga()           string
	OptionsRandomRanobe()          string

	OptionsOnlyPageLimitV2() string
	OptionsAnimeV2()         string
	OptionsMangaV2()         string
	OptionsRanobeV2()        string
	OptionsCalendarV2()      string
	OptionsAnimeRatesV2()    string
	OptionsMangaRatesV2()    string
	OptionsUserHistoryV2()   string
	OptionsMessagesV2()      string
	OptionsPeopleV2()        string
	OptionsTopicsV2()        string
	OptionsTopicsHotV2()     string
}

// TODO: (heycatch) abandon url.QueryEscape in the future.
func encodeParamEscaped(key, value string) string {
	return url.QueryEscape(key) + "=" + url.QueryEscape(value)
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

func (o *Options) OptionsOnlyPageLimitV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 2)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

func (o *Options) OptionsTopicsV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 5)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("forum", o.Forum))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Linked_id), 10)
	pairs = append(pairs, encodeParamEscaped("linked_id", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("linked_type", o.Linked_type))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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
		v.Add("type", consts.MESSAGE_TYPE_NEWS)
	} else {
		v.Add("type", o.Type)
	}

	return v.Encode()
}

func (o *Options) OptionsMessagesV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 3)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("type", o.Type))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

func (o *Options) OptionsUserHistoryV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 4)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("target_type", o.Target_type))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Target_id), 10)
	pairs = append(pairs, encodeParamEscaped("target_id", string(numBuf)))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

func (o *Options) OptionsAnimeV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 12)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Score), 10)
	pairs = append(pairs, encodeParamEscaped("score", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("order", o.Order))
	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	pairs = append(pairs, encodeParamEscaped("season", o.Season))
	pairs = append(pairs, encodeParamEscaped("rating", o.Rating))
	pairs = append(pairs, encodeParamEscaped("duration", o.Duration))
	pairs = append(pairs, encodeParamEscaped("mylist", o.Mylist))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}
	if genre := concat.MapGenresAnime(o.Genre_v2); genre != "" {
		pairs = append(pairs, encodeParamEscaped("genre_v2", genre))
	}

	totalLength := 0
	for _, p := range pairs {
		totalLength += len(p) + 1
	}
	if totalLength > 0 {
		totalLength--
	}
	sb.Grow(totalLength)

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

func (o *Options) OptionsMangaV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 10)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Score), 10)
	pairs = append(pairs, encodeParamEscaped("score", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("order", o.Order))
	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	pairs = append(pairs, encodeParamEscaped("season", o.Season))
	pairs = append(pairs, encodeParamEscaped("mylist", o.Mylist))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}
	if genre := concat.MapGenresAnime(o.Genre_v2); genre != "" {
		pairs = append(pairs, encodeParamEscaped("genre_v2", genre))
	}

	totalLength := 0
	for _, p := range pairs {
		totalLength += len(p) + 1
	}
	if totalLength > 0 {
		totalLength--
	}
	sb.Grow(totalLength)

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

func (o *Options) OptionsRanobeV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 10)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Score), 10)
	pairs = append(pairs, encodeParamEscaped("score", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("order", o.Order))
	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	pairs = append(pairs, encodeParamEscaped("season", o.Season))
	pairs = append(pairs, encodeParamEscaped("mylist", o.Mylist))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}
	if genre := concat.MapGenresAnime(o.Genre_v2); genre != "" {
		pairs = append(pairs, encodeParamEscaped("genre_v2", genre))
	}

	totalLength := 0
	for _, p := range pairs {
		totalLength += len(p) + 1
	}
	if totalLength > 0 {
		totalLength--
	}
	sb.Grow(totalLength)

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsCalendar() string {
	v := url.Values{}

	v.Add("censored", strconv.FormatBool(o.Censored))

	return v.Encode()
}

func (o *Options) OptionsCalendarV2() string {
	var sb strings.Builder
	pairs := make([]string, 0, 1)

	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

func (o *Options) OptionsAnimeRatesV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 4)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

// FIXME: (heycatch) The manga has no status, ranobe is missing.
// https://shikimori.one/api/doc/1.0/users/manga_rates.html
func (o *Options) OptionsMangaRatesV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 3)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsPeople() string {
	v := url.Values{}

	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}

	return v.Encode()
}

func (o *Options) OptionsPeopleV2() string {
	var sb strings.Builder
	pairs := make([]string, 0, 1)

	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsTopicsHot() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 10 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}

	return v.Encode()
}

func (o *Options) OptionsTopicsHotV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 1)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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

// NOTE: (heycatch) DEPRECATED and will be removed from future versions.
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
