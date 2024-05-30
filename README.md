# port-ninja (PortKiller)

port-ninja is a powerful cross-platform command-line tool to terminate processes listening on specified ports. It helps developers and system administrators manage and free up ports effectively.

## Features

- Terminate processes by specifying port numbers
- Support for Linux, macOS, and Windows
- Simple and intuitive command-line interface


## Usage Notes for Each OS

### Linux

#### Step-by-Step Instructions

1. **Download the binary**
    - Navigate to the [releases page](https://github.com/tijo-thomaz/port-ninja/releases).
    - Download the `port_ninja_linux_amd64` binary.

2. **Make the binary executable**
    ```sh
    chmod +x port_ninja_linux_amd64
    ```

3. **Move the binary to a directory in your PATH**
    ```sh
    sudo mv port_ninja_linux_amd64 /usr/local/bin/port_ninja
    ```

4. **Using Port Ninja**
    - To kill a process running on port 8080:
        ```sh
        port_ninja 8080
        ```
    - To kill processes running on multiple ports (e.g., 8080, 9090):
        ```sh
        port_ninja 8080 9090
        ```

### macOS

#### Step-by-Step Instructions

1. **Download the binary**
    - Navigate to the [releases page](https://github.com/yourrepo/releases).
    - Download the `port_ninja_darwin_amd64` binary.

2. **Make the binary executable**
    ```sh
    chmod +x port_ninja_darwin_amd64
    ```

3. **Move the binary to a directory in your PATH**
    ```sh
    sudo mv port_ninja_darwin_amd64 /usr/local/bin/port_ninja
    ```

4. **Using Port Ninja**
    - To kill a process running on port 8080:
        ```sh
        port_ninja 8080
        ```
    - To kill processes running on multiple ports (e.g., 8080, 9090):
        ```sh
        port_ninja 8080 9090
        ```

### Windows

#### Step-by-Step Instructions

1. **Download the binary**
    - Navigate to the [releases page](https://github.com/tijo-thomaz/port-ninja/releases/).
    - Download the `port_ninja_windows_amd64.exe` binary.

2. **Move the binary to a directory in your PATH**
    - You can place the binary in a directory like `C:\Program Files\PortNinja\` and add this directory to your system PATH:
        - Right-click on 'This PC' or 'Computer' on the desktop or in File Explorer.
        - Click on 'Properties'.
        - Click on 'Advanced system settings'.
        - Click on 'Environment Variables'.
        - Under 'System variables', find the PATH variable, select it, and click 'Edit'.
        - Click 'New' and add `C:\Program Files\PortNinja\` to the list.
        - Click 'OK' to close all dialog boxes.

3. **Using Port Ninja**
    - To kill a process running on port 8080:
        ```sh
        port_ninja_windows_amd64.exe 8080
        ```
    - To kill processes running on multiple ports (e.g., 8080, 9090):
        ```sh
        port_ninja_windows_amd64.exe 8080 9090
        ```
