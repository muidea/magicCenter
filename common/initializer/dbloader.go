package initializer

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func loadDatabase(server, name, account, password, dbfile string) error {
	items := strings.Split(server, ":")
	if len(items) != 2 {
		msg := fmt.Sprintf("illegal database server address, address:%s", server)
		return errors.New(msg)
	}

	host := fmt.Sprintf("-h%s", items[0])
	port := fmt.Sprintf("-P%s", items[1])
	user := fmt.Sprintf("-u%s", account)
	pwd := fmt.Sprintf("-p%s", password)
	args := []string{
		"mysql",
		user,
		pwd,
		host,
		port,
		"<",
		dbfile,
	}

	param := strings.Join(args, " ")
	importer := exec.Command("/bin/sh", "-c", param)
	var stdError bytes.Buffer
	importer.Stderr = &stdError

	err := importer.Start()
	if err != nil {
		log.Printf("start import failed, err:%s", err.Error())
		return err
	}

	importer.Wait()
	if err != nil {
		log.Printf("wait import finish failed, err:%s", err.Error())
		return err
	}

	errVal := stdError.String()
	if len(errVal) > 0 {
		msg := fmt.Sprintf("load data failed, cmd:%s, errVal:%s", param, errVal)
		return errors.New(msg)
	}

	return nil
}
