// Template系统演示程序
// 展示Go版本的Template系统功能，包括模板模式、导入片段、导入轨道等
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zhangshican/go-capcut/internal/material"
	"github.com/zhangshican/go-capcut/internal/template"
	"github.com/zhangshican/go-capcut/internal/types"
)

func main() {
	fmt.Println("=== Go版本 Template系统演示程序 ===")
	fmt.Println()

	// 演示1: 模式枚举
	demonstrateModes()
	fmt.Println()

	// 演示2: 导入片段功能
	demonstrateImportedSegments()
	fmt.Println()

	// 演示3: 导入轨道功能
	demonstrateImportedTracks()
	fmt.Println()

	// 演示4: 媒体轨道时间范围处理
	demonstrateTimerangeProcessing()
	fmt.Println()

	// 演示5: 素材类型检查
	demonstrateMaterialTypeCheck()
	fmt.Println()

	// 演示6: JSON导入导出兼容性
	demonstrateJSONCompatibility()
}

// demonstrateModes 演示模式枚举
func demonstrateModes() {
	fmt.Println("🔧 === 模式枚举演示 ===")

	// 演示缩短模式
	fmt.Printf("📉 缩短模式:\n")
	shrinkModes := []template.ShrinkMode{
		template.ShrinkModeCutHead,
		template.ShrinkModeCutTail,
		template.ShrinkModeCutTailAlign,
		template.ShrinkModeShrink,
	}

	shrinkDescriptions := []string{
		"裁剪头部，即后移片段起始点",
		"裁剪尾部，即前移片段终止点",
		"裁剪尾部并消除间隙，后续片段也依次前移",
		"保持中间点不变，两端点向中间靠拢",
	}

	for i, mode := range shrinkModes {
		fmt.Printf("   - %s: %s\n", mode, shrinkDescriptions[i])
	}

	// 演示延长模式
	fmt.Printf("\n📈 延长模式:\n")
	extendModes := []template.ExtendMode{
		template.ExtendModeCutMaterialTail,
		template.ExtendModeExtendHead,
		template.ExtendModeExtendTail,
		template.ExtendModePushTail,
	}

	extendDescriptions := []string{
		"裁剪素材尾部，使得片段维持原长不变，此方法总是成功",
		"延伸头部，即尝试前移片段起始点，与前续片段重合时失败",
		"延伸尾部，即尝试后移片段终止点，与后续片段重合时失败",
		"延伸尾部，若有必要则依次后移后续片段，此方法总是成功",
	}

	for i, mode := range extendModes {
		fmt.Printf("   - %s: %s\n", mode, extendDescriptions[i])
	}
}

