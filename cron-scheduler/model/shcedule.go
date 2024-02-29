package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

var ScheduleDB Schedule

type Schedule struct {
	ID                   int    `gorm:"primary_key;auto_increment;comment:'主键'"`
	TaskID               int    `gorm:"default:0;not null;comment:'task的id'"`
	Cron                 string `gorm:"size:32;default:'';not null;comment:'cron表达式'"`
	Type                 string `gorm:"size:32;default:'';not null;comment:'任务类型'"`
	NextScheduleTaskTime int    `gorm:"default:0;not null;comment:'下次计划定份任务时间'"`
	Cuser                int    `gorm:"default:0;not null;comment:'创建用户uid'"`
	Ctime                int    `gorm:"default:0;not null;comment:'计划创建时间'"`
	Utime                int    `gorm:"default:0;not null;comment:'最近修改时间'"`
	Dtime                int    `gorm:"default:0;not null;comment:'删除时间'"`
}

func (*Schedule) TableName() string {
	return "backup_schedule"
}

var (
	appdb  *gorm.DB
	sqldb  *sql.DB
	dbOnce sync.Once
)

// Initdb 初始化数据库连接
func Initdb() error {
	dsn := "root:liuxiao123@tcp(127.0.0.1:3306)/schedule_task?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return err
	}

	appdb = db

	if !db.Migrator().HasTable(&Schedule{}) {
		err = db.AutoMigrate(&Schedule{})
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
			return err
		}
	}
	return nil
}

func GetDB() *gorm.DB {
	return appdb
}

func (*Schedule) GetById(id int) (*Schedule, error) {
	var bs Schedule
	if err := appdb.Where("id = ?", id).First(&bs).Error; err != nil {
		return nil, err
	}
	return &bs, nil
}

// GetSchedules 获取所有定时计划
func (*Schedule) GetSchedules() ([]*Schedule, error) {
	var bs []*Schedule
	if err := appdb.Where("dtime = 0").Order("id DESC").Find(&bs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*Schedule{}, nil
		}
		return nil, err
	}
	return bs, nil
}

// UpdateScheduleByMap 更新定时计划
func (*Schedule) UpdateScheduleByMap(id int, data map[string]interface{}) error {
	data["utime"] = time.Now().Unix()
	if err := appdb.Model(&Schedule{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (*Schedule) Insert(bs *Schedule) error {
	if err := appdb.Create(bs).Error; err != nil {
		return err
	}
	return nil
}
