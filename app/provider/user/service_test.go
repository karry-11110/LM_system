package user

import (
	"context"
	"github.com/karry-11110/LM_system/test"
	"github.com/karry-11110/ginx/framework/contract"
	"github.com/karry-11110/ginx/framework/provider/cache"
	"github.com/karry-11110/ginx/framework/provider/config"
	"github.com/karry-11110/ginx/framework/provider/log"
	"github.com/karry-11110/ginx/framework/provider/orm"
	"github.com/karry-11110/ginx/framework/provider/redis"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 测试正常的注册登录流程
func Test_UserRegisterLogin(t *testing.T) {
	container := test.InitBaseContainer()
	container.Bind(&config.GinxConfigProvider{})
	container.Bind(&log.GinxLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.GinxCacheProvider{})

	ormService := container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("truncate table users").Error; err != nil {
		t.Fatal(err)
	}

	tmp, err := NewUserService(container)
	if err != nil {
		t.Fatal(err)
	}
	us := tmp.(*UserService)
	ctx := context.Background()

	user1 := &User{
		UserName: "wangkun",
		Password: "123456",
		Email:    "wangkun110@gmail.com",
	}

	Convey("正常流程", t, func() {

		Convey("注册用户", func() {
			userWithToken, err := us.Register(ctx, user1)
			So(err, ShouldBeNil)
			user1.Token = userWithToken.Token
		})

		Convey("发送邮件", func() {
			err := us.SendRegisterMail(ctx, user1)
			So(err, ShouldBeNil)
		})

		Convey("验证注册信息", func() {
			isExist, err := us.VerifyRegister(ctx, user1.Token)
			So(err, ShouldBeNil)
			So(isExist, ShouldBeTrue)
			// 数据库有数据
			userDB := &User{}
			err = db.Where("username=?", user1.UserName).First(userDB).Error
			So(err, ShouldBeNil)
			So(userDB.ID, ShouldNotBeZeroValue)
		})

		Convey("用户登录", func() {
			userDB, err := us.Login(ctx, user1)
			So(err, ShouldBeNil)
			So(userDB, ShouldNotBeNil)
			user1.Token = userDB.Token
		})

		Convey("用户验证", func() {
			userSession, err := us.VerifyLogin(ctx, user1.Token)
			So(err, ShouldBeNil)
			So(userSession, ShouldNotBeNil)
		})

		Convey("用户登出", func() {
			err := us.Logout(ctx, user1)
			So(err, ShouldBeNil)
			//重新验证为失败
			userSession, err := us.VerifyLogin(ctx, user1.Token)
			So(err, ShouldNotBeNil)
			So(userSession, ShouldBeNil)
		})
	})
}
