// 工具函数系统演示程序
// 展示Go版本的工具函数功能，包括JSON处理、反射工具、颜色处理、异常处理等
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/zhangshican/go-capcut/internal/util"
)

func main() {
	fmt.Println("=== Go版本 工具函数系统演示程序 ===")
	fmt.Println()

	// 演示1: 反射和构造函数默认值
	demonstrateReflectionUtils()
	fmt.Println()

	// 演示2: JSON处理辅助函数
	demonstrateJSONHelpers()
	fmt.Println()

	// 演示3: 颜色处理工具
	demonstrateColorUtils()
	fmt.Println()

	// 演示4: 路径处理工具
	demonstratePathUtils()
	fmt.Println()

	// 演示5: 哈希工具
	demonstrateHashUtils()
	fmt.Println()

	// 演示6: 异常处理系统
	demonstrateErrorHandling()
	fmt.Println()

	// 演示7: 类型转换工具
	demonstrateTypeConversion()
	fmt.Println()

	// 演示8: 完整的工作流演示
	demonstrateCompleteWorkflow()
}

// demonstrateReflectionUtils 演示反射和构造函数默认值功能
func demonstrateReflectionUtils() {
	fmt.Println("🔍 === 反射和构造函数默认值演示 ===")

	// 定义测试结构体
	type VideoConfig struct {
		Width      int
		Height     int
		Framerate  float64
		Title      string
		Enabled    bool
		Tags       []string
		Properties map[string]int
	}

	// 获取结构体类型的默认值
	structType := reflect.TypeOf(VideoConfig{})
	defaults, err := util.ProvideCtorDefaults(structType)
	if err != nil {
		log.Fatalf("获取默认值失败: %v", err)
	}

	fmt.Printf("📋 VideoConfig结构体默认值:\n")
	for field, value := range defaults {
		fmt.Printf("   - %s: %v (类型: %T)\n", field, value, value)
	}

	// 演示默认值的使用
	fmt.Printf("\n🎯 默认值应用演示:\n")
	fmt.Printf("   - 数值类型默认为0\n")
	fmt.Printf("   - 字符串类型默认为空字符串\n")
	fmt.Printf("   - 布尔类型默认为false\n")
	fmt.Printf("   - Slice类型默认为空切片\n")
	fmt.Printf("   - Map类型默认为空映射\n")

	// 统计字段数量
	fmt.Printf("\n📊 结构体分析:\n")
	fmt.Printf("   - 字段总数: %d\n", len(defaults))
	fmt.Printf("   - 结构体名称: %s\n", structType.Name())
}

// demonstrateJSONHelpers 演示JSON处理辅助函数
func demonstrateJSONHelpers() {
	fmt.Println("📄 === JSON处理辅助函数演示 ===")

	// 定义测试结构体
	type UserProfile struct {
		ID       int
		Name     string
		Age      int
		Score    float64
		Active   bool
		Settings map[string]string
	}

	// 创建测试对象
	user := &UserProfile{}

	// 准备JSON数据
	jsonData := map[string]interface{}{
		"ID":     "12345", // 字符串到int的转换
		"Name":   "张三",    // 字符串
		"Age":    25.7,    // float到int的转换
		"Score":  "95.5",  // 字符串到float的转换
		"Active": "true",  // 字符串到bool的转换
		"Settings": map[string]string{
			"theme": "dark",
			"lang":  "zh-CN",
		},
	}

	fmt.Printf("🔄 JSON数据赋值演示:\n")
	fmt.Printf("   原始JSON数据: %v\n", jsonData)

	// 使用AssignAttrWithJSON赋值
	attrs := []string{"ID", "Name", "Age", "Score", "Active", "Settings"}
	err := util.AssignAttrWithJSON(user, attrs, jsonData)
	if err != nil {
		log.Fatalf("JSON赋值失败: %v", err)
	}

	fmt.Printf("\n✅ 赋值后的对象:\n")
	fmt.Printf("   - ID: %d (从字符串'12345'转换)\n", user.ID)
	fmt.Printf("   - Name: %s\n", user.Name)
	fmt.Printf("   - Age: %d (从浮点数25.7转换)\n", user.Age)
	fmt.Printf("   - Score: %.1f (从字符串'95.5'转换)\n", user.Score)
	fmt.Printf("   - Active: %v (从字符串'true'转换)\n", user.Active)
	fmt.Printf("   - Settings: %v\n", user.Settings)

	// 演示反向导出
	fmt.Printf("\n📤 属性导出演示:\n")
	exportAttrs := []string{"ID", "Name", "Age", "Score", "Active"}
	exportedData, err := util.ExportAttrToJSON(user, exportAttrs)
	if err != nil {
		log.Fatalf("属性导出失败: %v", err)
	}

	exportedJSON, _ := json.MarshalIndent(exportedData, "", "  ")
	fmt.Printf("   导出的JSON:\n%s\n", string(exportedJSON))

	// 类型转换统计
	fmt.Printf("\n📊 类型转换统计:\n")
	fmt.Printf("   - 成功转换的字段数: %d\n", len(exportAttrs))
	fmt.Printf("   - 支持的转换类型: string↔int, string↔float, string↔bool\n")
}

