<div align="center">
    <img src="build/appicon.png" align="center" width="256"/>

# Kuddle

</div>

Kuddle is a cross-platform application that enables you to connect to your S3-compatible buckets and upload images, or screenshots simply by copy-pasting them right into the application. This is incredibly handy for bloggers, or 
simply people who frequently uploads images to their buckets to share to their blog, or related.

## Installation

To install Kuddle, simply head to our [GitHub Releases](/releases) and download the latest binary for your 
specific platform.

## Demo
| Decryption                                                           | Upload Images                                                     |
|----------------------------------------------------------------------|-------------------------------------------------------------------|
| ![image](https://blog-images.mihou.pw/kd/WPaEeBivQe.png)             | ![upload](https://blog-images.mihou.pw/kd/jkad3RPxOg.png)         |
| Successful Upload                                                    | Recent Uploads                                                    |
| ![successful upload](https://blog-images.mihou.pw/kd/nnTrv4qKce.png) | ![recent uploads](https://blog-images.mihou.pw/kd/6fYiFt3iYW.png) |

## Compiling

For those who wants to compile, or contribute to the project. You  can compile Kuddle on your own machines by simply 
following the steps indicated for your platform below. Please note that we have not tested anything for MacOS.

### Linux

Requirements:
1. [`Golang ^1.21.0`](https://go.dev)
2. [`Node & NPM`](https://nodejs.org)
3. [`UPX`](https://upx.github.io/)
4. [`nfpm`](https://nfpm.goreleaser.com/)

To build Kuddle for yourself on Linux, simply run the following commands:
```shell
git clone https://github.com/Qucy-Studios/kuddle && 
cd kuddle &&
cd frontend && npm install &&
cd .. && go install &&
./build-linux
```

### Windows

Requirements:
1. [`Golang ^1.21.0`](https://go.dev)
2. [`Node & NPM`](https://nodejs.org)
3. [`UPX`](https://upx.github.io/)
4. [`NSIS`](https://wails.io/docs/guides/windows-installer/)

To build Kuddle for yourself on Windows, simply run the following commands:
```shell
git clone https://github.com/Qucy-Studios/kuddle && 
cd kuddle &&
cd frontend && npm install &&
cd .. && go install &&
wails build --upx --nsis
```

## License

Kuddle is a **permanently** an MIT-licensed tool. No one shall have the right to override, modify or change this 
license at any later point of time. It shall be free-to-use and source-available for anyone to see.
