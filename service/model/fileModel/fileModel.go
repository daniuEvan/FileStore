/**
 * @date: 2022/4/13
 * @desc: ...
 */

package fileModel

import "gorm.io/gorm"

type TblFile struct {
	gorm.Model
	FileSha1 string `gorm:"unique"`
	FileName string
	FileSize int
	FileAddr string
	Status   int
	Ext1     string
	Ext2     string
}
