/************************************************************
 * 描述：此文件提供文件的上传、撤销上传、删除服务
 *      文件的保存位置取决于环境变量STATIC_PATH（参考.env.example）
 * 作者：zhangshaos
 ***********************************************************/

package service

import (
    "errors"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strconv"
    "time"
)

type FileService struct {
    Name    string
    Size    int64
    Read    io.Reader
    Path    string
    tmpFile *os.File
}

// 输入上传文件名（不包括路径），文件大小和文件描述符，创建对象
// 如果输入filename为空，或者sz<0，函数会抛出一个panic
func NewFileService(filename string, sz int64, r io.ReadWriter) *FileService {
    switch {
    case filename == "":
        panic("文件名为空")
    case sz < 0:
        panic("文件大小为负")
    }
    return &FileService{
        Name: filename,
        Size: sz,
        Read: r,
    }
}

// 上传文件到指定dest（不包含文件名）处
// 如果dest为""，则返回error
// 如果dest路径不存在，则创造路径
// 如果目的文件存在，则直接覆盖
func (p *FileService) Upload(dest string) error {
    if dest == "" {
        return errors.New("目标路径为空")
    }
    if _, err := os.Stat(dest); err != nil && os.IsNotExist(err) {
        os.MkdirAll(dest, 0755)
    }
    p.Path = dest
    // 先创建临时文件
    prefix := strconv.FormatInt(time.Now().Unix(), 10)
    tmp, err := ioutil.TempFile(dest, prefix)
    //if tmp != nil {
    //    defer tmp.Close() 临时文件的关闭操作留到提交/回滚时候执行
    //}
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

func (p *FileService) UploadRollBack()  {
    if p.tmpFile != nil {
        p.tmpFile.Close()
        os.Remove(p.tmpFile.Name())
        p.tmpFile = nil
    }
    p.Path = ""
}

func (p *FileService) UploadCommit()  {
    if err := os.Rename(p.tmpFile.Name(), filepath.Join(p.Path, p.Name));
        err != nil {
        p.UploadCommit()
    } else {
        p.tmpFile.Close()
        p.tmpFile = nil
    }
}