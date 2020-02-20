package models


type Employee struct{
		ID        int32 `bson:"_id,omitempty" json:"id,omitempty"`
		Firstname string   `bson:"firstname" json:"firstname"`
		Lastname  string   `bson:"lastname" json:"lastname"`
		Salary    int32    `bson:"salary" json:"salary"`
		Gender    string   `bson:"gender" json:"gender"`
}
var Employees []Employee