package morm


type TestModel struct {
	Model `bson:",inline"`
	Field1 string
	Field2 int
	TestModel2 *TestModel2
}

type TestModel2 struct {
	Model `bson:",inline"`
	Field3 string
	Field4 int
}