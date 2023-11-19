package youtube

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

type Params struct {
	ID      string
	Type    string // video, playlist, channel
	Audio   bool
	FullUrl string
}

type Command struct {
	Name string
	Args []string
}

func Main(params Params) {
	fmt.Println("Download video from " + params.FullUrl)
	if params.Audio {
		fmt.Println("Download audio only")
	}

	parts := strings.Split(params.FullUrl, "/")
	params.ID = parts[len(parts)-1]

	if strings.Contains(params.FullUrl, "=") {
		parts = strings.Split(params.FullUrl, "=")
		params.ID = parts[len(parts)-1]
	}

	command := buildCommand(params)
	execCommand(command)
}

func execCommand(command Command) {
	cmd := exec.Command(command.Name, command.Args...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		panic(err)
	}
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

}

func buildCommand(params Params) Command {
	saveTo := downloadDirectory()
	fmt.Print(strings.Contains(params.FullUrl, "playlist"))
	fmt.Print("id", params.ID)
	if strings.Contains(params.FullUrl, "playlist") {
		saveTo = filepath.Join(saveTo, params.ID)
	}
	path := filepath.Join(saveTo, "%(title)s.%(ext)s")

	url := params.FullUrl
	args := []string{}
	cmd := ""

	if params.Audio {
		args = []string{
			"-f",
			"ba",
			url,
			"-x",
			"--audio-format",
			"mp3",
			"-o",
			path,
		}
		cmd = "-f ba " + url + " -x --audio-format mp3 -o " + "'" + path + "'"
	} else {
		qbest := "bv*+ba/b"
		q1080 := "bv*[height=1080]+ba"
		// q720 := "bv*[height=720]+ba"
		quality := q1080 + "/" + qbest

		args = []string{
			"-f",
			quality,
			url,
			"-S",
			"res,ext:mp4:m4a",
			"--recode",
			"mp4",
			"-o",
			path,
		}
		cmd = "-f '" + quality + "' " + url + " -S res,ext:mp4:m4a --recode mp4 -o " + "'" + path + "'"
	}

	command := "yt-dlp"
	fmt.Println("")
	fmt.Println("Command:")
	fmt.Print(command + " " + cmd)
	fmt.Println("\n")

	return Command{
		Name: "yt-dlp",
		Args: args,
	}
}

func downloadDirectory() string {
	os := runtime.GOOS
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username

	switch os {
	case "windows":
		return filepath.Join("C:", "Users", username, "Downloads")
	case "darwin":
		return filepath.Join("/Users", username, "Downloads")
	case "linux":
		return filepath.Join("/home", username, "Downloads")
	default:
		return filepath.Join("/home", username, "Downloads")
	}
}
