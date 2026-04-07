package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ridwantaufk/psi-tech-test/config"
)

type UserCompanyResult struct {
	UserID      string  `json:"user_id"`
	CompanyID   *string `json:"company_id"`
	Nama        string  `json:"nama"`
	Email       *string `json:"email"`
	Telp        string  `json:"telp"`
	CompanyCode *string `json:"company_code"`
	CompanyName *string `json:"company_name"`
}

func GetUsers(c *gin.Context) {
	var results []UserCompanyResult

	query := `
		SELECT 
			u.id as user_id,
			c.id as company_id,
			u.nama,
			u.email,
			u.telp,
			c.company_code,
			c.company_name
		FROM users u
		LEFT JOIN companies c ON c.user_id = u.id
	`

	rows, err := config.DB.Raw(query).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r UserCompanyResult
		rows.Scan(
			&r.UserID,
			&r.CompanyID,
			&r.Nama,
			&r.Email,
			&r.Telp,
			&r.CompanyCode,
			&r.CompanyName,
		)
		results = append(results, r)
	}

	// 
	
	c.JSON(http.StatusOK, gin.H{
		"data":  results,
		"total": len(results),
	})
}
