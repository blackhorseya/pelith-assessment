package web

import (
	"embed"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

//go:embed templates/**/*
var f embed.FS

// SetHTMLTemplate sets up the HTML templates with multitemplate and handles nested directories
func SetHTMLTemplate(r *gin.Engine) {
	renderer := multitemplate.NewRenderer()

	// 基础模板路径
	baseTemplate := "templates/layout/base.templ"

	// 递归读取所有模板文件
	err := walkTemplates("templates", func(templatePath string) {
		if strings.HasSuffix(templatePath, "base.templ") {
			return // 跳过基础模板
		}

		// 生成模板名称，去掉路径和扩展名
		relativePath := strings.TrimPrefix(templatePath, "templates/")
		templateName := strings.TrimSuffix(relativePath, filepath.Ext(relativePath))

		// 判断是否需要嵌套 base.templ
		if strings.HasPrefix(templatePath, "templates/layout/") {
			// 如果是部分模板（如 partials 文件夹中的模板），不嵌套 base.templ
			renderer.AddFromFS(templateName, f, templatePath)
		} else {
			// 默认嵌套 base.templ
			renderer.AddFromFS(templateName, f, baseTemplate, templatePath)
		}
	})
	if err != nil {
		panic("failed to walk through templates: " + err.Error())
	}

	// 使用 multitemplate 渲染器
	r.HTMLRender = renderer
}

// walkTemplates recursively walks through the template directory
func walkTemplates(root string, fn func(string)) error {
	entries, err := f.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(root, entry.Name())
		if entry.IsDir() {
			// 如果是子目录，递归处理
			if err = walkTemplates(fullPath, fn); err != nil {
				return err
			}
		} else {
			// 如果是文件，调用处理函数
			fn(fullPath)
		}
	}
	return nil
}
