package speech

import (
	"fmt"
	"sync"

	ort "github.com/yalue/onnxruntime_go"
)

// OnnxConfig 只负责一件事：
// 在当前机器上，初始化一个 CPU-only 的 ONNX Runtime 环境
type OnnxConfig struct {
	// 必填：ONNX Runtime 动态库路径
	// 例如：/opt/homebrew/lib/libonnxruntime.dylib
	OnnxRuntimeLibPath string

	// 可选：CPU 推理线程数
	// <=0 表示使用 ONNX Runtime 默认值
	NumThreads int

	// 运行时生成
	SessionOptions *ort.SessionOptions
}

var (
	initErr error
	once    sync.Once
)

// New 初始化 ONNX Runtime（全局只会执行一次）
func (cfg *OnnxConfig) New() error {
	if cfg.OnnxRuntimeLibPath == "" {
		return fmt.Errorf("OnnxRuntimeLibPath 不能为空")
	}

	// 全局初始化 ONNX Runtime（线程安全）
	once.Do(func() {
		ort.SetSharedLibraryPath(cfg.OnnxRuntimeLibPath)
		initErr = ort.InitializeEnvironment()
	})

	if initErr != nil {
		return fmt.Errorf("初始化 ONNX Runtime 失败: %w", initErr)
	}

	// 创建 SessionOptions
	opts, err := ort.NewSessionOptions()
	if err != nil {
		return fmt.Errorf("创建 SessionOptions 失败: %w", err)
	}

	// 设置 CPU 推理线程数（如果指定）
	if cfg.NumThreads > 0 {
		if err := opts.SetIntraOpNumThreads(cfg.NumThreads); err != nil {
			return fmt.Errorf("设置 NumThreads 失败: %w", err)
		}
	}

	cfg.SessionOptions = opts
	return nil
}
