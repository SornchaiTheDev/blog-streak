---
Title: I don't know yet lol 
PublishedDate: 03/01/2025
---

### Getting Started making Blog Streak

Today I haved started code this website using go + Templ.I think use only these tools is enough to make a blog website for now.Maybe if I want to implement other things such as reactions or comments HTMX would be a great choice.

### Implement Markdown -> HTML

I use `github.com/yuin/goldmark` as a markdown parser tool (found this package on [templ.guide](https://templ.guide/) website) and custom some CSS styling using TailwindCSS `prose` class from [Typography](https://github.com/tailwindlabs/tailwindcss-typography) plugin

### Test Markdown Rendering

# Heading 1

## Heading 2

### Heading 3

#### Heading 4

##### Heading 5

###### Heading 6

Paragraph

Unordered list

- first
- second
- third

Ordered list

1. first
2. second
3. third

```go
		goldmark.WithExtensions(highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithWrapperRenderer(func(w util.BufWriter, context highlighting.CodeBlockContext, entering bool) {
				if entering {
					_, _ = w.WriteString(`<div class="not-prose">`)

				} else {
					_, _ = w.WriteString(`</div>`)
				}
			}),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
			),
		),
```

Horizontal rule

----------

| Column1 | Column2 | Column3 |
| ------------- | -------------- | -------------- |
| Item1 | Item1 | Item1 |

[Link](https://github.com/SornchaiTheDev) 

> [!NOTE]  
> Highlights information that users should take into account, even when skimming.

Todo List

- [ ] text

**Bold Text**

***Italic Text***

Hello {#id}

![Gopher Mascot](https://dwglogo.com/wp-content/uploads/2017/08/muscles-clipart-ghoper.gif) { style="width:100px" }

##### Mermaid.js

```mermaid
graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
```

