// Package youtube-dl is a CLI to use yt-dlp (https://github.com/yt-dlp/yt-dlp) easily.
//
// `get` command download a video, a playlist or a channel. Argument can be a URL or an YouTube ID, it will be detected automatically.
// With URL, it possible to use another website than YouTube, like DailyMotion or PeerTube for example, ID only accept YouTube ID (video, playlist or channel).
//
// Command `video` is a shortcut to download a video from URL or ID.
// Command `playlist` is a shortcut to download a playlist from URL or ID.
// Command `channel` is a shortcut to download a channel from URL or ID.
// Command `chapters` is to download a video with chapters splitted from URL or ID (flag `-d` will be ignored).
//
// Videos will be downloaded in current directory by default or in `Downloads` directory with flag `-d` or `--downloads`.
// You can execute command without argument, it will ask you to enter an URL or an ID.
//
// Options:
//
// - `-a` or `--audio`: download only audio
// - `-d` or `--downloads`: save downloaded video in Downloads folder. Default is current folder
//
// Examples:
//
//	dlp get <URL_OR_ID>
//	dlp get <URL_OR_ID> -a
//
// You can skip argument and enter URL or ID after command execution:
//
//	dlp get
//
// You can also use it with another website than YouTube:
//
//	dlp get https://www.dailymotion.com/video/x3lnp8j
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ewilan-riviere/dlp/pkg/youtube"
	"github.com/spf13/cobra"
)

var (
	version = "v0.0.24"
)

func main() {
	cmdChannel := createCommand(
		"channel [id]",
		"Download channel from URL or Youtube from ID",
		"Download channel from URL or Youtube from ID",
		"channel",
	)
	addFlags(cmdChannel)

	cmdChapters := createCommand(
		"chapters [id]",
		"Download video with chapters splitted from URL or Youtube from ID",
		"Download video with chapters splitted from URL or Youtube from ID",
		"chapters",
	)
	addFlags(cmdChapters)

	cmdGet := createCommand(
		"get [id]",
		"Download video from URL or Youtube from ID",
		"Download video from URL or Youtube from ID, it can be a video, a playlist or a channel",
		"auto",
	)
	addFlags(cmdGet)

	cmdPlaylist := createCommand(
		"playlist [id]",
		"Download playlist from URL or Youtube from ID",
		"Download playlist from URL or Youtube from ID",
		"playlist",
	)
	addFlags(cmdPlaylist)

	cmdVideo := createCommand(
		"video [id]",
		"Download video from URL or Youtube from ID",
		"Download video from URL or Youtube from ID",
		"video",
	)
	addFlags(cmdVideo)

	var rootCmd = &cobra.Command{
		Use:     "dlp",
		Version: version,
		Short:   "Download video from URL or Youtube from ID",
		Long:    "Download video from URL or Youtube from ID, it can be a video, a playlist or a channel",
	}
	rootCmd.AddCommand(cmdGet)
	rootCmd.AddCommand(cmdVideo)
	rootCmd.AddCommand(cmdPlaylist)
	rootCmd.AddCommand(cmdChannel)
	rootCmd.AddCommand(cmdChapters)
	rootCmd.Execute()
}

func createCommand(use string, short string, long string, origin string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			audio, _ := cmd.Flags().GetBool("audio")
			downloadsDir, _ := cmd.Flags().GetBool("downloads")
			id := ""

			if len(args) > 0 {
				id = args[0]
			}

			if id == "" {
				fmt.Println("")
				fmt.Println("Please enter a video URL or a Youtube ID")
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("URL or ID: ")
				text, _ := reader.ReadString('\n')
				id = strings.Trim(text, "\n")
			}

			url := ""
			isUrl := false
			if strings.Contains(id, "http") {
				url = id
				isUrl = true
			}

			if !isUrl {
				url = youtubeUrl(id, origin)
			}

			youtube.Main(youtube.Params{
				Audio:              audio,
				FullUrl:            url,
				SaveToDownloadsDir: downloadsDir,
				Type:               origin,
				Chapters:           origin == "chapters",
			})
		},
	}
}

func addFlags(cmdGet *cobra.Command) {
	cmdGet.Flags().BoolP("audio", "a", false, "To convert downloaded video to audio (works with video, playlist, channel)")
	cmdGet.Flags().BoolP("downloads", "d", false, "Save downloaded video in Downloads folder. Default is current folder")
}

func youtubeUrl(id string, origin string) string {
	const YoutubeBaseUrl = "https://www.youtube.com"
	const YoutubeTypeVideo = "watch?v="
	const YoutubeTypePlaylist = "playlist?list="
	const YoutubeTypeChannel = "/"

	if origin == "auto" {
		if strings.Contains(id, "@") {
			return YoutubeBaseUrl + "/" + YoutubeTypeChannel + id
		}

		length := len(id)
		if length > 15 {
			return YoutubeBaseUrl + "/" + YoutubeTypePlaylist + id
		}

		return YoutubeBaseUrl + "/" + YoutubeTypeVideo + id
	}

	if origin == "video" {
		return YoutubeBaseUrl + "/" + YoutubeTypeVideo + id
	}

	if origin == "playlist" {
		return YoutubeBaseUrl + "/" + YoutubeTypePlaylist + id
	}

	if origin == "channel" {
		return YoutubeBaseUrl + "/" + YoutubeTypeChannel + id
	}

	return YoutubeBaseUrl + "/" + YoutubeTypeVideo + id
}
