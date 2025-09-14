// Draft文件夹系统演示程序
// 展示Go版本的Draft文件夹系统功能，包括草稿管理、文件夹操作、模板复制等
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/zhangshican/go-capcut/internal/draft"
)

func main() {
	fmt.Println("=== Go版本 Draft文件夹系统演示程序 ===")
	fmt.Println()

	// 演示1: 创建和管理草稿文件夹
	demonstrateFolderManagement()
	fmt.Println()

	// 演示2: 草稿列表和信息
	demonstrateDraftListing()
	fmt.Println()

	// 演示3: 草稿复制和模板功能
	demonstrateDraftDuplication()
	fmt.Println()

	// 演示4: 草稿删除和清理
	demonstrateDraftDeletion()
	fmt.Println()

	// 演示5: 素材检查功能
	demonstrateMaterialInspection()
	fmt.Println()

	// 演示6: 完整的草稿管理工作流
	demonstrateCompleteWorkflow()
}

// demonstrateFolderManagement 演示文件夹管理
func demonstrateFolderManagement() {
	fmt.Println("📁 === 草稿文件夹管理演示 ===")

	// 创建临时草稿文件夹
	tempDir, err := os.MkdirTemp("", "draft_demo")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fmt.Printf("📂 创建演示文件夹: %s\n", tempDir)

	// 创建DraftFolder管理器
	df, err := draft.NewDraftFolder(tempDir)
	if err != nil {
		log.Fatalf("创建DraftFolder失败: %v", err)
	}

	fmt.Printf("✅ 成功创建草稿文件夹管理器\n")
	fmt.Printf("   - 管理路径: %s\n", df.FolderPath)

	// 测试不存在的路径
	fmt.Printf("\n🚫 测试不存在的路径:\n")
	nonExistentPath := filepath.Join(tempDir, "non_existent")
	_, err = draft.NewDraftFolder(nonExistentPath)
	if err != nil {
		fmt.Printf("   ✅ 正确处理不存在的路径: %v\n", err)
	}

	// 创建一些示例草稿文件夹
	draftNames := []string{"项目A", "项目B", "测试草稿", "备份草稿"}
	fmt.Printf("\n📋 创建示例草稿:\n")

	for i, name := range draftNames {
		draftPath := filepath.Join(tempDir, name)
		if err := os.Mkdir(draftPath, 0755); err != nil {
			log.Fatalf("创建草稿文件夹失败: %v", err)
		}

		// 为部分草稿创建draft_info.json
		if i < 3 {
			createSampleDraftInfo(draftPath, name)
		}

		fmt.Printf("   - %s: %s\n", name, getStatusText(i < 3))
	}

	// 验证草稿存在性
	fmt.Printf("\n🔍 验证草稿存在性:\n")
	for _, name := range draftNames {
		exists := df.DraftExists(name)
		fmt.Printf("   - %s: %v\n", name, exists)
	}

	fmt.Printf("\n📊 文件夹管理总结:\n")
	fmt.Printf("   - 管理路径: %s\n", df.FolderPath)
	fmt.Printf("   - 创建草稿: %d个\n", len(draftNames))
	fmt.Printf("   - 有效草稿: 3个 (包含draft_info.json)\n")
}

