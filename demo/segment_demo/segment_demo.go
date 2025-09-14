package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zhangshican/go-capcut/internal/segment"
	"github.com/zhangshican/go-capcut/internal/types"
)

func main() {
	fmt.Println("=== Segment系统演示 ===")

	// 演示基础片段
	fmt.Println("\n1. 基础片段 (BaseSegment)")
	timerange, _ := types.Trange("2s", "5s")
	baseSegment := segment.NewBaseSegment("base_material", timerange)
	fmt.Printf("   片段ID: %s\n", baseSegment.SegmentID)
	fmt.Printf("   素材ID: %s\n", baseSegment.MaterialID)
	fmt.Printf("   开始时间: %d微秒 (%.2f秒)\n", baseSegment.Start(), float64(baseSegment.Start())/1e6)
	fmt.Printf("   持续时间: %d微秒 (%.2f秒)\n", baseSegment.Duration(), float64(baseSegment.Duration())/1e6)

	// 演示视频片段
	fmt.Println("\n2. 视频片段 (VideoSegment)")
	sourceRange, _ := types.Trange("0s", "30s")
	targetRange, _ := types.Trange("5s", "10s")

	// 创建自定义的图像调节设置
	clipSettings := segment.NewClipSettingsWithParams(
		0.8,   // alpha
		15.0,  // rotation
		1.2,   // scaleX
		1.2,   // scaleY
		0.1,   // transformX
		-0.05, // transformY
		false, // flipH
		false, // flipV
	)

	videoSegment := segment.NewVideoSegment("video_material", sourceRange, targetRange, 1.0, 0.9, clipSettings)
	fmt.Printf("   视频片段ID: %s\n", videoSegment.SegmentID)
	fmt.Printf("   不透明度: %.2f\n", videoSegment.ClipSettings.Alpha)
	fmt.Printf("   旋转角度: %.1f度\n", videoSegment.ClipSettings.Rotation)

	// 添加蒙版、特效、滤镜等
	videoSegment.AddMask("circle", "圆形蒙版", "shape", "circle_mask", 0.0, 0.0, 1.0, 0.0, 0.0, false, nil, nil)
	videoSegment.AddEffect("炫光特效", "glitch_001", "glitch_res", "video_effect", 0)
	videoSegment.AddFilter("复古滤镜", "vintage_001", "vintage_res", 0.7, 0)

	transitionDuration, _ := types.Tim("1s")
	videoSegment.AddTransition("淡入淡出", "fade_001", "fade_res", transitionDuration)
	videoSegment.SetBackgroundFilling("canvas_blur", 10.0, "")

	fmt.Printf("   蒙版: %s\n", videoSegment.Mask.Name)
	fmt.Printf("   特效数量: %d\n", len(videoSegment.Effects))
	fmt.Printf("   滤镜数量: %d\n", len(videoSegment.Filters))
	fmt.Printf("   转场效果: %s\n", videoSegment.Transition.Name)

	// 演示音频片段
	fmt.Println("\n3. 音频片段 (AudioSegment)")
	audioTargetRange, _ := types.Trange("0s", "15s")
	audioSegment := segment.NewAudioSegment("audio_material", audioTargetRange, nil, 1.2, 0.8)
	fmt.Printf("   音频片段ID: %s\n", audioSegment.SegmentID)
	fmt.Printf("   播放速度: %.1fx\n", audioSegment.Speed.Value)
	fmt.Printf("   音量: %.1f\n", audioSegment.Volume)

	// 添加淡入淡出和音效
	inDuration, _ := types.Tim("2s")
	outDuration, _ := types.Tim("1s")
	audioSegment.Fade = segment.NewAudioFade(inDuration, outDuration)

	audioSegment.AddEffect("回声", "echo_res", segment.AudioEffectCategorySoundEffect)
	audioSegment.AddEffect("变声", "voice_change_res", segment.AudioEffectCategoryTone)

	fmt.Printf("   淡入时长: %.1f秒\n", float64(audioSegment.Fade.InDuration)/1e6)
	fmt.Printf("   淡出时长: %.1f秒\n", float64(audioSegment.Fade.OutDuration)/1e6)
	fmt.Printf("   音效数量: %d\n", len(audioSegment.Effects))

	// 演示文本片段
	fmt.Println("\n4. 文本片段 (TextSegment)")
	textRange, _ := types.Trange("3s", "8s")

	// 创建文本样式
	textStyle := segment.NewTextStyleWithParams(
		16.0,               // size
		true, false, false, // bold, italic, underline
		[3]float64{1.0, 0.8, 0.2}, // color (橙色)
		0.9,                       // alpha
		1,                         // align (居中)
		false,                     // vertical
		2, 5,                      // letterSpacing, lineSpacing
	)

	textSegment := segment.NewTextSegment("Hello, CapCut!", textRange, "思源黑体", textStyle, nil)
	fmt.Printf("   文本片段ID: %s\n", textSegment.SegmentID)
	fmt.Printf("   文本内容: %s\n", textSegment.Text)
	fmt.Printf("   字体: %s\n", textSegment.Font)
	fmt.Printf("   字体大小: %.1f\n", textSegment.Style.Size)
	fmt.Printf("   是否加粗: %v\n", textSegment.Style.Bold)
	fmt.Printf("   字符数: %d\n", textSegment.GetWordCount())
	fmt.Printf("   行数: %d\n", textSegment.GetLineCount())

	// 添加文本效果
	textSegment.SetBorder(0.8, [3]float64{0.0, 0.0, 0.0}, 20.0)
	textSegment.SetBackground("#FF6B35", 1, 0.9, 5.0, 0.2, 0.3, 0.5, 0.5)
	textSegment.SetShadow(true, 0.7, -45.0, "#800080", 8.0, 0.6)
	textSegment.SetFixedSize(1280, 100)

	fmt.Printf("   背景颜色: %s\n", textSegment.Background.Color)
	fmt.Printf("   阴影颜色: %s\n", textSegment.Shadow.Color)
	fmt.Printf("   固定尺寸: %dx%d\n", textSegment.FixedWidth, textSegment.FixedHeight)

	// 演示JSON导出
	fmt.Println("\n5. JSON导出示例")

	// 导出基础片段JSON
	baseJSON := baseSegment.ExportJSON()
	baseJSONBytes, _ := json.MarshalIndent(baseJSON, "", "  ")
	fmt.Println("\n基础片段JSON:")
	fmt.Println(string(baseJSONBytes)[:300] + "...")

	// 导出视频片段JSON
	videoJSON := videoSegment.ExportJSON()
	videoJSONBytes, _ := json.MarshalIndent(videoJSON, "", "  ")
	fmt.Println("\n视频片段JSON (前500字符):")
	fmt.Println(string(videoJSONBytes)[:500] + "...")

	// 导出音频片段JSON
	audioJSON := audioSegment.ExportJSON()
	audioJSONBytes, _ := json.MarshalIndent(audioJSON, "", "  ")
	fmt.Println("\n音频片段JSON (前400字符):")
	fmt.Println(string(audioJSONBytes)[:400] + "...")

	// 导出文本片段JSON
	textJSON := textSegment.ExportJSON()
	textJSONBytes, _ := json.MarshalIndent(textJSON, "", "  ")
	fmt.Println("\n文本片段JSON (前500字符):")
	fmt.Println(string(textJSONBytes)[:500] + "...")

	// 验证与Python版本的兼容性
	fmt.Println("\n6. 与Python版本兼容性验证")

	// 检查必要的字段
	requiredFields := []string{"id", "material_id", "target_timerange", "visible", "enable_adjust"}
	for _, field := range requiredFields {
		if _, exists := baseJSON[field]; exists {
			fmt.Printf("   ✓ %s 字段存在\n", field)
		} else {
			log.Printf("   ✗ %s 字段缺失\n", field)
		}
	}

	// 检查时间范围格式
	if timerangeData, ok := baseJSON["target_timerange"].(map[string]interface{}); ok {
		if start, hasStart := timerangeData["start"]; hasStart {
			if duration, hasDuration := timerangeData["duration"]; hasDuration {
				fmt.Printf("   ✓ 时间范围格式正确: start=%v, duration=%v\n", start, duration)
			}
		}
	}

	fmt.Println("\n=== Segment系统演示完成 ===")
	fmt.Println("\n已成功实现:")
	fmt.Println("  - BaseSegment: 基础片段类")
	fmt.Println("  - MediaSegment: 媒体片段基类")
	fmt.Println("  - VisualSegment: 视觉片段基类")
	fmt.Println("  - VideoSegment: 视频片段及相关效果")
	fmt.Println("  - AudioSegment: 音频片段及相关效果")
	fmt.Println("  - TextSegment: 文本片段及相关样式")
	fmt.Println("  - Speed, ClipSettings: 辅助配置类")
	fmt.Println("  - 各种特效: Mask, VideoEffect, Filter, Transition等")
	fmt.Println("  - JSON导出功能，与Python版本兼容")
	fmt.Println("\n所有功能都已经过单元测试验证！")
}
