package router

import (
	"context"
	"embed"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

//go:embed dist
var distFS embed.FS

// Register registers the routes for static files
func RegisterStatic(r *server.Hertz) {
	// Create a temp directory in the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory:", err)
	}

	tempDir := filepath.Join(cwd, "tmp")
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		log.Fatal("Failed to create temp directory:", err)
	}

	// Extract embedded files to the temp directory
	if err := extractEmbeddedFiles(distFS, "dist", tempDir); err != nil {
		log.Fatal("Failed to extract embedded files:", err)
	}

	log.Printf("Static files extracted to: %s", tempDir)

	// Configure Hertz to serve files from the extracted directory
	r.StaticFS("/", &app.FS{
		Root:        "./tmp",
		PathRewrite: app.NewPathSlashesStripper(1),
		PathNotFound: func(_ context.Context, ctx *app.RequestContext) {
			// For SPA - serve index.html for any unknown routes
			indexPath := filepath.Join(tempDir, "index.html")
			content, err := os.ReadFile(indexPath)
			if err != nil {
				log.Printf("Error reading index.html: %v", err)
				ctx.JSON(consts.StatusNotFound, "The requested resource does not exist")
				return
			}

			ctx.Response.Header.SetContentType("text/html; charset=utf-8")
			ctx.Response.SetStatusCode(consts.StatusOK)
			ctx.Response.SetBody(content)
		},
		IndexNames:      []string{"index.html"},
		AcceptByteRange: true,
	})

	// Assets directory still served normally
	r.StaticFS("assets", &app.FS{
		Root:        "./tmp/assets",
		PathRewrite: app.NewPathSlashesStripper(1),
		PathNotFound: func(_ context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusNotFound, "The requested resource does not exist")
		},
		AcceptByteRange: true,
	})
}

// extractEmbeddedFiles extracts files from an embedded filesystem to a target directory
func extractEmbeddedFiles(embeddedFS embed.FS, sourceDir, targetDir string) error {
	// Walk through the embedded filesystem
	return fs.WalkDir(embeddedFS, sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == sourceDir {
			return nil
		}

		// Get the relative path from the source directory
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		// Create the target path
		targetPath := filepath.Join(targetDir, relPath)

		// If it's a directory, create it
		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		// If it's a file, copy its content
		sourceFile, err := embeddedFS.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		// Ensure the directory exists
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		// Create the target file
		targetFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		// Copy the content
		_, err = io.Copy(targetFile, sourceFile)
		return err
	})
}
