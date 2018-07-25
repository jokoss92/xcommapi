package controller

import (
	"fmt"
	"strconv"
	"xcommapi/config"
	"xcommapi/model"

	"github.com/gin-gonic/gin"
)

func GetAllDepartment(c *gin.Context) {
	// initialize the DbMap
	dbmap := config.InitDb()
	defer dbmap.Db.Close()

	// fetch all data department
	var departments []model.Department
	_, err := dbmap.Select(&departments, "select * from tbl_m_department order by id")

	if err == nil {
		c.JSON(200, departments)
	} else {
		c.JSON(404, gin.H{"error": "no instruction(s) into the table"})
	}
}

func GetDepartmentByID(c *gin.Context) {
	id := c.Params.ByName("id")

	// initialize the DbMap
	dbmap := config.InitDb()
	defer dbmap.Db.Close()

	var department model.Department

	err := dbmap.SelectOne(&department, "SELECT * FROM tbl_m_department WHERE id=?", id)

	if err == nil {
		DepartmentID, _ := strconv.ParseInt(id, 0, 64)

		content := &model.Department{
			DepartmentID:          DepartmentID,
			DepartmentName:        department.DepartmentName,
			DepartmentDescription: department.DepartmentDescription,
		}

		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "no instruction to get data"})
	}
}

func CreateNewDepartment(c *gin.Context) {
	// initialize the DbMap
	dbmap := config.InitDb()
	defer dbmap.Db.Close()

	var department model.Department

	department.DepartmentName = c.PostForm("DepartmentName")
	department.DepartmentDescription = c.PostForm("DepartmentDescription")

	if department.DepartmentName != "" && department.DepartmentDescription != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO tbl_m_department (name, description) VALUES (?, ?)`, department.DepartmentName, department.DepartmentDescription); insert != nil {
			DepartmentID, err := insert.LastInsertId()

			if err == nil {
				content := &model.Department{
					DepartmentID:          DepartmentID,
					DepartmentName:        department.DepartmentName,
					DepartmentDescription: department.DepartmentDescription,
				}

				c.JSON(201, content)
			} else {
				config.CheckErr(err, "Insert failed")
			}
		}
	} else {
		fmt.Println("99")
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

func UpdateExistingDepartment(c *gin.Context) {
	// initialize the DbMap
	dbmap := config.InitDb()
	defer dbmap.Db.Close()

	var department model.Department

	department.DepartmentID, _ = strconv.ParseInt(c.PostForm("DepartmentID"), 0, 64)
	department.DepartmentName = c.PostForm("DepartmentName")
	department.DepartmentDescription = c.PostForm("DepartmentDescription")

	if department.DepartmentName != "" && department.DepartmentDescription != "" {
		_, err := dbmap.Update(&department)

		if err == nil {
			c.JSON(200, department)
		} else {
			config.CheckErr(err, "Updated failed")
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

func DeleteExistingDepartment(c *gin.Context) {
	// initialize the DbMap
	dbmap := config.InitDb()
	defer dbmap.Db.Close()

	id := c.Params.ByName("id")
	var department model.Department
	err := dbmap.SelectOne(&department, "SELECT id FROM tbl_m_department WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&department)
		if err == nil {
			c.JSON(200, gin.H{"id #" + id: " deleted"})
		} else {
			config.CheckErr(err, "Delete failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "Department not found"})
	}
}
