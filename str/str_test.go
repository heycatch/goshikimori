package str

import "testing"

func TestConvertAchievements(t *testing.T) {
  if ConvertAchievements(1337) == "achievements?user_id=1337" {
    t.Log("ConvertAchievements passed")
  } else {
    t.Error("ConvertAchievements failed")
  }
}

func TestConvertAnime(t *testing.T) {
  if ConvertAnime(1337, "test") == "animes/1337/test" {
    t.Log("ConvertAnime passed")
  } else {
    t.Error("ConvertAnime failed")
  }
}

func TestConvertRoles(t *testing.T) {
  if ConvertRoles(1337, "test") == "test/1337/roles" {
    t.Log("ConvertRoles passed")
  } else {
    t.Error("ConvertRoles failed")
  }
}

func TestConvertSimilar(t *testing.T) {
  if ConvertSimilar(1337, "test") == "test/1337/similar" {
    t.Log("ConvertSimilar passed")
  } else {
    t.Error("ConvertSimilar failed")
  }
}

func TestConvertRelated(t *testing.T) {
  if ConvertRelated(1337, "test") == "test/1337/related" {
    t.Log("ConvertRelated passed")
  } else {
    t.Error("ConvertRelated failed")
  }
}

func TestConvertFranchise(t *testing.T) {
  if ConvertFranchise(1337, "test") == "test/1337/franchise" {
    t.Log("ConvertFranchise passed")
  } else {
    t.Error("ConvertFranchise failed")
  }
}

func TestConvertCalendar(t *testing.T) {
  if ConvertCalendar("test") == "calendar?test" {
    t.Log("ConvertCalendar passed")
  } else {
    t.Error("ConvertCalendar failed")
  }
}

func TestConvertUser(t *testing.T) {
  if ConvertUser(1337, "test") == "users/1337/test" {
    t.Log("ConvertUser passed")
  } else {
    t.Error("ConvertUser failed")
  }
}

func TestConvertExternalLinks(t *testing.T) {
  if ConvertExternalLinks(1337, "test") == "test/1337/external_links" {
    t.Log("ConvertExternalLinks passed")
  } else {
    t.Error("ConvertExternalLinks failed")
  }
}

func TestConvertUserRates(t *testing.T) {
  if ConvertUserRates(1337, "test", "bob") == "users/1337/test?bob" {
    t.Log("ConvertUserRates passed")
  } else {
    t.Error("ConvertUserRates failed")
  }
}

func TestConvertFriend(t *testing.T) {
  if ConvertFriend(1337) == "friends/1337" {
    t.Log("ConvertFriend passed")
  } else {
    t.Error("ConvertFriend failed")
  }
}

func TestConvertMessages(t *testing.T) {
  if ConvertMessages(1337, "test") == "users/1337/messages?test" {
    t.Log("ConvertMessages passed")
  } else {
    t.Error("ConvertMessages failed")
  }
}

func TestConvertConstants(t *testing.T) {
  if ConvertConstants("test") == "constants/test" {
    t.Log("ConvertConstants passed")
  } else {
    t.Error("ConvertConstants failed")
  }
}

func TestConvertPeople(t *testing.T) {
  if ConvertPeople(1337) == "people/1337" {
    t.Log("ConvertPeople passed")
  } else {
    t.Error("ConvertPeople failed")
  }
}

func TestConvertClub(t *testing.T) {
  if ConvertClub(1337, "test") == "clubs/1337/test" {
    t.Log("ConvertClub passed")
  } else {
    t.Error("ConvertClub failed")
  }
}

func TestConvertFavorites(t *testing.T) {
  if ConvertFavorites("test", 1337, "test") == "favorites/test/1337/test" {
    t.Log("ConvertFavorites passed")
  } else {
    t.Error("ConvertFavorites failed")
  }
}

func TestConvertFavoritesReorder(t *testing.T) {
  if ConvertFavoritesReorder(1337) == "favorites/1337/reorder" {
    t.Log("ConvertFavoritesReorder passed")
  } else {
    t.Error("ConvertFavoritesReorder failed")
  }
}

func TestConvertDialogs(t *testing.T) {
  if ConvertDialogs(1337) == "dialogs/1337" {
    t.Log("ConvertDialogs passed")
  } else {
    t.Error("ConvertDialogs failed")
  }
}

func TestRandom(t *testing.T) {
  if ConvertRandom("animes", 10) == "animes?order=random&limit=10" {
    t.Log("ConvertRandom passed")
  } else {
    t.Error("ConvertRandom failed")
  }
}

func TestConvertIgnoreUser(t *testing.T) {
  if ConvertIgnoreUser(1337) == "v2/users/1337/ignore" {
    t.Log("ConvertIgnoreUser passed")
  } else {
    t.Error("ConvertIgnoreUser failed")
  }
}

func TestConvertUserBriefInfo(t *testing.T) {
  if ConvertUserBriefInfo(1337) == "users/1337/info" {
    t.Log("ConvertUserBriefInfo passed")
  } else {
    t.Error("ConvertUserBriefInfo failed")
  }
}

func TestConvertTopicsId(t *testing.T) {
  if ConvertTopicsId(1337) == "topics/1337" {
    t.Log("ConvertTopicsId passed")
  } else {
    t.Error("ConvertTopicsId failed")
  }
}

func TestConvertIgnoreTopic(t *testing.T) {
  if ConvertIgnoreTopic(1337) == "v2/topics/1337/ignore" {
    t.Log("ConvertIgnoreTopic passed")
  } else {
    t.Error("ConvertIgnoreTopic failed")
  }
}