// demonstrateDraftListing 演示草稿列表功能
func demonstrateDraftListing() {
	fmt.Println("📋 === 草稿列表演示 ===")

	// 创建临时草稿文件夹
	tempDir, err := os.MkdirTemp("", "draft_demo")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 创建DraftFolder管理器
	df, err := draft.NewDraftFolder(tempDir)
	if err != nil {
		log.Fatalf("创建DraftFolder失败: %v", err)
	}

	// 创建不同类型的草稿
	draftConfigs := []struct {
		name     string
		hasInfo  bool
		fps      int
		width    int
		height   int
		duration int64
	}{
		{"高清视频项目", true, 30, 1920, 1080, 15000000},
		{"竖屏短视频", true, 25, 720, 1280, 8000000},
		{"4K电影项目", true, 24, 3840, 2160, 120000000},
		{"测试项目", false, 0, 0, 0, 0},
		{"空草稿", false, 0, 0, 0, 0},
	}

	fmt.Printf("📝 创建多样化草稿:\n")
	for _, config := range draftConfigs {
		draftPath := filepath.Join(tempDir, config.name)
		if err := os.Mkdir(draftPath, 0755); err != nil {
			log.Fatalf("创建草稿文件夹失败: %v", err)
		}

		if config.hasInfo {
			createDetailedDraftInfo(draftPath, config.name, config.fps, config.width, config.height, config.duration)
			fmt.Printf("   ✅ %s: %dx%d@%dfps, %.1f秒\n",
				config.name, config.width, config.height, config.fps, float64(config.duration)/1e6)
		} else {
			fmt.Printf("   📁 %s: 空草稿文件夹\n", config.name)
		}

		// 添加一些随机延迟以确保修改时间不同
		time.Sleep(10 * time.Millisecond)
	}

	// 基本列表功能
	fmt.Printf("\n📂 基本草稿列表:\n")
	drafts, err := df.ListDrafts()
	if err != nil {
		log.Fatalf("列出草稿失败: %v", err)
	}

	sort.Strings(drafts) // 排序以便展示
	for i, name := range drafts {
		fmt.Printf("   [%d] %s\n", i+1, name)
	}

	// 详细信息列表
	fmt.Printf("\n📊 详细草稿信息:\n")
	infos, err := df.ListDraftsWithInfo()
	if err != nil {
		log.Fatalf("获取草稿详细信息失败: %v", err)
	}

	// 按修改时间排序
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].ModTime.After(infos[j].ModTime)
	})

	for i, info := range infos {
		fmt.Printf("   [%d] %s\n", i+1, info.String())
		fmt.Printf("       路径: %s\n", info.Path)
		fmt.Printf("       年龄: %v\n", info.Age().Truncate(time.Second))

		if info.IsValid() {
			// 尝试加载并显示项目信息
			if scriptFile, err := df.LoadTemplate(info.Name); err == nil {
				fmt.Printf("       项目: %dx%d@%dfps, %.1f秒\n",
					scriptFile.Width, scriptFile.Height, scriptFile.FPS, float64(scriptFile.Duration)/1e6)
			}
		}
		fmt.Println()
	}

	// 统计信息
	validCount := 0
	for _, info := range infos {
		if info.IsValid() {
			validCount++
		}
	}

	fmt.Printf("📈 统计信息:\n")
	fmt.Printf("   - 总草稿数: %d\n", len(infos))
	fmt.Printf("   - 有效草稿: %d\n", validCount)
	fmt.Printf("   - 无效草稿: %d\n", len(infos)-validCount)
}

