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

## 📖 Usage
### 1. Basic Example: Archiving Files
```golang
	err := rar.ArchiveFiles("my_archive.rar", []string{"file1.txt", "file2.txt"})
	if err != nil {
		log.Fatalf("Failed to create archive: %v", err)
	}
	log.Println("Archive created successfully!")
```
### 2. Extracting Files
```golang
archive, err := rar.OpenArchive("my_archive.rar")
if err != nil {
	log.Fatalf("Failed to open archive: %v", err)
}
defer archive.Close()

files, err := archive.Extract()
if err != nil {
	log.Fatalf("Failed to extract archive: %v", err)
}


```
### 3. Listing Archive Contents
```golang
files, err := rar.ListArchiveContents("my_archive.rar")
if err != nil {
	log.Fatalf("Failed to list archive contents: %v", err)
}

for _, file := range files {
	log.Printf("File: %s, Size: %d bytes", file.Name, file.Size)
}
```
## 🛠️ Configuration

The package allows customization of the following parameters through exported configuration functions:

- Compression Levels: Control the RAR compression using parameters like `-m5` (maximum compression).
    ```golang
    archive.SetCompression(rar.CompressionLVL)
    ```
- Custom Command Paths: If `rar` or `unrar` is not in your PATH, you can specify the binary location.
    ```golang
    rar.SetRarPath("/custom/path/to/rar")
    rar.SetUnrarPath("/custom/path/to/unrar")
    ```
---
## ⚖️ License

> ❗ Important:
This package is a wrapper around the `rar` and `unrar` utilities, which are proprietary software by Alexander Roshal.
While this package is MIT-licensed, ensure compliance with the licensing terms of the RAR utilities when redistributing or deploying.

---
## 💬 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

---
## 🌟 Acknowledgements

Special thanks to the developers of rar and unrar utilities for their amazing tools!

---
