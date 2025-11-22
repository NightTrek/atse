package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	sizeFlag    int
	langFlag    string
	outputFlag  string
	workersFlag int
	seedFlag    int64
)

func init() {
	flag.IntVar(&sizeFlag, "size", 100, "Number of files to generate")
	flag.StringVar(&langFlag, "lang", "go", "Language (go, typescript, python)")
	flag.StringVar(&outputFlag, "output", "", "Output directory")
	flag.IntVar(&workersFlag, "workers", 8, "Number of concurrent workers")
	flag.Int64Var(&seedFlag, "seed", 42, "Random seed")
	flag.Parse()
}

func main() {
	if outputFlag == "" {
		fmt.Println("Error: --output directory is required")
		os.Exit(1)
	}

	fmt.Printf("Generating %d %s files in %s...\n", sizeFlag, langFlag, outputFlag)
	rand.Seed(seedFlag)

	// Ensure output directory exists
	if err := os.MkdirAll(outputFlag, 0755); err != nil {
		panic(err)
	}

	// Generate file paths first to ensure even distribution
	// We use a logarithmic depth structure to mimic real repos
	// e.g., root/pkg1/subpkg2/file.go
	filePaths := generateFilePaths(sizeFlag, langFlag)

	// Workers pool
	var wg sync.WaitGroup
	pathChan := make(chan string, sizeFlag)

	for i := 0; i < workersFlag; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range pathChan {
				fullPath := filepath.Join(outputFlag, path)
				content := generateContent(langFlag)

				if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
					fmt.Printf("Error creating dir: %v\n", err)
					continue
				}

				if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
					fmt.Printf("Error writing file: %v\n", err)
				}
			}
		}()
	}

	start := time.Now()
	for _, p := range filePaths {
		pathChan <- p
	}
	close(pathChan)
	wg.Wait()

	fmt.Printf("Done in %s\n", time.Since(start))
}

func generateFilePaths(count int, lang string) []string {
	ext := ".go"
	switch lang {
	case "typescript":
		ext = ".ts"
	case "python":
		ext = ".py"
	}

	paths := make([]string, count)

	// Heuristic: ~10-20 files per directory
	// Determine roughly how many directories we need
	dirCount := count / 15
	if dirCount < 1 {
		dirCount = 1
	}

	directories := make([]string, dirCount)
	directories[0] = "." // Root

	// Generate directory tree
	for i := 1; i < dirCount; i++ {
		// each new directory is a child of an existing random directory
		parent := directories[rand.Intn(i)]
		name := fmt.Sprintf("pkg_%d", i)
		directories[i] = filepath.Join(parent, name)
	}

	// Distribute files
	for i := 0; i < count; i++ {
		dir := directories[rand.Intn(dirCount)]
		name := fmt.Sprintf("file_%d%s", i, ext)
		paths[i] = filepath.Join(dir, name)
	}

	return paths
}

func generateContent(lang string) string {
	switch lang {
	case "go":
		return generateGoContent()
	case "typescript":
		return generateTSContent()
	case "python":
		return generatePythonContent()
	default:
		return ""
	}
}

// --- Content Generators ---

var commonPatterns = []string{"ProcessRequest", "HandleError", "ValidateInput", "TransformData", "LogEvent"}
var rarePatterns = []string{"DepreciatedLegacyFunctionXYZ", "RareEdgeCaseHandler", "SecretBackdoorV2"}

func pickPattern() string {
	r := rand.Float32()
	if r < 0.05 { // 5% chance of rare pattern
		return rarePatterns[rand.Intn(len(rarePatterns))]
	}
	if r < 0.3 { // 25% chance of common pattern
		return commonPatterns[rand.Intn(len(commonPatterns))]
	}
	// 70% chance of random name
	return fmt.Sprintf("Func_%d", rand.Intn(100000))
}

func generateGoContent() string {
	return fmt.Sprintf(`package main

import (
	"fmt"
	"os"
)

type GenericAuthService struct {
	ID string
}

func (s *GenericAuthService) Authenticate() bool {
	return true
}

func %s() {
	fmt.Println("Running")
}

func %s() {
	// TODO: Implement
}
`, pickPattern(), pickPattern())
}

func generateTSContent() string {
	return fmt.Sprintf(`
import * as fs from 'fs';

class GenericAuthService {
  id: string;
  constructor(id: string) {
    this.id = id;
  }

  authenticate(): boolean {
    return true;
  }
}

export function %s() {
  console.log("Running");
}

function %s() {
   // TODO
}
`, pickPattern(), pickPattern())
}

func generatePythonContent() string {
	return fmt.Sprintf(`
import os
import sys

class GenericAuthService:
    def __init__(self, id):
        self.id = id

    def authenticate(self):
        return True

def %s():
    print("Running")

def %s():
    pass
`, pickPattern(), pickPattern())
}
