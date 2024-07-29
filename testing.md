# Testing

```bash
go build -o dlp && ./dlp
```

```bash
go test
```

```bash
rm -rf testing
mkdir testing
cd testing
../dlp video https://www.youtube.com/watch?v=dQw4w9WgXcQ
../dlp video -c https://www.youtube.com/watch?v=TLV2IqSIr44
../dlp playlist https://www.youtube.com/playlist?list=PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC
../dlp channel https://www.youtube.com/@ewilanriviere2804
cd ..
```

```bash
./dlp video https://www.youtube.com/watch?v=dQw4w9WgXcQ
./dlp video dQw4w9WgXcQ

./dlp playlist -d https://www.youtube.com/playlist?list=PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC
./dlp playlist -d PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC

./dlp channel -da https://www.youtube.com/@ewilanriviere2804
./dlp channel -da @ewilanriviere2804
```

```bash
git tag v0.0.23
git push origin v0.0.23
```

Click on "Request" on <https://pkg.go.dev/github.com/ewilan-riviere/dlp@v0.0.23>

```bash
go build -o dlp && ./dlp get
```
