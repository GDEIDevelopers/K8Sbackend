package apputils

import (
	"fmt"
	"time"

	"github.com/GDEIDevelopers/K8Sbackend/config"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServerUtils struct {
	DB      *gorm.DB
	RedisDB *redis.Client
	Config  *config.Config
	Token   token.TokenGenerate
	Class   *Class
}

func newSQL(DBAddr, DBase, DBUser, DBPass string, isParseTime ...bool) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", DBUser, DBPass, DBAddr, DBase)
	if len(isParseTime) > 0 && isParseTime[0] {
		dsn += "&parseTime=True"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errhandle.Log.Fatal(err)
	}
	d, _ := db.DB()
	d.SetConnMaxIdleTime(120 * time.Second)
	d.SetMaxOpenConns(1000)
	d.SetMaxIdleConns(10)

	return db
}

func NewServerUtils(cfg *config.Config) *ServerUtils {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisTokenAddr,
		DB:   cfg.RedisTokenDB,
	})
	utils := &ServerUtils{
		DB:      newSQL(cfg.DatabaseAddr, cfg.DatabaseDB, cfg.DatabaseUser, cfg.DatabasePass),
		RedisDB: rdb,
		Config:  cfg,
		Token:   token.NewJWTAccessGenerate(rdb, jwt.SigningMethodHS256),
	}
	utils.Class = NewClass(utils)
	return utils
}
