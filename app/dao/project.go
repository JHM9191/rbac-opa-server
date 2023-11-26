package dao

type Project struct {
	Id   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type User struct {
	Id   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type Role struct {
	Id   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type ProjectUserRole struct {
	ProjectId string `gorm:"column:project_id;primaryKey"`
	UserId    string `gorm:"column:user_id;primaryKey"`
	RoleId    string `gorm:"column:role_id;primaryKey"`
}

type Permission struct {
	Id   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type RolePermission struct {
	RoleId       string `gorm:"column:role_id;primaryKey"`
	PermissionId string `gorm:"column:permission_id;primaryKey"`
}
