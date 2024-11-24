package single_responsibility_principle

import "fmt"

type User struct {
	Name  string
	Email string
	Phone string
}

type UserService interface {
	AddUser(user User)
	NotifyUser(user User)
}

type UserServiceImpl struct {
	users []User
}

func (s *UserServiceImpl) AddUser(user User) {
	s.users = append(s.users, user)
}

func (s *UserServiceImpl) NotifyUser(user User) {
	fmt.Printf("Sending email to %s\n", user.Email)
}

//--------------------第二种实现方式--------------------

// 拆为两个服务：用户服务，通知服务
type UserService2 interface {
	AddUser(user User)
}
type NotificationService interface {
	NotifyUser(user User)
}

// 在用户服务中，组合一个通知服务
type UserServiceImpl2 struct {
	users    []User
	notifier NotificationService
}

func (s *UserServiceImpl2) AddUser(user User) {
	s.users = append(s.users, user)
}

type EmailNotifyService struct{}

func (e *EmailNotifyService) NotifyUser(user User) {
	fmt.Println("用邮件通知")
}

// 主函数
func Main_UserService2() {
	userService := &UserServiceImpl2{
		notifier: &EmailNotifyService{},
	}
	zhangsan := User{
		Name:  "张三",
		Email: "zhangsan@163.com",
		Phone: "123456789",
	}
	userService.AddUser(zhangsan)

	userService.notifier.NotifyUser(zhangsan)

}
