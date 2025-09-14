// Script文件系统演示程序
// 展示Go版本的Script文件系统功能，包括草稿创建、素材管理、轨道管理、JSON导入导出等
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/zhangshican/go-capcut/internal/material"
	"github.com/zhangshican/go-capcut/internal/script"
	"github.com/zhangshican/go-capcut/internal/track"
)

func main() {
	fmt.Println("=== Go版本 Script文件系统演示程序 ===")
	fmt.Println()

	// 演示1: 创建草稿文件
	demonstrateScriptFileCreation()
	fmt.Println()

	// 演示2: 素材管理
	demonstrateMaterialManagement()
	fmt.Println()

	// 演示3: 轨道管理
	demonstrateTrackManagement()
	fmt.Println()

	// 演示4: JSON导入导出
	demonstrateJSONOperations()
	fmt.Println()

	// 演示5: 模板加载功能
	demonstrateTemplateLoading()
	fmt.Println()

	// 演示6: 完整的草稿文件工作流
	demonstrateCompleteWorkflow()
}

// demonstrateScriptFileCreation 演示草稿文件创建
func demonstrateScriptFileCreation() {
	fmt.Println("🎬 === 草稿文件创建演示 ===")

	// 创建1920x1080的草稿文件
	sf, err := script.NewScriptFile(1920, 1080, 30)
	if err != nil {
		log.Fatalf("创建草稿文件失败: %v", err)
	}

	fmt.Printf("✅ 创建草稿文件成功:\n")
	fmt.Printf("   - 分辨率: %dx%d\n", sf.Width, sf.Height)
	fmt.Printf("   - 帧率: %d FPS\n", sf.FPS)
	fmt.Printf("   - 初始时长: %.2f秒\n", float64(sf.Duration)/1e6)
	fmt.Printf("   - 轨道数量: %d\n", len(sf.Tracks))
	fmt.Printf("   - 导入轨道数量: %d\n", len(sf.ImportedTracks))

	// 创建不同规格的草稿文件
	fmt.Printf("\n📱 创建竖屏草稿文件 (720x1280):\n")
	verticalSF, err := script.NewScriptFile(720, 1280, 25)
	if err != nil {
		log.Fatalf("创建竖屏草稿文件失败: %v", err)
	}
	fmt.Printf("   - 分辨率: %dx%d\n", verticalSF.Width, verticalSF.Height)
	fmt.Printf("   - 帧率: %d FPS\n", verticalSF.FPS)

	// 创建使用默认帧率的草稿文件
	fmt.Printf("\n🎥 创建4K草稿文件 (3840x2160, 默认30FPS):\n")
	fourKSF, err := script.NewScriptFile(3840, 2160)
	if err != nil {
		log.Fatalf("创建4K草稿文件失败: %v", err)
	}
	fmt.Printf("   - 分辨率: %dx%d\n", fourKSF.Width, fourKSF.Height)
	fmt.Printf("   - 默认帧率: %d FPS\n", fourKSF.FPS)
}

