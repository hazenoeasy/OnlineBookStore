/************************************************************
 * 描述：此文件提供文件的上传、撤销上传、删除服务
 *      文件的保存位置取决于环境变量STATIC_PATH（参考.env.example）
 * 作者：zhangshaos
 * TODO：目前这个服务还没构思好，文件上传请参考seller_addbook_service.go
 ***********************************************************/

package service

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileService struct {
	Name    string
	Size    int64
	Read    io.Reader
	path    string
	tmpFile *os.File
}

// 输入上传文件名dest（包括路径），文件大小sz和文件描述符r，创建对象
// 如果输入dest为空，或者sz<0，函数会抛出一个panic
func NewFileService(dest string, sz int64, r io.Reader) *FileService {
	if dest == "" {
		panic("文件名为空")
	}
	if sz < 0 {
		panic("文件大小为负")
	}
	return &FileService{
		Name: dest,
		Size: sz,
		Read: r,
		path: filepath.Dir(dest),
	}
}

// 上传文件
// 如果dest路径不存在，则创造路径
// 如果目的文件存在，则直接覆盖
func (p *FileService) Upload() error {
	if _, err := os.Stat(p.path); err != nil && os.IsNotExist(err) {
		os.MkdirAll(p.path, 0755)
	}
	// 先创建临时文件
	// 临时文件的关闭操作留到提交/回滚时候执行
	tmp, err := ioutil.TempFile(p.path, "tmp")
	if tmp != nil {
		defer tmp.Close()
	}
	if err != nil {
		return err
	}
	p.tmpFile = tmp
	// 从reader中读取数据，写入临时文件中
	buf := make([]byte, p.Size)
	if _, err := p.Read.Read(buf); err != nil {
		return err
	}
	if _, err := p.tmpFile.Write(buf); err != nil {
		return err
	}
	return nil
}

func (p *FileService) UploadRollBack() {
	if p.tmpFile != nil {
		os.Remove(p.tmpFile.Name())
		p.tmpFile = nil
	}
}

func (p *FileService) UploadCommit() {
	if err := os.Rename(p.tmpFile.Name(), p.Name); err != nil {
		p.UploadRollBack()
	} else {
		p.tmpFile = nil
	}
}
