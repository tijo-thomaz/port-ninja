# build.ps1

# Define build targets
$targets = @(
    @{ GOOS = "linux"; GOARCH = "amd64"; Output = "port_ninja_linux_amd64" },
    @{ GOOS = "darwin"; GOARCH = "amd64"; Output = "port_ninja_darwin_amd64" },
    @{ GOOS = "windows"; GOARCH = "amd64"; Output = "port_ninja_windows_amd64.exe" }
)

# Iterate over each target and build the application
foreach ($target in $targets) {
    $env:GOOS = $target.GOOS
    $env:GOARCH = $target.GOARCH
    go build -o $target.Output ./cmd/main.go
}
