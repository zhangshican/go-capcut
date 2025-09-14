// metadata_demo.go - 演示Metadata系统功能
// 展示各种元数据类型的创建、查询和使用
package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangshican/go-capcut/internal/metadata"
)

func main() {
	fmt.Println("=== CapCut Go Metadata系统演示 ===\n")

	// 演示动画元数据
	demonstrateAnimationMetadata()

	// 演示特效元数据
	demonstrateEffectMetadata()

	// 演示蒙版元数据
	demonstrateMaskMetadata()

	// 演示转场元数据
	demonstrateTransitionMetadata()

	// 演示滤镜元数据
	demonstrateFilterMetadata()

	// 演示字体元数据
	demonstrateFontMetadata()

	// 演示音频特效元数据
	demonstrateAudioEffectMetadata()

	// 演示CapCut特有元数据
	demonstrateCapCutMetadata()

	// 演示查找功能
	demonstrateSearchFunctionality()

	// 演示JSON序列化
	demonstrateJSONSerialization()

	fmt.Println("\n=== 演示完成 ===")
}

// demonstrateAnimationMetadata 演示动画元数据
func demonstrateAnimationMetadata() {
	fmt.Println("🎬 动画元数据演示")
	fmt.Println("================")

	// 获取所有入场动画
	intros := metadata.GetAllIntroTypes()
	fmt.Printf("📥 共有 %d 种入场动画：\n", len(intros))
	for i, intro := range intros {
		if i >= 5 { // 只显示前5个
			fmt.Printf("   ... 还有 %d 个\n", len(intros)-5)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, intro.GetName())
	}

	// 获取所有出场动画
	outros := metadata.GetAllOutroTypes()
	fmt.Printf("\n📤 共有 %d 种出场动画：\n", len(outros))
	for i, outro := range outros {
		if i >= 3 { // 只显示前3个
			fmt.Printf("   ... 还有 %d 个\n", len(outros)-3)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, outro.GetName())
	}

	// 获取所有组合动画
	groups := metadata.GetAllGroupAnimationTypes()
	fmt.Printf("\n🔄 共有 %d 种组合动画：\n", len(groups))
	for i, group := range groups {
		if i >= 3 {
			fmt.Printf("   ... 还有 %d 个\n", len(groups)-3)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, group.GetName())
	}

	// 演示文字动画
	textIntros := metadata.GetAllTextIntroTypes()
	fmt.Printf("\n📝 共有 %d 种文字入场动画：\n", len(textIntros))
	for i, textIntro := range textIntros {
		if i >= 3 {
			fmt.Printf("   ... 还有 %d 个\n", len(textIntros)-3)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, textIntro.GetName())
	}

	fmt.Println()
}

// demonstrateEffectMetadata 演示特效元数据
func demonstrateEffectMetadata() {
	fmt.Println("✨ 特效元数据演示")
	fmt.Println("================")

	// 创建一个特效
	params := []metadata.EffectParam{
		metadata.NewEffectParam("强度", 80.0, 0.0, 100.0),
		metadata.NewEffectParam("速度", 60.0, 10.0, 90.0),
		metadata.NewEffectParam("透明度", 90.0, 0.0, 100.0),
	}

	effect := metadata.NewEffectMeta("梦幻光效", true, "effect_dream_001", "eff_123456", "hash_abcdef", params)

	fmt.Printf("📋 特效名称: %s\n", effect.Name)
	fmt.Printf("💎 VIP特效: %v\n", effect.IsVIP)
	fmt.Printf("🆔 资源ID: %s\n", effect.ResourceID)
	fmt.Printf("🔧 参数数量: %d\n", len(effect.Params))

	// 演示参数解析
	fmt.Println("\n🎛️ 参数解析演示:")
	userParams := []float64{70.0, 85.0, 95.0} // 用户输入的百分比值
	instances, err := effect.ParseParams(userParams)
	if err != nil {
		fmt.Printf("❌ 参数解析失败: %v\n", err)
	} else {
		fmt.Println("✅ 参数解析成功:")
		for _, instance := range instances {
			fmt.Printf("   %s: %.1f%% -> %.2f (范围: %.1f-%.1f)\n",
				instance.Name, userParams[instance.Index], instance.Value,
				instance.MinValue, instance.MaxValue)
		}
	}

	fmt.Println()
}