// demonstrateMaterialManagement 演示素材管理
func demonstrateMaterialManagement() {
	fmt.Println("🎞️ === 素材管理演示 ===")

	sf, err := script.NewScriptFile(1920, 1080)
	if err != nil {
		log.Fatalf("创建草稿文件失败: %v", err)
	}

	// 创建视频素材
	videoMaterial1 := &material.VideoMaterial{
		MaterialID: "video_001",
		Path:       "/Users/demo/Videos/intro.mp4",
		Width:      1920,
		Height:     1080,
		Duration:   5000000, // 5秒
		CropSettings: &material.CropSettings{
			UpperLeftX:  0.0,
			UpperLeftY:  0.0,
			LowerRightX: 1.0,
			LowerRightY: 1.0,
		},
	}

	videoMaterial2 := &material.VideoMaterial{
		MaterialID: "video_002",
		Path:       "/Users/demo/Videos/main_content.mp4",
		Width:      1920,
		Height:     1080,
		Duration:   15000000, // 15秒
		CropSettings: &material.CropSettings{
			UpperLeftX:  0.0,
			UpperLeftY:  0.0,
			LowerRightX: 1.0,
			LowerRightY: 1.0,
		},
	}

	// 创建音频素材
	audioMaterial1 := &material.AudioMaterial{
		MaterialID: "audio_001",
		Path:       "/Users/demo/Audio/background_music.mp3",
		Duration:   20000000, // 20秒
	}

	audioMaterial2 := &material.AudioMaterial{
		MaterialID: "audio_002",
		Path:       "/Users/demo/Audio/voice_over.wav",
		Duration:   18000000, // 18秒
	}

	fmt.Printf("📹 添加视频素材:\n")
	sf.AddMaterial(videoMaterial1)
	fmt.Printf("   - %s: %s (%.1f秒)\n", videoMaterial1.MaterialID, filepath.Base(videoMaterial1.Path), float64(videoMaterial1.Duration)/1e6)

	sf.AddMaterial(videoMaterial2)
	fmt.Printf("   - %s: %s (%.1f秒)\n", videoMaterial2.MaterialID, filepath.Base(videoMaterial2.Path), float64(videoMaterial2.Duration)/1e6)

	fmt.Printf("\n🎵 添加音频素材:\n")
	sf.AddMaterial(audioMaterial1)
	fmt.Printf("   - %s: %s (%.1f秒)\n", audioMaterial1.MaterialID, filepath.Base(audioMaterial1.Path), float64(audioMaterial1.Duration)/1e6)

	sf.AddMaterial(audioMaterial2)
	fmt.Printf("   - %s: %s (%.1f秒)\n", audioMaterial2.MaterialID, filepath.Base(audioMaterial2.Path), float64(audioMaterial2.Duration)/1e6)

	// 测试素材包含检查
	fmt.Printf("\n🔍 素材包含检查:\n")
	fmt.Printf("   - 包含video_001: %v\n", sf.Materials.Contains(videoMaterial1))
	fmt.Printf("   - 包含audio_001: %v\n", sf.Materials.Contains(audioMaterial1))

	// 测试重复添加
	fmt.Printf("\n🔄 重复添加测试:\n")
	originalVideoCount := len(sf.Materials.Videos)
	sf.AddMaterial(videoMaterial1) // 重复添加
	newVideoCount := len(sf.Materials.Videos)
	fmt.Printf("   - 重复添加前视频数量: %d\n", originalVideoCount)
	fmt.Printf("   - 重复添加后视频数量: %d\n", newVideoCount)
	fmt.Printf("   - 重复添加被正确忽略: %v\n", originalVideoCount == newVideoCount)

	// 显示最终素材统计
	fmt.Printf("\n📊 最终素材统计:\n")
	fmt.Printf("   - 视频素材: %d个\n", len(sf.Materials.Videos))
	fmt.Printf("   - 音频素材: %d个\n", len(sf.Materials.Audios))
	fmt.Printf("   - 贴纸素材: %d个\n", len(sf.Materials.Stickers))
	fmt.Printf("   - 文本素材: %d个\n", len(sf.Materials.Texts))
	fmt.Printf("   - 特效素材: %d个\n", len(sf.Materials.VideoEffects))
	fmt.Printf("   - 动画素材: %d个\n", len(sf.Materials.Animations))
}

