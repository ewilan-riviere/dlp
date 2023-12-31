# dlp

[![go][go-version-src]][go-version-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

CLI to use [`yt-dlp`](https://github.com/yt-dlp/dlp) easily. `yt-dlp` have many options and it can be hard to do simple things. This CLI is here to help you.

## Install

```bash
go install github.com/ewilan-riviere/dlp@latest
```

Dependencies:

- [`yt-dlp`](https://github.com/yt-dlp/yt-dlp): min 2023.11.16
- [`ffmpeg`](https://github.com/FFmpeg/FFmpeg): min v6.0

## Usage

### Get

`get` command download a video, a playlist or a channel. Argument can be a URL or an YouTube ID, it will be detected automatically. With URL, it possible to use another website than YouTube, like DailyMotion or PeerTube for example, ID only accept YouTube ID (video, playlist or channel).

- Video quality is set to `best` with maximum to `1080p`
- Videos will be downloaded in `Downloads` folder
- You can execute command without argument, it will ask you to enter an URL or an ID.

Options:

- `-a` or `--audio`: download only audio

```bash
dlp get <URL_OR_ID>
```

or

```bash
dlp get
```

Will ask:

```bash
Please enter a video URL or a Youtube ID
URL or ID:
```

You can use `--audio` option to download only audio:

```bash
dlp get -a <URL_OR_ID> # `dlp get -a` works too
```

#### Example

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
