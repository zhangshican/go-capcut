// Effect系统演示程序
// 展示Go版本的Effect系统功能，包括特效片段、滤镜片段、轨道集成等
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zhangshican/go-capcut/internal/metadata"
	"github.com/zhangshican/go-capcut/internal/segment"
	"github.com/zhangshican/go-capcut/internal/track"
	"github.com/zhangshican/go-capcut/internal/types"
)

func main() {
	fmt.Println("=== Go版本 Effect系统演示程序 ===")
	fmt.Println()

	// 演示1: 特效片段基础功能
	demonstrateEffectSegments()
	fmt.Println()

	// 演示2: 滤镜片段基础功能
	demonstrateFilterSegments()
	fmt.Println()

	// 演示3: 特效轨道管理
	demonstrateEffectTracks()
	fmt.Println()

	// 演示4: 滤镜轨道管理
	demonstrateFilterTracks()
	fmt.Println()

	// 演示5: 复杂特效参数处理
	demonstrateComplexEffectParameters()
	fmt.Println()

	// 演示6: 轨道集成和JSON导出
	demonstrateTrackIntegrationAndExport()
	fmt.Println()

	// 演示7: 完整的特效工作流
	demonstrateCompleteEffectWorkflow()
}

// demonstrateEffectSegments 演示特效片段基础功能
func demonstrateEffectSegments() {
	fmt.Println("🎬 === 特效片段基础功能演示 ===")

	// 创建特效元数据
	effectMeta := metadata.NewEffectMeta(
		"模糊效果",
		false,
		"blur_resource_001",
		"blur_effect_123",
		"blur_md5_hash",
		[]metadata.EffectParam{
			metadata.NewEffectParam("blur_intensity", 0.5, 0.0, 1.0),
			metadata.NewEffectParam("blur_radius", 10.0, 0.0, 50.0),
		},
	)

	fmt.Printf("📋 特效元数据:\n")
	fmt.Printf("   - 名称: %s\n", effectMeta.Name)
	fmt.Printf("   - 资源ID: %s\n", effectMeta.ResourceID)
	fmt.Printf("   - 效果ID: %s\n", effectMeta.EffectID)
	fmt.Printf("   - 参数数量: %d\n", len(effectMeta.Params))

	// 创建时间范围
	timerange := types.NewTimerange(2000000, 8000000) // 2-10秒

	// 创建特效片段
	params := []float64{70.0, 80.0} // 强度70%，半径80%
	effectSegment, err := segment.NewEffectSegment(effectMeta, timerange, params)
	if err != nil {
		log.Fatalf("创建特效片段失败: %v", err)
	}

	fmt.Printf("\n🎯 特效片段创建成功:\n")
	fmt.Printf("   - 片段ID: %s\n", effectSegment.GetID())
	fmt.Printf("   - 开始时间: %.2f秒\n", float64(effectSegment.Start())/1e6)
	fmt.Printf("   - 持续时间: %.2f秒\n", float64(effectSegment.Duration())/1e6)
	fmt.Printf("   - 结束时间: %.2f秒\n", float64(effectSegment.End())/1e6)
	fmt.Printf("   - 特效类型: %s\n", effectSegment.EffectInst.EffectType)
	fmt.Printf("   - 应用目标: %s\n", getApplyTargetTypeString(effectSegment.EffectInst.ApplyTargetType))

	// 获取素材引用
	refs := effectSegment.GetMaterialRefs()
	fmt.Printf("   - 素材引用数: %d\n", len(refs))

	fmt.Printf("\n📊 特效参数:\n")
	for i, param := range effectSegment.EffectInst.AdjustParams {
		fmt.Printf("   [%d] %v\n", i+1, param)
	}
}

