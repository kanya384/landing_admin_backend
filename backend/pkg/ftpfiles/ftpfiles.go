package ftpfiles

import (
	"io/ioutil"
	"time"

	"github.com/jlaffaye/ftp"
)

type FtpFiles interface {
	GetFiles() (entries [][]string, err error)
	CloseConnection() (err error)
}

type ftpFiles struct {
	client *ftp.ServerConn
}

func NewFtpFiles(adress string, login string, pass string, pathToFiles string) (FtpFiles, error) {
	client, err := ftp.Dial(adress, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	err = client.Login(login, pass)
	if err != nil {
		return nil, err
	}
	err = client.ChangeDir(pathToFiles)
	if err != nil {
		return nil, err
	}

	return &ftpFiles{
		client: client,
	}, nil

}

func (f *ftpFiles) GetFiles() (entries [][]string, err error) {
	/*files, err := f.client.List("./")
	if err != nil {
		return
	}
	for _, file := range files {
		file, err := f.getFile(file.Name)
		if err != nil {
			return entries, err
		}
		entries = append(entries, file)
	}*/
	return entries, nil
}

func (f *ftpFiles) getFile(path string) (file []byte, err error) {
	r, err := f.client.Retr(path)
	if err != nil {
		return
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	return buf, nil
}

func (f *ftpFiles) CloseConnection() (err error) {
	err = f.client.Quit()
	if err != nil {
		return
	}
	return nil
}