// demonstrateDraftDuplication 演示草稿复制功能
func demonstrateDraftDuplication() {
	fmt.Println("📋 === 草稿复制演示 ===")

	// 创建临时草稿文件夹
	tempDir, err := os.MkdirTemp("", "draft_demo")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 创建DraftFolder管理器
	df, err := draft.NewDraftFolder(tempDir)
	if err != nil {
		log.Fatalf("创建DraftFolder失败: %v", err)
	}

	// 创建原始模板草稿
	templateName := "原始模板"
	templatePath := filepath.Join(tempDir, templateName)
	if err := os.Mkdir(templatePath, 0755); err != nil {
		log.Fatalf("创建模板文件夹失败: %v", err)
	}

	// 创建详细的模板内容
	createDetailedDraftInfo(templatePath, templateName, 30, 1920, 1080, 20000000)

	// 创建一些额外文件来测试完整复制
	extraFiles := map[string]string{
		"README.txt": "这是项目说明文件",
		"notes.md":   "# 项目笔记\n\n这是一个示例项目",
		"config.ini": "[settings]\nquality=high\nformat=mp4",
	}

	for fileName, content := range extraFiles {
		filePath := filepath.Join(templatePath, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			log.Fatalf("创建额外文件失败: %v", err)
		}
	}

	// 创建子文件夹
	subDir := filepath.Join(templatePath, "assets")
	if err := os.Mkdir(subDir, 0755); err != nil {
		log.Fatalf("创建子文件夹失败: %v", err)
	}

	assetFile := filepath.Join(subDir, "logo.txt")
	if err := os.WriteFile(assetFile, []byte("LOGO内容"), 0644); err != nil {
		log.Fatalf("创建资源文件失败: %v", err)
	}

	fmt.Printf("📁 创建原始模板: %s\n", templateName)
	fmt.Printf("   - 包含draft_info.json\n")
	fmt.Printf("   - 包含%d个额外文件\n", len(extraFiles))
	fmt.Printf("   - 包含assets子文件夹\n")

	// 测试基本复制
	fmt.Printf("\n📋 基本复制测试:\n")
	newDraftName := "复制的项目"
	scriptFile, err := df.DuplicateAsTemplate(templateName, newDraftName, false)
	if err != nil {
		log.Fatalf("复制草稿失败: %v", err)
	}

	fmt.Printf("   ✅ 成功复制草稿: %s -> %s\n", templateName, newDraftName)
	fmt.Printf("   - 项目规格: %dx%d@%dfps\n", scriptFile.Width, scriptFile.Height, scriptFile.FPS)
	fmt.Printf("   - 项目时长: %.1f秒\n", float64(scriptFile.Duration)/1e6)

	// 验证文件完整性
	fmt.Printf("\n🔍 验证复制完整性:\n")
	newDraftPath := df.GetDraftPath(newDraftName)

	// 检查额外文件
	for fileName := range extraFiles {
		filePath := filepath.Join(newDraftPath, fileName)
		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("   ✅ %s 已复制\n", fileName)
		} else {
			fmt.Printf("   ❌ %s 复制失败\n", fileName)
		}
	}

	// 检查子文件夹
	subDirPath := filepath.Join(newDraftPath, "assets")
	if _, err := os.Stat(subDirPath); err == nil {
		fmt.Printf("   ✅ assets子文件夹已复制\n")

		assetFilePath := filepath.Join(subDirPath, "logo.txt")
		if _, err := os.Stat(assetFilePath); err == nil {
			fmt.Printf("   ✅ assets/logo.txt已复制\n")
		}
	}

	// 测试重复复制（不允许覆盖）
	fmt.Printf("\n🚫 重复复制测试（不允许覆盖）:\n")
	_, err = df.DuplicateAsTemplate(templateName, newDraftName, false)
	if err != nil {
		fmt.Printf("   ✅ 正确拒绝重复复制: %v\n", err)
	}

	// 测试重复复制（允许覆盖）
	fmt.Printf("\n🔄 重复复制测试（允许覆盖）:\n")
	_, err = df.DuplicateAsTemplate(templateName, newDraftName, true)
	if err != nil {
		log.Fatalf("允许覆盖的复制失败: %v", err)
	}
	fmt.Printf("   ✅ 成功覆盖复制\n")

	// 批量复制测试
	fmt.Printf("\n📚 批量复制测试:\n")
	copyNames := []string{"项目副本1", "项目副本2", "项目副本3"}
	for i, copyName := range copyNames {
		_, err := df.DuplicateAsTemplate(templateName, copyName, false)
		if err != nil {
			log.Fatalf("批量复制失败: %v", err)
		}
		fmt.Printf("   [%d] ✅ %s\n", i+1, copyName)
	}

	// 显示最终状态
	fmt.Printf("\n📊 复制完成统计:\n")
	finalDrafts, _ := df.ListDrafts()
	fmt.Printf("   - 总草稿数: %d\n", len(finalDrafts))
	fmt.Printf("   - 原始模板: 1个\n")
	fmt.Printf("   - 复制草稿: %d个\n", len(finalDrafts)-1)
}

