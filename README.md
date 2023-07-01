# Google Fonts Directory Downloader

> Simple Google Fonts TTF downloader

<img width="1090" alt="image" src="https://github.com/matisiekpl/google-fonts-directory/assets/21008961/cc705d7f-fa05-4c30-aa2e-e56277df801b">

## Usage
```bash
git clone https://github.com/matisiekpl/google-fonts-directory.git .
go build github.com/matisiekpl/google-fonts-directory
./google-fonts-directory
```

Also `fonts.json` will be created:
```json
[
 {
  "Family": "Zilla Slab",
  "Weight": 600,
  "Italic": true,
  "Filename": "Zilla Slab-600-normal.ttf"
 },
 {
  "Family": "Zilla Slab",
  "Weight": 700,
  "Italic": false,
  "Filename": "Zilla Slab-700-italic.ttf"
 },
 {
  "Family": "Zilla Slab",
  "Weight": 700,
  "Italic": true,
  "Filename": "Zilla Slab-700-normal.ttf"
 },
 {
  "Family": "Zilla Slab Highlight",
  "Weight": 400,
  "Italic": false,
  "Filename": "Zilla Slab Highlight-400-italic.ttf"
 }
]
```
## Motivation
I wanted to download all Google Fonts `.ttf`s but splited on `italic/normal` and `weight`.

## Resources
Protobuf file: `fonts.proto`
