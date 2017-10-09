package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	row = 7
	col = 50
)

var (
	argStrength     = flag.Int("s", 1, "commit strength")
	argGrassFile    = flag.String("g", "grass.txt", "the file name of grass")
	argMessageFile  = flag.String("m", "message.txt", "the file name of messages")
	argContrastFile = flag.String("c", "contrast.txt", "the file name of contrast pattern")
)

func calculateStartDate() time.Time {
	date := time.Now().AddDate(0, 0, -(row * col))
	for {
		if date.Weekday() == time.Sunday {
			return date
		}
		date = date.AddDate(0, 0, -1)
	}
}

func getGrassData() ([][]byte, error) {
	grassFile, err := os.Open(*argGrassFile)
	if err != nil {
		return nil, fmt.Errorf("Coudln't read %s", *argGrassFile)
	}
	defer grassFile.Close()

	lines := make([][]byte, row)

	grassFileReader := bufio.NewReaderSize(grassFile, 128)
	for r := 0; r < row; r++ {
		line, _, err := grassFileReader.ReadLine()

		if len(line) < col {
			return nil, fmt.Errorf("Cols in the row[%d] in %s < %d", r, *argGrassFile, col)
		}

		if err == io.EOF {
			if r < (row - 1) {
				return nil, fmt.Errorf("Rows in %s < %d", *argGrassFile, row)
			}
			break
		} else {
			if err != nil {
				return nil, err
			}
			lines[r] = make([]byte, len(line))
			copy(lines[r], line)
		}
	}
	return lines, nil
}

func getMessages() ([]string, error) {
	messageFile, err := os.Open(*argMessageFile)
	if err != nil {
		return nil, fmt.Errorf("Coudln't read %s", *argMessageFile)
	}
	defer messageFile.Close()

	messages := make([]string, 0, 32)
	messageFileReader := bufio.NewReaderSize(messageFile, 128)
	for {
		line, _, err := messageFileReader.ReadLine()

		if err == io.EOF {
			break
		} else {
			if err != nil {
				return nil, err
			}
			message := make([]byte, len(line))
			copy(message, line)
			messages = append(messages, string(message))
		}
	}

	if len(messages) < 1 {
		return nil, fmt.Errorf("The message file %s must contain at least one message", *argMessageFile)
	}

	return messages, nil
}

func getContrastData(filename string) (map[interface{}]interface{}, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Couldn't get contrast data from %s", *argContrastFile)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		return nil, fmt.Errorf("The file %s is not yaml format", *argContrastFile)
	}
	return m, nil
}

func execGitCommit(commitDate time.Time, messages []string) error {
	message := messages[rand.Int()%len(messages)]

	content := fmt.Sprintf("# %s\n", message)
	file, err := os.OpenFile("README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	cmd := exec.Command("git", "add", "README.md")
	cmdMessage, err := cmd.CombinedOutput()

	dateString := fmt.Sprintf("%04d/%02d/%02d 00:00:00", commitDate.Year(), int(commitDate.Month()), commitDate.Day())

	commitMessage := fmt.Sprintf("'%s'", message)
	cmd = exec.Command("git", "commit", "-m", commitMessage, "README.md", "--date", dateString)
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	return nil
}

func kusa(lines [][]byte, messages []string, contrast map[interface{}]interface{}) error {
	commitDate := calculateStartDate()
	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			data := string(lines[r][c])
			c := contrast[data]
			if c != nil {
				if c_f64, ok := c.(float64); ok {
					count := int(float64(*argStrength) * c_f64)
					for s := 0; s < count; s++ {
						err := execGitCommit(commitDate, messages)
						if err != nil {
							return err
						}
					}
				}
			}
			commitDate = commitDate.AddDate(0, 0, 1)
		}
	}

	return nil
}

func main() {
	flag.Parse()

	lines, err := getGrassData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	messages, err := getMessages()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	contrast, err := getContrastData(*argContrastFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = kusa(lines, messages, contrast)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