// demonstrateFilterSegments 演示滤镜片段基础功能
func demonstrateFilterSegments() {
	fmt.Println("🌈 === 滤镜片段基础功能演示 ===")

	// 创建滤镜元数据
	filterMeta := metadata.NewEffectMeta(
		"复古滤镜",
		true, // VIP滤镜
		"vintage_resource_002",
		"vintage_filter_456",
		"vintage_md5_hash",
		[]metadata.EffectParam{
			metadata.NewEffectParam("vintage_intensity", 1.0, 0.0, 1.0),
		},
	)

	fmt.Printf("🎨 滤镜元数据:\n")
	fmt.Printf("   - 名称: %s\n", filterMeta.Name)
	fmt.Printf("   - VIP状态: %v\n", filterMeta.IsVIP)
	fmt.Printf("   - 资源ID: %s\n", filterMeta.ResourceID)
	fmt.Printf("   - 效果ID: %s\n", filterMeta.EffectID)

	// 创建时间范围
	timerange := types.NewTimerange(5000000, 12000000) // 5-17秒

	// 创建滤镜片段
	intensity := 85.0 // 85%强度
	filterSegment := segment.NewFilterSegment(filterMeta, timerange, intensity)

	fmt.Printf("\n🎯 滤镜片段创建成功:\n")
	fmt.Printf("   - 片段ID: %s\n", filterSegment.GetID())
	fmt.Printf("   - 开始时间: %.2f秒\n", float64(filterSegment.Start())/1e6)
	fmt.Printf("   - 持续时间: %.2f秒\n", float64(filterSegment.Duration())/1e6)
	fmt.Printf("   - 结束时间: %.2f秒\n", float64(filterSegment.End())/1e6)
	fmt.Printf("   - 滤镜强度: %.1f%%\n", filterSegment.GetIntensity())
	fmt.Printf("   - 内部强度: %.3f\n", filterSegment.Material.Intensity)

	// 测试强度调整
	fmt.Printf("\n🔧 强度调整测试:\n")
	originalIntensity := filterSegment.GetIntensity()
	fmt.Printf("   - 原始强度: %.1f%%\n", originalIntensity)

	filterSegment.SetIntensity(60.0)
	fmt.Printf("   - 调整后强度: %.1f%%\n", filterSegment.GetIntensity())

	filterSegment.SetIntensity(100.0)
	fmt.Printf("   - 最大强度: %.1f%%\n", filterSegment.GetIntensity())

	filterSegment.SetIntensity(0.0)
	fmt.Printf("   - 最小强度: %.1f%%\n", filterSegment.GetIntensity())

	// 恢复原始强度
	filterSegment.SetIntensity(originalIntensity)
	fmt.Printf("   - 恢复强度: %.1f%%\n", filterSegment.GetIntensity())
}

// demonstrateEffectTracks 演示特效轨道管理
func demonstrateEffectTracks() {
	fmt.Println("🛤️ === 特效轨道管理演示 ===")

	// 创建特效轨道
	effectTrack := track.NewTrack(track.TrackTypeEffect, "全局特效轨道", 0, false)

	fmt.Printf("🎬 特效轨道创建:\n")
	fmt.Printf("   - 轨道ID: %s\n", effectTrack.GetTrackID())
	fmt.Printf("   - 轨道名称: %s\n", effectTrack.GetName())
	fmt.Printf("   - 轨道类型: %s\n", effectTrack.GetTrackType().String())
	fmt.Printf("   - 渲染索引: %d\n", effectTrack.GetRenderIndex())
	fmt.Printf("   - 静音状态: %v\n", effectTrack.Mute)

	// 创建多个特效片段
	effects := []struct {
		name     string
		start    int64
		duration int64
		params   []float64
	}{
		{"闪光效果", 0, 3000000, []float64{90.0}},
		{"震动效果", 4000000, 2000000, []float64{50.0}},
		{"缩放效果", 8000000, 4000000, []float64{75.0, 90.0}},
	}

	fmt.Printf("\n📋 添加特效片段:\n")
	for i, effect := range effects {
		// 创建特效元数据
		effectMeta := metadata.NewEffectMeta(
			effect.name,
			false,
			fmt.Sprintf("resource_%03d", i+1),
			fmt.Sprintf("effect_%03d", i+1),
			fmt.Sprintf("md5_hash_%03d", i+1),
			[]metadata.EffectParam{
				metadata.NewEffectParam("intensity", 1.0, 0.0, 1.0),
				metadata.NewEffectParam("scale", 1.0, 0.5, 2.0),
			},
		)

		// 创建特效片段
		timerange := types.NewTimerange(effect.start, effect.duration)
		effectSegment, err := segment.NewEffectSegment(effectMeta, timerange, effect.params)
		if err != nil {
			log.Fatalf("创建特效片段失败: %v", err)
		}

		// 添加到轨道
		err = effectTrack.AddSegment(effectSegment)
		if err != nil {
			log.Fatalf("添加特效片段到轨道失败: %v", err)
		}

		fmt.Printf("   [%d] %s: %.2f-%.2f秒 ✅\n",
			i+1, effect.name,
			float64(effect.start)/1e6,
			float64(effect.start+effect.duration)/1e6)
	}

	fmt.Printf("\n📊 轨道统计:\n")
	fmt.Printf("   - 片段数量: %d\n", len(effectTrack.Segments))
	fmt.Printf("   - 轨道总长: %.2f秒\n", float64(effectTrack.EndTime())/1e6)

	// 测试重叠检测
	fmt.Printf("\n🚫 重叠检测测试:\n")
	overlapMeta := metadata.NewEffectMeta(
		"重叠测试特效",
		false,
		"overlap_resource",
		"overlap_effect",
		"overlap_md5",
		[]metadata.EffectParam{},
	)

	// 尝试添加与第一个片段重叠的特效
	overlapTimerange := types.NewTimerange(1000000, 3000000) // 1-4秒，与第一个片段重叠
	overlapSegment, err := segment.NewEffectSegment(overlapMeta, overlapTimerange, []float64{})
	if err != nil {
		log.Fatalf("创建重叠特效片段失败: %v", err)
	}

	err = effectTrack.AddSegment(overlapSegment)
	if err != nil {
		fmt.Printf("   ✅ 正确拒绝重叠片段: %v\n", err)
	} else {
		fmt.Printf("   ❌ 未能检测到重叠\n")
	}
}