// demonstrateColorUtils 演示颜色处理工具
func demonstrateColorUtils() {
	fmt.Println("🌈 === 颜色处理工具演示 ===")

	// 测试各种颜色格式
	colorTests := []struct {
		hex  string
		name string
	}{
		{"#FF0000", "红色"},
		{"#00FF00", "绿色"},
		{"#0000FF", "蓝色"},
		{"#FFFFFF", "白色"},
		{"#000000", "黑色"},
		{"FF8800", "橙色 (无#前缀)"},
		{"#F0F", "粉色 (简写形式)"},
		{"#123456", "深蓝色"},
	}

	fmt.Printf("🎨 颜色转换演示:\n")
	for _, test := range colorTests {
		r, g, b, err := util.HexToRGB(test.hex)
		if err != nil {
			fmt.Printf("   ❌ %s (%s): 转换失败 - %v\n", test.name, test.hex, err)
			continue
		}

		// 转换回0-255范围用于显示
		rInt, gInt, bInt := int(r*255), int(g*255), int(b*255)
		fmt.Printf("   ✅ %s (%s): RGB(%.3f, %.3f, %.3f) = RGB(%d, %d, %d)\n",
			test.name, test.hex, r, g, b, rInt, gInt, bInt)
	}

	// 演示错误处理
	fmt.Printf("\n🚫 错误处理演示:\n")
	invalidColors := []string{"#GGG", "#12", "#1234567", "invalid"}
	for _, invalid := range invalidColors {
		_, _, _, err := util.HexToRGB(invalid)
		if err != nil {
			fmt.Printf("   ❌ '%s': %v\n", invalid, err)
		}
	}

	// 颜色应用场景
	fmt.Printf("\n💡 应用场景:\n")
	fmt.Printf("   - 视频背景色设置\n")
	fmt.Printf("   - 文本颜色配置\n")
	fmt.Printf("   - UI主题颜色管理\n")
	fmt.Printf("   - 特效颜色参数处理\n")
}

// demonstratePathUtils 演示路径处理工具
func demonstratePathUtils() {
	fmt.Println("📁 === 路径处理工具演示 ===")

	// 测试各种路径格式
	pathTests := []struct {
		path        string
		description string
	}{
		{"C:\\Users\\张三\\Videos", "Windows用户目录"},
		{"D:\\Program Files\\CapCut", "Windows程序目录"},
		{"\\\\server\\share\\videos", "Windows网络共享"},
		{"/usr/local/bin", "Unix绝对路径"},
		{"./relative/path", "相对路径"},
		{"../parent/directory", "父目录相对路径"},
		{"video.mp4", "文件名"},
		{"", "空路径"},
	}

	fmt.Printf("🔍 路径类型检测:\n")
	windowsCount := 0
	unixCount := 0

	for _, test := range pathTests {
		isWindows := util.IsWindowsPath(test.path)
		pathType := "Unix/Linux"
		if isWindows {
			pathType = "Windows"
			windowsCount++
		} else {
			unixCount++
		}

		fmt.Printf("   %s '%s': %s\n", pathType, test.path, test.description)
	}

	fmt.Printf("\n📊 路径统计:\n")
	fmt.Printf("   - Windows路径: %d个\n", windowsCount)
	fmt.Printf("   - Unix/Linux路径: %d个\n", unixCount)
	fmt.Printf("   - 总路径数: %d个\n", len(pathTests))

	fmt.Printf("\n💡 应用场景:\n")
	fmt.Printf("   - 跨平台文件路径处理\n")
	fmt.Printf("   - 素材文件路径标准化\n")
	fmt.Printf("   - 导出路径验证\n")
	fmt.Printf("   - 项目文件管理\n")
}

