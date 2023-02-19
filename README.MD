# 🔢 dec2ascii 🔤

**dec2ascii** is a command-line tool for converting decimal numbers to ASCII text. It can read decimal numbers from the command-line arguments or from a file, and can write the output to the console or to a file. It's written in [Go](https://golang.org/) and can be compiled to run on various operating systems.  

## 🚀 Getting Started

### Prerequisites

To compile and run **dec2ascii**, you need to have [Go](https://golang.org/) installed on your system. You can download and install Go from the official website: https://golang.org/dl/

## 🐹 Installation

1. Clone this repository:
```bash
git clone https://github.com/your-username/dec2ascii.git
```
2. Navigate to the project directory:
```bash
cd dec2ascii
```

3. Compile the source code:
```bash
go build -o dec2ascii main.go
```

OR 

## 💾 Makefile

`dec2ascii` comes with a Makefile to simplify building and installing the program. Here are the available Makefile targets:

- `build` - Compiles the `dec2ascii` program.
- `install` - Compiles and installs the `dec2ascii` program.
- `clean` - Cleans up the build artifacts.

To build and install `dec2ascii`, simply run the following command:
```bash
 sudo make install
```

This will compile the program and install it to the system-wide `bin` directory (`/usr/local/bin` by default).

To customize the installation directory, modify the `INSTALL_DIR` variable in the Makefile.

Note that the `install` target requires root privileges, as it installs the executable to the system-wide `bin` directory.


### Usage
```bash
dec2ascii [-i input-file] [-o output-file] 
dec2ascii decimal-numbers
```
Here's an overview of the command-line arguments:
```bash
-i input-file: specify an input file to read decimal numbers from

-o output-file: specify an output file to write the converted text to

decimal-numbers...: one or more decimal numbers to convert to ASCII text
```
### If the `-i` flag is not specified, the program will read decimal numbers from the command-line arguments. If the `-o` flag is not specified, the program will print the converted text to the console.

<br>

# 💡 Examples

Here are some examples of how to use **dec2ascii**:

### Convert decimal numbers from the command-line
```bash
dec2ascii 72 101 108 108 111 32 87 111 114 108 100
```


This will print the ASCII text equivalent of the decimal numbers to the console:
```bash
Hello World
```


### Convert decimal numbers from a file and write the output to a file


This will read the decimal numbers from the `input.txt` file and write the converted text to the `output.txt` file.

```bash
dec2ascii -i input.txt -o output.txt
```
This will read the decimal numbers from the `input.txt` file and write the converted text to the `output.txt` file.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
