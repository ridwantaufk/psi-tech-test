package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ManipulatedUser struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Email    string   `json:"email"`
	Age      int      `json:"age"`
	Phone    string   `json:"phone"`
	Cell     string   `json:"cell"`
	Picture  []string `json:"picture"`
}

func GetExternalUsers(c *gin.Context) {
	results := c.DefaultQuery("results", "10")
	page := c.DefaultQuery("page", "1")

	if _, err := strconv.Atoi(results); err != nil {
		results = "10"
	}
	if _, err := strconv.Atoi(page); err != nil {
		page = "1"
	}

	url := fmt.Sprintf("https://randomuser.me/api?results=%s&page=%s", results, page)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal fetch external api"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var raw map[string]interface{}
	
	if err := json.Unmarshal(body, &raw); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal parse response"})
		return
	}

	rawResults, ok := raw["results"].([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "format response tidak sesuai"})
		return
	}

	var manipulated []ManipulatedUser
	for _, item := range rawResults {
		u, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		mu := manipulateUser(u)
		manipulated = append(manipulated, mu)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    manipulated,
		"results": len(manipulated),
		"page":    page,
	})
}

func manipulateUser(u map[string]interface{}) ManipulatedUser {
	nameMap, _ := u["name"].(map[string]interface{})
	title, _ := nameMap["title"].(string)
	first, _ := nameMap["first"].(string)
	last, _ := nameMap["last"].(string)
	fullName := fmt.Sprintf("%s, %s %s", title, first, last)

	locMap, _ := u["location"].(map[string]interface{})
	streetMap, _ := locMap["street"].(map[string]interface{})
	streetNum := fmt.Sprintf("%v", streetMap["number"])
	streetName, _ := streetMap["name"].(string)
	city, _ := locMap["city"].(string)
	state, _ := locMap["state"].(string)
	country, _ := locMap["country"].(string)
	location := fmt.Sprintf("%s,%s, %s,%s, %s", streetNum, streetName, city, state, country)

	email, _ := u["email"].(string)

	registered, _ := u["registered"].(map[string]interface{})
	age := 0
	if ageFloat, ok := registered["age"].(float64); ok {
		age = int(ageFloat)
	}
	phone, _ := u["phone"].(string)
	cell, _ := u["cell"].(string)
	picMap, _ := u["picture"].(map[string]interface{})
	large, _ := picMap["large"].(string)
	medium, _ := picMap["medium"].(string)
	thumb, _ := picMap["thumbnail"].(string)


	return ManipulatedUser{
		Name:     fullName,
		Location: location,
		Email:    email,
		Age:      age,
		Phone:    phone,
		Cell:     cell,
		Picture:  []string{large, medium, thumb},
	}
}
