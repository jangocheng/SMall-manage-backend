package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Menu_classify struct {
	M_Id       int       `json:"m_id"`
	M_name     string    `json:"m_name"`
	M_status   string    `json:"m_status"`
	M_other    string    `json:"m_other"`
	M_datetime time.Time `json:"m_datetime"`
}

func GetMenu_classifysByPage(c *gin.Context) {
	var (
		menu_classify  Menu_classify
		menu_classifys []Menu_classify
		count          int64
	)
	page := c.Param("pageid")
	rows, err := db.Query("select m_id,m_other,m_datetime,m_name,m_status from menu_classify;")
	if err != nil {
		fmt.Print(err.Error())
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err)
	}
	pagestart := (pageInt-1)*10 + 1
	pageend := pageInt * 10
	count = 0
	for rows.Next() {
		count++
		if count >= int64(pagestart) && count <= int64(pageend) {
			err = rows.Scan(&menu_classify.M_Id, &menu_classify.M_other, &menu_classify.M_datetime, &menu_classify.M_name, &menu_classify.M_status)
			menu_classifys = append(menu_classifys, menu_classify)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": menu_classifys,
		"count":  len(menu_classifys),
	})
}

func GetCountMenu_classifys(c *gin.Context) {
	var count int64
	err := db.QueryRow("select count(*) from menu_classify").Scan(&count)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})

}
func GetMenu_classify(c *gin.Context) {
	var (
		menu_classify Menu_classify
		result        gin.H
	)
	id := c.Param("id")
	row := db.QueryRow("select m_id,m_other,m_datetime,m_name,m_status from menu_classify where m_id = ?;", id)
	err := row.Scan(&menu_classify.M_Id, &menu_classify.M_other, &menu_classify.M_datetime, &menu_classify.M_name, &menu_classify.M_status)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": menu_classify,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}
func GetMenu_classifys(c *gin.Context) {
	var (
		menu_classify  Menu_classify
		menu_classifys []Menu_classify
	)
	rows, err := db.Query("select m_id,m_other,m_datetime,m_name,m_status from menu_classify;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&menu_classify.M_Id, &menu_classify.M_other, &menu_classify.M_datetime, &menu_classify.M_name, &menu_classify.M_status)
		menu_classifys = append(menu_classifys, menu_classify)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": menu_classifys,
		"count":  len(menu_classifys),
	})

}

func PostMenu_classify(c *gin.Context) {
	var buffer bytes.Buffer
	m_datetime := c.PostForm("m_datetime")
	m_name := c.PostForm("m_name")
	m_status := c.PostForm("m_status")
	m_other := c.PostForm("m_other")

	stmt, err := db.Prepare("insert into menu_classify (m_datetime,m_name,m_status,m_other) values(?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(m_datetime, m_name, m_status, m_other)

	if err != nil {
		fmt.Print(err.Error())
	}

	buffer.WriteString(m_datetime)
	buffer.WriteString("  ")
	buffer.WriteString(m_name)
	buffer.WriteString("  ")
	buffer.WriteString(m_status)
	buffer.WriteString("  ")
	buffer.WriteString(m_other)

	defer stmt.Close()
	_name := buffer.String()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", _name),
	})
}
func PutMenu_classify(c *gin.Context) {
	var buffer bytes.Buffer
	id := c.Param("id")
	m_datetime := c.PostForm("m_datetime")
	m_name := c.PostForm("m_name")
	m_status := c.PostForm("m_status")
	m_other := c.PostForm("m_other")

	stmt, err := db.Prepare("update menu_classify set m_datetime=?,m_name=?,m_status=?,m_other=? where m_id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(m_datetime, m_name, m_status, m_other, id)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Fastest way to append strings
	buffer.WriteString(m_datetime)
	buffer.WriteString("  ")
	buffer.WriteString(m_name)
	buffer.WriteString("  ")
	buffer.WriteString(m_status)
	buffer.WriteString("  ")

	defer stmt.Close()
	_name := buffer.String()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated to %s", _name),
	})

}
func DeleteMenu_classify(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Prepare("delete from menu_classify where m_id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted user: %s", id),
	})

}
