package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Problem ...
type Problem struct {
	ProblemID     string `gorm:"column:problem_id"`
	Title         string `gorm:"column:title"`
	Description   string `gorm:"column:description"`
	ESIndex       string `gorm:"column:es_index"`
	StandardQuery string `gorm:"column:standard_query"`
}

// ProblemDao ...
type ProblemDao struct {
	db *gorm.DB
}

// NewProblemDao ...
func NewProblemDao() (*ProblemDao, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:@tcp(127.0.0.1:3306)/elastic_oj?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}))
	if err != nil {
		return nil, err
	}

	return &ProblemDao{
		db: db,
	}, nil
}

// ListProblems ...
func (pd *ProblemDao) ListProblems() ([]*Problem, error) {
	var problems []*Problem
	if err := pd.db.Find(&problems).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return problems, nil
}

// GetProblemByID ...
func (pd *ProblemDao) GetProblemByID(problemID string) (*Problem, error) {
	problem := &Problem{}
	if err := pd.db.First(problem, "problem_id = ?", problemID).Error; err != nil {
		return nil, err
	}
	return problem, nil
}
