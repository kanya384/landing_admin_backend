package ftpfiles

import (
	"io"
	"os"
	"strings"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

type sftpFiles struct {
	client      *sftp.Client
	pathToFiles string
}

func NewSFtpFiles(adress string, login string, pass string, pathToFiles string) (FtpFiles, error) {
	config := &ssh.ClientConfig{
		User: login,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//Ciphers: []string{"3des-cbc", "aes256-cbc", "aes192-cbc", "aes128-cbc"},
	}
	conn, err := ssh.Dial("tcp", adress, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	client, err := sftp.NewClient(conn)
	if err != nil {
		panic("Failed to create client: " + err.Error())
	}
	return &sftpFiles{
		client:      client,
		pathToFiles: pathToFiles,
	}, nil
}

func (f *sftpFiles) GetFiles() (entries [][]string, err error) {
	files, err := f.client.ReadDir(f.pathToFiles)
	if err != nil {
		return
	}
	for _, file := range files {
		fileSftp, err := f.client.OpenFile(f.pathToFiles+"/"+file.Name(), os.O_RDONLY)
		if err != nil {
			return entries, err
		}
		fileBytes, err := io.ReadAll(fileSftp)
		if err != nil {
			return entries, err
		}
		fileString := string(fileBytes)
		strings := strings.Split(fileString, "\n")
		for index, stringItem := range strings {
			if index != 0 {
				entries = append(entries, []string{stringItem})
			}
		}
	}

	return
}

func (f *sftpFiles) CloseConnection() (err error) {

	// удаляем обработанные файлы
	files, err := f.client.ReadDir(f.pathToFiles)
	if err != nil {
		return
	}
	for _, file := range files {
		f.client.Remove(f.pathToFiles + "/" + file.Name())
	}

	//закрываем соединение
	err = f.client.Close()
	if err != nil {
		return
	}
	return nil
}
