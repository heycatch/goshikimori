package str

import "testing"

func TestConvertAchievements(t *testing.T) {
  if "achievements?user_id=1337" == ConvertAchievements(1337) {
    t.Log("ConvertAchievements passed")
  } else {
    t.Error("ConvertAchievements failed")
  }
}

func TestConvertAnime(t *testing.T) {
  if "animes/1337/test" == ConvertAnime(1337, "test") {
    t.Log("ConvertAnime passed")
  } else {
    t.Error("ConvertAnime failed")
  }
}

func TestConvertRoles(t *testing.T) {
  if "test/1337/roles" == ConvertRoles(1337, "test") {
    t.Log("ConvertRoles passed")
  } else {
    t.Error("ConvertRoles failed")
  }
}

func TestConvertSimilar(t *testing.T) {
  if "test/1337/similar" == ConvertSimilar(1337, "test") {
    t.Log("ConvertSimilar passed")
  } else {
    t.Error("ConvertSimilar failed")
  }
}

func TestConvertRelated(t *testing.T) {
  if "test/1337/related" == ConvertRelated(1337, "test") {
    t.Log("ConvertRelated passed")
  } else {
    t.Error("ConvertRelated failed")
  }
}

func TestConvertFranchise(t *testing.T) {
  if "test/1337/franchise" == ConvertFranchise(1337, "test") {
    t.Log("ConvertFranchise passed")
  } else {
    t.Error("ConvertFranchise failed")
  }
}

func TestConvertCalendar(t *testing.T) {
  if "calendar?test" == ConvertCalendar("test") {
    t.Log("ConvertCalendar passed")
  } else {
    t.Error("ConvertCalendar failed")
  }
}

func TestConvertUser(t *testing.T) {
  if "users/1337/test" == ConvertUser(1337, "test") {
    t.Log("ConvertUser passed")
  } else {
    t.Error("ConvertUser failed")
  }
}

func TestConvertExternalLinks(t *testing.T) {
  if "test/1337/external_links" == ConvertExternalLinks(1337, "test") {
    t.Log("ConvertExternalLinks passed")
  } else {
    t.Error("ConvertExternalLinks failed")
  }
}

func TestConvertUserRates(t *testing.T) {
  if "users/1337/test?bob" == ConvertUserRates(1337, "test", "bob") {
    t.Log("ConvertUserRates passed")
  } else {
    t.Error("ConvertUserRates failed")
  }
}

func TestConvertFriend(t *testing.T) {
  if "friends/1337" == ConvertFriend(1337) {
    t.Log("ConvertFriend passed")
  } else {
    t.Error("ConvertFriend failed")
  }
}

func TestConvertMessages(t *testing.T) {
  if "users/1337/messages?test" == ConvertMessages(1337, "test") {
    t.Log("ConvertMessages passed")
  } else {
    t.Error("ConvertMessages failed")
  }
}

func TestConvertConstants(t *testing.T) {
  if "constants/test" == ConvertConstants("test") {
    t.Log("ConvertConstants passed")
  } else {
    t.Error("ConvertConstants failed")
  }
}

func TestConvertPeople(t *testing.T) {
  if "people/1337" == ConvertPeople(1337) {
    t.Log("ConvertPeople passed")
  } else {
    t.Error("ConvertPeople failed")
  }
}

func TestConvertClub(t *testing.T) {
  if "clubs/1337/test" == ConvertClub(1337, "test") {
    t.Log("ConvertClub passed")
  } else {
    t.Error("ConvertClub failed")
  }
}

func TestConvertFavorites(t *testing.T) {
  if "favorites/test/1337/test" == ConvertFavorites("test", 1337, "test") {
    t.Log("ConvertFavorites passed")
  } else {
    t.Error("ConvertFavorites failed")
  }
}

func TestConvertFavoritesReorder(t *testing.T) {
  if "favorites/1337/reorder" == ConvertFavoritesReorder(1337) {
    t.Log("ConvertFavoritesReorder passed")
  } else {
    t.Error("ConvertFavoritesReorder failed")
  }
}

func TestConvertDialogs(t *testing.T) {
  if "dialogs/1337" == ConvertDialogs(1337) {
    t.Log("ConvertDialogs passed")
  } else {
    t.Error("ConvertDialogs failed")
  }
}

func TestConvertIgnoreUser(t *testing.T) {
  if "v2/users/1337/ignore" == ConvertIgnoreUser(1337) {
    t.Log("ConvertIgnoreUser passed")
  } else {
    t.Error("ConvertIgnoreUser failed")
  }
}