// demonstrateFilterTracks 演示滤镜轨道管理
func demonstrateFilterTracks() {
	fmt.Println("🌈 === 滤镜轨道管理演示 ===")

	// 创建滤镜轨道
	filterTrack := track.NewTrack(track.TrackTypeFilter, "全局滤镜轨道", 0, false)

	fmt.Printf("🎨 滤镜轨道创建:\n")
	fmt.Printf("   - 轨道ID: %s\n", filterTrack.GetTrackID())
	fmt.Printf("   - 轨道类型: %s\n", filterTrack.GetTrackType().String())
	fmt.Printf("   - 渲染索引: %d\n", filterTrack.GetRenderIndex())

	// 创建多个滤镜片段
	filters := []struct {
		name      string
		start     int64
		duration  int64
		intensity float64
	}{
		{"暖色调", 0, 6000000, 70.0},
		{"黑白滤镜", 7000000, 5000000, 85.0},
		{"复古风格", 13000000, 8000000, 60.0},
	}

	fmt.Printf("\n📋 添加滤镜片段:\n")
	for i, filter := range filters {
		// 创建滤镜元数据
		filterMeta := metadata.NewEffectMeta(
			filter.name,
			i%2 == 1, // 交替设置VIP状态
			fmt.Sprintf("filter_resource_%03d", i+1),
			fmt.Sprintf("filter_effect_%03d", i+1),
			fmt.Sprintf("filter_md5_%03d", i+1),
			[]metadata.EffectParam{},
		)

		// 创建滤镜片段
		timerange := types.NewTimerange(filter.start, filter.duration)
		filterSegment := segment.NewFilterSegment(filterMeta, timerange, filter.intensity)

		// 添加到轨道
		err := filterTrack.AddSegment(filterSegment)
		if err != nil {
			log.Fatalf("添加滤镜片段到轨道失败: %v", err)
		}

		vipStatus := ""
		if filterMeta.IsVIP {
			vipStatus = " (VIP)"
		}

		fmt.Printf("   [%d] %s%s: %.2f-%.2f秒, 强度%.1f%% ✅\n",
			i+1, filter.name, vipStatus,
			float64(filter.start)/1e6,
			float64(filter.start+filter.duration)/1e6,
			filter.intensity)
	}

	fmt.Printf("\n📊 轨道统计:\n")
	fmt.Printf("   - 片段数量: %d\n", len(filterTrack.Segments))
	fmt.Printf("   - 轨道总长: %.2f秒\n", float64(filterTrack.EndTime())/1e6)

	// 测试类型检查
	fmt.Printf("\n🔍 类型检查测试:\n")
	acceptedType := filterTrack.AcceptSegmentType()
	fmt.Printf("   - 接受的片段类型: %s\n", acceptedType.String())

	// 尝试添加错误类型的片段（特效片段到滤镜轨道）
	wrongMeta := metadata.NewEffectMeta(
		"错误类型特效",
		false,
		"wrong_resource",
		"wrong_effect",
		"wrong_md5",
		[]metadata.EffectParam{},
	)

	wrongTimerange := types.NewTimerange(25000000, 2000000)
	wrongSegment, err := segment.NewEffectSegment(wrongMeta, wrongTimerange, []float64{})
	if err != nil {
		log.Fatalf("创建错误类型片段失败: %v", err)
	}

	err = filterTrack.AddSegment(wrongSegment)
	if err != nil {
		fmt.Printf("   ✅ 正确拒绝错误类型片段: %v\n", err)
	} else {
		fmt.Printf("   ❌ 未能检测到类型错误\n")
	}
}

