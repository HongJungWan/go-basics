package main

import (
	"bufio"
	"fmt"
	"os"
)

const FILE_NAME string = "data.txt"

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	readDocs := bufio.NewReader(file)
	line, err := readDocs.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		return "", err
	}

	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	return err
}

func main() {
	line, err := ReadFile(FILE_NAME)
	if err != nil || line == "" {
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다:", err)
		} else {
			fmt.Println("파일이 비어있습니다. 새로 생성합니다.")
		}

		err := WriteFile(FILE_NAME, "기본 내용")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}

		line, err = ReadFile(FILE_NAME)
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}

	fmt.Println("파일 내용:", line)
}

/*
📌 에러 반환
- 예를 들어 ReadFile() 함수로 파일을 읽을 때 해당하는 파일이 없어 에러가 발생했다고 하자.
- 이럴 떄 프로그램이 강제 종료되는 것보다는 적절한 메시지를 출력하고 다른 파일을 읽거나 임시 파일을 생성한다면 훨씬 사용자 경험이 좋을 것이다.
*/
