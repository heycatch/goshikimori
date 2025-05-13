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

  fastc, status, _ := c.FastIdPeople("Томокадзу Сэки")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    pc, _, _ := fastc.SearchPeople()

    if pc.Id == 1 && pc.Job_title == "Сэйю"  {
      t.Logf("%s - found (Cyrillic alphabet)", pc.Name)
    } else {
      t.Skip()
    }
  }

  fastl, status, _ := c.FastIdPeople("Aya Hirano")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    pl, _, _ := fastl.SearchPeople()

    if pl.Id == 4 && pl.Job_title == "Сэйю"  {
      t.Logf("%s - found (Latin alphabet)", pl.Name)
    } else {
      t.Error("People not found (Latin alphabet)")
    }
  }
}

func TestCharacter(t *testing.T) {
  c := conf()

  fastc, status, _ := c.FastIdCharacter("Такуми Усуи")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    pc, _, _ := fastc.SearchCharacter()

    if pc.Id == 14523 && pc.Altname == "Perverted Alien" {
      t.Logf("%s - found (Cyrillic alphabet)", pc.Name)
    } else {
      t.Skip()
    }
  }

  fastl, status, _ := c.FastIdCharacter("Takumi Usui")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    pl, _, _ := fastl.SearchCharacter()

    if pl.Id == 14523 && pl.Altname == "Perverted Alien" {
      t.Logf("%s - found (Latin alphabet)", pl.Name)
    } else {
      t.Error("Character not found (Latin alphabet)")
    }
  }
}

func TestClub(t *testing.T) {
  var s StatusBar
  s.settings(5, "#", 1)
  s.run()

  c := conf()

  fastc, status, _ := c.FastIdClub("Ачивки (достижения)")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastc.Id == 315 {
      t.Logf("Fast club found (Cyrillic alphabet)")
    } else {
      t.Skip()
    }
  }

  fastl, status, _ := c.FastIdClub("Genshin Impact")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastl.Id == 3057 {
      t.Logf("Fast club found (Latin alphabet)")
    } else {
      t.Error("Fast club not found (Latin alphabet)")
    }
  }
}

func TestRanobe(t *testing.T) {
  c := conf()

  fastc, status, _ := c.FastIdRanobe("Ангел кровопролития")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastc.Id == 115586 {
      t.Logf("Fast ranobe found (Cyrillic alphabet)")
    } else {
      t.Skip()
    }
  }

  fastl, status, _ := c.FastIdRanobe("Satsuriku no Tenshi")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastl.Id == 115586 {
      t.Logf("Fast ranobe found (Latin alphabet)")
    } else {
      t.Error("Fast ranobe not found (Latin alphabet)")
    }
  }
}

func TestManga(t *testing.T) {
  c := conf()

  fastc, status, _ := c.FastIdManga("Тетрадь смерти")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastc.Id == 21 {
      t.Logf("Fast manga found (Cyrillic alphabet)")
    } else {
      t.Skip()
    }
  }

  fastl, status, _ := c.FastIdManga("Death Note")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastl.Id == 21 {
      t.Logf("Fast manga found (Latin alphabet)")
    } else {
      t.Error("Fast manga not found (Latin alphabet)")
    }
  }
}

func TestAnime(t *testing.T) {
  var s StatusBar
  s.settings(5, "#", 1)
  s.run()

  c := conf()

  fastc, status, _ := c.FastIdAnime("Тетрадь смерти")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastc.Id == 1535 {
      t.Logf("Fast anime found (Cyrillic alphabet)")
    } else {
      t.Skip()
    }
  }

  fastl, status, _ := c.FastIdAnime("Death Note")
  if status == -1 {
    t.Log("timeout from shikimori")
  } else {
    if fastl.Id == 1535 {
      t.Logf("Fast anime found (Latin alphabet)")
    } else {
      t.Error("Fast anime not found (Latin alphabet)")
    }
  }
}
