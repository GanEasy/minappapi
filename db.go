package minappapi

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Fans 粉丝
type Fans struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"type:varchar(255);unique_index"` // 微信文章地址
	SubNum    int64  // 订阅次数 用户每提交一次+1
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Post 提交的url
type Post struct {
	ID               uint   `gorm:"primary_key"`
	Title            string `gorm:"type:varchar(1024);"`             // 微信文章地址
	URL              string `gorm:"type:varchar(1024);unique_index"` // 微信文章地址
	SubNum           int64  // 订阅人次 用户每提交一次+1
	FolNum           int64  // 当前关注人数 注，如果有人关注，每过8小时检查更新 没人关注不再推送
	ChapterFragments string `gorm:"type:text;"` // 章节片段
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
}

// Subscribe 粉丝
type Subscribe struct {
	ID        uint   `gorm:"primary_key"`
	FansID    uint   `sql:"index"`               //粉丝 ID
	OpenID    string `gorm:"type:varchar(255);"` //提交的openid
	PostID    uint   `sql:"index"`               //post ID
	FormID    string `gorm:"type:varchar(255);"` //订阅formID，一次订阅只能推送一次通知
	Push      bool   //是否推送
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"` //删除后不再推送
}

var db *gorm.DB

//DB 返回 *gorm.DB
func DB() *gorm.DB {
	if db == nil {

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.LogMode(false)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {

	sqlConnection := fmt.Sprintf(
		"host=%v user=%v port=%v dbname=%v sslmode=%v password=%v",
		config.Database.Host,
		config.Database.User,
		config.Database.Port,
		config.Database.Dbname,
		config.Database.Sslmode,
		config.Database.Password,
	)
	db, err := gorm.Open(config.Database.Type, sqlConnection)
	// db, err := gorm.Open("sqlite3", "notice.db")

	if err != nil {
		return nil, err
	}
	return db, nil
}
