package goshikimori

import "testing"

func TestLanguageCheck(t *testing.T) {
  if languageCheck("Ая Хирано") == "Ая Хирано" {
    t.Log("Cyrillic passed")
  } else {
    t.Error("Cyrillic failed")
  }

  if languageCheck("Aya Hirano") == "Aya+Hirano" {
    t.Log("Latin passed")
  } else {
    t.Error("Latin failed")
  }
}

func TestPeople(t *testing.T) {
  c := conf()

  fastc, _, _ := c.FastIdPeople("Ая Хирано")
  pc, _, _ := fastc.SearchPeople()

  if pc.Id == 4 && pc.Job_title == "Сэйю"  {
    t.Logf("%s - found (Cyrillic alphabet)", pc.Name)
  } else {
    t.Error("People not found (Cyrillic alphabet)")
  }

  fastl, _, _ := c.FastIdPeople("Aya Hirano")
  pl, _, _ := fastl.SearchPeople()

  if pl.Id == 4 && pl.Job_title == "Сэйю"  {
    t.Logf("%s - found (Latin alphabet)", pl.Name)
  } else {
    t.Error("People not found (Latin alphabet)")
  }
}

func TestCharacter(t *testing.T) {
  c := conf()

  fastc, _, _ := c.FastIdCharacter("Такуми Усуи")
  pc, _, _ := fastc.SearchCharacter()

  if pc.Id == 14523 && pc.Altname == "Perverted Alien" {
    t.Logf("%s - found (Cyrillic alphabet)", pc.Name)
  } else {
    t.Error("Character not found (Cyrillic alphabet)")
  }

  fastl, _, _ := c.FastIdCharacter("Takumi Usui")
  pl, _, _ := fastl.SearchCharacter()

  if pl.Id == 14523 && pl.Altname == "Perverted Alien" {
    t.Logf("%s - found (Latin alphabet)", pl.Name)
  } else {
    t.Error("Character not found (Latin alphabet)")
  }
}

func TestClub(t *testing.T) {
  var s StatusBar
  s.settings(5, "#", 1)
  s.run()

  c := conf()

  fastc, _, _ := c.FastIdClub("Ачивки (достижения)")
  if fastc.Id == 315 {
    t.Logf("Fast club found (Cyrillic alphabet)")
  } else {
    t.Error("Fast club not found (Cyrillic alphabet)")
  }

  fastl, _, _ := c.FastIdClub("Genshin Impact")
  if fastl.Id == 3057 {
    t.Logf("Fast club found (Latin alphabet)")
  } else {
    t.Error("Fast club not found (Latin alphabet)")
  }
}

func TestRanobe(t *testing.T) {
  c := conf()

  fastc, _, _ := c.FastIdRanobe("Ангел кровопролития")
  if fastc.Id == 115586 {
    t.Logf("Fast ranobe found (Cyrillic alphabet)")
  } else {
    t.Error("Fast ranobe not found (Cyrillic alphabet)")
  }

  fastl, _, _ := c.FastIdRanobe("Satsuriku no Tenshi")
  if fastl.Id == 115586 {
    t.Logf("Fast ranobe found (Latin alphabet)")
  } else {
    t.Error("Fast ranobe not found (Latin alphabet)")
  }
}

func TestManga(t *testing.T) {
  c := conf()

  fastc, _, _ := c.FastIdManga("Тетрадь смерти")
  if fastc.Id == 21 {
    t.Logf("Fast manga found (Cyrillic alphabet)")
  } else {
    t.Error("Fast manga not found (Cyrillic alphabet)")
  }

  fastl, _, _ := c.FastIdManga("Death Note")
  if fastl.Id == 21 {
    t.Logf("Fast manga found (Latin alphabet)")
  } else {
    t.Error("Fast manga not found (Latin alphabet)")
  }
}

func TestAnime(t *testing.T) {
  var s StatusBar
  s.settings(5, "#", 1)
  s.run()

  c := conf()

  fastc, _, _ := c.FastIdAnime("Тетрадь смерти")
  if fastc.Id == 1535 {
    t.Logf("Fast anime found (Cyrillic alphabet)")
  } else {
    t.Error("Fast anime not found (Cyrillic alphabet)")
  }

  fastl, _, _ := c.FastIdAnime("Death Note")
  if fastl.Id == 1535 {
    t.Logf("Fast anime found (Latin alphabet)")
  } else {
    t.Error("Fast anime not found (Latin alphabet)")
  }
}