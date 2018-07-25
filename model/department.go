package model

type Department struct {
	// id int(6) UN AI PK
	// name varchar(100)
	// description varchar(500)

	DepartmentID          int64  `db:"id" json:"id"`
	DepartmentName        string `db:"name" json:"name"`
	DepartmentDescription string `db:"description" json:"description"`
}