// demonstrateDraftDeletion 演示草稿删除功能
func demonstrateDraftDeletion() {
	fmt.Println("🗑️ === 草稿删除演示 ===")

	// 创建临时草稿文件夹
	tempDir, err := os.MkdirTemp("", "draft_demo")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 创建DraftFolder管理器
	df, err := draft.NewDraftFolder(tempDir)
	if err != nil {
		log.Fatalf("创建DraftFolder失败: %v", err)
	}

	// 创建多个测试草稿
	draftNames := []string{"待删除草稿1", "待删除草稿2", "保留草稿", "待删除草稿3"}
	fmt.Printf("📁 创建测试草稿:\n")

	for i, name := range draftNames {
		draftPath := filepath.Join(tempDir, name)
		if err := os.Mkdir(draftPath, 0755); err != nil {
			log.Fatalf("创建草稿文件夹失败: %v", err)
		}

		// 创建一些内容
		createSampleDraftInfo(draftPath, name)

		// 添加一些文件
		testFile := filepath.Join(draftPath, fmt.Sprintf("file_%d.txt", i))
		content := fmt.Sprintf("这是%s的测试文件", name)
		if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
			log.Fatalf("创建测试文件失败: %v", err)
		}

		fmt.Printf("   [%d] %s ✅\n", i+1, name)
	}

	// 显示删除前状态
	fmt.Printf("\n📊 删除前状态:\n")
	beforeDrafts, _ := df.ListDrafts()
	fmt.Printf("   - 总草稿数: %d\n", len(beforeDrafts))
	for i, name := range beforeDrafts {
		fmt.Printf("     [%d] %s\n", i+1, name)
	}

	// 执行删除操作
	fmt.Printf("\n🗑️ 执行删除操作:\n")
	toDelete := []string{"待删除草稿1", "待删除草稿3"}

	for _, name := range toDelete {
		fmt.Printf("   删除: %s ... ", name)
		if err := df.Remove(name); err != nil {
			fmt.Printf("❌ 失败: %v\n", err)
		} else {
			fmt.Printf("✅ 成功\n")
		}
	}

	// 测试删除不存在的草稿
	fmt.Printf("   删除不存在的草稿: 不存在草稿 ... ")
	if err := df.Remove("不存在草稿"); err != nil {
		fmt.Printf("✅ 正确处理: %v\n", err)
	}

	// 显示删除后状态
	fmt.Printf("\n📊 删除后状态:\n")
	afterDrafts, _ := df.ListDrafts()
	fmt.Printf("   - 总草稿数: %d\n", len(afterDrafts))
	for i, name := range afterDrafts {
		fmt.Printf("     [%d] %s\n", i+1, name)
	}

	// 验证文件确实被删除
	fmt.Printf("\n🔍 验证删除结果:\n")
	for _, name := range toDelete {
		exists := df.DraftExists(name)
		fmt.Printf("   - %s 存在: %v\n", name, exists)
	}

	// 验证保留的草稿仍然完整
	fmt.Printf("\n✅ 验证保留草稿完整性:\n")
	remainingDraft := "保留草稿"
	if df.DraftExists(remainingDraft) {
		info, err := df.GetDraftInfo(remainingDraft)
		if err == nil {
			fmt.Printf("   - %s: %s\n", remainingDraft, info.String())
			fmt.Printf("   - 有效性: %v\n", info.IsValid())
		}
	}

	fmt.Printf("\n📈 删除操作总结:\n")
	fmt.Printf("   - 删除前: %d个草稿\n", len(beforeDrafts))
	fmt.Printf("   - 删除后: %d个草稿\n", len(afterDrafts))
	fmt.Printf("   - 成功删除: %d个草稿\n", len(beforeDrafts)-len(afterDrafts))
}