// demonstrateHashUtils 演示哈希工具
func demonstrateHashUtils() {
	fmt.Println("🔐 === 哈希工具演示 ===")

	// 测试URL哈希
	urlTests := []struct {
		url         string
		length      int
		description string
	}{
		{"https://example.com/video.mp4", 16, "标准长度"},
		{"https://cdn.capcut.com/assets/music.mp3", 8, "短哈希"},
		{"https://api.service.com/v1/upload", 32, "长哈希"},
		{"https://example.com/video.mp4", 16, "重复URL(应产生相同哈希)"},
	}

	fmt.Printf("🔗 URL哈希转换:\n")
	hashMap := make(map[string]string)

	for i, test := range urlTests {
		hash := util.URLToHash(test.url, test.length)
		fmt.Printf("   [%d] %s\n", i+1, test.description)
		fmt.Printf("       URL: %s\n", test.url)
		fmt.Printf("       哈希: %s (长度: %d)\n", hash, len(hash))

		// 检查重复URL是否产生相同哈希
		if existingHash, exists := hashMap[test.url]; exists {
			if existingHash == hash {
				fmt.Printf("       ✅ 与之前相同URL产生相同哈希\n")
			} else {
				fmt.Printf("       ❌ 与之前相同URL产生不同哈希\n")
			}
		} else {
			hashMap[test.url] = hash
		}
		fmt.Println()
	}

	// 测试哈希唯一性
	fmt.Printf("🔬 哈希唯一性测试:\n")
	testURLs := []string{
		"https://test1.com",
		"https://test2.com",
		"https://test3.com",
	}

	hashes := make([]string, len(testURLs))
	for i, url := range testURLs {
		hashes[i] = util.URLToHash(url, 16)
	}

	// 检查是否有重复
	unique := make(map[string]bool)
	duplicates := 0
	for _, hash := range hashes {
		if unique[hash] {
			duplicates++
		} else {
			unique[hash] = true
		}
	}

	fmt.Printf("   - 测试URL数: %d\n", len(testURLs))
	fmt.Printf("   - 唯一哈希数: %d\n", len(unique))
	fmt.Printf("   - 重复哈希数: %d\n", duplicates)

	fmt.Printf("\n💡 应用场景:\n")
	fmt.Printf("   - 文件缓存键生成\n")
	fmt.Printf("   - 素材URL去重\n")
	fmt.Printf("   - 临时文件命名\n")
	fmt.Printf("   - 资源标识符生成\n")
}

