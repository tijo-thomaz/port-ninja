# build.ps1

# Create directories for storing built binaries
$directories = @(".build/release/linux", ".build/release/darwin", ".build/release/windows")
foreach ($directory in $directories) {
    if (-not (Test-Path $directory -PathType Container)) {
        New-Item -ItemType Directory -Path $directory | Out-Null
    }
}

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
    go build -o ".build/release/$($target.GOOS)/$($target.Output)" ./cmd/portkiller/main.go
}
