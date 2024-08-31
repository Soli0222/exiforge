# Picture Arrange

[日本語版 README はこちら](README-ja.md)

This Go program extracts EXIF date information from photos and organizes them into folders by date. It supports various file formats such as JPG and DNG and can process all files within a specified directory.

## Features

- Extracts date information from photo EXIF data
- Organizes photos into folders by date
- Supports multiple file extensions (e.g., JPG, DNG, PNG)
- Displays processing status with a progress bar

For example, if your folder looks like this before organization:

```
/photos
    ├── IMG_0001.JPG //2024-08-30
    ├── IMG_0001.DNG //2024-08-30
    ├── IMG_0002.JPG //2024-08-30
    ├── IMG_0002.DNG //2024-08-30
    ├── IMG_0003.JPG //2024-09-01
    ├── IMG_0003.DNG //2024-09-01
    ├── DSC_1234.DNG //2024-08-31
    └── vacation_2024.JPG //2024-08-30
```

Using the following command will organize it like this:

```bash
./picture-arrange -dir="/photos"
```

```
/photos
    ├── JPG
    │   ├── 2024-08-30
    │   │   ├── IMG_0001.JPG
    │   │   ├── IMG_0002.JPG
    │   │   └── vacation_2024.JPG
    │   └── 2024-09-01
    │       └── IMG_0003.JPG
    └── DNG
        ├── 2024-08-30
        │   ├── IMG_0001.DNG
        │   └── IMG_0002.DNG
        ├── 2024-08-31
        │   └── DSC_1234.DNG
        └── 2024-09-01
            └── IMG_0003.DNG
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

- `-ext`: Specify file extensions to process, separated by commas (default is `JPG,DNG`).
- `-dir`: Specify the directory to process (default is the current directory).

### Examples

1. **Organize photos in the current directory with default extensions:**

   ```bash
   ./picture-arrange
   ```

2. **Organize photos in a specific directory and process PNG files:**

   ```bash
   ./picture-arrange -dir="/path/to/photos" -ext="PNG"
   ```

3. **Organize JPG and PNG files in the current directory:**

   ```bash
   ./picture-arrange -ext="JPG,PNG"
   ```

## How It Works

1. The program scans the specified directory for files matching the specified extensions.
2. It extracts the original date from each file's EXIF metadata.
3. Files are moved into folders by date (e.g., `2024-08-31`).
4. The progress is displayed with a progress bar.

## Dependencies

This project uses the following Go packages:

- [`github.com/rwcarlsen/goexif/exif`](https://pkg.go.dev/github.com/rwcarlsen/goexif) - Library for reading EXIF metadata from image files
- [`github.com/schollz/progressbar/v3`](https://pkg.go.dev/github.com/schollz/progressbar/v3) - Library for displaying a progress bar during file processing

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
