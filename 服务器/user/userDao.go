package user

import (
	"client/model"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//工厂模式
func NewUserDao(pool *redis.Pool) *UserDao {
	userDao := &UserDao{
		pool: pool,
	}
	return userDao
}

func (userDao *UserDao) GetUser(admin int, password string) (*model.User, error) {
	conn := userDao.pool.Get()
	defer func() {
		getUserConnClose := conn.Close()
		if getUserConnClose != nil {
			fmt.Printf("getUserConnClose：%v\n", getUserConnClose)
			return
		}
		fmt.Println("getUserConnClose suc")
	}()

	user, getUserIdErr := userDao.getUserId(conn, admin)
	if getUserIdErr != nil {
		fmt.Printf("getUserErr：%v\n", getUserIdErr)
		return nil, getUserIdErr
	}

	if password != user.Password {
		fmt.Printf("用户%v密码错误！\n", admin)
		return nil, ERROR_USER_PWD_ERROR
	}

	return user, nil
}

func (userDao *UserDao) AddUser(user *model.User) error {

	conn := userDao.pool.Get()
	defer func() {
		addUserConnClose := conn.Close()
		if addUserConnClose != nil {
			fmt.Printf("addUserConnClose：%v\n", addUserConnClose)
			return
		}
		fmt.Println("addUserConnClose suc")
	}()

	_, getUserIdErr := userDao.getUserId(conn, user.Admin)
	if getUserIdErr == nil {
		fmt.Printf("用户%v已存在！\n", user.Admin)
		return ERROR_USER_EXISTS
	} else {
		userJson, userJsonErr := json.Marshal(user)
		if userJsonErr != nil {
			fmt.Printf("userJsonErr：%v\n", userJsonErr)
			return userJsonErr
		}
		_, addUserIdErr := conn.Do("hset", "user", user.Admin, string(userJson))
		if addUserIdErr != nil {
			fmt.Printf("addUserIdErr：%v\n", addUserIdErr)
			return addUserIdErr
		}
	}
	return nil
}

func (userDao *UserDao) getUserId(conn redis.Conn, admin int) (*model.User, error) {
	res, getUserIdErr := redis.String(conn.Do("hget", "user", admin))
	if getUserIdErr != nil {
		if getUserIdErr == redis.ErrNil {
			fmt.Printf("%v不存在！\n", admin)
			return nil, ERROR_USER_NOTEXISTS
		}
		fmt.Printf("getUserIdErr:%v\n", getUserIdErr)
		return nil, getUserIdErr
	}

	user := &model.User{}
	userUnmarshalErr := json.Unmarshal([]byte(res), &user)
	if userUnmarshalErr != nil {
		fmt.Printf("userUnmarshalErr：%v\n", userUnmarshalErr)
		return nil, userUnmarshalErr
	}

	return user, nil
}
