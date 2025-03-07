# Picture Arrange

[日本語版 README はこちら](README-ja.md)

This Go program extracts EXIF date information from photos and organizes them into folders by date. It supports various file formats such as JPG and ARW and can process all files within a specified directory.

## Features

- Extracts date information from photo EXIF data
- Organizes photos into folders by date
- Supports multiple file extensions (e.g., JPG, ARW, RAF)
- Displays processing status with a progress bar
- Optional organization by camera model information

For example, if your folder looks like this before organization:

```
/photos
    ├── DSC00001.JPG //2024-08-30
    ├── DSC00001.ARW //2024-08-30
    ├── DSC00002.JPG //2024-08-30
    ├── DSC00002.ARW //2024-08-30
    ├── DSC00003.JPG //2024-09-01
    ├── DSC00003.ARW //2024-09-01
    ├── DSC00004.ARW //2024-08-31
    └── DSCF0590.JPG //2024-08-30
```

Using the following command will organize it like this:

```bash
picture-arrange -d="/photos"
```

```
/photos
    ├── JPG
    │   ├── 2024-08-30
    │   │   ├── DSC00001.JPG
    │   │   ├── DSC00002.JPG
    │   │   └── DSCF0590.JPG
    │   └── 2024-09-01
    │       └── DSC00003.JPG
    └── ARW
        ├── 2024-08-30
        │   ├── DSC00001.ARW
        │   └── DSC00002.ARW
        ├── 2024-08-31
        │   └── DSC00004.ARW
        └── 2024-09-01
            └── DSC00003.ARW
```

When organizing by camera model, use the following command:

```bash
picture-arrange -d="/photos" -m
```

```
/photos
    ├── ILCE-7CM2
    │   ├── JPG
    │   │   ├── 2024-08-30
    │   │   │   ├── DSC00001.JPG
    │   │   │   └── DSC00002.JPG
    │   │   └── 2024-09-01
    │   │       └── DSC00003.JPG
    │   └── ARW
    │       ├── 2024-08-30
    │       │   ├── DSC00001.ARW
    │       │   └── DSC00002.ARW
    │       ├── 2024-08-31
    │       │   └── DSC00004.ARW
    │       └── 2024-09-01
    │           └── DSC00003.ARW
    └── X-T5
        └── JPG
            └── 2024-08-31
                └── DSCF0590.JPG
```

## Installation

You can run the program using the pre-built binary without needing to build the source code.

1. Download the latest binary for your OS from the [releases page](https://github.com/Soli0222/picture-arrange/releases).
2. Extract the downloaded binary and make it executable (if necessary).

   ```bash
   chmod +x picture-arrange
   ```

3. Move it to a directory in your PATH or use it directly as a command.

## Usage

Run the downloaded binary from the command line with the following options:

- `-e, --extensions`: Specify file extensions to process, separated by commas (default is `JPG,ARW`).
- `-d, --directory`: Specify the directory to process (default is the current directory).
- `-m, --model`: Organize files using camera model information.
- `version`: Display the program's version information.

### Examples

1. **Organize photos in the current directory with default extensions:**

   ```bash
   picture-arrange
   ```

2. **Organize photos in a specific directory and process RAF files:**

   ```bash
   picture-arrange -d="/path/to/photos" -e="RAF"
   ```

3. **Organize JPG and RAF files in the current directory:**

   ```bash
   picture-arrange -e="JPG,RAF"
   ```

4. **Organize photos by camera model:**

   ```bash
   picture-arrange -m
   ```

5. **Display version information:**

   ```bash
   picture-arrange version
   ```

## How It Works

1. The program scans the specified directory for files matching the specified extensions.
2. It extracts the original date from each file's EXIF metadata.
3. Files are organized based on the selected mode:
   - Normal mode: Files are moved into folders by extension and date (e.g., `JPG/2024-08-31`).
   - Model mode: Files are organized in a hierarchy of camera model, extension, and date (e.g., `ILCE-7CM2/JPG/2024-08-31`).
4. The progress is displayed with a progress bar.

## Dependencies

This project uses the following Go packages:

- [`github.com/rwcarlsen/goexif/exif`](https://pkg.go.dev/github.com/rwcarlsen/goexif) - Library for reading EXIF metadata from image files
- [`github.com/schollz/progressbar/v3`](https://pkg.go.dev/github.com/schollz/progressbar/v3) - Library for displaying a progress bar during file processing
- [`github.com/spf13/cobra`](https://pkg.go.dev/github.com/spf13/cobra) - Library for building command line interfaces

## License

This project is licensed under the MIT License. See the LICENSE file for details.