// demonstrateMaterialInspection 演示素材检查功能
func demonstrateMaterialInspection() {
	fmt.Println("🔍 === 素材检查演示 ===")

	// 创建临时草稿文件夹
	tempDir, err := os.MkdirTemp("", "draft_demo")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 创建DraftFolder管理器
	df, err := draft.NewDraftFolder(tempDir)
	if err != nil {
		log.Fatalf("创建DraftFolder失败: %v", err)
	}

	// 创建包含丰富素材的草稿
	draftName := "素材丰富项目"
	draftPath := filepath.Join(tempDir, draftName)
	if err := os.Mkdir(draftPath, 0755); err != nil {
		log.Fatalf("创建草稿文件夹失败: %v", err)
	}

	// 创建包含各种素材的draft_info.json
	draftInfo := map[string]interface{}{
		"fps":      float64(30),
		"duration": float64(15000000),
		"canvas_config": map[string]interface{}{
			"width":  float64(1920),
			"height": float64(1080),
		},
		"materials": map[string]interface{}{
			"stickers": []interface{}{
				map[string]interface{}{
					"resource_id": "sticker_001",
					"name":        "可爱表情包",
				},
				map[string]interface{}{
					"resource_id": "sticker_002",
					"name":        "动态贴纸",
				},
				map[string]interface{}{
					"resource_id": "sticker_003",
					"name":        "节日贴纸",
				},
			},
			"effects": []interface{}{
				map[string]interface{}{
					"type":        "text_shape",
					"effect_id":   "effect_001",
					"resource_id": "bubble_001",
					"name":        "圆形气泡",
				},
				map[string]interface{}{
					"type":        "text_shape",
					"effect_id":   "effect_002",
					"resource_id": "bubble_002",
					"name":        "方形气泡",
				},
				map[string]interface{}{
					"type":        "text_effect",
					"resource_id": "flower_001",
					"name":        "花朵文字效果",
				},
				map[string]interface{}{
					"type":        "text_effect",
					"resource_id": "flower_002",
					"name":        "星光文字效果",
				},
			},
		},
		"tracks": []interface{}{},
	}

	draftInfoPath := filepath.Join(draftPath, "draft_info.json")
	jsonBytes, err := json.Marshal(draftInfo)
	if err != nil {
		log.Fatalf("序列化draft_info失败: %v", err)
	}

	if err := os.WriteFile(draftInfoPath, jsonBytes, 0644); err != nil {
		log.Fatalf("写入draft_info.json失败: %v", err)
	}

	fmt.Printf("📁 创建素材丰富项目: %s\n", draftName)
	fmt.Printf("   - 贴纸素材: 3个\n")
	fmt.Printf("   - 文字气泡: 2个\n")
	fmt.Printf("   - 花字效果: 2个\n")

	// 执行素材检查
	fmt.Printf("\n🔍 执行素材检查:\n")
	fmt.Println(strings.Repeat("=", 50))
	if err := df.InspectMaterial(draftName); err != nil {
		log.Fatalf("素材检查失败: %v", err)
	}
	fmt.Println(strings.Repeat("=", 50))

	// 创建另一个空素材的草稿进行对比
	emptyDraftName := "空素材项目"
	emptyDraftPath := filepath.Join(tempDir, emptyDraftName)
	if err := os.Mkdir(emptyDraftPath, 0755); err != nil {
		log.Fatalf("创建空草稿文件夹失败: %v", err)
	}

	emptyDraftInfo := map[string]interface{}{
		"fps":      float64(30),
		"duration": float64(5000000),
		"canvas_config": map[string]interface{}{
			"width":  float64(1920),
			"height": float64(1080),
		},
		"materials": map[string]interface{}{
			"stickers": []interface{}{},
			"effects":  []interface{}{},
		},
		"tracks": []interface{}{},
	}

	emptyDraftInfoPath := filepath.Join(emptyDraftPath, "draft_info.json")
	emptyJsonBytes, err := json.Marshal(emptyDraftInfo)
	if err != nil {
		log.Fatalf("序列化空draft_info失败: %v", err)
	}

	if err := os.WriteFile(emptyDraftInfoPath, emptyJsonBytes, 0644); err != nil {
		log.Fatalf("写入空draft_info.json失败: %v", err)
	}

	fmt.Printf("\n📁 创建空素材项目: %s\n", emptyDraftName)
	fmt.Printf("🔍 执行空素材检查:\n")
	fmt.Println(strings.Repeat("-", 30))
	if err := df.InspectMaterial(emptyDraftName); err != nil {
		log.Fatalf("空素材检查失败: %v", err)
	}
	fmt.Println(strings.Repeat("-", 30))

	// 测试检查不存在的草稿
	fmt.Printf("\n🚫 测试检查不存在的草稿:\n")
	if err := df.InspectMaterial("不存在的草稿"); err != nil {
		fmt.Printf("   ✅ 正确处理不存在的草稿: %v\n", err)
	}

	fmt.Printf("\n📊 素材检查总结:\n")
	fmt.Printf("   - 检查了2个有效草稿\n")
	fmt.Printf("   - 发现丰富素材项目包含多种素材类型\n")
	fmt.Printf("   - 空素材项目无特殊素材\n")
	fmt.Printf("   - 正确处理错误情况\n")
}

