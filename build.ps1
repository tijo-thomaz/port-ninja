# build.ps1

$targets = @(
    @{ GOOS = "linux"; GOARCH = "amd64"; Output = "port_ninja_linux_amd64" },
    @{ GOOS = "darwin"; GOARCH = "amd64"; Output = "port_ninja_darwin_amd64" },
    @{ GOOS = "windows"; GOARCH = "amd64"; Output = "port_ninja_windows_amd64.exe" }
)

foreach ($target in $targets) {
    $env:GOOS = $target.GOOS
    $env:GOARCH = $target.GOARCH
    
    $outputDir = "./build/release/$($target.GOOS)"
    if (-not (Test-Path -Path $outputDir)) {
        New-Item -ItemType Directory -Path $outputDir | Out-Null
    }

    go build -o "$outputDir/$($target.Output)" ./cmd/portkiller/main.go
}
