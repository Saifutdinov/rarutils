# 📦 RARUtils — Go Package for RAR/UnRAR Operations

A powerful Go package to handle `.rar` archives with the help of the `rar` and `unrar` utilities. This library allows you to create and extract RAR archives efficiently, as well as interact with their contents programmatically.

---

## 🚀 Features

- 📁 **Create RAR archives** with specified compression parameters.
- 🗂️ **Extract RAR archives** and access file information programmatically.
- 🛠️ Support for listing archive contents (`unrar l` and `unrar lt`).
- 🔄 **Cross-platform compatibility** (with installed `rar`/`unrar` utilities).
- ✅ Fully customizable via exported configuration.

---

## 📋 Prerequisites

Before using the package, ensure the `rar` and `unrar` utilities are installed on your system.  
You can check their availability by running:

```bash
rar --version
unrar --version
```

---

## 📦 Installation of the Package
```bash
go get github.com/Saifutdinov/rarutils
```
---

## RAR

### Usage

#### Create file and compress

```golang
archive := rar.NewArchive("my_archive.rar")

archive.AddFile("/path/to/file1.pdf")
archive.AddFile("/path/to/file2.pdf")

// rar a my_archive.rar @fileslist.txt
err := archive.Compress()

if err != nil {
	fmt.Printf("Failed to create archive: %v\n", err)
	return
}
fmt.Println("Archive created successfully!")
```
---

### Methods

The package allows customization of the following parameters through exported configuration functions:

#### Compression

Control the RAR compression using parameters like `-m5` (maximum compression).

```golang
archive.SetCompression(rar.CompressionLVL5)
```
---

#### Volumes

Split archive to selected size volumes. Will run as `-v10MB`

```golang
archive.SetVolumes("10MB")
```
---

#### Password

Protect your files by setting password. `-p secretpassword`

```golang
archive.SetPassord("secretpassword")
```
---

#### Solid archive

Make archive solid and compress it better. `-s`

```golang
archive.ToggleSolid(true)
```

#### Files
You can set up any type of files source 
```golang
archive.SetSourceDir("/path/to/directory")

archive.SetFilePattern("file*.pdf")

// creates @fileslist temp file and use it as list
archive.AddFile("/path/to/file1.pdf")
```
---

#### Stream (WIP)
Retuns you stream of file

```golang
stream, err := archive.Stream()
if err != nil {
	return
}
```

## UNRAR

### 📖 Usage

#### 1. Extracting Files
```golang
archive := unrar.NewArchive("/path/to/archive.rar")

// unrar x /path/to/archive.rar ./
files, err := archive.Extract()
if err != nil {
	fmt.Printf("Failed to extract archive: %v", err)
	return
}

fmt.Println("Archive extracted successfully!")
```

--
#### 2. Listing Archive Contents
```golang
archive := unrar.NewArchive("/path/to/archive.rar")

// unrar laadesw  /path/to/archive.rar
files, err := archive.List()
if err != nil {
	fmt.Printf("Failed list archive files: %v\n", err)
}

for _, file := range files {
	fmt.Printf("File: %s, Size: %d bytes", file.Name, file.Size)
}
```
--
### Methods

####  Destination
Specify extraction path to store your files locally

```golang
archive.SetDestination("/path/to/extract")
```

---
####  Password
Use password to extract files from archive. `-p{secretpassword}`

```golang
archive.SetPassword("secretpassword")
```
---

####  Overwrite mode
If you need to overwrite (or don't) you can toggle overwrite mode. `-o-` or `-o+`.
 
```golang
archive.SetOverwriteMode(true)
```
---

## ⚖️ License

> ❗ Important:
> This package is a wrapper around the `rar` and `unrar` utilities, which are proprietary software by Alexander Roshal.
> While this package is MIT-licensed, ensure compliance with the licensing terms of the RAR utilities when redistributing or deploying.

---
## 💬 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

---
## 🌟 Acknowledgements

Special thanks to the developers of `rar` and `unrar` utilities for their amazing tools!

---