// demonstrateTrackManagement 演示轨道管理
func demonstrateTrackManagement() {
	fmt.Println("🎚️ === 轨道管理演示 ===")

	sf, err := script.NewScriptFile(1920, 1080, 30)
	if err != nil {
		log.Fatalf("创建草稿文件失败: %v", err)
	}

	// 添加主视频轨道
	mainVideoTrack := "主视频轨道"
	sf.AddTrack(track.TrackTypeVideo, &mainVideoTrack)
	fmt.Printf("✅ 添加主视频轨道: %s\n", mainVideoTrack)

	// 添加覆盖视频轨道（相对位置+1）
	overlayVideoTrack := "覆盖视频轨道"
	sf.AddTrack(track.TrackTypeVideo, &overlayVideoTrack, script.WithRelativeIndex(1))
	fmt.Printf("✅ 添加覆盖视频轨道: %s (相对层级+1)\n", overlayVideoTrack)

	// 添加背景音乐轨道
	bgMusicTrack := "背景音乐"
	sf.AddTrack(track.TrackTypeAudio, &bgMusicTrack)
	fmt.Printf("✅ 添加背景音乐轨道: %s\n", bgMusicTrack)

	// 添加静音的语音轨道
	voiceTrack := "语音轨道"
	sf.AddTrack(track.TrackTypeAudio, &voiceTrack, script.WithMute(true))
	fmt.Printf("✅ 添加语音轨道: %s (静音状态)\n", voiceTrack)

	// 添加文本轨道
	textTrack := "字幕轨道"
	sf.AddTrack(track.TrackTypeText, &textTrack, script.WithRelativeIndex(2))
	fmt.Printf("✅ 添加字幕轨道: %s (相对层级+2)\n", textTrack)

	// 添加特效轨道（使用绝对层级）
	effectTrack := "特效轨道"
	sf.AddTrack(track.TrackTypeEffect, &effectTrack, script.WithAbsoluteIndex(20000))
	fmt.Printf("✅ 添加特效轨道: %s (绝对层级20000)\n", effectTrack)

	// 显示轨道信息
	fmt.Printf("\n📋 轨道详细信息:\n")
	for name, t := range sf.Tracks {
		fmt.Printf("   - %s:\n", name)
		fmt.Printf("     类型: %s\n", t.TrackType)
		fmt.Printf("     渲染层级: %d\n", t.RenderIndex)
		fmt.Printf("     静音状态: %v\n", t.Mute)
		fmt.Printf("     片段数量: %d\n", len(t.Segments))
	}

	// 测试轨道获取
	fmt.Printf("\n🔍 轨道获取测试:\n")
	foundTrack, err := sf.GetTrack("video", &mainVideoTrack)
	if err != nil {
		fmt.Printf("   - 获取轨道失败: %v\n", err)
	} else {
		fmt.Printf("   - 成功获取轨道: %s (类型: %s)\n", foundTrack.Name, foundTrack.TrackType)
	}

	// 测试获取不存在的轨道
	nonExistentTrack := "不存在的轨道"
	_, err = sf.GetTrack("video", &nonExistentTrack)
	if err != nil {
		fmt.Printf("   - 正确处理不存在的轨道: %v\n", err)
	}

	// 测试重复添加轨道
	fmt.Printf("\n🔄 重复添加轨道测试:\n")
	originalCount := len(sf.Tracks)
	sf.AddTrack(track.TrackTypeVideo, &mainVideoTrack) // 重复添加
	newCount := len(sf.Tracks)
	fmt.Printf("   - 重复添加前轨道数量: %d\n", originalCount)
	fmt.Printf("   - 重复添加后轨道数量: %d\n", newCount)
	fmt.Printf("   - 重复添加被正确忽略: %v\n", originalCount == newCount)

	fmt.Printf("\n📊 最终轨道统计: %d个轨道\n", len(sf.Tracks))
}

