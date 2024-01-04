package morm

import "github.com/devsamahd/morm"

type TestModel struct {
	morm.Model `bson:",inline"`
	Field1 string
	Field2 int
	TestModel2 *TestModel2
}

type TestModel2 struct {
	morm.Model `bson:",inline"`
	Field3 string
	Field4 int
}