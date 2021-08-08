package run

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

func RunExe(exe string, parms []string) ([]byte, []byte, error) {
	// fmt.Printf("======== parms: %v\n\n", parms)
	cmd := exec.Command(exe, parms...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return nil, nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return nil, nil, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return nil, nil, err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return nil, nil, err
	}

	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return nil, nil, err
	}

	if err := cmd.Wait(); err != nil {
		// fmt.Println("Wait: ", err.Error())
		return output, bytesErr, err
	}

	// fmt.Printf("stdout: %s", bytes)
	return output, bytesErr, nil
}

func RunExe2(name string, arg ...string) {
	// fmt.Printf("run, name=%v\n", name)
	cmd := exec.Command(name, arg...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}

	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}

	// fmt.Printf("stdout: %s", bytes)
}

func RunAndChangeDir(stdOutput chan string, errOutput chan string, dir string, exe string, params ...string) error {
	cmd := exec.Command(exe, params...)

	o, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	e, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if len(dir) > 0 {
		cmd.Dir = dir
	}
	err = cmd.Start()
	if err != nil {
		return err
	}

	go func(o io.ReadCloser) {
		readerO := bufio.NewReader(o)
		for {
			line, err2 := readerO.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			stdOutput <- line
		}
	}(o)

	go func(e io.ReadCloser) {
		readerE := bufio.NewReader(e)
		for {
			line, err2 := readerE.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			errOutput <- line
		}
	}(e)

	err = cmd.Wait()
	return err
}
