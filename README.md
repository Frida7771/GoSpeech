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

### Speech-to-Text (ASR)

Recognize speech from a WAV file:

```bash
./go-speech asr <wav-file>
```

Example:

```bash
./go-speech asr ./audio.wav
# Output: Recognized text: 你好，这是一个测试
```

### Text-to-Speech (TTS)

Generate speech from text:

```bash
./go-speech tts "<text>" [--out output.wav]
```

Examples:

```bash
# Default output location (assets/output.wav)
./go-speech tts "你好，这是一个语音合成测试"

# Custom output location
./go-speech tts "2019年12月30日，中国人口突破14亿人" --out hello.wav
```

---

## Project Structure

```
go-speech/
├── asr/                    # Speech recognition module
│   └── paraformer/         # Paraformer ASR implementation
├── tts/                    # Text-to-speech module
│   └── melotts/            # MeloTTS implementation
├── cmd/
│   └── go-speech/          # CLI application
├── examples/               # Example code
├── onnx.go                # ONNX Runtime wrapper
├── melo_weights/           # TTS model files
├── paraformer_weights/     # ASR model files
└── lib/                    # ONNX Runtime libraries
```

---

## Configuration

The CLI tool uses hardcoded paths for ONNX Runtime and model files. To customize:

1. Edit `cmd/go-speech/main.go`
2. Update the `onnxRuntimePath` constant (default: `/opt/homebrew/lib/libonnxruntime.dylib`)
3. Adjust model paths in the config structs

---

## Development

### Running tests

```bash
cd examples
go test -v -run TestMeloTTS
go test -v -run TestParaformer
```

### Using as a library

```go
package main

import (
    "github.com/getcharzp/go-speech/asr/paraformer"
    "github.com/getcharzp/go-speech/tts/melotts"
)

// TTS example
ttsEngine, _ := melotts.NewEngine(melotts.DefaultConfig())
defer ttsEngine.Destroy()
wavData, _ := ttsEngine.SynthesizeToWav("Hello", 1.0)

// ASR example
asrEngine, _ := paraformer.NewEngine(paraformer.DefaultConfig())
defer asrEngine.Destroy()
text, _ := asrEngine.RecognizeFile("./audio.wav")
```

---

## Troubleshooting

### ONNX Runtime not found

Ensure ONNX Runtime is installed via Homebrew and the path is correct:

```bash
brew install onnxruntime
ls /opt/homebrew/lib/libonnxruntime.dylib
```

### Model files missing

Verify model files are downloaded and in the correct locations:

```bash
ls melo_weights/model.onnx
ls paraformer_weights/model.int8.onnx
```

### Memory issues

Model loading requires significant memory (recommended: 2GB+ available). Consider adjusting `EnableCpuMemArena` in the ONNX configuration.

---

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