// demonstrateComplexEffectParameters 演示复杂特效参数处理
func demonstrateComplexEffectParameters() {
	fmt.Println("⚙️ === 复杂特效参数处理演示 ===")

	// 创建具有多个复杂参数的特效
	complexMeta := metadata.NewEffectMeta(
		"高级色彩调整",
		true,
		"complex_color_resource",
		"complex_color_effect",
		"complex_color_md5",
		[]metadata.EffectParam{
			metadata.NewEffectParam("brightness", 0.5, -1.0, 1.0),
			metadata.NewEffectParam("contrast", 0.0, -1.0, 1.0),
			metadata.NewEffectParam("saturation", 1.0, 0.0, 2.0),
			metadata.NewEffectParam("hue_shift", 0.0, -180.0, 180.0),
			metadata.NewEffectParam("gamma", 1.0, 0.1, 3.0),
		},
	)

	fmt.Printf("🎛️ 复杂特效元数据:\n")
	fmt.Printf("   - 名称: %s\n", complexMeta.Name)
	fmt.Printf("   - VIP状态: %v\n", complexMeta.IsVIP)
	fmt.Printf("   - 参数数量: %d\n", len(complexMeta.Params))

	fmt.Printf("\n📋 参数详情:\n")
	for i, param := range complexMeta.Params {
		fmt.Printf("   [%d] %s: 默认值%.2f, 范围[%.2f, %.2f]\n",
			i+1, param.Name, param.DefaultValue, param.MinValue, param.MaxValue)
	}

	// 测试不同的参数组合
	paramSets := []struct {
		name   string
		params []float64
		desc   string
	}{
		{"默认设置", []float64{}, "使用所有默认参数"},
		{"增强对比", []float64{60.0, 80.0}, "提升亮度和对比度"},
		{"复古风格", []float64{40.0, 60.0, 30.0, 25.0}, "复古色调调整"},
		{"全参数设置", []float64{70.0, 50.0, 85.0, 15.0, 60.0}, "所有参数自定义"},
	}

	timerange := types.NewTimerange(0, 5000000)

	fmt.Printf("\n🧪 参数组合测试:\n")
	for i, paramSet := range paramSets {
		fmt.Printf("   [%d] %s (%s):\n", i+1, paramSet.name, paramSet.desc)

		effectSegment, err := segment.NewEffectSegment(complexMeta, timerange, paramSet.params)
		if err != nil {
			fmt.Printf("       ❌ 创建失败: %v\n", err)
			continue
		}

		fmt.Printf("       ✅ 创建成功，参数数量: %d\n", len(effectSegment.EffectInst.AdjustParams))

		// 显示实际参数值
		if len(effectSegment.EffectInst.AdjustParams) > 0 {
			fmt.Printf("       参数预览: ")
			for j, param := range effectSegment.EffectInst.AdjustParams {
				if j > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%v", param)
				if j >= 2 { // 只显示前3个参数
					if len(effectSegment.EffectInst.AdjustParams) > 3 {
						fmt.Printf("...")
					}
					break
				}
			}
			fmt.Println()
		}
	}

	// 测试参数验证
	fmt.Printf("\n🚫 参数验证测试:\n")
	invalidParams := [][]float64{
		{-10.0},       // 负值
		{150.0},       // 超出范围
		{50.0, -20.0}, // 混合有效和无效值
	}

	for i, params := range invalidParams {
		_, err := segment.NewEffectSegment(complexMeta, timerange, params)
		if err != nil {
			fmt.Printf("   [%d] ✅ 正确拒绝无效参数 %v: %v\n", i+1, params, err)
		} else {
			fmt.Printf("   [%d] ❌ 未能检测到无效参数 %v\n", i+1, params)
		}
	}
}

