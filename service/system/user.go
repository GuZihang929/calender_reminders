package system

import (
	"calender_reminders/global"
	"calender_reminders/model/system"
	"time"
)

type UserService struct {
}

func (s *UserService) Add(userId int, context string, time time.Time) error {
	u := &system.Calender{UserId: userId, Context: context, RemindTime: time}
	err := global.FAST_DB.Create(u).Error
	return err
}

func (s *UserService) Del(id int) error {
	err := global.FAST_DB.Delete("id = ?", id).Error
	return err
}

func (s *UserService) Check(userId int) ([]*system.Calender, error) {
	res := []*system.Calender{}
	err := global.FAST_DB.Where("user_id = ?", userId).Find(res).Error
	return res, err
}

func (s *UserService) Update(id int, context string, time time.Time) error {
	res := &system.Calender{}

	result := global.FAST_DB.First(res, id)
	if result.Error != nil {
		return result.Error
	}

	res.Context = context
	res.RemindTime = time

	if err := global.FAST_DB.Save(res).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetPendingTasks() ([]*system.Calender, error) {
	res := []*system.Calender{}
	err := global.FAST_DB.Where("remind_time <= ? and status = ?", time.Now().Unix(), 1).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) UpdateStatus(id int, status int) error {
	err := global.FAST_DB.Model(&system.Calender{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUser(userId int) *system.User {
	res := &system.User{}
	err := global.FAST_DB.Where("user_id = ?", userId).Find(&res).Error
	if err != nil {
		return nil
	}
	return res
}