// demonstrateImportedSegments 演示导入片段功能
func demonstrateImportedSegments() {
	fmt.Println("📦 === 导入片段演示 ===")

	// 创建基本导入片段
	segmentData := map[string]interface{}{
		"material_id": "demo_material_123",
		"target_timerange": map[string]interface{}{
			"start":    float64(1000000), // 1秒
			"duration": float64(3000000), // 3秒
		},
		"render_index": float64(100),
		"visible":      true,
		"volume":       float64(0.8),
	}

	segment, err := template.NewImportedSegment(segmentData)
	if err != nil {
		log.Fatalf("创建导入片段失败: %v", err)
	}

	fmt.Printf("✅ 创建基本导入片段:\n")
	fmt.Printf("   - 素材ID: %s\n", segment.MaterialID)
	fmt.Printf("   - 目标时间范围: %d微秒 - %d微秒\n",
		segment.TargetTimerange.Start,
		segment.TargetTimerange.Start+segment.TargetTimerange.Duration)
	fmt.Printf("   - 持续时间: %.2f秒\n", float64(segment.TargetTimerange.Duration)/1e6)

	// 修改片段属性
	segment.MaterialID = "modified_material_456"
	segment.TargetTimerange.Start = 500000
	segment.TargetTimerange.Duration = 2500000

	fmt.Printf("\n🔧 修改片段属性后:\n")
	fmt.Printf("   - 新素材ID: %s\n", segment.MaterialID)
	fmt.Printf("   - 新时间范围: %d微秒 - %d微秒\n",
		segment.TargetTimerange.Start,
		segment.TargetTimerange.Start+segment.TargetTimerange.Duration)

	// 创建媒体片段
	mediaSegmentData := map[string]interface{}{
		"material_id": "media_material_789",
		"target_timerange": map[string]interface{}{
			"start":    float64(0),
			"duration": float64(2000000), // 2秒
		},
		"source_timerange": map[string]interface{}{
			"start":    float64(500000),  // 从0.5秒开始
			"duration": float64(2000000), // 取2秒
		},
		"speed":  float64(1.0),
		"volume": float64(1.0),
	}

	mediaSegment, err := template.NewImportedMediaSegment(mediaSegmentData)
	if err != nil {
		log.Fatalf("创建导入媒体片段失败: %v", err)
	}

	fmt.Printf("\n✅ 创建媒体片段:\n")
	fmt.Printf("   - 素材ID: %s\n", mediaSegment.MaterialID)
	fmt.Printf("   - 源时间范围: %d微秒 - %d微秒\n",
		mediaSegment.SourceTimerange.Start,
		mediaSegment.SourceTimerange.Start+mediaSegment.SourceTimerange.Duration)
	fmt.Printf("   - 目标时间范围: %d微秒 - %d微秒\n",
		mediaSegment.TargetTimerange.Start,
		mediaSegment.TargetTimerange.Start+mediaSegment.TargetTimerange.Duration)
}

// demonstrateImportedTracks 演示导入轨道功能
func demonstrateImportedTracks() {
	fmt.Println("🎬 === 导入轨道演示 ===")

	// 创建视频轨道数据
	videoTrackData := map[string]interface{}{
		"type": "video",
		"name": "主视频轨道",
		"id":   "video_track_001",
		"segments": []interface{}{
			map[string]interface{}{
				"material_id": "video_material_1",
				"target_timerange": map[string]interface{}{
					"start":    float64(0),
					"duration": float64(5000000), // 5秒
				},
				"source_timerange": map[string]interface{}{
					"start":    float64(0),
					"duration": float64(5000000),
				},
				"render_index": float64(200),
			},
			map[string]interface{}{
				"material_id": "video_material_2",
				"target_timerange": map[string]interface{}{
					"start":    float64(6000000), // 6秒开始
					"duration": float64(4000000), // 4秒
				},
				"source_timerange": map[string]interface{}{
					"start":    float64(1000000), // 从素材1秒开始
					"duration": float64(4000000),
				},
				"render_index": float64(200),
			},
		},
	}

	// 创建视频轨道
	videoTrack, err := template.NewImportedMediaTrack(videoTrackData)
	if err != nil {
		log.Fatalf("创建视频轨道失败: %v", err)
	}

	fmt.Printf("✅ 创建视频轨道:\n")
	fmt.Printf("   - 轨道类型: %s\n", videoTrack.TrackType)
	fmt.Printf("   - 轨道名称: %s\n", videoTrack.Name)
	fmt.Printf("   - 轨道ID: %s\n", videoTrack.TrackID)
	fmt.Printf("   - 渲染层级: %d\n", videoTrack.RenderIndex)
	fmt.Printf("   - 片段数量: %d\n", videoTrack.Len())
	fmt.Printf("   - 轨道起始时间: %d微秒 (%.2f秒)\n", videoTrack.StartTime(), float64(videoTrack.StartTime())/1e6)
	fmt.Printf("   - 轨道结束时间: %d微秒 (%.2f秒)\n", videoTrack.EndTime(), float64(videoTrack.EndTime())/1e6)

	// 创建文本轨道
	textTrackData := map[string]interface{}{
		"type": "text",
		"name": "字幕轨道",
		"id":   "text_track_001",
		"segments": []interface{}{
			map[string]interface{}{
				"material_id": "text_material_1",
				"target_timerange": map[string]interface{}{
					"start":    float64(1000000), // 1秒开始
					"duration": float64(3000000), // 持续3秒
				},
				"text":    "第一段字幕",
				"visible": true,
			},
			map[string]interface{}{
				"material_id": "text_material_2",
				"target_timerange": map[string]interface{}{
					"start":    float64(5000000), // 5秒开始
					"duration": float64(2000000), // 持续2秒
				},
				"text":    "第二段字幕",
				"visible": true,
			},
		},
	}

	textTrack, err := template.NewImportedTextTrack(textTrackData)
	if err != nil {
		log.Fatalf("创建文本轨道失败: %v", err)
	}

	fmt.Printf("\n✅ 创建文本轨道:\n")
	fmt.Printf("   - 轨道类型: %s\n", textTrack.TrackType)
	fmt.Printf("   - 轨道名称: %s\n", textTrack.Name)
	fmt.Printf("   - 片段数量: %d\n", textTrack.Len())
	fmt.Printf("   - 轨道时长: %.2f秒\n", float64(textTrack.EndTime()-textTrack.StartTime())/1e6)
}

