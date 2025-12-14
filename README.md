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
git clone https://github.com/getcharzp/go-speech.git
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

### Option 1: Run the compiled binary

After building, run from the `cmd/go-speech` directory:

```bash
cd cmd/go-speech
./go-speech asr <wav-file>
./go-speech tts "<text>" [--out output.wav]
```

### Option 2: Run directly with `go run`

From the project root:

```bash
go run ./cmd/go-speech asr <wav-file>
go run ./cmd/go-speech tts "<text>" [--out output.wav]
```

Or from the `cmd/go-speech` directory:

```bash
cd cmd/go-speech
go run . asr <wav-file>
go run . tts "<text>" [--out output.wav]
```

**Note:** Do not run `go run .` from the project root, as the root package is `package speech`, not `package main`. The main package is located in `cmd/go-speech/`.




## License

MIT License - see [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [MeloTTS](https://github.com/myshell-ai/MeloTTS) - Text-to-speech model
- [Paraformer](https://github.com/alibaba-damo-academy/FunASR) - Speech recognition model
- [ONNX Runtime](https://github.com/microsoft/onnxruntime) - Model inference engine

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
