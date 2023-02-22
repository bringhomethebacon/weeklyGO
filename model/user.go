package model

type Student struct {
	ID                string `gorm:"primarykey;" json:"id"`
	TeacherID         string `json:"teacher_id"`
	StudentName       string `json:"student_name"` //用户名
	Phone             string `json:"phone"`
	Grade             string `json:"grade"`                                      //年级
	Category          string `json:"category"`                                   //类别
	Specialized       string `json:"specialized"`                                //专业
	ResearchDirection string `json:"research_direction"`                         //研究方向
	Gender            string `json:"gender"`                                     //性别
	NativePlace       string `json:"native_place"`                               //籍贯
	Password          string `gorm:"column:password;type:varchar(32);" json:"-"` // 密码
	Role              string `json:"role"`
}

type Teacher struct {
	ID          string `gorm:"primarykey;" json:"id"`
	TeacherName string `json:"teacher_name"`                               //用户名
	Password    string `gorm:"column:password;type:varchar(32);" json:"-"` // 密码
	Role        string `json:"role"`
}
