package belajargolanggorm

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:Bismillah@123@tcp(127.0.0.1:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	sql := "INSERT INTO sample(id,name) values(?,?)"
	err := db.Exec(sql, "1", "Nur").Error
	assert.Nil(t, err)

	sql = "INSERT INTO sample(id,name) values(?,?)"
	err = db.Exec(sql, "2", "HUMI").Error
	assert.Nil(t, err)

	sql = "INSERT INTO sample(id,name) values(?,?)"
	err = db.Exec(sql, "3", "GIAA").Error
	assert.Nil(t, err)

	sql = "INSERT INTO sample(id,name) values(?,?)"
	err = db.Exec(sql, "4", "BBCA").Error
	assert.Nil(t, err)

}

type Sample struct {
	Id   int
	Name string
}

func TestRawSql(t *testing.T) {
	var sample Sample
	sql := "select id,name from sample where id = ?"
	err := db.Raw(sql, "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, sample.Id)

	var samples []Sample
	sql = "select id,name from sample"
	err = db.Raw(sql).Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))

}

func TestSqlRow(t *testing.T) { // cara sulit
	sql := "select id,name from sample"
	rows, err := db.Raw(sql).Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})

	}
	assert.Equal(t, 4, len(samples))

}

func TestScanRow(t *testing.T) { // cara mudah
	sql := "select id,name from sample"
	rows, err := db.Raw(sql).Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)

	}
	assert.Equal(t, 4, len(samples))

}

func TestCreateUser(t *testing.T) {
	user := User{
		ID:       "1",
		Password: "rahasia",
		Name: Name{
			FirstName:  "Nur",
			MiddleName: "Firdaus",
			LastName:   "Ramandani",
		},
		Information: "Column ignore",
	}

	resp := db.Create(&user)
	assert.Nil(t, resp.Error)
	assert.Equal(t, int64(1), resp.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "rahasia" + strconv.Itoa(i),
			Name: Name{
				FirstName:  "Jon",
				MiddleName: "doe" + strconv.Itoa(i),
				LastName:   strconv.Itoa(i),
			},
			Information: "Will be ignore",
		})
	}
	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))

}

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "10", Password: "rahasia", Name: Name{FirstName: "user 10"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "user 11"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "12", Password: "rahasia", Name: Name{FirstName: "user 12"}}).Error
		if err != nil {
			return err
		}
		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionRollback(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "user 13"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "user 11"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "user 13"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "14", Password: "rahasia", Name: Name{FirstName: "user 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}

	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)

	user = User{}

	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.ID)
}

func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := User{}

	err := db.First(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	// assert.Equal(t, "5", user.ID)
}

func TestQueryAllObjects(t *testing.T) {
	var users []User

	err := db.Find(&users, "id in ?", []string{"1", "2", "3"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

}

func TestQueryCondition(t *testing.T) {
	var users []User

	err := db.Where("firstFirstName like ?", "%user%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))

}

func TestSelectSpecificFields(t *testing.T) {
	var users []User
	err := db.Select("id", "firstFirstName").Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}
	assert.Equal(t, 14, len(users))

}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	err := db.Order("id asc, firstFirstName desc").Limit(5).Offset(5).Find(&users).Error
	assert.Nil(t, err)

}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	err := db.Model(&User{}).Select("id", "firstFirstName", "lastName").Find(&users).Error
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	user := User{}
	result := db.First(&user, "id = ?", "1")

	assert.Nil(t, result.Error)

	user.Name.FirstName = "Annajwa"
	user.Name.MiddleName = "Lyra"
	user.Name.LastName = ""
	user.Password = "new123"

	err := db.Save(&user).Error
	assert.Nil(t, err)
}

func TestSelectColumns(t *testing.T) {
	// cara 1
	err := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middleName": "",
		"lastName":   "Moore",
	}).Error
	assert.Nil(t, err)

	// cara 2
	err = db.Model(&User{}).Where("id = ?", "1").Update("password", "rubah lagi").Error
	assert.Nil(t, err)

	// cara 3
	err = db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName:  "Nurrr",
			MiddleName: "Firdausss",
		},
	}).Error

	assert.Nil(t, err)

}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 100000,
	}

	err := db.Save(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
	var user User

	err := db.Model(&User{}).Preload("Wallet").First(&user).Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestRetrieveRelationJoin(t *testing.T) {
	var user []User

	err := db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ? ", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, 1, len(user))
}
