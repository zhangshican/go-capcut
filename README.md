# GoCapcut

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhangshican/go-capcut)](https://goreportcard.com/report/github.com/zhangshican/go-capcut)

一个功能完整的Go语言库，用于创建和操作CapCut/剪映草稿文件。支持视频编辑的所有核心功能，包括片段管理、轨道系统、动画效果、特效滤镜等。

## 🚀 特性

- **完整的视频编辑支持**: 视频、音频、文本片段管理
- **轨道系统**: 支持多种轨道类型和渲染层级管理
- **动画系统**: 丰富的动画效果，支持剪映和CapCut两套动画体系
- **特效滤镜**: 完整的特效和滤镜处理系统
- **关键帧系统**: 精确的关键帧动画控制
- **类型安全**: Go的强类型系统确保数据正确性
- **高性能**: 优化的JSON序列化和内存管理
- **易于使用**: 简洁的API设计和详细的中文文档

## 📦 安装

```bash
go get -u github.com/zhangshican/go-capcut
```


## 🎨 完整示例

### 创建一个简单的视频项目

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/zhangshican/go-capcut/internal/track"
    "github.com/zhangshican/go-capcut/internal/segment"
    "github.com/zhangshican/go-capcut/internal/types"
    "github.com/zhangshican/go-capcut/internal/animation"
    "github.com/zhangshican/go-capcut/internal/metadata"
)

func main() {
    // 创建视频轨道
    videoTrack := track.NewTrack(track.TrackTypeVideo, "主视频轨道", 0, false)
    
    // 创建视频片段
    videoSegment, _ := segment.NewVideoSegment(
        "video_001",
        types.NewTimerange(0, 5000000), // 0-5秒
        1.0, 100.0,
    )
    
    // 添加片段到轨道
    videoTrack.AddSegment(videoSegment)
    
    // 创建文本轨道和字幕
    textTrack := track.NewTrack(track.TrackTypeText, "字幕轨道", 0, false)
    
    textSegment, _ := segment.NewTextSegment(
        "欢迎观看！",
        types.NewTimerange(1000000, 3000000), // 1-4秒
    )
    textSegment.SetFontSize(48.0)
    textSegment.SetColor("#FFFFFF")
    
    textTrack.AddSegment(textSegment)
    
    // 添加动画效果
    segmentAnims := animation.NewSegmentAnimations()
    segmentAnims.AddVideoAnimation(metadata.IntroType渐显, 0, 500000)
    segmentAnims.AddTextAnimation(metadata.TextIntroType打字机, 1000000, 800000)
    
    // 导出JSON
    videoJSON := videoTrack.ExportJSON()
    textJSON := textTrack.ExportJSON()
    animJSON := segmentAnims.ExportJSON()
    
    // 打印结果
    videoBytes, _ := json.MarshalIndent(videoJSON, "", "  ")
    textBytes, _ := json.MarshalIndent(textJSON, "", "  ")
    animBytes, _ := json.MarshalIndent(animJSON, "", "  ")
    
    fmt.Println("视频轨道JSON:")
    fmt.Println(string(videoBytes))
    
    fmt.Println("\n文本轨道JSON:")
    fmt.Println(string(textBytes))
    
    fmt.Println("\n动画JSON:")
    fmt.Println(string(animBytes))
}
```

## 📖 演示程序

项目包含丰富的演示程序，展示各种功能的使用方法：

- `demo/simple_demo/` - 基础功能演示
- `demo/animation_demo/` - 动画系统演示
- `demo/segment_demo/` - 片段系统演示
- `demo/track_demo/` - 轨道系统演示
- `demo/effect_demo/` - 特效系统演示
- `demo/keyframe_demo/` - 关键帧系统演示

运行演示程序：
```bash
cd demo/simple_demo
go run simple_demo.go
```

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- CapCut/剪映团队提供的优秀视频编辑工具
- Go语言社区的支持和贡献

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 提交 [Issue](https://github.com/zhangshican/go-capcut/issues)
- 发送邮件至 [shicanzhang@gmail.com]

---

**注意**: 请将 `zhangshican` 替换为你的实际GitHub用户名，并更新相关链接。
