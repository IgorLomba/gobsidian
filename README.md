<p align="center">

<h3 align="center">Gobsidian CLI</h3>

<p align="center">
    Golang based CLI tool to download obsidian publish sites. Supports parallel downloads.
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
    <li><a href="#output">Output</a></li>
  </ol>
</details>

## Compile

To be able to build the code you should have:

* Go - You can download and install Go using this [link](https://golang.org/doc/install).

#### Windows

``` powershell
setx GOOS=windows 
setx GOARCH=amd64
go build -o gobsidian.exe .
```

#### Linux

``` bash
export GOARCH=amd64
export GOOS=linux
go build -o gobsidian .
```

#### Macintosh

``` bash
export GOOS=darwin 
export GOARCH=amd64
go build -o gobsidian-mac .
```

## Usage

``` bash
./gobsidian https://publish.obsidian.md/addielamarr/ directory
```

## Output

Progress output:
``` bash
[##################################################] 100%
```

Directory tree:
``` bash
directory/
├── 00 Home MOC.md
├── 01 Cybersecurity Mastery.md
...
└── Zero Knowledge Proof-based cryptography.md
```
