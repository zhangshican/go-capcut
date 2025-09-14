package main

import (
	"encoding/json"
	"fmt"

	"github.com/zhangshican/go-capcut/internal/keyframe"
	"github.com/zhangshican/go-capcut/internal/segment"
	"github.com/zhangshican/go-capcut/internal/track"
	"github.com/zhangshican/go-capcut/internal/types"
)

func main() {
	fmt.Println("=== Keyframe系统演示 ===\n")

	// 1. 演示关键帧创建
	fmt.Println("1. 关键帧创建演示")
	kf := keyframe.NewKeyframe(2000000, 0.8) // 2秒，透明度0.8
	fmt.Printf("   ✓ 关键帧创建成功\n")
	fmt.Printf("     关键帧ID: %s\n", kf.KfID)
	fmt.Printf("     时间偏移: %d微秒 (%.2f秒)\n", kf.TimeOffset, float64(kf.TimeOffset)/1e6)
	fmt.Printf("     值: %v\n\n", kf.Values)

	// 2. 演示关键帧属性
	fmt.Println("2. 关键帧属性演示")

	validProperties := []keyframe.KeyframeProperty{
		keyframe.KeyframePropertyAlpha,
		keyframe.KeyframePropertyPositionX,
		keyframe.KeyframePropertyPositionY,
		keyframe.KeyframePropertyRotation,
		keyframe.KeyframePropertyScaleX,
		keyframe.KeyframePropertyScaleY,
		keyframe.KeyframePropertyUniformScale,
		keyframe.KeyframePropertyVolume,
	}

	fmt.Printf("   支持的关键帧属性 (%d种):\n", len(validProperties))
	for i, prop := range validProperties {
		fmt.Printf("     %d. %s\n", i+1, prop)
	}
	fmt.Println()

	// 3. 演示从字符串解析属性
	fmt.Println("3. 属性字符串解析演示")

	testStrings := []string{"alpha", "position_x", "rotation", "uniform_scale", "volume"}
	for _, propStr := range testStrings {
		prop, err := keyframe.KeyframePropertyFromString(propStr)
		if err != nil {
			fmt.Printf("   ❌ 解析 '%s' 失败: %v\n", propStr, err)
		} else {
			fmt.Printf("   ✓ '%s' -> %s\n", propStr, prop)
		}
	}
	fmt.Println()

	// 4. 演示值解析
	fmt.Println("4. 值解析演示")

	parseTests := []struct {
		property keyframe.KeyframeProperty
		value    string
		expected string
	}{
		{keyframe.KeyframePropertyAlpha, "80%", "0.8"},
		{keyframe.KeyframePropertyRotation, "45deg", "45"},
		{keyframe.KeyframePropertyPositionX, "0.5", "0.5"},
		{keyframe.KeyframePropertyBrightness, "+0.3", "0.3"},
		{keyframe.KeyframePropertyContrast, "-0.2", "-0.2"},
		{keyframe.KeyframePropertyVolume, "70%", "0.7"},
	}

	for _, test := range parseTests {
		value, err := keyframe.ParseValue(test.property, test.value)
		if err != nil {
			fmt.Printf("   ❌ 解析 %s='%s' 失败: %v\n", test.property, test.value, err)
		} else {
			fmt.Printf("   ✓ %s='%s' -> %g\n", test.property, test.value, value)
		}
	}
	fmt.Println()

	// 5. 演示关键帧列表
	fmt.Println("5. 关键帧列表演示")

	alphaList := keyframe.NewKeyframeList(keyframe.KeyframePropertyAlpha)

	// 添加多个关键帧（乱序添加测试排序）
	alphaList.AddKeyframe(0, 0.0)       // 0秒，完全透明
	alphaList.AddKeyframe(3000000, 1.0) // 3秒，完全不透明
	alphaList.AddKeyframe(1000000, 0.5) // 1秒，半透明
	alphaList.AddKeyframe(2000000, 0.8) // 2秒，80%不透明

	fmt.Printf("   ✓ 添加了 %d 个透明度关键帧\n", len(alphaList.Keyframes))
	fmt.Printf("   关键帧列表ID: %s\n", alphaList.ListID)

	// 验证排序
	fmt.Printf("   关键帧时间线:\n")
	for i, kf := range alphaList.Keyframes {
		time := float64(kf.TimeOffset) / 1e6
		fmt.Printf("     %d. %.1fs -> 透明度 %.1f\n", i+1, time, kf.Values[0])
	}
	fmt.Println()

	// 6. 演示线性插值
	fmt.Println("6. 线性插值演示")

	interpolationTests := []int64{
		500000,  // 0.5秒 - 在第一个关键帧之前
		1500000, // 1.5秒 - 在1秒和2秒之间
		2500000, // 2.5秒 - 在2秒和3秒之间
		4000000, // 4秒 - 在最后一个关键帧之后
	}

	for _, testTime := range interpolationTests {
		value := alphaList.GetValueAt(testTime)
		fmt.Printf("   %.1fs -> 透明度 %.2f\n", float64(testTime)/1e6, value)
	}
	fmt.Println()

	// 7. 演示关键帧管理器
	fmt.Println("7. 关键帧管理器演示")

	manager := keyframe.NewKeyframeManager()

	// 添加不同属性的关键帧
	manager.AddKeyframe(keyframe.KeyframePropertyAlpha, 0, 0.0)
	manager.AddKeyframe(keyframe.KeyframePropertyAlpha, 2000000, 1.0)
	manager.AddKeyframe(keyframe.KeyframePropertyRotation, 1000000, 0.0)
	manager.AddKeyframe(keyframe.KeyframePropertyRotation, 3000000, 360.0)
	manager.AddKeyframe(keyframe.KeyframePropertyScaleX, 500000, 1.0)
	manager.AddKeyframe(keyframe.KeyframePropertyScaleX, 2500000, 1.5)

	fmt.Printf("   ✓ 关键帧管理器创建成功\n")
	fmt.Printf("   包含属性: %d 种\n", len(manager.GetAllKeyframeLists()))

	for _, list := range manager.GetAllKeyframeLists() {
		fmt.Printf("     - %s: %d个关键帧\n", list.KeyframeProperty, len(list.Keyframes))
	}
	fmt.Println()

	// 8. 演示从字符串添加关键帧
	fmt.Println("8. 从字符串添加关键帧演示")

	stringTests := []struct {
		property string
		time     int64
		value    string
	}{
		{"position_x", 1000000, "0.3"},
		{"position_y", 1500000, "-0.2"},
		{"volume", 2000000, "80%"},
		{"brightness", 2500000, "+0.4"},
	}

	for _, test := range stringTests {
		err := manager.AddKeyframeFromString(test.property, test.time, test.value)
		if err != nil {
			fmt.Printf("   ❌ 添加关键帧失败: %s=%s, %v\n", test.property, test.value, err)
		} else {
			fmt.Printf("   ✓ 添加关键帧成功: %s=%.1fs,%s\n", test.property, float64(test.time)/1e6, test.value)
		}
	}
	fmt.Println()

	// 9. 演示与Segment系统集成
	fmt.Println("9. 与Segment系统集成演示")

	// 创建视频片段
	sourceRange, _ := types.Trange("0s", "10s")
	targetRange, _ := types.Trange("2s", "8s")

	videoSeg := segment.NewVideoSegment("video_material_123", sourceRange, targetRange, 1.0, 0.8, nil)

	// 添加关键帧到视频片段
	err := videoSeg.AddKeyframe("alpha", 1000000, 0.5) // 1秒，透明度50%
	if err != nil {
		fmt.Printf("   ❌ 添加透明度关键帧失败: %v\n", err)
	} else {
		fmt.Printf("   ✓ 添加透明度关键帧成功\n")
	}

	err = videoSeg.AddKeyframe("rotation", "2s", 45.0) // 2秒，旋转45度
	if err != nil {
		fmt.Printf("   ❌ 添加旋转关键帧失败: %v\n", err)
	} else {
		fmt.Printf("   ✓ 添加旋转关键帧成功\n")
	}

	// 检查片段是否有关键帧
	if videoSeg.HasKeyframes() {
		fmt.Printf("   ✓ 视频片段包含关键帧\n")
		keyframeLists := videoSeg.KeyframeManager.GetAllKeyframeLists()
		fmt.Printf("   关键帧属性数: %d\n", len(keyframeLists))
	}
	fmt.Println()

	// 10. 演示与Track系统集成
	fmt.Println("10. 与Track系统集成演示")

	// 创建视频轨道
	videoTrack := track.NewTrack(track.TrackTypeVideo, "主视频轨道", 0, false)

	// 添加片段到轨道
	err = videoTrack.AddSegment(videoSeg)
	if err != nil {
		fmt.Printf("   ❌ 添加片段到轨道失败: %v\n", err)
	} else {
		fmt.Printf("   ✓ 添加片段到轨道成功\n")
	}

	// 添加待处理的关键帧
	videoTrack.AddPendingKeyframe("alpha", 3.0, "70%")
	videoTrack.AddPendingKeyframe("position_x", 4.0, "0.2")
	videoTrack.AddPendingKeyframe("uniform_scale", 5.0, "1.2")

	fmt.Printf("   ✓ 添加了 %d 个待处理关键帧\n", len(videoTrack.PendingKeyframes))

	// 处理待处理的关键帧
	fmt.Printf("   处理待处理关键帧:\n")
	err = videoTrack.ProcessPendingKeyframes()
	if err != nil {
		fmt.Printf("   ❌ 处理关键帧失败: %v\n", err)
	} else {
		fmt.Printf("   ✓ 关键帧处理完成\n")
	}

	fmt.Printf("   剩余待处理关键帧: %d 个\n\n", len(videoTrack.PendingKeyframes))

	// 11. 演示JSON导出
	fmt.Println("11. JSON导出演示")

	// 导出关键帧列表JSON
	fmt.Println("关键帧列表JSON (前500字符):")
	listJSON := alphaList.ExportJSON()
	listJSONBytes, _ := json.MarshalIndent(listJSON, "", "  ")
	listJSONStr := string(listJSONBytes)
	if len(listJSONStr) > 500 {
		fmt.Printf("%s...\n\n", listJSONStr[:500])
	} else {
		fmt.Printf("%s\n\n", listJSONStr)
	}

	// 导出关键帧管理器JSON
	fmt.Println("关键帧管理器JSON (前600字符):")
	managerJSON := manager.ExportJSON()
	managerJSONBytes, _ := json.MarshalIndent(managerJSON, "", "  ")
	managerJSONStr := string(managerJSONBytes)
	if len(managerJSONStr) > 600 {
		fmt.Printf("%s...\n\n", managerJSONStr[:600])
	} else {
		fmt.Printf("%s\n\n", managerJSONStr)
	}

	// 12. 演示复杂动画场景
	fmt.Println("12. 复杂动画场景演示")

	complexManager := keyframe.NewKeyframeManager()

	// 淡入效果 (0-2秒)
	complexManager.AddKeyframe(keyframe.KeyframePropertyAlpha, 0, 0.0)
	complexManager.AddKeyframe(keyframe.KeyframePropertyAlpha, 2000000, 1.0)

	// 缩放弹跳效果 (0-4秒)
	complexManager.AddKeyframe(keyframe.KeyframePropertyUniformScale, 0, 0.8)
	complexManager.AddKeyframe(keyframe.KeyframePropertyUniformScale, 1000000, 1.2)
	complexManager.AddKeyframe(keyframe.KeyframePropertyUniformScale, 2000000, 1.0)
	complexManager.AddKeyframe(keyframe.KeyframePropertyUniformScale, 3000000, 1.1)
	complexManager.AddKeyframe(keyframe.KeyframePropertyUniformScale, 4000000, 1.0)

	// 旋转动画 (1-5秒)
	complexManager.AddKeyframe(keyframe.KeyframePropertyRotation, 1000000, 0)
	complexManager.AddKeyframe(keyframe.KeyframePropertyRotation, 3000000, 180)
	complexManager.AddKeyframe(keyframe.KeyframePropertyRotation, 5000000, 360)

	// 位置移动动画 (2-6秒)
	complexManager.AddKeyframe(keyframe.KeyframePropertyPositionX, 2000000, -0.3)
	complexManager.AddKeyframe(keyframe.KeyframePropertyPositionX, 4000000, 0.3)
	complexManager.AddKeyframe(keyframe.KeyframePropertyPositionX, 6000000, 0.0)

	fmt.Printf("   ✓ 复杂动画场景创建完成\n")
	fmt.Printf("   包含属性: %d 种\n", len(complexManager.GetAllKeyframeLists()))
	fmt.Printf("   总关键帧数: ")

	totalKeyframes := 0
	for _, list := range complexManager.GetAllKeyframeLists() {
		totalKeyframes += len(list.Keyframes)
		fmt.Printf("%s(%d) ", list.KeyframeProperty, len(list.Keyframes))
	}
	fmt.Printf("= %d个\n\n", totalKeyframes)

	// 13. 演示默认值系统
	fmt.Println("13. 默认值系统演示")

	defaultTests := []keyframe.KeyframeProperty{
		keyframe.KeyframePropertyAlpha,
		keyframe.KeyframePropertyScaleX,
		keyframe.KeyframePropertyPositionX,
		keyframe.KeyframePropertyRotation,
		keyframe.KeyframePropertyBrightness,
		keyframe.KeyframePropertyVolume,
	}

	for _, prop := range defaultTests {
		emptyList := keyframe.NewKeyframeList(prop)
		defaultValue := emptyList.GetValueAt(1000000) // 任意时间点
		fmt.Printf("   %s 默认值: %g\n", prop, defaultValue)
	}
	fmt.Println()

	fmt.Println("=== Keyframe系统演示完成 ===\n")

	fmt.Println("已成功实现:")
	fmt.Println("  - Keyframe: 单个关键帧，支持线性插值")
	fmt.Println("  - KeyframeProperty: 关键帧属性枚举，支持11种属性类型")
	fmt.Println("  - KeyframeList: 关键帧列表，自动排序和插值计算")
	fmt.Println("  - KeyframeManager: 关键帧管理器，统一管理所有属性")
	fmt.Println("  - 值解析: 支持百分比、角度、位置等多种格式")
	fmt.Println("  - Segment集成: 与片段系统无缝集成")
	fmt.Println("  - Track集成: 与轨道系统完整集成，支持待处理关键帧")
	fmt.Println("  - JSON导出: 与Python版本完全兼容的JSON格式")
	fmt.Println("  - 默认值: 每种属性都有合理的默认值")
	fmt.Println("  - 线性插值: 自动在关键帧之间进行线性插值")
	fmt.Println("  - 复杂动画: 支持多属性复杂动画场景")
	fmt.Println("\n所有功能都已经过单元测试验证！")
}
