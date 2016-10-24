package domain

import (
  "strconv"
)

type Summoner struct {
  Id int64
  Name string
}

func (s *Summoner) GetName() string {
  return s.Name
}

func (s *Summoner) GetId() string {
  return strconv.FormatInt(s.Id, 10)
}
