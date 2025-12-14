# GoSpeech

GoSpeech is a Go-based command-line speech processing tool that supports both
speech-to-text (ASR) and text-to-speech (TTS) using ONNX Runtime.

This project refactors and extends an open-source speech project into a
production-style CLI application, with macOS-native ONNX Runtime integration
and a pluggable architecture for multilingual speech synthesis.

---

## Features

- 🎙 Speech-to-Text (ASR) using Paraformer  
  - Supports Chinese and English speech recognition
- 🔊 Text-to-Speech (TTS)
  - Mandarin Chinese TTS using MeloTTS
  - Architecture prepared for English TTS backend (e.g. Piper)
- 🖥 CLI-first design with simple subcommands
- ⚙️ Native ONNX Runtime integration on macOS (CGO)
- 🧩 Clean and extensible project structure

---

## Prerequisites

- Go 1.20+
- macOS (Apple Silicon tested)
- ONNX Runtime (installed via Homebrew)

```bash
brew install onnxruntime
```

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/Frida7771/GoSpeech
cd go-speech
```

### 2. Download model files

Model files need to be downloaded from Hugging Face (requires git-lfs):

```bash
# Install git-lfs if not already installed
brew install git-lfs
git lfs install

# Clone model repository
git clone https://huggingface.co/getcharzp/go-speech ./temp_models

# Move files to correct locations
mv ./temp_models/lib ./lib
mv ./temp_models/melo_weights ./melo_weights
mv ./temp_models/paraformer_weights ./paraformer_weights

# Clean up
rm -rf ./temp_models
```

### 3. Build the CLI tool

```bash
cd cmd/go-speech
go build -o go-speech
```

---

## Usage


After building, run from the `cmd/go-speech` directory:

```bash
cd cmd/go-speech
./go-speech asr <wav-file>
./go-speech tts "<text>" [--out output.wav]
```


**Note:** Do not run `go run .` from the project root, as the root package is `package speech`, not `package main`. The main package is located in `cmd/go-speech/`.




## License

MIT License - see [LICENSE](LICENSE) file for details.

---

## Acknowledgments

This project is based on the open-source project
[getcharzp/go-speech](https://huggingface.co/getcharzp/go-speech),
with significant refactoring and CLI restructuring.

---


