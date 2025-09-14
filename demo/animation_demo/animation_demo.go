// Animation系统演示程序
// 展示Go版本的Animation系统功能，包括视频动画、文本动画和CapCut动画
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zhangshican/go-capcut/internal/animation"
	"github.com/zhangshican/go-capcut/internal/metadata"
)

func main() {
	fmt.Println("=== Go版本 Animation系统演示程序 ===")
	fmt.Println()

	// 演示1: 视频动画
	demonstrateVideoAnimations()
	fmt.Println()

	// 演示2: 文本动画
	demonstrateTextAnimations()
	fmt.Println()

	// 演示3: CapCut动画
	demonstrateCapCutAnimations()
	fmt.Println()

	// 演示4: 复杂动画序列
	demonstrateComplexAnimationSequence()
	fmt.Println()

	// 演示5: 动画约束验证
	demonstrateAnimationConstraints()
	fmt.Println()

	// 演示6: JSON导出兼容性
	demonstrateJSONCompatibility()
}

// demonstrateVideoAnimations 演示视频动画功能
func demonstrateVideoAnimations() {
	fmt.Println("📹 === 视频动画演示 ===")

	// 创建视频入场动画
	videoIntro, err := animation.NewVideoAnimation(metadata.IntroType渐显, 0, 500000)
	if err != nil {
		log.Fatalf("创建视频入场动画失败: %v", err)
	}

	fmt.Printf("✅ 创建视频入场动画: %s\n", videoIntro.Animation.Name)
	fmt.Printf("   - 类型: %s\n", videoIntro.Animation.AnimationType)
	fmt.Printf("   - 开始时间: %d微秒\n", videoIntro.Animation.Start)
	fmt.Printf("   - 持续时间: %d微秒\n", videoIntro.Animation.Duration)
	fmt.Printf("   - 资源ID: %s\n", videoIntro.Animation.ResourceID)

	// 创建视频出场动画
	videoOutro, err := animation.NewVideoAnimation(metadata.OutroType缩小, 1500000, 500000)
	if err != nil {
		log.Fatalf("创建视频出场动画失败: %v", err)
	}

	fmt.Printf("✅ 创建视频出场动画: %s\n", videoOutro.Animation.Name)

	// 创建组合动画
	groupAnim, err := animation.NewVideoAnimation(metadata.GroupAnimationType三分割, 0, 2000000)
	if err != nil {
		log.Fatalf("创建组合动画失败: %v", err)
	}

	fmt.Printf("✅ 创建组合动画: %s\n", groupAnim.Animation.Name)
	fmt.Printf("   - 类型: %s\n", groupAnim.Animation.AnimationType)
}

// demonstrateTextAnimations 演示文本动画功能
func demonstrateTextAnimations() {
	fmt.Println("📝 === 文本动画演示 ===")

	// 创建文本入场动画
	textIntro, err := animation.NewTextAnimation(metadata.TextIntroType打字机, 0, 800000)
	if err != nil {
		log.Fatalf("创建文本入场动画失败: %v", err)
	}

	fmt.Printf("✅ 创建文本入场动画: %s\n", textIntro.Animation.Name)
	fmt.Printf("   - 类型: %s\n", textIntro.Animation.AnimationType)
	fmt.Printf("   - 是否视频动画: %v\n", textIntro.Animation.IsVideoAnimation)

	// 创建文本出场动画
	textOutro, err := animation.NewTextAnimation(metadata.TextOutroType渐隐, 2000000, 500000)
	if err != nil {
		log.Fatalf("创建文本出场动画失败: %v", err)
	}

	fmt.Printf("✅ 创建文本出场动画: %s\n", textOutro.Animation.Name)

	// 创建文本循环动画
	textLoop, err := animation.NewTextAnimation(metadata.TextLoopAnimType跳动, 0, 0)
	if err != nil {
		log.Fatalf("创建文本循环动画失败: %v", err)
	}

	fmt.Printf("✅ 创建文本循环动画: %s\n", textLoop.Animation.Name)
	fmt.Printf("   - 类型: %s\n", textLoop.Animation.AnimationType)
	fmt.Printf("   - 持续时间: %d微秒 (0表示无限循环)\n", textLoop.Animation.Duration)
}

