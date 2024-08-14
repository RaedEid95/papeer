package book

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ScrapeAndConvert(args []string, configs []*ScrapeConfig, outputDir string, c chapter) {
	// Create the output directory dynamically if it doesn't exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755) // 0755 gives read and execute permissions to everyone, and write permissions to the owner
		if err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
	}

	for _, u := range args {
		newChapter := NewChapterFromURL(u, "", configs, 0, func(index int, name string) {})
		c.AddSubChapter(newChapter)

		subChapters := newChapter.SubChapters()
		if len(subChapters) > 0 {
			newChapter.SetName(subChapters[0].Name())
			// Process subchapters
			for _, subChapter := range subChapters {
				// Create the full file path (directory + filename)
				filename := filepath.Join(outputDir, fmt.Sprintf("%s.md", Filename(subChapter.Name())))
				ToMarkdown(subChapter, filename)
				fmt.Printf("Markdown saved to \"%s\"\n", filename)
			}
		} else {
			// If there are no subchapters, process the main chapter
			newChapter.SetName("Main Chapter")
			filename := filepath.Join(outputDir, fmt.Sprintf("%s.md", Filename(newChapter.Name())))
			ToMarkdown(newChapter, filename)
			fmt.Printf("Markdown saved to \"%s\"\n", filename)
		}
	}

	// Set the name of the root chapter based on its first subchapter
	if len(c.SubChapters()) > 0 {
		c.SetName(c.SubChapters()[0].Name())
	} else {
		c.SetName("Default Root Chapter Name")
	}
}