// demonstrateCompleteWorkflow 演示完整的草稿管理工作流
func demonstrateCompleteWorkflow() {
	fmt.Println("🎯 === 完整草稿管理工作流演示 ===")

	// 创建临时草稿文件夹
	tempDir, err := os.MkdirTemp("", "draft_demo")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fmt.Printf("🎬 草稿管理工作流 - 模拟真实使用场景:\n")

	// 步骤1: 初始化草稿文件夹管理器
	fmt.Printf("   📂 步骤1: 初始化草稿文件夹管理器\n")
	df, err := draft.NewDraftFolder(tempDir)
	if err != nil {
		log.Fatalf("创建DraftFolder失败: %v", err)
	}
	fmt.Printf("     ✅ 管理器创建成功: %s\n", df.FolderPath)

	// 步骤2: 创建项目模板
	fmt.Printf("   🎨 步骤2: 创建项目模板\n")
	templateName := "标准视频模板"
	templatePath := filepath.Join(tempDir, templateName)
	if err := os.Mkdir(templatePath, 0755); err != nil {
		log.Fatalf("创建模板失败: %v", err)
	}

	createDetailedDraftInfo(templatePath, templateName, 30, 1920, 1080, 10000000)

	// 添加模板资源
	templateResources := map[string]string{
		"template_guide.md": "# 模板使用指南\n\n这是标准视频模板，适用于大多数项目。",
		"settings.json":     `{"quality": "high", "format": "mp4", "bitrate": 8000}`,
	}

	for fileName, content := range templateResources {
		filePath := filepath.Join(templatePath, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			log.Fatalf("创建模板资源失败: %v", err)
		}
	}

	fmt.Printf("     ✅ 模板创建完成: %s (1920x1080@30fps, 10秒)\n", templateName)

	// 步骤3: 基于模板创建多个项目
	fmt.Printf("   📋 步骤3: 基于模板创建项目\n")
	projects := []struct {
		name        string
		description string
	}{
		{"客户A宣传视频", "为客户A制作的品牌宣传视频"},
		{"产品演示视频", "新产品功能演示视频"},
		{"培训教程视频", "内部培训使用的教程视频"},
	}

	for i, project := range projects {
		scriptFile, err := df.DuplicateAsTemplate(templateName, project.name, false)
		if err != nil {
			log.Fatalf("创建项目失败: %v", err)
		}

		// 为每个项目定制一些属性
		scriptFile.Duration = int64((i + 2) * 5000000) // 10秒, 15秒, 20秒

		fmt.Printf("     [%d] ✅ %s: %dx%d@%dfps, %.1f秒\n",
			i+1, project.name, scriptFile.Width, scriptFile.Height, scriptFile.FPS,
			float64(scriptFile.Duration)/1e6)
	}

	// 步骤4: 查看所有项目
	fmt.Printf("   📊 步骤4: 查看项目列表\n")
	infos, err := df.ListDraftsWithInfo()
	if err != nil {
		log.Fatalf("获取项目列表失败: %v", err)
	}

	// 按类型分类显示
	templates := make([]*draft.DraftInfo, 0)
	activeProjects := make([]*draft.DraftInfo, 0)

	for _, info := range infos {
		if info.Name == templateName {
			templates = append(templates, info)
		} else {
			activeProjects = append(activeProjects, info)
		}
	}

	fmt.Printf("     📋 模板 (%d个):\n", len(templates))
	for _, info := range templates {
		fmt.Printf("       - %s\n", info.Name)
	}

	fmt.Printf("     🎬 活跃项目 (%d个):\n", len(activeProjects))
	for i, info := range activeProjects {
		fmt.Printf("       [%d] %s (修改于 %s)\n",
			i+1, info.Name, info.ModTime.Format("15:04:05"))
	}

	// 步骤5: 项目管理操作
	fmt.Printf("   🔧 步骤5: 项目管理操作\n")

	// 5a: 检查项目素材
	fmt.Printf("     🔍 检查项目素材:\n")
	for i, project := range projects[:2] { // 只检查前两个
		fmt.Printf("       [%d] %s的素材:\n", i+1, project.name)
		if err := df.InspectMaterial(project.name); err != nil {
			fmt.Printf("         检查失败: %v\n", err)
		}
	}

	// 5b: 创建项目备份
	fmt.Printf("     💾 创建项目备份:\n")
	backupName := projects[0].name + "_备份"
	_, err = df.DuplicateAsTemplate(projects[0].name, backupName, false)
	if err != nil {
		log.Fatalf("创建备份失败: %v", err)
	}
	fmt.Printf("       ✅ 备份创建: %s -> %s\n", projects[0].name, backupName)

	// 5c: 删除测试项目
	fmt.Printf("     🗑️ 清理测试项目:\n")
	testProject := projects[2].name // 删除培训教程视频
	if err := df.Remove(testProject); err != nil {
		log.Fatalf("删除项目失败: %v", err)
	}
	fmt.Printf("       ✅ 已删除: %s\n", testProject)

	// 步骤6: 最终状态统计
	fmt.Printf("   📈 步骤6: 最终状态统计\n")
	finalDrafts, _ := df.ListDrafts()
	finalInfos, _ := df.ListDraftsWithInfo()

	validProjects := 0
	oldestProject := time.Now()
	newestProject := time.Time{}

	for _, info := range finalInfos {
		if info.IsValid() {
			validProjects++
		}
		if info.ModTime.Before(oldestProject) {
			oldestProject = info.ModTime
		}
		if info.ModTime.After(newestProject) {
			newestProject = info.ModTime
		}
	}

	fmt.Printf("     📊 工作流完成统计:\n")
	fmt.Printf("       - 总项目数: %d\n", len(finalDrafts))
	fmt.Printf("       - 有效项目: %d\n", validProjects)
	fmt.Printf("       - 模板数: 1\n")
	fmt.Printf("       - 活跃项目: %d\n", len(finalDrafts)-1)
	fmt.Printf("       - 项目时间跨度: %v\n", newestProject.Sub(oldestProject).Truncate(time.Second))

	fmt.Printf("\n🎉 完整工作流演示完成!\n")
	fmt.Printf("   - 成功演示了从模板创建到项目管理的完整流程\n")
	fmt.Printf("   - 展示了草稿复制、备份、删除等核心功能\n")
	fmt.Printf("   - 验证了素材检查和项目信息管理功能\n")
	fmt.Printf("   - 模拟了真实的视频项目管理场景\n")
}

