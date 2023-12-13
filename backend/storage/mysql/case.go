package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Case ...
type Case struct {
	CaseID    string `gorm:"column:case_id"`
	ProblemID string `gorm:"column:problem_id"`
	Docs      string `gorm:"column:docs"`
}

// CaseDao ...
type CaseDao struct {
	db *gorm.DB
}

// NewCaseDao ...
func NewCaseDao() (*CaseDao, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:@tcp(127.0.0.1:3306)/elastic_oj?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}))
	if err != nil {
		return nil, err
	}

	return &CaseDao{
		db: db,
	}, nil
}

// GetCasesByProblemID ...
func (cd *CaseDao) GetCasesByProblemID(problemID string) ([]*Case, error) {
	var cases []*Case
	if err := cd.db.Where("problem_id", problemID).Find(&cases).Error; err != nil {
		return nil, err
	}
	return cases, nil
}
