package api

type Num struct {
  Id string `json:":id"`
}

func FoundID(s string) string {
  create := new(Num)
  create.Id = s
  return s
}
