# dlp

[![go][go-version-src]][go-version-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

CLI to use [`yt-dlp`](https://github.com/yt-dlp/yt-dlp) easily. `yt-dlp` have many options and it can be hard to do simple things. This CLI is here to help you.

> [!NOTE]
> This package use informations from tutorials
>
> - <https://www.linuxfordevices.com/tutorials/linux/yt-dlp-download-youtube-videos>
> - <https://www.linuxtricks.fr/wiki/yt-dlp-telecharger-des-videos-sur-internet-youtube-dl>
> - <https://write.corbpie.com/downloading-youtube-videos-and-playlists-with-yt-dlp/>

## Install

```bash
go install github.com/ewilan-riviere/dlp@latest
```

Dependencies:

- [`yt-dlp`](https://github.com/yt-dlp/yt-dlp): min 2023.11.16
- [`ffmpeg`](https://github.com/FFmpeg/FFmpeg): min v6.0

### Update

Update to the latest version

```bash
go install github.com/ewilan-riviere/dlp@latest
```

## Usage

- `channel`: download a channel
- `chapters`: download a video with chapters (don't work with `downloads` flag)
- `get`: download a video, a playlist or a channel
- `playlist`: download a playlist
- `video`: download a video

Parameter can be set directly:

```bash
dlp video https://www.youtube.com/watch?v=dQw4w9WgXcQ
dlp video dQw4w9WgXcQ # only for YouTube
```

Or after the command:

```bash
dlp video

Please enter a video URL or a Youtube ID
URL or ID: dQw4w9WgXcQ
```

`get` command will download a video, a playlist or a channel, it will guess from the URL or the ID. If you want to use a direct command, you can use `video`, `chapters`, `playlist` or `channel` command.

Argument can be a URL or an YouTube ID, it will be detected automatically. With URL, it possible to use another website than YouTube, like DailyMotion or PeerTube for example, ID only accept YouTube ID (video, playlist or channel).

> [!NOTE]
> Playlist and channel videos will be downloaded in a folder with the name of the playlist or the channel. And index will be added to the name of the video.

- Video quality is set to `best` with maximum to `1080p`
- Videos will be downloaded in `Downloads` folder
- You can execute command without argument, it will ask you to enter an URL or an ID.

Options:

- `-a` or `--audio`: download only audio
- `-d` or `--downloads`: save videos to Downloads folder, default is current folder (don't work with `chapters` command)
- `-c` or `--cookies`: add cookies option to download via cookies website (value is cookies path)

### Example

Here URL used are from YouTube but it works with other websites.

Video full URL

```bash
dlp get https://www.youtube.com/watch?v=dQw4w9WgXcQ
```

ID (only for YouTube)

```bash
dlp get dQw4w9WgXcQ
```

Short command

```bash
dlp get
```

Will ask:

```bash
Please enter a video URL or a Youtube ID
URL or ID: dQw4w9WgXcQ
```

Playlist full URL

```bash
dlp get https://www.youtube.com/playlist?list=PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC
```

ID (only for YouTube)

```bash
dlp get PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC
```

```bash
dlp get
```

Will ask:

```bash
Please enter a video URL or a Youtube ID
URL or ID: PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC
```

Channel full URL

```bash
dlp get https://www.youtube.com/@FranceInter
```

ID (only for YouTube)

```bash
dlp get @FranceInter
```

```bash
dlp get
```

Will ask:

```bash
Please enter a video URL or a Youtube ID
URL or ID: @FranceInter
```

## Roadmap

- [ ] add quality option
- [ ] add format option

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[tests-src]: https://img.shields.io/github/actions/workflow/status/ewilan-riviere/dlp/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://github.com/ewilan-riviere/dlp/actions
[license-src]: https://img.shields.io/github/license/ewilan-riviere/dlp.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/dlp/blob/main/LICENSE