// demonstrateJSONOperations 演示JSON操作
func demonstrateJSONOperations() {
	fmt.Println("📄 === JSON操作演示 ===")

	sf, err := script.NewScriptFile(1920, 1080, 25)
	if err != nil {
		log.Fatalf("创建草稿文件失败: %v", err)
	}

	// 添加一些内容
	videoMaterial := &material.VideoMaterial{
		MaterialID: "demo_video_123",
		Path:       "/demo/video.mp4",
		Width:      1920,
		Height:     1080,
		Duration:   10000000, // 10秒
		CropSettings: &material.CropSettings{
			UpperLeftX:  0.0,
			UpperLeftY:  0.0,
			LowerRightX: 1.0,
			LowerRightY: 1.0,
		},
	}
	sf.AddMaterial(videoMaterial)

	audioMaterial := &material.AudioMaterial{
		MaterialID: "demo_audio_456",
		Path:       "/demo/audio.mp3",
		Duration:   12000000, // 12秒
	}
	sf.AddMaterial(audioMaterial)

	// 添加轨道
	videoTrack := "演示视频轨道"
	audioTrack := "演示音频轨道"
	textTrack := "演示文本轨道"

	sf.AddTrack(track.TrackTypeVideo, &videoTrack)
	sf.AddTrack(track.TrackTypeAudio, &audioTrack, script.WithMute(true))
	sf.AddTrack(track.TrackTypeText, &textTrack, script.WithRelativeIndex(1))

	// 设置草稿时长
	sf.Duration = 15000000 // 15秒

	fmt.Printf("🎬 草稿文件内容:\n")
	fmt.Printf("   - 分辨率: %dx%d @ %d FPS\n", sf.Width, sf.Height, sf.FPS)
	fmt.Printf("   - 时长: %.1f秒\n", float64(sf.Duration)/1e6)
	fmt.Printf("   - 素材数量: %d个视频 + %d个音频\n", len(sf.Materials.Videos), len(sf.Materials.Audios))
	fmt.Printf("   - 轨道数量: %d个\n", len(sf.Tracks))

	// 导出JSON字符串
	fmt.Printf("\n📤 导出JSON字符串:\n")
	jsonStr, err := sf.Dumps()
	if err != nil {
		log.Fatalf("JSON导出失败: %v", err)
	}

	// 验证JSON格式
	var parsedData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsedData); err != nil {
		log.Fatalf("JSON格式验证失败: %v", err)
	}

	fmt.Printf("   ✅ JSON导出成功，字符串长度: %d字符\n", len(jsonStr))

	// 验证关键字段
	fmt.Printf("\n🔍 JSON内容验证:\n")
	if fps, ok := parsedData["fps"].(float64); ok {
		fmt.Printf("   - FPS: %.0f\n", fps)
	}
	if duration, ok := parsedData["duration"].(float64); ok {
		fmt.Printf("   - 时长: %.1f秒\n", duration/1e6)
	}
	if canvasConfig, ok := parsedData["canvas_config"].(map[string]interface{}); ok {
		if width, ok := canvasConfig["width"].(float64); ok {
			if height, ok := canvasConfig["height"].(float64); ok {
				fmt.Printf("   - 画布: %.0fx%.0f\n", width, height)
			}
		}
	}

	// 验证平台信息
	if platform, ok := parsedData["platform"].(map[string]interface{}); ok {
		if appSource, ok := platform["app_source"].(string); ok {
			if appVersion, ok := platform["app_version"].(string); ok {
				fmt.Printf("   - 平台: %s v%s\n", appSource, appVersion)
			}
		}
	}

	// 验证素材和轨道
	if materials, ok := parsedData["materials"].(map[string]interface{}); ok {
		if videos, ok := materials["videos"].([]interface{}); ok {
			fmt.Printf("   - 视频素材: %d个\n", len(videos))
		}
		if audios, ok := materials["audios"].([]interface{}); ok {
			fmt.Printf("   - 音频素材: %d个\n", len(audios))
		}
	}

	if tracks, ok := parsedData["tracks"].([]interface{}); ok {
		fmt.Printf("   - 轨道: %d个\n", len(tracks))
	}

	// 保存到临时文件
	tempFile := filepath.Join(os.TempDir(), "go_script_demo.json")
	fmt.Printf("\n💾 保存到文件: %s\n", tempFile)
	if err := sf.Dump(tempFile); err != nil {
		log.Fatalf("文件保存失败: %v", err)
	}

	// 验证文件
	if fileInfo, err := os.Stat(tempFile); err == nil {
		fmt.Printf("   ✅ 文件保存成功，大小: %.1f KB\n", float64(fileInfo.Size())/1024)
	} else {
		fmt.Printf("   ❌ 文件验证失败: %v\n", err)
	}

	// 清理临时文件
	defer func() {
		if err := os.Remove(tempFile); err != nil {
			fmt.Printf("   ⚠️ 清理临时文件失败: %v\n", err)
		} else {
			fmt.Printf("   🗑️ 临时文件已清理\n")
		}
	}()
}

