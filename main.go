package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
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
	argStrength    = flag.Int("s", 1, "commit strength")
	argGrassFile   = flag.String("g", "grass.txt", "the file name of grass")
	argMessageFile = flag.String("m", "message.txt", "the file name of messages")
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

func main() {

	flag.Parse()

	grassFile, err := os.Open(*argGrassFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Coudln't read %s\n", *argGrassFile)
		os.Exit(1)
	}
	defer grassFile.Close()

	lines := make([][]byte, row)

	grassFileReader := bufio.NewReaderSize(grassFile, 128)
	for r := 0; r < row; r++ {
		line, _, err := grassFileReader.ReadLine()

		if len(line) < col {
			fmt.Fprintf(os.Stderr, "Cols in the row[%d] in %s < %d\n", r, *argGrassFile, col)
			os.Exit(1)
		}

		if err == io.EOF {
			if r < (row - 1) {
				fmt.Fprintf(os.Stderr, "Rows in %s < %d\n", *argGrassFile, row)
				os.Exit(1)
			}
			break
		} else {
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
				os.Exit(1)
			}
			lines[r] = make([]byte, len(line))
			copy(lines[r], line)
		}
	}

	messageFile, err := os.Open(*argMessageFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Coudln't read %s\n", *argMessageFile)
		os.Exit(1)
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
				fmt.Fprint(os.Stderr, err.Error())
				os.Exit(1)
			}
			message := make([]byte, len(line))
			copy(message, line)
			messages = append(messages, string(message))
		}
	}

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	commitDate := calculateStartDate()
	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			data := lines[r][c]
			if data != '0' {
				for s := 0; s < *argStrength; s++ {
					err := execGitCommit(commitDate, messages)
					if err != nil {
						fmt.Fprintf(os.Stderr, "%s\n", err.Error())
						os.Exit(1)
					}
				}
			}
			commitDate = commitDate.AddDate(0, 0, 1)
		}
	}
}
