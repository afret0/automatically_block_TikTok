package user

var role *RoleList

func init() {
	role = new(RoleList)
	role.Waiter = "waiter"
	role.Customer = "customer"
	role.Boss = "boss"
}

type User struct {
	Id           string
	Name         string
	Phone        string
	WXName       string
	Role         string
	RegisterTime float64
	UpdateTime   float64
}

type RoleList struct {
	Boss     string
	Waiter   string
	Customer string
}
