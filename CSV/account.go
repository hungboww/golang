package CSV

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type AccountData interface {
	InsertDataAccount(rows [][]string)
}
type accountCSVData struct {
	connection *gorm.DB
}

func NewAccountCSVData(db *gorm.DB) AccountData {
	return &accountCSVData{
		connection: db,
	}
}

type valueAccount struct {
	ID          string `csv:"id"`
	FirstName   string `csv:"firstName"`
	LastName    string `csv:"last_name"`
	Email       string `csv:"email"`
	Password    string `csv:"string"`
	Role        string `csv:"role"`
	About       string `csv:"about"`
	Avatar      string `csv:"avatar"`
	PhoneNumber string `csv:"phone_number"`

	Address   string `csv:"address"`
	City      string `csv:"city"`
	Country   string `csv:"country"`
	Gender    string `csv:"gender"`
	Postcode  string `csv:"postcode"`
	Birthday  string `csv:"birthday"`
	CreatedAt string `csv:"created_at"`
	Updated   string `csv:"updated_at"`
}

func (csv *accountCSVData) InsertDataAccount(rows [][]string) {
	roleData := []valueAccount{}
	for _, value := range rows {
		roleData = append(roleData, valueAccount{
			ID:          value[0],
			FirstName:   value[1],
			LastName:    value[2],
			Email:       value[3],
			Password:    value[4],
			Role:        value[5],
			About:       value[6],
			Avatar:      value[7],
			PhoneNumber: value[8],
			Address:     value[9],
			City:        value[10],
			Country:     value[11],
			Gender:      value[12],
			Postcode:    value[13],
			Birthday:    value[14],
			CreatedAt:   value[15],
			Updated:     value[16],
		})
	}
	s := ""
	for i := 1; i < len(roleData); i++ {
		id, _ := strconv.Atoi(roleData[i].ID)
		post_code, _ := strconv.Atoi(roleData[i].Postcode)
		s += fmt.Sprintf("(%d, %s, %s, %s,%s, %s,%s, %s,%s, %s,%s, %s, %s,%d, %s,%s, %s),", id, roleData[i].FirstName, roleData[i].LastName, roleData[i].Email, roleData[i].Password, roleData[i].Role, roleData[i].About, roleData[i].Avatar, roleData[i].PhoneNumber, roleData[i].Address, roleData[i].City, roleData[i].Country, roleData[i].Gender, post_code, roleData[i].Birthday, roleData[i].CreatedAt, roleData[i].Updated)
	}
	sql := fmt.Sprintf(`INSERT INTO users(user_id,first_name,last_name, email,password,role , about, avatar, phone_number, address, city, country, gender, postcode, birthday, created_at, updated_at) values  %s`, strings.TrimSuffix(s, ","))
	err := csv.connection.Exec(sql)
	if err != nil {
		fmt.Println("Can't insert CSV file into database. Please try again!!!", err)
	}

}