// demonstrateTemplateLoading 演示模板加载功能
func demonstrateTemplateLoading() {
	fmt.Println("📋 === 模板加载演示 ===")

	// 创建一个测试模板文件
	tempDir := os.TempDir()
	templateFile := filepath.Join(tempDir, "test_template.json")

	// 创建模板内容
	templateContent := map[string]interface{}{
		"fps":      float64(24),
		"duration": float64(8000000), // 8秒
		"canvas_config": map[string]interface{}{
			"width":  float64(1280),
			"height": float64(720),
			"ratio":  "16:9",
		},
		"materials": map[string]interface{}{
			"videos": []interface{}{
				map[string]interface{}{
					"id":       "template_video_1",
					"path":     "/template/intro.mp4",
					"width":    float64(1280),
					"height":   float64(720),
					"duration": float64(5000000),
				},
			},
			"audios": []interface{}{
				map[string]interface{}{
					"id":       "template_audio_1",
					"path":     "/template/bg_music.mp3",
					"duration": float64(8000000),
				},
			},
			"texts": []interface{}{
				map[string]interface{}{
					"id":      "template_text_1",
					"content": "模板标题文字",
					"font":    "Arial",
				},
			},
		},
		"tracks": []interface{}{
			map[string]interface{}{
				"type": "video",
				"name": "主视频轨道",
				"id":   "main_video_track",
				"segments": []interface{}{
					map[string]interface{}{
						"material_id": "template_video_1",
						"target_timerange": map[string]interface{}{
							"start":    float64(0),
							"duration": float64(5000000),
						},
						"render_index": float64(0),
					},
				},
			},
			map[string]interface{}{
				"type": "audio",
				"name": "背景音乐轨道",
				"id":   "bg_music_track",
				"segments": []interface{}{
					map[string]interface{}{
						"material_id": "template_audio_1",
						"target_timerange": map[string]interface{}{
							"start":    float64(0),
							"duration": float64(8000000),
						},
						"render_index": float64(0),
					},
				},
			},
			map[string]interface{}{
				"type": "text",
				"name": "标题轨道",
				"id":   "title_track",
				"segments": []interface{}{
					map[string]interface{}{
						"material_id": "template_text_1",
						"target_timerange": map[string]interface{}{
							"start":    float64(1000000), // 1秒开始
							"duration": float64(3000000), // 持续3秒
						},
						"render_index": float64(15000),
					},
				},
			},
		},
	}

	// 写入模板文件
	jsonBytes, err := json.MarshalIndent(templateContent, "", "    ")
	if err != nil {
		log.Fatalf("创建模板JSON失败: %v", err)
	}

	if err := os.WriteFile(templateFile, jsonBytes, 0644); err != nil {
		log.Fatalf("写入模板文件失败: %v", err)
	}

	fmt.Printf("📝 创建测试模板: %s\n", templateFile)

	// 加载模板
	fmt.Printf("\n📂 加载模板文件:\n")
	sf, err := script.LoadTemplate(templateFile)
	if err != nil {
		log.Fatalf("加载模板失败: %v", err)
	}

	fmt.Printf("   ✅ 模板加载成功\n")
	fmt.Printf("   - 分辨率: %dx%d\n", sf.Width, sf.Height)
	fmt.Printf("   - 帧率: %d FPS\n", sf.FPS)
	fmt.Printf("   - 时长: %.1f秒\n", float64(sf.Duration)/1e6)
	fmt.Printf("   - 保存路径: %s\n", *sf.SavePath)

	// 显示导入的素材
	fmt.Printf("\n📦 导入的素材:\n")
	for materialType, materials := range sf.ImportedMaterials {
		if len(materials) > 0 {
			fmt.Printf("   - %s: %d个\n", materialType, len(materials))
			for i, mat := range materials {
				if i < 2 { // 只显示前两个
					if id, ok := mat["id"].(string); ok {
						if path, ok := mat["path"].(string); ok {
							fmt.Printf("     [%d] %s: %s\n", i+1, id, path)
						} else {
							fmt.Printf("     [%d] %s\n", i+1, id)
						}
					}
				}
			}
		}
	}

	// 显示导入的轨道
	fmt.Printf("\n🎚️ 导入的轨道:\n")
	for i, importedTrack := range sf.ImportedTracks {
		fmt.Printf("   [%d] %s (%s)\n", i+1, importedTrack.Name, importedTrack.TrackType)
		fmt.Printf("       - 渲染层级: %d\n", importedTrack.RenderIndex)
		fmt.Printf("       - 片段数量: %d\n", len(importedTrack.Segments))
	}

	// 测试保存功能
	fmt.Printf("\n💾 测试保存功能:\n")
	sf.Duration = 10000000 // 修改时长为10秒
	if err := sf.Save(); err != nil {
		fmt.Printf("   ❌ 保存失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 保存成功\n")
	}

	// 清理临时文件
	defer func() {
		if err := os.Remove(templateFile); err != nil {
			fmt.Printf("   ⚠️ 清理模板文件失败: %v\n", err)
		} else {
			fmt.Printf("   🗑️ 模板文件已清理\n")
		}
	}()

	// 测试素材检查功能
	fmt.Printf("\n🔍 素材检查功能:\n")
	sf.InspectMaterial()
}

