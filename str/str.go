package str

import "fmt"

func ConvertAchievements(id int) string {
  return fmt.Sprintf("achievements?user_id=%d", id)
}

func ConvertSearchById(name string, id int) string {
  return fmt.Sprintf("%s/%d", name, id)
}

func ConvertAnime(id int, name string) string {
  return fmt.Sprintf("animes/%d/%s", id, name)
}

func ConvertRoles(id int, name string) string {
  return fmt.Sprintf("%s/%d/roles", name, id)
}

func ConvertSimilar(id int, name string) string {
  return fmt.Sprintf("%s/%d/similar", name, id)
}

func ConvertRelated(id int, name string) string {
  return fmt.Sprintf("%s/%d/related", name, id)
}

func ConvertFranchise(id int, name string) string {
  return fmt.Sprintf("%s/%d/franchise", name, id)
}

func ConvertCalendar(name string) string {
  return fmt.Sprintf("calendar?%s", name)
}

func ConvertUser(id int, name string) string {
  return fmt.Sprintf("users/%d/%s", id, name)
}

func ConvertExternalLinks(id int, name string) string {
  return fmt.Sprintf("%s/%d/external_links", name, id)
}

func ConvertUserRates(id int, name, options string) string {
  return fmt.Sprintf("users/%d/%s?%s", id, name, options)
}

func ConvertFriend(id int) string {
  return fmt.Sprintf("friends/%d", id)
}

func ConvertMessages(id int, name string) string {
  return fmt.Sprintf("users/%d/messages?%s", id, name)
}

func ConvertConstants(name string) string {
  return fmt.Sprintf("constants/%s", name)
}

func ConvertPeople(id int) string {
  return fmt.Sprintf("people/%d", id)
}

func ConvertClub(id int, name string) string {
  return fmt.Sprintf("clubs/%d/%s", id, name)
}

func ConvertFavorites(linked_type string, id int, kind string) string {
  return fmt.Sprintf("favorites/%s/%d/%s", linked_type, id, kind)
}

func ConvertFavoritesReorder(id int) string {
  return fmt.Sprintf("favorites/%d/reorder", id)
}

func ConvertDialogs(id int) string {
  return fmt.Sprintf("dialogs/%d", id)
}

func ConvertRandom(name string, limit int) string {
  return fmt.Sprintf("%s?order=random&limit=%d", name, limit)
}

func ConvertIgnoreUser(id int) string {
  return fmt.Sprintf("v2/users/%d/ignore", id)
}

func ConvertUserBriefInfo(id int) string {
  return fmt.Sprintf("users/%d/info", id)
}

func ConvertTopicsId(id int) string {
  return fmt.Sprintf("topics/%d", id)
}

func ConvertIgnoreTopic(id int) string {
  return fmt.Sprintf("v2/topics/%d/ignore", id)
}

func ConvertTopicsType(id int, name string) string {
  return fmt.Sprintf("%s/%d/topics", name, id)
}