// 辅助函数

func createSampleDraftInfo(draftPath, name string) {
	draftInfo := map[string]interface{}{
		"fps":      float64(30),
		"duration": float64(5000000),
		"canvas_config": map[string]interface{}{
			"width":  float64(1920),
			"height": float64(1080),
		},
		"materials": map[string]interface{}{
			"videos": []interface{}{},
			"audios": []interface{}{},
		},
		"tracks": []interface{}{},
	}

	draftInfoPath := filepath.Join(draftPath, "draft_info.json")
	jsonBytes, _ := json.Marshal(draftInfo)
	os.WriteFile(draftInfoPath, jsonBytes, 0644)
}

func createDetailedDraftInfo(draftPath, name string, fps, width, height int, duration int64) {
	draftInfo := map[string]interface{}{
		"fps":      float64(fps),
		"duration": float64(duration),
		"canvas_config": map[string]interface{}{
			"width":  float64(width),
			"height": float64(height),
		},
		"materials": map[string]interface{}{
			"videos": []interface{}{},
			"audios": []interface{}{},
		},
		"tracks": []interface{}{},
	}

	draftInfoPath := filepath.Join(draftPath, "draft_info.json")
	jsonBytes, _ := json.Marshal(draftInfo)
	os.WriteFile(draftInfoPath, jsonBytes, 0644)
}

func getStatusText(hasInfo bool) string {
	if hasInfo {
		return "有效草稿"
	}
	return "空文件夹"
}