// demonstrateCompleteWorkflow 演示完整的草稿文件工作流
func demonstrateCompleteWorkflow() {
	fmt.Println("🎯 === 完整工作流演示 ===")

	fmt.Printf("🎬 创建新项目 - 制作一个简单的视频:\n")

	// 第1步: 创建草稿文件
	sf, err := script.NewScriptFile(1920, 1080, 30)
	if err != nil {
		log.Fatalf("创建草稿文件失败: %v", err)
	}
	fmt.Printf("   ✅ 步骤1: 创建1920x1080@30fps草稿\n")

	// 第2步: 准备素材
	fmt.Printf("   📁 步骤2: 准备素材\n")

	introVideo := &material.VideoMaterial{
		MaterialID: "intro_clip",
		Path:       "/project/assets/intro.mp4",
		Width:      1920,
		Height:     1080,
		Duration:   3000000, // 3秒
		CropSettings: &material.CropSettings{
			UpperLeftX:  0.0,
			UpperLeftY:  0.0,
			LowerRightX: 1.0,
			LowerRightY: 1.0,
		},
	}

	mainVideo := &material.VideoMaterial{
		MaterialID: "main_content",
		Path:       "/project/assets/main_video.mp4",
		Width:      1920,
		Height:     1080,
		Duration:   20000000, // 20秒
		CropSettings: &material.CropSettings{
			UpperLeftX:  0.0,
			UpperLeftY:  0.0,
			LowerRightX: 1.0,
			LowerRightY: 1.0,
		},
	}

	outroVideo := &material.VideoMaterial{
		MaterialID: "outro_clip",
		Path:       "/project/assets/outro.mp4",
		Width:      1920,
		Height:     1080,
		Duration:   2000000, // 2秒
		CropSettings: &material.CropSettings{
			UpperLeftX:  0.0,
			UpperLeftY:  0.0,
			LowerRightX: 1.0,
			LowerRightY: 1.0,
		},
	}

	backgroundMusic := &material.AudioMaterial{
		MaterialID: "bg_music",
		Path:       "/project/assets/background.mp3",
		Duration:   30000000, // 30秒
	}

	voiceOver := &material.AudioMaterial{
		MaterialID: "voice_narration",
		Path:       "/project/assets/narration.wav",
		Duration:   18000000, // 18秒
	}

	// 添加素材到草稿
	sf.AddMaterial(introVideo)
	sf.AddMaterial(mainVideo)
	sf.AddMaterial(outroVideo)
	sf.AddMaterial(backgroundMusic)
	sf.AddMaterial(voiceOver)

	fmt.Printf("     - 添加了%d个视频素材\n", len(sf.Materials.Videos))
	fmt.Printf("     - 添加了%d个音频素材\n", len(sf.Materials.Audios))

	// 第3步: 创建轨道结构
	fmt.Printf("   🎚️ 步骤3: 创建轨道结构\n")

	mainVideoTrack := "主视频轨道"
	bgMusicTrack := "背景音乐轨道"
	voiceTrack := "语音轨道"
	titleTrack := "标题轨道"
	effectTrack := "特效轨道"

	sf.AddTrack(track.TrackTypeVideo, &mainVideoTrack)                                // 主视频轨道
	sf.AddTrack(track.TrackTypeAudio, &bgMusicTrack)                                  // 背景音乐
	sf.AddTrack(track.TrackTypeAudio, &voiceTrack, script.WithRelativeIndex(1))       // 语音轨道 (层级+1)
	sf.AddTrack(track.TrackTypeText, &titleTrack, script.WithRelativeIndex(2))        // 标题轨道 (层级+2)
	sf.AddTrack(track.TrackTypeEffect, &effectTrack, script.WithAbsoluteIndex(20000)) // 特效轨道 (绝对层级)

	fmt.Printf("     - 创建了%d个轨道\n", len(sf.Tracks))

	// 第4步: 计算总时长
	totalDuration := introVideo.Duration + mainVideo.Duration + outroVideo.Duration
	sf.Duration = totalDuration
	fmt.Printf("   ⏱️ 步骤4: 设置总时长 %.1f秒\n", float64(totalDuration)/1e6)

	// 第5步: 显示项目概览
	fmt.Printf("   📊 步骤5: 项目概览\n")
	fmt.Printf("     - 项目规格: %dx%d @ %dFPS\n", sf.Width, sf.Height, sf.FPS)
	fmt.Printf("     - 项目时长: %.1f秒\n", float64(sf.Duration)/1e6)
	fmt.Printf("     - 素材统计:\n")
	fmt.Printf("       * 视频: %d个 (总时长%.1f秒)\n",
		len(sf.Materials.Videos),
		float64(introVideo.Duration+mainVideo.Duration+outroVideo.Duration)/1e6)
	fmt.Printf("       * 音频: %d个 (总时长%.1f秒)\n",
		len(sf.Materials.Audios),
		float64(backgroundMusic.Duration+voiceOver.Duration)/1e6)

	// 第6步: 轨道详情
	fmt.Printf("     - 轨道详情:\n")
	for name, t := range sf.Tracks {
		fmt.Printf("       * %s: %s (层级%d)\n", name, t.TrackType, t.RenderIndex)
	}

	// 第7步: 导出项目
	fmt.Printf("   💾 步骤6: 导出项目文件\n")

	outputFile := filepath.Join(os.TempDir(), "complete_project.json")
	if err := sf.Dump(outputFile); err != nil {
		fmt.Printf("     ❌ 导出失败: %v\n", err)
	} else {
		if fileInfo, err := os.Stat(outputFile); err == nil {
			fmt.Printf("     ✅ 导出成功: %s (%.1f KB)\n", outputFile, float64(fileInfo.Size())/1024)
		}
	}

	// 第8步: 验证项目文件
	fmt.Printf("   🔍 步骤7: 验证项目文件\n")

	// 重新加载项目验证
	loadedSF, err := script.LoadTemplate(outputFile)
	if err != nil {
		fmt.Printf("     ❌ 验证失败: %v\n", err)
	} else {
		fmt.Printf("     ✅ 验证成功:\n")
		fmt.Printf("       - 分辨率匹配: %v\n", loadedSF.Width == sf.Width && loadedSF.Height == sf.Height)
		fmt.Printf("       - 帧率匹配: %v\n", loadedSF.FPS == sf.FPS)
		fmt.Printf("       - 时长匹配: %v\n", loadedSF.Duration == sf.Duration)
		fmt.Printf("       - 导入素材: %d类型\n", len(loadedSF.ImportedMaterials))
		fmt.Printf("       - 导入轨道: %d个\n", len(loadedSF.ImportedTracks))
	}

	// 清理
	defer func() {
		if err := os.Remove(outputFile); err != nil {
			fmt.Printf("   ⚠️ 清理项目文件失败: %v\n", err)
		} else {
			fmt.Printf("   🗑️ 项目文件已清理\n")
		}
	}()

	fmt.Printf("\n🎉 完整工作流演示完成！\n")
	fmt.Printf("   - 成功创建了包含%d个素材和%d个轨道的项目\n",
		len(sf.Materials.Videos)+len(sf.Materials.Audios), len(sf.Tracks))
	fmt.Printf("   - 项目总时长: %.1f秒\n", float64(sf.Duration)/1e6)
	fmt.Printf("   - JSON导出和重新加载验证通过\n")
}
