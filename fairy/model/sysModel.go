package model

type SysUser struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SysMenu struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	ParentId int         `json:"parent_id"`
	Sort     int         `json:"sort"`
	Path     string      `json:"path"`
	Type     int         `json:"type"`
	Icon     string      `json:"icon"`
	Role     string      `json:"role"`
	Children []*SysMenu  `gorm:"foreignKey:ParentId" json:"children"`
}

func (u SysUser) TableName() string {
	return "sys_user"
}
func (m SysMenu) TableName() string {
	return "sys_menu"
}