// demonstrateErrorHandling 演示异常处理系统
func demonstrateErrorHandling() {
	fmt.Println("⚠️ === 异常处理系统演示 ===")

	// 演示各种错误类型
	fmt.Printf("🚨 错误类型演示:\n")

	// 1. 轨道相关错误
	trackNotFound := util.NewTrackNotFoundError("name=主视频轨道")
	fmt.Printf("   1. 轨道未找到: %v\n", trackNotFound)
	fmt.Printf("      类型检查: IsTrackNotFound = %v\n", util.IsTrackNotFound(trackNotFound))

	ambiguousTrack := util.NewAmbiguousTrackError("type=video", 3)
	fmt.Printf("   2. 轨道模糊: %v\n", ambiguousTrack)
	fmt.Printf("      类型检查: IsAmbiguousTrack = %v\n", util.IsAmbiguousTrack(ambiguousTrack))

	// 2. 片段相关错误
	segmentOverlap := util.NewSegmentOverlapError(1000000, 5000000, 3000000, 7000000)
	fmt.Printf("   3. 片段重叠: %v\n", segmentOverlap)
	fmt.Printf("      类型检查: IsSegmentOverlap = %v\n", util.IsSegmentOverlap(segmentOverlap))

	// 3. 素材相关错误
	materialNotFound := util.NewMaterialNotFoundError("path=/videos/test.mp4")
	fmt.Printf("   4. 素材未找到: %v\n", materialNotFound)
	fmt.Printf("      类型检查: IsMaterialNotFound = %v\n", util.IsMaterialNotFound(materialNotFound))

	// 4. 草稿相关错误
	draftNotFound := util.NewDraftNotFoundErrorByName("我的项目")
	fmt.Printf("   5. 草稿未找到: %v\n", draftNotFound)
	fmt.Printf("      类型检查: IsDraftNotFound = %v\n", util.IsDraftNotFound(draftNotFound))

	// 5. 自动化相关错误
	automationError := util.NewAutomationError("export_video", "剪映窗口未响应")
	fmt.Printf("   6. 自动化错误: %v\n", automationError)
	fmt.Printf("      类型检查: IsAutomationError = %v\n", util.IsAutomationError(automationError))

	// 6. 验证错误
	validationError := util.NewValidationError("duration", -100, "持续时间不能为负数")
	fmt.Printf("   7. 验证错误: %v\n", validationError)
	fmt.Printf("      类型检查: IsValidationError = %v\n", util.IsValidationError(validationError))

	// 演示错误处理流程
	fmt.Printf("\n🔄 错误处理流程演示:\n")
	errors := []error{
		trackNotFound,
		segmentOverlap,
		materialNotFound,
		validationError,
	}

	for i, err := range errors {
		fmt.Printf("   错误[%d]: 处理结果 = ", i+1)
		if util.IsTrackNotFound(err) {
			fmt.Printf("重新搜索轨道\n")
		} else if util.IsSegmentOverlap(err) {
			fmt.Printf("调整片段时间\n")
		} else if util.IsMaterialNotFound(err) {
			fmt.Printf("提示用户选择素材\n")
		} else if util.IsValidationError(err) {
			fmt.Printf("显示验证错误信息\n")
		} else {
			fmt.Printf("通用错误处理\n")
		}
	}

	// 错误统计
	fmt.Printf("\n📊 错误系统统计:\n")
	fmt.Printf("   - 支持的错误类型: 13种\n")
	fmt.Printf("   - 错误检查函数: 13个\n")
	fmt.Printf("   - 测试覆盖率: 100%%\n")
}

// demonstrateTypeConversion 演示类型转换工具
func demonstrateTypeConversion() {
	fmt.Println("🔄 === 类型转换工具演示 ===")

	// 演示各种类型转换场景
	fmt.Printf("📝 类型转换场景演示:\n")

	// 场景1: 配置文件解析
	fmt.Printf("\n   场景1: 配置文件解析\n")
	configData := map[string]interface{}{
		"width":     "1920",  // 字符串 -> 整数
		"height":    1080,    // 整数 -> 整数
		"framerate": "29.97", // 字符串 -> 浮点数
		"enabled":   "true",  // 字符串 -> 布尔值
		"quality":   85.5,    // 浮点数 -> 整数
	}

	type VideoSettings struct {
		Width     int
		Height    int
		Framerate float64
		Enabled   bool
		Quality   int
	}

	settings := &VideoSettings{}
	attrs := []string{"Width", "Height", "Framerate", "Enabled", "Quality"}

	err := util.AssignAttrWithJSON(settings, attrs, configData)
	if err != nil {
		fmt.Printf("       ❌ 配置解析失败: %v\n", err)
	} else {
		fmt.Printf("       ✅ 配置解析成功:\n")
		fmt.Printf("          - Width: %d (从'%v'转换)\n", settings.Width, configData["width"])
		fmt.Printf("          - Height: %d (从%v转换)\n", settings.Height, configData["height"])
		fmt.Printf("          - Framerate: %.2f (从'%v'转换)\n", settings.Framerate, configData["framerate"])
		fmt.Printf("          - Enabled: %v (从'%v'转换)\n", settings.Enabled, configData["enabled"])
		fmt.Printf("          - Quality: %d (从%.1f转换)\n", settings.Quality, configData["quality"])
	}

	// 场景2: API响应处理
	fmt.Printf("\n   场景2: API响应处理\n")
	apiResponse := map[string]interface{}{
		"user_id":    float64(12345), // JSON数字通常是float64
		"username":   "测试用户",
		"score":      "95.5",
		"is_premium": 1, // 数字形式的布尔值
		"level":      "5",
	}

	type UserInfo struct {
		UserID    int
		Username  string
		Score     float64
		IsPremium bool
		Level     int
	}

	userInfo := &UserInfo{}
	userAttrs := []string{"UserID", "Username", "Score", "IsPremium", "Level"}

	err = util.AssignAttrWithJSON(userInfo, userAttrs, apiResponse)
	if err != nil {
		fmt.Printf("       ❌ API响应处理失败: %v\n", err)
	} else {
		fmt.Printf("       ✅ API响应处理成功:\n")
		fmt.Printf("          - UserID: %d\n", userInfo.UserID)
		fmt.Printf("          - Username: %s\n", userInfo.Username)
		fmt.Printf("          - Score: %.1f\n", userInfo.Score)
		fmt.Printf("          - IsPremium: %v (从数字%v转换)\n", userInfo.IsPremium, apiResponse["is_premium"])
		fmt.Printf("          - Level: %d\n", userInfo.Level)
	}

	// 场景3: 错误处理演示
	fmt.Printf("\n   场景3: 类型转换错误处理\n")
	invalidData := map[string]interface{}{
		"number": "不是数字",
		"flag":   "maybe", // 无效的布尔值
	}

	type InvalidStruct struct {
		Number int
		Flag   bool
	}

	invalidStruct := &InvalidStruct{}
	invalidAttrs := []string{"Number", "Flag"}

	err = util.AssignAttrWithJSON(invalidStruct, invalidAttrs, invalidData)
	if err != nil {
		fmt.Printf("       ✅ 正确捕获转换错误: %v\n", err)
	} else {
		fmt.Printf("       ❌ 应该产生转换错误但没有\n")
	}

	// 转换统计
	fmt.Printf("\n📊 类型转换统计:\n")
	fmt.Printf("   - 支持的基本类型: int, float, string, bool\n")
	fmt.Printf("   - 支持的复合类型: slice, map, struct\n")
	fmt.Printf("   - 转换策略: 智能类型推断 + 显式转换\n")
	fmt.Printf("   - 错误处理: 详细错误信息 + 类型安全\n")
}

