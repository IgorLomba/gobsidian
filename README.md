<p align="center">

<h3 align="center">Gobisidian CLI</h3>

<p align="center">
    Golang based CLI tool to download obsidian published sites. Supports parallel downloads.
</p>

<p align="center">
    Inspired by <a href="https://github.com/Saghetti0/obsidian-publish-downloader">obsidian-publish-downloader</a>
</p>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
   <li><a href="#compile">Compile</a></li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>

## Compile

To be able to build the code you should have:

* Go - You can download and install Go using this [link](https://golang.org/doc/install).

#### Windows

``` powershell
setx GOOS=windows 
setx GOARCH=amd64
go build -o gobisidian.exe .
```

#### Linux

``` bash
export GOARCH=amd64
export GOOS=linux
go build -o gobisidian .
```

#### Macintosh

``` bash
export GOOS=darwin 
export GOARCH=amd64
go build -o gobisidian-mac .
```

## Usage

``` bash
./gobisidian https://publish.obsidian.md/sitename directory
```