// demonstrateCapCutAnimations 演示CapCut动画功能
func demonstrateCapCutAnimations() {
	fmt.Println("🎬 === CapCut动画演示 ===")

	// 创建CapCut视频动画
	capCutVideo, err := animation.NewVideoAnimation(metadata.CapCutIntroTypeFadeIn, 0, 500000)
	if err != nil {
		log.Fatalf("创建CapCut视频动画失败: %v", err)
	}

	fmt.Printf("✅ 创建CapCut视频动画: %s\n", capCutVideo.Animation.Name)
	fmt.Printf("   - 效果ID: %s\n", capCutVideo.Animation.EffectID)

	// 创建CapCut文本动画
	capCutText, err := animation.NewTextAnimation(metadata.CapCutTextIntroTypeTypewriter, 0, 1000000)
	if err != nil {
		log.Fatalf("创建CapCut文本动画失败: %v", err)
	}

	fmt.Printf("✅ 创建CapCut文本动画: %s\n", capCutText.Animation.Name)

	// 创建CapCut组合动画
	capCutGroup, err := animation.NewVideoAnimation(metadata.CapCutGroupAnimationTypeRotation, 0, 1500000)
	if err != nil {
		log.Fatalf("创建CapCut组合动画失败: %v", err)
	}

	fmt.Printf("✅ 创建CapCut组合动画: %s\n", capCutGroup.Animation.Name)
}

// demonstrateComplexAnimationSequence 演示复杂动画序列
func demonstrateComplexAnimationSequence() {
	fmt.Println("🎭 === 复杂动画序列演示 ===")

	// 创建视频片段动画序列
	videoSegmentAnims := animation.NewSegmentAnimations()
	fmt.Printf("✅ 创建视频片段动画序列，ID: %s\n", videoSegmentAnims.AnimationID)

	// 添加入场动画
	err := videoSegmentAnims.AddVideoAnimation(metadata.IntroType渐显, 0, 500000)
	if err != nil {
		log.Fatalf("添加入场动画失败: %v", err)
	}
	fmt.Printf("✅ 添加入场动画: 渐显\n")

	// 添加出场动画
	err = videoSegmentAnims.AddVideoAnimation(metadata.OutroType缩小, 1500000, 500000)
	if err != nil {
		log.Fatalf("添加出场动画失败: %v", err)
	}
	fmt.Printf("✅ 添加出场动画: 缩小\n")

	fmt.Printf("📊 当前动画序列包含 %d 个动画\n", len(videoSegmentAnims.Animations))

	// 获取入场动画的时间范围
	introTimerange, err := videoSegmentAnims.GetAnimationTimerange(animation.AnimationTypeIn)
	if err != nil {
		log.Fatalf("获取入场动画时间范围失败: %v", err)
	}
	if introTimerange != nil {
		fmt.Printf("📅 入场动画时间范围: %d微秒 - %d微秒\n",
			introTimerange.Start, introTimerange.Start+introTimerange.Duration)
	}

	// 创建文本片段动画序列
	textSegmentAnims := animation.NewSegmentAnimations()
	fmt.Printf("\n✅ 创建文本片段动画序列，ID: %s\n", textSegmentAnims.AnimationID)

	// 添加文本入场、出场和循环动画
	err = textSegmentAnims.AddTextAnimation(metadata.TextIntroType打字机, 0, 800000)
	if err != nil {
		log.Fatalf("添加文本入场动画失败: %v", err)
	}

	err = textSegmentAnims.AddTextAnimation(metadata.TextOutroType渐隐, 2000000, 500000)
	if err != nil {
		log.Fatalf("添加文本出场动画失败: %v", err)
	}

	err = textSegmentAnims.AddTextAnimation(metadata.TextLoopAnimType跳动, 0, 0)
	if err != nil {
		log.Fatalf("添加文本循环动画失败: %v", err)
	}

	fmt.Printf("✅ 文本动画序列完成，包含 %d 个动画\n", len(textSegmentAnims.Animations))
}

// demonstrateAnimationConstraints 演示动画约束验证
func demonstrateAnimationConstraints() {
	fmt.Println("⚠️  === 动画约束验证演示 ===")

	segmentAnims := animation.NewSegmentAnimations()

	// 添加第一个入场动画
	err := segmentAnims.AddVideoAnimation(metadata.IntroType渐显, 0, 500000)
	if err != nil {
		log.Fatalf("添加第一个入场动画失败: %v", err)
	}
	fmt.Printf("✅ 成功添加入场动画: 渐显\n")

	// 尝试添加第二个入场动画（应该失败）
	err = segmentAnims.AddVideoAnimation(metadata.IntroType放大, 0, 500000)
	if err != nil {
		fmt.Printf("❌ 预期失败：%v\n", err)
	} else {
		fmt.Printf("⚠️  意外成功：应该禁止添加重复类型的动画\n")
	}

	// 测试组合动画约束
	groupSegmentAnims := animation.NewSegmentAnimations()

	// 先添加组合动画
	err = groupSegmentAnims.AddVideoAnimation(metadata.GroupAnimationType三分割, 0, 2000000)
	if err != nil {
		log.Fatalf("添加组合动画失败: %v", err)
	}
	fmt.Printf("✅ 成功添加组合动画: 三分割\n")

	// 尝试再添加其他动画（应该失败）
	err = groupSegmentAnims.AddVideoAnimation(metadata.IntroType渐显, 0, 500000)
	if err != nil {
		fmt.Printf("❌ 预期失败：%v\n", err)
	} else {
		fmt.Printf("⚠️  意外成功：应该禁止在组合动画后添加其他动画\n")
	}
}