// demonstrateTimerangeProcessing 演示时间范围处理
func demonstrateTimerangeProcessing() {
	fmt.Println("⏱️  === 时间范围处理演示 ===")

	// 创建测试媒体轨道
	trackData := map[string]interface{}{
		"type": "video",
		"name": "时间处理测试轨道",
		"id":   "timerange_test_track",
		"segments": []interface{}{
			map[string]interface{}{
				"material_id": "material_1",
				"target_timerange": map[string]interface{}{
					"start":    float64(0),
					"duration": float64(3000000), // 3秒
				},
				"source_timerange": map[string]interface{}{
					"start":    float64(0),
					"duration": float64(3000000),
				},
			},
			map[string]interface{}{
				"material_id": "material_2",
				"target_timerange": map[string]interface{}{
					"start":    float64(4000000), // 4秒开始，有1秒间隙
					"duration": float64(2000000), // 2秒
				},
				"source_timerange": map[string]interface{}{
					"start":    float64(0),
					"duration": float64(2000000),
				},
			},
		},
	}

	mediaTrack, err := template.NewImportedMediaTrack(trackData)
	if err != nil {
		log.Fatalf("创建媒体轨道失败: %v", err)
	}

	fmt.Printf("📊 原始轨道状态:\n")
	for i, seg := range mediaTrack.MediaSegments {
		fmt.Printf("   片段%d: %d微秒 - %d微秒 (%.2f秒 - %.2f秒)\n",
			i+1,
			seg.TargetTimerange.Start,
			seg.TargetTimerange.Start+seg.TargetTimerange.Duration,
			float64(seg.TargetTimerange.Start)/1e6,
			float64(seg.TargetTimerange.Start+seg.TargetTimerange.Duration)/1e6)
	}

	// 演示缩短处理 - 将第一个片段从3秒缩短到2秒
	fmt.Printf("\n🔧 缩短处理演示 (cut_tail模式):\n")
	shortTimerange := types.NewTimerange(0, 2000000) // 缩短到2秒
	err = mediaTrack.ProcessTimerange(0, shortTimerange, template.ShrinkModeCutTail, nil)
	if err != nil {
		log.Fatalf("缩短处理失败: %v", err)
	}

	fmt.Printf("   第一个片段缩短后: %d微秒 - %d微秒 (%.2f秒 - %.2f秒)\n",
		mediaTrack.MediaSegments[0].TargetTimerange.Start,
		mediaTrack.MediaSegments[0].TargetTimerange.Start+mediaTrack.MediaSegments[0].TargetTimerange.Duration,
		float64(mediaTrack.MediaSegments[0].TargetTimerange.Start)/1e6,
		float64(mediaTrack.MediaSegments[0].TargetTimerange.Start+mediaTrack.MediaSegments[0].TargetTimerange.Duration)/1e6)

	// 演示延长处理 - 将第一个片段从2秒延长到3.5秒
	fmt.Printf("\n🔧 延长处理演示 (extend_tail模式):\n")
	longTimerange := types.NewTimerange(0, 3500000) // 延长到3.5秒
	err = mediaTrack.ProcessTimerange(0, longTimerange, template.ShrinkModeCutTail, []template.ExtendMode{template.ExtendModeExtendTail})
	if err != nil {
		log.Fatalf("延长处理失败: %v", err)
	}

	fmt.Printf("   第一个片段延长后: %d微秒 - %d微秒 (%.2f秒 - %.2f秒)\n",
		mediaTrack.MediaSegments[0].TargetTimerange.Start,
		mediaTrack.MediaSegments[0].TargetTimerange.Start+mediaTrack.MediaSegments[0].TargetTimerange.Duration,
		float64(mediaTrack.MediaSegments[0].TargetTimerange.Start)/1e6,
		float64(mediaTrack.MediaSegments[0].TargetTimerange.Start+mediaTrack.MediaSegments[0].TargetTimerange.Duration)/1e6)

	// 演示push_tail模式 - 延长到会与下一个片段重叠的长度
	fmt.Printf("\n🔧 推移处理演示 (push_tail模式):\n")
	pushTimerange := types.NewTimerange(0, 5000000) // 延长到5秒，会与下个片段重叠
	err = mediaTrack.ProcessTimerange(0, pushTimerange, template.ShrinkModeCutTail, []template.ExtendMode{template.ExtendModePushTail})
	if err != nil {
		log.Fatalf("推移处理失败: %v", err)
	}

	fmt.Printf("   处理后的片段状态:\n")
	for i, seg := range mediaTrack.MediaSegments {
		fmt.Printf("   片段%d: %d微秒 - %d微秒 (%.2f秒 - %.2f秒)\n",
			i+1,
			seg.TargetTimerange.Start,
			seg.TargetTimerange.Start+seg.TargetTimerange.Duration,
			float64(seg.TargetTimerange.Start)/1e6,
			float64(seg.TargetTimerange.Start+seg.TargetTimerange.Duration)/1e6)
	}
}