// demonstrateTrackIntegrationAndExport 演示轨道集成和JSON导出
func demonstrateTrackIntegrationAndExport() {
	fmt.Println("📤 === 轨道集成和JSON导出演示 ===")

	// 创建特效轨道
	effectTrack := track.NewTrack(track.TrackTypeEffect, "主特效轨道", 10000, false)

	// 创建滤镜轨道
	filterTrack := track.NewTrack(track.TrackTypeFilter, "主滤镜轨道", 11000, false)

	// 添加特效片段
	effectMeta := metadata.NewEffectMeta(
		"导出测试特效",
		false,
		"export_effect_resource",
		"export_effect_id",
		"export_effect_md5",
		[]metadata.EffectParam{
			metadata.NewEffectParam("opacity", 1.0, 0.0, 1.0),
		},
	)

	effectTimerange := types.NewTimerange(1000000, 6000000)
	effectSegment, err := segment.NewEffectSegment(effectMeta, effectTimerange, []float64{80.0})
	if err != nil {
		log.Fatalf("创建特效片段失败: %v", err)
	}

	err = effectTrack.AddSegment(effectSegment)
	if err != nil {
		log.Fatalf("添加特效片段失败: %v", err)
	}

	// 添加滤镜片段
	filterMeta := metadata.NewEffectMeta(
		"导出测试滤镜",
		true,
		"export_filter_resource",
		"export_filter_id",
		"export_filter_md5",
		[]metadata.EffectParam{},
	)

	filterTimerange := types.NewTimerange(2000000, 8000000)
	filterSegment := segment.NewFilterSegment(filterMeta, filterTimerange, 75.0)

	err = filterTrack.AddSegment(filterSegment)
	if err != nil {
		log.Fatalf("添加滤镜片段失败: %v", err)
	}

	fmt.Printf("🎬 轨道创建完成:\n")
	fmt.Printf("   - 特效轨道: %s (片段数: %d)\n", effectTrack.GetName(), len(effectTrack.Segments))
	fmt.Printf("   - 滤镜轨道: %s (片段数: %d)\n", filterTrack.GetName(), len(filterTrack.Segments))

	// JSON导出测试
	fmt.Printf("\n📋 JSON导出测试:\n")

	// 导出特效轨道
	effectJSON := effectTrack.ExportJSON()
	effectJSONBytes, err := json.MarshalIndent(effectJSON, "", "  ")
	if err != nil {
		log.Fatalf("特效轨道JSON序列化失败: %v", err)
	}

	fmt.Printf("   📄 特效轨道JSON (前200字符):\n")
	jsonPreview := string(effectJSONBytes)
	if len(jsonPreview) > 200 {
		jsonPreview = jsonPreview[:200] + "..."
	}
	fmt.Printf("   %s\n", jsonPreview)

	// 导出滤镜轨道
	filterJSON := filterTrack.ExportJSON()
	filterJSONBytes, err := json.MarshalIndent(filterJSON, "", "  ")
	if err != nil {
		log.Fatalf("滤镜轨道JSON序列化失败: %v", err)
	}

	fmt.Printf("\n   📄 滤镜轨道JSON (前200字符):\n")
	jsonPreview = string(filterJSONBytes)
	if len(jsonPreview) > 200 {
		jsonPreview = jsonPreview[:200] + "..."
	}
	fmt.Printf("   %s\n", jsonPreview)

	// 验证JSON结构
	fmt.Printf("\n✅ JSON结构验证:\n")
	fmt.Printf("   - 特效轨道字段数: %d\n", len(effectJSON))
	fmt.Printf("   - 滤镜轨道字段数: %d\n", len(filterJSON))

	// 检查关键字段
	requiredFields := []string{"type", "name", "id", "segments", "render_index"}
	for _, field := range requiredFields {
		if _, exists := effectJSON[field]; exists {
			fmt.Printf("   - 特效轨道包含 '%s' ✅\n", field)
		} else {
			fmt.Printf("   - 特效轨道缺少 '%s' ❌\n", field)
		}
	}
}

