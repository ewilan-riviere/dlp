package youtube

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

type Params struct {
	ID                 string
	Type               string // auto, video, playlist, channel
	Audio              bool
	FullUrl            string
	SaveToDownloadsDir bool
	Chapters           bool
}

type Command struct {
	Name string
	Args []string
}

func Main(params Params) {
	if params.Type == "" {
		params.Type = "auto"
	}
	fmt.Println("Download video from " + params.FullUrl)
	if params.Audio {
		fmt.Println("Download audio only")
	}
	if params.Chapters {
		fmt.Println("Download with chapters")
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
	// print params with label
	fmt.Println("Params:")
	fmt.Println("ID: " + params.ID)
	fmt.Println("Type: " + params.Type)
	fmt.Println("Audio: " + fmt.Sprint(params.Audio))
	fmt.Println("FullUrl: " + params.FullUrl)
	fmt.Println("SaveToDownloadsDir: " + fmt.Sprint(params.SaveToDownloadsDir))
	fmt.Println("Chapters: " + fmt.Sprint(params.Chapters))
	fmt.Println("\n")
	isPlaylist := false
	isChannel := false
	if params.Type == "auto" {
		isPlaylist = strings.Contains(params.FullUrl, "playlist")
		isChannel = strings.Contains(params.FullUrl, "channel")
		if !isChannel {
			isChannel = strings.Contains(params.FullUrl, "@")
		}
	}

	if params.Type == "playlist" {
		isPlaylist = true

	}

	if params.Type == "channel" {
		isChannel = true
	}

	saveTo := currentDirectory()

	if params.SaveToDownloadsDir {
		saveTo = downloadDirectory()
	}

	if isPlaylist || isChannel {
		saveTo = filepath.Join(saveTo, params.ID)
	}

	path := filepath.Join(saveTo, "%(title)s.%(ext)s")
	if params.SaveToDownloadsDir {
		path = filepath.Join(saveTo, "%(playlist_index)s-%(title)s.%(ext)s")
	}
	if params.Chapters {
		path = filepath.Join(saveTo, params.ID, "%(title)s.%(ext)s")
	}

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
		}
		if params.Chapters {
			args = append(args, "--split-chapters")
		}
		args = append(args, "-o", path)

		cmd = "-f ba " + url + " -x --audio-format mp3 "
		if params.Chapters {
			cmd += "--split-chapters "
		}
		cmd += "-o " + "'" + path + "'"
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
		}
		if params.Chapters {
			args = append(args, "--split-chapters")
		}
		args = append(args, "-o", path)

		cmd = "-f '" + quality + "' " + url + " -S res,ext:mp4:m4a --recode mp4 "
		if params.Chapters {
			cmd += "--split-chapters "
		}
		cmd += "-o " + "'" + path + "'"
	}

	command := "yt-dlp"
	fmt.Println("\n")
	fmt.Println("Command:")
	fmt.Print(command + " " + cmd)
	fmt.Println("\n")

	return Command{
		Name: "yt-dlp",
		Args: args,
	}
}

func currentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return dir
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
