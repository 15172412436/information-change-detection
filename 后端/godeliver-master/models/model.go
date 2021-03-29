package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"godeliver/pkg/setting"
	"time"
)

var (
	DB *gorm.DB
)

type Model struct {
	UID       uint      `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// * 意思是在何时删除
	DeletedAt *time.Time `json:"deleted_at"`
}

// 初始化 Database
func SetUp() {
	var err error
	DB, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		// todo 记录log log
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(setting.DatabaseSetting.LogMode)
	//db.SetLogger(logging.GetLogger())
	// todo 设置更新处理的回调
	// 删除事件
	//DB.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

// deleteCallback will set `DeletedAt` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, _ := scope.FieldByName("DeletedAt")

		//if !scope.Search.Unscoped  {
		scope.Raw(fmt.Sprintf(
			"UPDATE %v SET %v=%v%v%v",
			scope.QuotedTableName(),
			scope.Quote(deletedAtField.DBName),
			scope.AddToVars(time.Now().Unix()),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
		//}
		//} else {
		//	scope.Raw(fmt.Sprintf(
		//		"DELETE FROM %v%v%v",
		//		scope.QuotedTableName(),
		//		addExtraSpaceIfExist(scope.CombinedConditionSql()),
		//		addExtraSpaceIfExist(extraOption),
		//	)).Exec()
		//}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

func closeDB() {
	defer DB.Close()
}
