# build.ps1

$targets = @(
    @{ GOOS = "linux"; GOARCH = "amd64"; Output = "kill_port_linux_amd64" },
    @{ GOOS = "darwin"; GOARCH = "amd64"; Output = "kill_port_darwin_amd64" },
    @{ GOOS = "windows"; GOARCH = "amd64"; Output = "kill_port_windows_amd64.exe" }
)

foreach ($target in $targets) {
    $env:GOOS = $target.GOOS
    $env:GOARCH = $target.GOARCH
    go build -o $target.Output kill_port.go
}