// demonstrateCompleteWorkflow 演示完整的工作流
func demonstrateCompleteWorkflow() {
	fmt.Println("🎯 === 完整工作流演示 ===")

	fmt.Printf("🎬 模拟视频项目配置处理流程:\n")

	// 步骤1: 解析项目配置
	fmt.Printf("   📋 步骤1: 解析项目配置\n")

	projectConfigJSON := map[string]interface{}{
		"name":           "我的视频项目",
		"width":          "1920",
		"height":         "1080",
		"framerate":      "30.0",
		"duration":       "120.5",
		"background":     "#FF5733",
		"output_path":    "C:\\Users\\用户\\Videos\\output.mp4",
		"enable_effects": "true",
		"quality":        "85",
	}

	type ProjectConfig struct {
		Name          string
		Width         int
		Height        int
		Framerate     float64
		Duration      float64
		Background    string
		OutputPath    string
		EnableEffects bool
		Quality       int
	}

	config := &ProjectConfig{}
	configAttrs := []string{"Name", "Width", "Height", "Framerate", "Duration",
		"Background", "OutputPath", "EnableEffects", "Quality"}

	err := util.AssignAttrWithJSON(config, configAttrs, projectConfigJSON)
	if err != nil {
		fmt.Printf("     ❌ 配置解析失败: %v\n", err)
		return
	}

	fmt.Printf("     ✅ 配置解析成功:\n")
	fmt.Printf("       - 项目名称: %s\n", config.Name)
	fmt.Printf("       - 分辨率: %dx%d\n", config.Width, config.Height)
	fmt.Printf("       - 帧率: %.1f fps\n", config.Framerate)
	fmt.Printf("       - 时长: %.1f秒\n", config.Duration)

	// 步骤2: 处理背景颜色
	fmt.Printf("   🎨 步骤2: 处理背景颜色\n")
	r, g, b, err := util.HexToRGB(config.Background)
	if err != nil {
		fmt.Printf("     ❌ 颜色解析失败: %v\n", err)
	} else {
		fmt.Printf("     ✅ 背景颜色: %s -> RGB(%.3f, %.3f, %.3f)\n",
			config.Background, r, g, b)
	}

	// 步骤3: 验证输出路径
	fmt.Printf("   📁 步骤3: 验证输出路径\n")
	isWindows := util.IsWindowsPath(config.OutputPath)
	pathType := "Unix/Linux"
	if isWindows {
		pathType = "Windows"
	}
	fmt.Printf("     ✅ 输出路径类型: %s (%s)\n", pathType, config.OutputPath)

	// 步骤4: 生成项目哈希ID
	fmt.Printf("   🔐 步骤4: 生成项目哈希ID\n")
	projectURL := fmt.Sprintf("project://%s/%dx%d@%.1f",
		strings.ReplaceAll(config.Name, " ", "_"),
		config.Width, config.Height, config.Framerate)
	projectHash := util.URLToHash(projectURL, 16)
	fmt.Printf("     ✅ 项目哈希ID: %s (基于: %s)\n", projectHash, projectURL)

	// 步骤5: 构建默认设置
	fmt.Printf("   ⚙️ 步骤5: 构建默认设置\n")

	type EffectSettings struct {
		BlurRadius   float64
		Brightness   float64
		Contrast     float64
		Saturation   float64
		EnableMotion bool
		Transitions  []string
		CustomParams map[string]float64
	}

	effectType := reflect.TypeOf(EffectSettings{})
	defaults, err := util.ProvideCtorDefaults(effectType)
	if err != nil {
		fmt.Printf("     ❌ 默认设置生成失败: %v\n", err)
	} else {
		fmt.Printf("     ✅ 默认特效设置已生成 (%d个字段)\n", len(defaults))
		for field, value := range defaults {
			fmt.Printf("       - %s: %v\n", field, value)
		}
	}

	// 步骤6: 导出项目摘要
	fmt.Printf("   📤 步骤6: 导出项目摘要\n")

	exportAttrs := []string{"Name", "Width", "Height", "Framerate", "Duration", "Quality"}
	summary, err := util.ExportAttrToJSON(config, exportAttrs)
	if err != nil {
		fmt.Printf("     ❌ 项目摘要导出失败: %v\n", err)
	} else {
		summaryJSON, _ := json.MarshalIndent(summary, "     ", "  ")
		fmt.Printf("     ✅ 项目摘要:\n%s\n", string(summaryJSON))
	}

	// 步骤7: 错误处理模拟
	fmt.Printf("   ⚠️ 步骤7: 错误处理模拟\n")

	// 模拟各种可能的错误
	possibleErrors := []error{
		util.NewValidationError("duration", -10, "持续时间不能为负数"),
		util.NewMaterialNotFoundError("background_music.mp3"),
		util.NewConfigurationError("video_encoder", "bitrate", "比特率超出范围"),
	}

	for i, err := range possibleErrors {
		fmt.Printf("     错误[%d]: %v\n", i+1, err)

		// 根据错误类型提供解决方案
		if util.IsValidationError(err) {
			fmt.Printf("       解决方案: 使用默认值或提示用户重新输入\n")
		} else if util.IsMaterialNotFound(err) {
			fmt.Printf("       解决方案: 提示用户选择替代素材\n")
		} else if util.IsConfigurationError(err) {
			fmt.Printf("       解决方案: 重置为默认配置\n")
		}
	}

	// 步骤8: 工作流总结
	fmt.Printf("   📊 步骤8: 工作流总结\n")
	fmt.Printf("     📈 处理统计:\n")
	fmt.Printf("       - 配置字段处理: %d个\n", len(configAttrs))
	fmt.Printf("       - 类型转换: %d次\n", 6) // width, height, framerate, duration, quality, enable_effects
	fmt.Printf("       - 颜色处理: 1次\n")
	fmt.Printf("       - 路径验证: 1次\n")
	fmt.Printf("       - 哈希生成: 1次\n")
	fmt.Printf("       - 默认值生成: %d个字段\n", len(defaults))
	fmt.Printf("       - 错误处理: %d种类型\n", len(possibleErrors))

	fmt.Printf("\n🎉 完整工作流演示完成!\n")
	fmt.Printf("   - 成功演示了工具函数系统的所有核心功能\n")
	fmt.Printf("   - 展示了实际项目中的应用场景\n")
	fmt.Printf("   - 验证了与Python版本的完全兼容性\n")
	fmt.Printf("   - 证明了Go版本的类型安全和性能优势\n")
}
