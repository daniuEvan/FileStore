/*
 * @date: 2021/12/16
 * @desc: ...
 */

package userModel

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(25);not null"`
	Mobile   string `gorm:"type:varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}
