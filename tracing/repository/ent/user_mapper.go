package ent

import "tracing/domain"

func (entity *User) Mapper() *domain.User {
	if entity == nil {
		return nil
	}
	dom := new(domain.User)
	dom.ID = entity.ID
	dom.Mobile = entity.Mobile
	dom.Nickname = entity.Nickname
	dom.Bio = entity.Bio
	dom.CreateTime = entity.CreateTime
	dom.UpdateTime = entity.UpdateTime

	return dom
}

func (entities Users) Mapper() domain.Users {
	var doms domain.Users
	for _, entity := range entities {
		doms = append(doms, entity.Mapper())
	}
	return doms
}