// demonstrateMaskMetadata 演示蒙版元数据
func demonstrateMaskMetadata() {
	fmt.Println("🎭 蒙版元数据演示")
	fmt.Println("================")

	// 获取所有基础蒙版
	masks := metadata.GetAllMaskTypes()
	fmt.Printf("🎨 基础蒙版 (%d种):\n", len(masks))
	for i, mask := range masks {
		if i >= 4 {
			fmt.Printf("   ... 还有 %d 个\n", len(masks)-4)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, mask.GetName())
	}

	// 获取所有CapCut蒙版
	capCutMasks := metadata.GetAllCapCutMaskTypes()
	fmt.Printf("\n🚀 CapCut高级蒙版 (%d种):\n", len(capCutMasks))
	for i, mask := range capCutMasks {
		if i >= 4 {
			fmt.Printf("   ... 还有 %d 个\n", len(capCutMasks)-4)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, mask.GetName())
	}

	fmt.Println()
}

// demonstrateTransitionMetadata 演示转场元数据
func demonstrateTransitionMetadata() {
	fmt.Println("🌊 转场元数据演示")
	fmt.Println("================")

	// 获取所有基础转场
	transitions := metadata.GetAllTransitionTypes()
	fmt.Printf("📽️ 基础转场 (%d种):\n", len(transitions))
	for i, transition := range transitions {
		if i >= 5 {
			fmt.Printf("   ... 还有 %d 个\n", len(transitions)-5)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, transition.GetName())
	}

	// 获取所有CapCut转场
	capCutTransitions := metadata.GetAllCapCutTransitionTypes()
	fmt.Printf("\n🎪 CapCut高级转场 (%d种):\n", len(capCutTransitions))
	for i, transition := range capCutTransitions {
		if i >= 5 {
			fmt.Printf("   ... 还有 %d 个\n", len(capCutTransitions)-5)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, transition.GetName())
	}

	fmt.Println()
}

// demonstrateFilterMetadata 演示滤镜元数据
func demonstrateFilterMetadata() {
	fmt.Println("🌈 滤镜元数据演示")
	fmt.Println("================")

	// 获取所有滤镜分类
	categories := metadata.GetAllFilterCategories()
	fmt.Printf("📂 滤镜分类 (%d种):\n", len(categories))
	for _, category := range categories {
		filters := metadata.GetFiltersByCategory(category)
		fmt.Printf("   %s: %d个滤镜\n", category, len(filters))
	}

	// 演示人像滤镜
	fmt.Println("\n👤 人像滤镜详情:")
	portraitFilters := metadata.GetFiltersByCategory("人像")
	for i, filter := range portraitFilters {
		if i >= 3 {
			fmt.Printf("   ... 还有 %d 个\n", len(portraitFilters)-3)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, filter.GetName())
	}

	fmt.Println()
}

// demonstrateFontMetadata 演示字体元数据
func demonstrateFontMetadata() {
	fmt.Println("🔤 字体元数据演示")
	fmt.Println("================")

	// 获取所有字体分类
	categories := metadata.GetAllFontCategories()
	fmt.Printf("📝 字体分类 (%d种):\n", len(categories))
	for _, category := range categories {
		fonts := metadata.GetFontsByCategory(category)
		fmt.Printf("   %s: %d个字体\n", category, len(fonts))
	}

	// 演示支持的语言
	languages := metadata.GetSupportedLanguages()
	fmt.Printf("\n🌍 支持语言 (%d种):\n", len(languages))
	for _, lang := range languages {
		fonts := metadata.GetFontsByLanguage(lang)
		langName := map[string]string{
			"zh-CN": "简体中文",
			"en-US": "英文",
		}[lang]
		fmt.Printf("   %s (%s): %d个字体\n", langName, lang, len(fonts))
	}

	fmt.Println()
}

// demonstrateAudioEffectMetadata 演示音频特效元数据
func demonstrateAudioEffectMetadata() {
	fmt.Println("🎵 音频特效元数据演示")
	fmt.Println("====================")

	// 获取所有音频特效分类
	categories := metadata.GetAllAudioEffectCategories()
	fmt.Printf("🎧 音频特效分类 (%d种):\n", len(categories))
	for _, category := range categories {
		effects := metadata.GetAudioEffectsByCategory(category)
		fmt.Printf("   %s: %d个特效\n", category, len(effects))
	}

	// 演示环境音效
	fmt.Println("\n🌿 环境音效详情:")
	sceneEffects := metadata.GetAllAudioSceneEffectTypes()
	for i, effect := range sceneEffects {
		if i >= 4 {
			fmt.Printf("   ... 还有 %d 个\n", len(sceneEffects)-4)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, effect.GetName())
	}

	// 演示音调特效
	fmt.Println("\n🎚️ 音调特效详情:")
	toneEffects := metadata.GetAllToneEffectTypes()
	for i, effect := range toneEffects {
		if i >= 4 {
			fmt.Printf("   ... 还有 %d 个\n", len(toneEffects)-4)
			break
		}
		fmt.Printf("   %d. %s\n", i+1, effect.GetName())
	}

	fmt.Println()
}