// demonstrateJSONCompatibility 演示JSON导出兼容性
func demonstrateJSONCompatibility() {
	fmt.Println("📄 === JSON导出兼容性演示 ===")

	// 创建动画序列
	segmentAnims := animation.NewSegmentAnimations()

	// 添加多种动画
	err := segmentAnims.AddVideoAnimation(metadata.IntroType渐显, 0, 500000)
	if err != nil {
		log.Fatalf("添加动画失败: %v", err)
	}

	err = segmentAnims.AddVideoAnimation(metadata.OutroType缩小, 1500000, 500000)
	if err != nil {
		log.Fatalf("添加动画失败: %v", err)
	}

	// 导出JSON
	jsonData := segmentAnims.ExportJSON()
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("JSON序列化失败: %v", err)
	}

	fmt.Printf("✅ 动画序列JSON导出:\n%s\n", string(jsonBytes))

	// 验证JSON结构
	fmt.Printf("📊 JSON结构验证:\n")
	fmt.Printf("   - ID: %v\n", jsonData["id"])
	fmt.Printf("   - 类型: %v\n", jsonData["type"])
	fmt.Printf("   - 多语言: %v\n", jsonData["multi_language_current"])

	if animations, ok := jsonData["animations"].([]map[string]interface{}); ok {
		fmt.Printf("   - 动画数量: %d\n", len(animations))
		for i, anim := range animations {
			fmt.Printf("     [%d] %s (%s)\n", i, anim["name"], anim["type"])
		}
	}

	// 演示单个动画的JSON导出
	singleAnim, err := animation.NewVideoAnimation(metadata.CapCutIntroTypeFadeIn, 100000, 600000)
	if err != nil {
		log.Fatalf("创建单个动画失败: %v", err)
	}

	singleJSON := singleAnim.Animation.ExportJSON()
	singleBytes, err := json.MarshalIndent(singleJSON, "", "  ")
	if err != nil {
		log.Fatalf("单个动画JSON序列化失败: %v", err)
	}

	fmt.Printf("\n✅ 单个动画JSON导出:\n%s\n", string(singleBytes))
}

// demonstrateMetadataSearch 演示元数据搜索功能
func demonstrateMetadataSearch() {
	fmt.Println("🔍 === 元数据搜索演示 ===")

	// 搜索入场动画
	foundIntro, err := metadata.FindIntroTypeByName("渐显")
	if err != nil {
		log.Fatalf("搜索入场动画失败: %v", err)
	}
	fmt.Printf("✅ 找到入场动画: %s\n", foundIntro.GetName())

	// 搜索文本动画
	foundTextIntro, err := metadata.FindTextIntroTypeByName("打字机")
	if err != nil {
		log.Fatalf("搜索文本动画失败: %v", err)
	}
	fmt.Printf("✅ 找到文本动画: %s\n", foundTextIntro.GetName())

	// 搜索CapCut动画
	foundCapCut, err := metadata.FindCapCutIntroTypeByName("Fade In")
	if err != nil {
		log.Fatalf("搜索CapCut动画失败: %v", err)
	}
	fmt.Printf("✅ 找到CapCut动画: %s\n", foundCapCut.GetName())

	// 列出所有可用的动画类型
	fmt.Printf("\n📋 可用的动画类型:\n")

	allIntros := metadata.GetAllIntroTypes()
	fmt.Printf("   入场动画: %d个\n", len(allIntros))

	allOutros := metadata.GetAllOutroTypes()
	fmt.Printf("   出场动画: %d个\n", len(allOutros))

	allGroups := metadata.GetAllGroupAnimationTypes()
	fmt.Printf("   组合动画: %d个\n", len(allGroups))

	allTextIntros := metadata.GetAllTextIntroTypes()
	fmt.Printf("   文本入场动画: %d个\n", len(allTextIntros))

	allTextOutros := metadata.GetAllTextOutroTypes()
	fmt.Printf("   文本出场动画: %d个\n", len(allTextOutros))

	allTextLoops := metadata.GetAllTextLoopAnimTypes()
	fmt.Printf("   文本循环动画: %d个\n", len(allTextLoops))
}