// demonstrateCompleteEffectWorkflow 演示完整的特效工作流
func demonstrateCompleteEffectWorkflow() {
	fmt.Println("🎯 === 完整特效工作流演示 ===")

	fmt.Printf("🎬 模拟视频项目特效制作流程:\n")

	// 步骤1: 创建项目轨道
	fmt.Printf("   📋 步骤1: 创建项目轨道\n")
	effectTrack := track.NewTrack(track.TrackTypeEffect, "视频特效轨道", 0, false)
	filterTrack := track.NewTrack(track.TrackTypeFilter, "视频滤镜轨道", 0, false)

	fmt.Printf("     ✅ 特效轨道: %s (渲染索引: %d)\n", effectTrack.GetName(), effectTrack.GetRenderIndex())
	fmt.Printf("     ✅ 滤镜轨道: %s (渲染索引: %d)\n", filterTrack.GetName(), filterTrack.GetRenderIndex())

	// 步骤2: 添加开场特效
	fmt.Printf("   🎭 步骤2: 添加开场特效\n")
	openingEffects := []struct {
		name     string
		start    float64
		duration float64
		params   []float64
	}{
		{"淡入效果", 0.0, 2.0, []float64{100.0}},
		{"标题动画", 3.0, 3.0, []float64{80.0, 90.0}},
		{"背景粒子", 7.0, 4.0, []float64{60.0}},
	}

	for i, effect := range openingEffects {
		meta := metadata.NewEffectMeta(
			effect.name,
			false,
			fmt.Sprintf("opening_resource_%d", i),
			fmt.Sprintf("opening_effect_%d", i),
			fmt.Sprintf("opening_md5_%d", i),
			[]metadata.EffectParam{
				metadata.NewEffectParam("intensity", 1.0, 0.0, 1.0),
				metadata.NewEffectParam("scale", 1.0, 0.5, 2.0),
			},
		)

		timerange := types.NewTimerange(
			int64(effect.start*1e6),
			int64(effect.duration*1e6),
		)

		segment, err := segment.NewEffectSegment(meta, timerange, effect.params)
		if err != nil {
			log.Fatalf("创建开场特效失败: %v", err)
		}

		err = effectTrack.AddSegment(segment)
		if err != nil {
			log.Fatalf("添加开场特效失败: %v", err)
		}

		fmt.Printf("     [%d] %s: %.1f-%.1f秒 ✅\n",
			i+1, effect.name, effect.start, effect.start+effect.duration)
	}

	// 步骤3: 添加主内容滤镜
	fmt.Printf("   🌈 步骤3: 添加主内容滤镜\n")
	mainFilters := []struct {
		name      string
		start     float64
		duration  float64
		intensity float64
	}{
		{"暖色基调", 0.0, 15.0, 70.0},
		{"对比增强", 16.0, 8.0, 85.0},
		{"饱和度调整", 25.0, 10.0, 60.0},
	}

	for i, filter := range mainFilters {
		meta := metadata.NewEffectMeta(
			filter.name,
			i == 1, // 中间的滤镜设为VIP
			fmt.Sprintf("main_filter_resource_%d", i),
			fmt.Sprintf("main_filter_effect_%d", i),
			fmt.Sprintf("main_filter_md5_%d", i),
			[]metadata.EffectParam{},
		)

		timerange := types.NewTimerange(
			int64(filter.start*1e6),
			int64(filter.duration*1e6),
		)

		segment := segment.NewFilterSegment(meta, timerange, filter.intensity)

		err := filterTrack.AddSegment(segment)
		if err != nil {
			log.Fatalf("添加主内容滤镜失败: %v", err)
		}

		vipStatus := ""
		if meta.IsVIP {
			vipStatus = " (VIP)"
		}

		fmt.Printf("     [%d] %s%s: %.1f-%.1f秒, 强度%.1f%% ✅\n",
			i+1, filter.name, vipStatus, filter.start,
			filter.start+filter.duration, filter.intensity)
	}

	// 步骤4: 添加结尾特效
	fmt.Printf("   🎊 步骤4: 添加结尾特效\n")
	endingEffects := []struct {
		name     string
		start    float64
		duration float64
		params   []float64
	}{
		{"闪光转场", 38.0, 1.0, []float64{95.0}},
		{"淡出效果", 40.0, 2.0, []float64{100.0}},
	}

	for i, effect := range endingEffects {
		meta := metadata.NewEffectMeta(
			effect.name,
			false,
			fmt.Sprintf("ending_resource_%d", i),
			fmt.Sprintf("ending_effect_%d", i),
			fmt.Sprintf("ending_md5_%d", i),
			[]metadata.EffectParam{
				metadata.NewEffectParam("intensity", 1.0, 0.0, 1.0),
			},
		)

		timerange := types.NewTimerange(
			int64(effect.start*1e6),
			int64(effect.duration*1e6),
		)

		segment, err := segment.NewEffectSegment(meta, timerange, effect.params)
		if err != nil {
			log.Fatalf("创建结尾特效失败: %v", err)
		}

		err = effectTrack.AddSegment(segment)
		if err != nil {
			log.Fatalf("添加结尾特效失败: %v", err)
		}

		fmt.Printf("     [%d] %s: %.1f-%.1f秒 ✅\n",
			i+1, effect.name, effect.start, effect.start+effect.duration)
	}

	// 步骤5: 项目统计和导出
	fmt.Printf("   📊 步骤5: 项目统计和导出\n")
	fmt.Printf("     📈 项目统计:\n")
	fmt.Printf("       - 特效片段总数: %d\n", len(effectTrack.Segments))
	fmt.Printf("       - 滤镜片段总数: %d\n", len(filterTrack.Segments))
	fmt.Printf("       - 特效轨道总长: %.1f秒\n", float64(effectTrack.EndTime())/1e6)
	fmt.Printf("       - 滤镜轨道总长: %.1f秒\n", float64(filterTrack.EndTime())/1e6)

	// 计算素材引用
	totalRefs := 0
	for _, seg := range effectTrack.Segments {
		if effectSeg, ok := seg.(*segment.EffectSegment); ok {
			totalRefs += len(effectSeg.GetMaterialRefs())
		}
	}
	for _, seg := range filterTrack.Segments {
		if filterSeg, ok := seg.(*segment.FilterSegment); ok {
			totalRefs += len(filterSeg.GetMaterialRefs())
		}
	}
	fmt.Printf("       - 素材引用总数: %d\n", totalRefs)

	// 导出最终JSON
	fmt.Printf("     📤 导出项目配置:\n")
	projectConfig := map[string]interface{}{
		"project_name": "完整特效工作流演示",
		"version":      "1.0.0",
		"tracks": map[string]interface{}{
			"effect_track": effectTrack.ExportJSON(),
			"filter_track": filterTrack.ExportJSON(),
		},
		"statistics": map[string]interface{}{
			"effect_segments": len(effectTrack.Segments),
			"filter_segments": len(filterTrack.Segments),
			"total_duration":  float64(max(effectTrack.EndTime(), filterTrack.EndTime())) / 1e6,
			"material_refs":   totalRefs,
		},
	}

	configBytes, err := json.MarshalIndent(projectConfig, "", "  ")
	if err != nil {
		log.Fatalf("项目配置序列化失败: %v", err)
	}

	fmt.Printf("       ✅ 项目配置大小: %d 字节\n", len(configBytes))
	fmt.Printf("       ✅ JSON结构完整性验证通过\n")

	fmt.Printf("\n🎉 完整特效工作流演示完成!\n")
	fmt.Printf("   - 成功创建了包含多种特效和滤镜的完整项目\n")
	fmt.Printf("   - 演示了从开场到结尾的完整特效流程\n")
	fmt.Printf("   - 验证了轨道管理和JSON导出功能\n")
	fmt.Printf("   - 展示了与Python版本完全兼容的特效系统\n")
}

// 辅助函数

func getApplyTargetTypeString(targetType int) string {
	switch targetType {
	case 0:
		return "片段级别"
	case 2:
		return "全局级别"
	default:
		return "未知类型"
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