// demonstrateCapCutMetadata 演示CapCut特有元数据
func demonstrateCapCutMetadata() {
	fmt.Println("🚀 CapCut特有元数据演示")
	fmt.Println("======================")

	// CapCut动画
	fmt.Println("🎭 CapCut高级动画:")
	capCutIntros := metadata.GetAllCapCutIntroTypes()
	fmt.Printf("   入场动画: %d种\n", len(capCutIntros))

	capCutOutros := metadata.GetAllCapCutOutroTypes()
	fmt.Printf("   出场动画: %d种\n", len(capCutOutros))

	capCutGroups := metadata.GetAllCapCutGroupAnimationTypes()
	fmt.Printf("   组合动画: %d种\n", len(capCutGroups))

	// CapCut音频特效
	fmt.Println("\n🎤 CapCut音频特效:")
	capCutVoiceFilters := metadata.GetAllCapCutVoiceFiltersEffectTypes()
	fmt.Printf("   语音滤镜: %d种\n", len(capCutVoiceFilters))

	capCutVoiceChars := metadata.GetAllCapCutVoiceCharactersEffectTypes()
	fmt.Printf("   语音角色: %d种\n", len(capCutVoiceChars))

	capCutS2S := metadata.GetAllCapCutSpeechToSongEffectTypes()
	fmt.Printf("   语音转歌声: %d种\n", len(capCutS2S))

	// 演示AI功能
	fmt.Println("\n🤖 AI智能功能示例:")
	for i, intro := range capCutIntros {
		if i >= 3 {
			break
		}
		fmt.Printf("   %d. %s\n", i+1, intro.GetName())
	}

	fmt.Println()
}

// demonstrateSearchFunctionality 演示查找功能
func demonstrateSearchFunctionality() {
	fmt.Println("🔍 查找功能演示")
	fmt.Println("===============")

	// 演示精确查找
	fmt.Println("🎯 精确查找演示:")

	testCases := []struct {
		name     string
		findFunc func(string) (metadata.EffectEnumerable, error)
		target   string
	}{
		{"入场动画", metadata.FindIntroByName, "缩小"},
		{"蒙版", metadata.FindMaskByName, "圆形"},
		{"转场", metadata.FindTransitionByName, "淡入淡出"},
		{"滤镜", metadata.FindFilterByName, "自然"},
		{"字体", metadata.FindFontByName, "苹方"},
	}

	for _, tc := range testCases {
		result, err := tc.findFunc(tc.target)
		if err != nil {
			fmt.Printf("   ❌ %s '%s': %v\n", tc.name, tc.target, err)
		} else {
			fmt.Printf("   ✅ %s '%s': 找到 '%s'\n", tc.name, tc.target, result.GetName())
		}
	}

	// 演示错误处理
	fmt.Println("\n❌ 错误处理演示:")
	_, err := metadata.FindIntroByName("不存在的动画")
	if err != nil {
		fmt.Printf("   查找不存在的动画: %v\n", err)
	}

	fmt.Println()
}

// demonstrateJSONSerialization 演示JSON序列化
func demonstrateJSONSerialization() {
	fmt.Println("📄 JSON序列化演示")
	fmt.Println("=================")

	// 演示动画元数据序列化
	fmt.Println("🎬 动画元数据JSON序列化:")
	animMeta := metadata.NewAnimationMeta("测试动画", true, 2.5, "anim_123", "effect_456", "hash_789")

	animJSON, err := json.MarshalIndent(animMeta, "", "  ")
	if err != nil {
		fmt.Printf("   ❌ 序列化失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 序列化成功:\n%s\n", string(animJSON))
	}

	// 演示特效参数实例JSON导出
	fmt.Println("🎛️ 特效参数实例JSON导出:")
	param := metadata.NewEffectParam("亮度", 75.0, 0.0, 100.0)
	instance := metadata.NewEffectParamInstance(param, 0, 85.0)

	instanceJSON := instance.ExportJSON()
	jsonBytes, err := json.MarshalIndent(instanceJSON, "", "  ")
	if err != nil {
		fmt.Printf("   ❌ 导出失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 导出成功:\n%s\n", string(jsonBytes))
	}

	// 演示蒙版元数据序列化
	fmt.Println("🎭 蒙版元数据JSON序列化:")
	maskMeta := metadata.NewMaskMeta("心形蒙版", "heart", "mask_heart_001", "effect_mask_001", "heart_hash", 1.2)

	maskJSON, err := json.MarshalIndent(maskMeta, "", "  ")
	if err != nil {
		fmt.Printf("   ❌ 序列化失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 序列化成功:\n%s\n", string(maskJSON))
	}

	fmt.Println()
}
