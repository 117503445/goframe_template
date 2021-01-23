package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GA40mgexIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwNjoDdPQs+k/X6/DQSOf7c7+tcmPPeXv/1KaemcHJ4Vi8KqfHVUt6XGN646dUzZILDxRVNEwvsr0xfq7mHtEs5wdnJddue73Xfb+bWbpRgYNtd1GEZday9WtzBUuXSmtbmqL/ajgsPhXwdu8fY/XTUh+egvg9uhfRd+uWUIeS17YxCQfDh5T8cEacuWsNf151+X7La9ce+uw7yWr6cv3/60cWdlXM67cPnup2X/uGK693R/VnUUXNbH6mxhZXtzeY1q5+8pcpo3J+tfvPHHYW3b9Rcpi5+yXJ7jFD1nUSTHXBGFqUfW3Flx8cj/wmbWpg+zZlk9ctokZLBiumwnO/udPsOuK/d73qj07Zy5rHj/+rzrn9lPdGkHNIdpZ14M43s5Ydvzv7NauraX733gMzPjy/pFT8tn3naou3Cr0ORqHrOOpGLzJ96aWs/UZWvEpsZO+au69YkxA8P//wHe7BzqvhwbXBgZGJSYGRhg0cGAER3siOiAxwBIN7KaAG9GJhFmRHQimwyKThjY1ggi8UYuwijsToEAAYb/jv2MDFgcxsoGkmdiYGLoZGBgOMsI4gECAAD///+ZEiZsAgAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