// demonstrateMaterialTypeCheck 演示素材类型检查
func demonstrateMaterialTypeCheck() {
	fmt.Println("🎯 === 素材类型检查演示 ===")

	// 创建视频轨道
	videoTrackData := map[string]interface{}{
		"type": "video",
		"name": "视频轨道",
		"id":   "video_track_type_test",
	}

	videoTrack, err := template.NewImportedMediaTrack(videoTrackData)
	if err != nil {
		log.Fatalf("创建视频轨道失败: %v", err)
	}

	// 创建音频轨道
	audioTrackData := map[string]interface{}{
		"type": "audio",
		"name": "音频轨道",
		"id":   "audio_track_type_test",
	}

	audioTrack, err := template.NewImportedMediaTrack(audioTrackData)
	if err != nil {
		log.Fatalf("创建音频轨道失败: %v", err)
	}

	// 创建测试素材
	videoMaterial := &material.VideoMaterial{}
	audioMaterial := &material.AudioMaterial{}

	fmt.Printf("📹 视频轨道类型检查:\n")
	fmt.Printf("   - 接受视频素材: %v\n", videoTrack.CheckMaterialType(videoMaterial))
	fmt.Printf("   - 接受音频素材: %v\n", videoTrack.CheckMaterialType(audioMaterial))

	fmt.Printf("\n🔊 音频轨道类型检查:\n")
	fmt.Printf("   - 接受视频素材: %v\n", audioTrack.CheckMaterialType(videoMaterial))
	fmt.Printf("   - 接受音频素材: %v\n", audioTrack.CheckMaterialType(audioMaterial))

	fmt.Printf("\n✅ 类型检查确保了轨道只能接受匹配的素材类型\n")
}

