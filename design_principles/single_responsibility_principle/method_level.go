package single_responsibility_principle

import "fmt"

// ---------------实现一-------------------------------------

// 操作类型枚举
const ( // 未知操作
	UpdateUserName = iota + 1 // 更新用户名
	UpdatePassword            // 更新密码
)

type IUserOperate1 interface {
	UpdateUserInfo(t int)
}

type UserOperateImpl1 struct{}

func (u *UserOperateImpl1) UpdateUserInfo(t int) {
	switch t {
	case UpdateUserName:
		fmt.Println("更新用户名")
	case UpdatePassword:
		fmt.Println("更新密码")
	}
}

// ---------------实现二-------------------------------------

type IUserOperate2 interface {
	UpdateUserName()
	UpdatePassword()
}

type UserOperateImpl2 struct{}

func (u *UserOperateImpl2) UpdateUserName() {
	fmt.Println("更新用户名")
}

func (u *UserOperateImpl2) UpdatePassword() {
	fmt.Println("更新密码")
}
