package services

import (
	"apiTest.com/models"
	"sync"
	"time"
)

var once sync.Once

type testService struct {
	testRepository models.TestRepository
}

func NewTestService(r models.TestRepository) models.TestService {
	var instance models.TestService
	once.Do(func() {
		instance = &testService{
			testRepository: r,
		}
	})
	return instance
}

func (t testService) Create(input *models.CreateTestInput) (*models.Test, error) {
	test := models.Test{
		Title:    input.Title,
		Content:  input.Content,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	return t.testRepository.Create(&test)
}

func (t testService) FindAll() (*[]models.Test, error) {
	return t.testRepository.FindAll()
}

func (t testService) FindById(id string) (*models.Test, error) {
	return t.testRepository.FindById(id)
}

//func (t testService) Update(id int32, input *models.UpdateTestInput) (*models.Test, error) {
//	update := models.Test{
//		Title:    input.Title,
//		Content:  input.Content,
//		UpdateAt: time.Now(),
//	}
//	post, error := t.FindById(id)
//	if()
//
//	result, error = t.testRepository.Update(post, &update)
//
//	return
//}
//
//func (t testService) Delete(id int32) (*models.Test, error) {
//}

func (t testService) Update(id string, input *models.UpdateTestInput) (*models.Test, error) {
	post, err := t.FindById(id)
	if err != nil {
		return nil, err
	}
	update := models.Test{
		Title:    input.Title,
		Content:  input.Content,
		UpdateAt: time.Now(),
	}

	return t.testRepository.Update(post, &update)
}

func (t testService) Delete(id int32) (*models.Test, error) {
	//TODO implement me
	panic("implement me")
}
