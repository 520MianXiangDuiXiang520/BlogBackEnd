// Package sitemap builds the public XML sitemap from published articles.
package sitemap

import (
	"JuneBlog/internal/config"
	"JuneBlog/internal/db"
	"JuneBlog/internal/db/module"
	"JuneBlog/internal/db/opt"
	"context"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	baseURL  = "https://junebao.top"
	pageSize = int64(500)
)

type sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []url    `xml:"url"`
}

type url struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod"`
	ChangeFreq string `xml:"changefreq"`
	Priority   string `xml:"priority"`
}

// Generate writes a complete sitemap atomically. It includes the homepage and
// every article currently in the database, so updates and deletions are also
// reflected in the file.
func Generate(ctx context.Context) error {
	articles, err := allArticles(ctx)
	if err != nil {
		return err
	}

	today := time.Now().Format("2006-01-02")
	result := sitemap{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  []url{{Loc: baseURL + "/", LastMod: today, ChangeFreq: "daily", Priority: "1.0"}},
	}
	for _, article := range articles {
		ts := article.UpdateTs
		if ts == 0 {
			ts = article.CreateTs
		}
		if ts > 1e10 {
			ts /= 1000
		}
		result.URLs = append(result.URLs, url{
			Loc:        fmt.Sprintf("%s/detail/%d", baseURL, article.Id),
			LastMod:    time.Unix(ts, 0).Format("2006-01-02"),
			ChangeFreq: "weekly",
			Priority:   "0.8",
		})
	}

	data, err := xml.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	data = append([]byte(xml.Header), data...)

	path := config.G.GetStrWithDefault(config.CommonKeySitemapPath, "/data/juneblog/frontend/www/sitemap.xml")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".sitemap-*.xml")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if err := tmp.Chmod(0644); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, path)
}

func allArticles(ctx context.Context) ([]*module.ArticleHeader, error) {
	count, err := db.Db().FindArticleCount(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	pages := (int64(count) + pageSize - 1) / pageSize
	articles := make([]*module.ArticleHeader, 0, count)
	for page := int64(1); page <= pages; page++ {
		items, err := db.Db().FindSomeArticleInfo(ctx,
			opt.WithPage(page), opt.WithPageSize(pageSize), opt.WithOrderBy("id", true))
		if err != nil {
			return nil, err
		}
		articles = append(articles, items...)
	}
	return articles, nil
}
