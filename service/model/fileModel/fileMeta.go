/**
 * @date: 2022/4/13
 * @desc: ...
 */

package fileModel

//
// FileMeta
// @Description: 文件元数据信息
//
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//
// UpdateFileMeta
// @Description: 新增/更新文件元信息
// @param fileMeta:
//
func UpdateFileMeta(fileMeta FileMeta) {
	fileMetas[fileMeta.FileSha1] = fileMeta
}

//
// GetFileMeta
// @Description: 通过sha1值获取文件的元信息对象
// @param fileSha1:
// @return FileMeta:
//
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//
// RemoveFile
// @Description: 删除元信息
// @param fileSha1:
//
func RemoveFile(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
