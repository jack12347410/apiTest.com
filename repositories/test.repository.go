package repositories

import (
	"apiTest.com/models"
	"gorm.io/gorm"
)

type testRepository struct {
	DB *gorm.DB
}

func NewTestRepository(db *gorm.DB) models.TestRepository {
	return &testRepository{
		DB: db,
	}
}

func (t *testRepository) Create(create *models.Test) (*models.Test, error) {
	return create, t.DB.Create(create).Error
}

func (t *testRepository) FindAll() (*[]models.Test, error) {
	var tests *[]models.Test
	return tests, t.DB.Find(&tests).Error
}

func (t *testRepository) FindById(id string) (*models.Test, error) {
	var test *models.Test
	return test, t.DB.Where("id = ?", id).Find(&test).Error
}

func (t *testRepository) Update(post *models.Test, update *models.Test) (*models.Test, error) {
	return post, t.DB.Model(&post).Updates(update).Error
}

func (t *testRepository) Delete(delete *models.Test) (*models.Test, error) {
	return delete, t.DB.Delete(&delete).Error
}

func (t *testRepository) Migrate() error {
	return t.DB.AutoMigrate(&models.Test{})
}
