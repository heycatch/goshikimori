package str

import "fmt"

func ConvertAchievements(id int) string {
  return fmt.Sprintf("achievements?user_id=%d", id)
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