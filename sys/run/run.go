package run

import (
	"fmt"
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