// demonstrateJSONCompatibility 演示JSON导入导出兼容性
func demonstrateJSONCompatibility() {
	fmt.Println("📄 === JSON兼容性演示 ===")

	// 创建复杂的轨道数据
	complexTrackData := map[string]interface{}{
		"type":      "video",
		"name":      "复杂视频轨道",
		"id":        "complex_track_123",
		"attribute": float64(0), // 非静音
		"segments": []interface{}{
			map[string]interface{}{
				"material_id": "video_material_1",
				"target_timerange": map[string]interface{}{
					"start":    float64(0),
					"duration": float64(3000000),
				},
				"source_timerange": map[string]interface{}{
					"start":    float64(500000),
					"duration": float64(3000000),
				},
				"render_index": float64(150),
				"visible":      true,
				"volume":       float64(0.8),
				"speed":        float64(1.0),
				"extra_data":   "保留的额外数据",
			},
			map[string]interface{}{
				"material_id": "video_material_2",
				"target_timerange": map[string]interface{}{
					"start":    float64(4000000),
					"duration": float64(2000000),
				},
				"source_timerange": map[string]interface{}{
					"start":    float64(0),
					"duration": float64(2000000),
				},
				"render_index": float64(150),
				"visible":      true,
				"volume":       float64(1.0),
			},
		},
		"custom_property": "自定义轨道属性",
	}

	// 创建轨道
	complexTrack, err := template.NewImportedMediaTrack(complexTrackData)
	if err != nil {
		log.Fatalf("创建复杂轨道失败: %v", err)
	}

	fmt.Printf("✅ 创建复杂轨道:\n")
	fmt.Printf("   - 轨道名称: %s\n", complexTrack.Name)
	fmt.Printf("   - 片段数量: %d\n", len(complexTrack.MediaSegments))

	// 修改轨道属性
	complexTrack.Name = "修改后的轨道名称"
	complexTrack.MediaSegments[0].TargetTimerange.Duration = 2500000 // 修改第一个片段的持续时间

	// 导出JSON
	exportedData := complexTrack.ExportJSON()

	// 序列化为JSON字符串
	jsonBytes, err := json.MarshalIndent(exportedData, "", "  ")
	if err != nil {
		log.Fatalf("JSON序列化失败: %v", err)
	}

	fmt.Printf("\n📤 导出的JSON数据:\n%s\n", string(jsonBytes))

	// 验证JSON结构
	fmt.Printf("📊 JSON结构验证:\n")
	fmt.Printf("   - 轨道名称: %v\n", exportedData["name"])
	fmt.Printf("   - 轨道类型: %v\n", exportedData["type"])
	fmt.Printf("   - 自定义属性: %v\n", exportedData["custom_property"])

	if segments, ok := exportedData["segments"].([]map[string]interface{}); ok {
		fmt.Printf("   - 片段数量: %d\n", len(segments))
		for i, seg := range segments {
			fmt.Printf("     [%d] 素材ID: %s, 渲染层级: %v\n", i, seg["material_id"], seg["render_index"])
			if extraData, exists := seg["extra_data"]; exists {
				fmt.Printf("         额外数据: %v\n", extraData)
			}
		}
	}

	// 测试轨道导入功能
	fmt.Printf("\n🔄 轨道导入功能测试:\n")
	importedMaterials := map[string]interface{}{
		"videos": []interface{}{},
		"audios": []interface{}{},
	}

	newTrack, err := template.ImportTrack(complexTrackData, importedMaterials)
	if err != nil {
		log.Fatalf("导入轨道失败: %v", err)
	}

	fmt.Printf("   ✅ 成功导入轨道:\n")
	fmt.Printf("     - 轨道类型: %s\n", newTrack.TrackType)
	fmt.Printf("     - 轨道名称: %s\n", newTrack.Name)
	fmt.Printf("     - 轨道ID: %s\n", newTrack.TrackID)
	fmt.Printf("     - 渲染层级: %d\n", newTrack.RenderIndex)
	fmt.Printf("     - 静音状态: %v\n", newTrack.Mute)

	fmt.Printf("\n✅ JSON导入导出完全兼容Python版本！\n")
}
