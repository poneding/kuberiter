/*
Copyright 2023 The Kuberiter Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package article

import (
	"github.com/gin-gonic/gin"
	corev1alpha1 "github.com/poneding/kuberiter/pkg/apis/core/v1alpha1"
	"github.com/poneding/kuberiter/pkg/server"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var articles = []corev1alpha1.Article{
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Article",
			APIVersion: "core.kuberiter.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "article-0",
		},
		Spec: corev1alpha1.ArticleSpec{
			Title:       "Hello World",
			Description: "This is a sample article",
			Content:     "Hello World.",
			CreatedAt:   metav1.Now(),
			UpdatedAt:   metav1.Now(),
			Authors: []string{
				"kuberiter",
			},
			Likes: 0,
		},
	},
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Article",
			APIVersion: "core.kuberiter.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "article-1",
		},
		Spec: corev1alpha1.ArticleSpec{
			Title:       "Hello Kuberiter",
			Description: "This is a sample article",
			Content: `# Hello Kuberiter
			This is a sample article. You can write your article in ` + "`markdown`" + ` format.
			## Images
			![Kuberiter Logo](https://raw.githubusercontent.com/poneding/kuberiter/master/docs/images/writer.jpg)
			## Codes\n` +
				"```go" + `
			package main
			import \"fmt\"
			func main() {
			  # Print Hello Kuberiter
			  fmt.Println("Hello Kuberiter")
			}` + "```",
			CreatedAt: metav1.Now(),
			UpdatedAt: metav1.Now(),
			Authors: []string{
				"kuberiter",
			},
			Likes: 0,
		},
	},
}

func List(c *gin.Context) {
	c.JSONP(200, server.Response{
		Message: "success",
		Data:    articles,
	})
}
