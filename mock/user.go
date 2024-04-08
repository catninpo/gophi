package mock

import "github.com/catninpo/gophi"

type UserService struct {
	ByIDFn      func(id int) (*gophi.User, error)
	ByIDInvoked bool
}

func (s *UserService) UserByID(id int) (*gophi.User, error) {
	s.ByIDInvoked = true
	return s.ByIDFn(id)
}
