package services

import (
	"blogstreak/models"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/mermaid"
)

func parseMD(data []byte) (*models.Blog, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithExtensions(
			meta.Meta,
		),
		goldmark.WithExtensions(highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithWrapperRenderer(func(w util.BufWriter, context highlighting.CodeBlockContext, entering bool) {
				if entering {
					_, _ = w.WriteString(`<div class="not-prose rounded-md overflow-hidden">`)

				} else {
					_, _ = w.WriteString(`</div>`)
				}
			}),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
				chromahtml.WrapLongLines(true),
			),
		),
		),
		goldmark.WithExtensions(&mermaid.Extender{}),
	)

	var buf bytes.Buffer
	ctx := parser.NewContext()
	if err := markdown.Convert(data, &buf, parser.WithContext(ctx)); err != nil {
		return nil, err
	}

	metadata := meta.Get(ctx)

	title, ok := metadata["Title"]
	if !ok {
		return nil, fmt.Errorf("Title not found on this blog")
	}

	publishedDate, ok := metadata["PublishedDate"]
	if !ok {
		return nil, fmt.Errorf("publishedDate not found on this blog")
	}

	return &models.Blog{
		Title: title.(string),
		Body: templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			w.Write(buf.Bytes())
			return nil
		}),
		PublishedDate: publishedDate.(string),
	}, nil
}
