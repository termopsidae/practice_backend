package imitate

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"paractice/config"
	"paractice/database"
	"paractice/model"
	"strconv"
	"time"
)

func InitDataBase() {
	allModels := []interface{}{
		&model.User{},
		&model.Manager{},
	}
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	sqlLog := logger.New(log.New(os.Stdout, "[SQL] ", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	fmt.Println(dsn)
	if database.DB, err = gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true, // 开启自动更新UpdatedAt字段
			Logger:                                   sqlLog,
		}); err != nil {
		panic("failed to connect database")
	}
	//删表
	var tableList = []string{"currencies", "managers", "mining", "key_values"}
	err = dropExistsTable(tableList)
	if err != nil {
		log.Fatal(err)
	}
	//创表
	for _, m := range allModels {
		if !database.DB.Migrator().HasTable(m) {
			if err = database.DB.AutoMigrate(m); err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("xxx")
	err = database.DB.Transaction(func(tx *gorm.DB) error {

		err = insertImitateManager(tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func dropExistsTable(tables []string) error {
	for _, table := range tables {
		err := database.DB.Exec("DROP TABLE IF EXISTS " + table).Error
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
