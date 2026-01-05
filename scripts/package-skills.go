package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorRed    = "\033[31m"
)

type MarketplaceConfig struct {
	Name    string   `json:"name"`
	Owner   Owner    `json:"owner"`
	Plugins []Plugin `json:"plugins"`
}

type Owner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

type Plugin struct {
	Name        string   `json:"name"`
	Source      string   `json:"source"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
}

type PackageStats struct {
	SkillsPackaged int
	SkillsFailed   int
	FilesAdded     int
}

func main() {
	// Parse command-line flags
	outputFile := flag.String("output", "", "Output zip file path (default: claude-plugins-skills.zip)")
	marketplaceFile := flag.String("marketplace", "./.claude-plugin/marketplace.json", "Path to marketplace.json")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	dryRun := flag.Bool("dry-run", false, "Perform a dry run without creating zip file")
	usePrefix := flag.Bool("prefix", false, "Prefix skill names with plugin name (e.g., core-commit-messages)")
	flag.Parse()

	// Determine output file path
	zipPath := *outputFile
	if zipPath == "" {
		timestamp := time.Now().Format("20060102-150405")
		zipPath = fmt.Sprintf("claude-plugins-skills-%s.zip", timestamp)
	}

	// Convert to absolute path
	absZipPath, err := filepath.Abs(zipPath)
	if err != nil {
		fatal("Failed to resolve output path: %v", err)
	}

	// Print configuration
	printHeader("Package Skills to Zip")
	fmt.Printf("%sOutput file:%s %s\n", colorBlue, colorReset, absZipPath)
	if *dryRun {
		fmt.Printf("%sDry run mode: No files will be created%s\n", colorYellow, colorReset)
	}
	fmt.Println()

	// Read marketplace.json
	marketplace, err := readMarketplace(*marketplaceFile)
	if err != nil {
		fatal("Failed to read marketplace.json: %v", err)
	}

	// Create zip file
	stats := &PackageStats{}
	if !*dryRun {
		if err := createSkillsZip(absZipPath, marketplace, *verbose, *usePrefix, stats); err != nil {
			fatal("Failed to create zip file: %v", err)
		}
	} else {
		// Dry run - just validate skills
		for _, plugin := range marketplace.Plugins {
			validatePlugin(plugin, *verbose, *usePrefix, stats)
		}
	}

	// Print summary
	printSummary(stats, absZipPath, *dryRun)
}

func readMarketplace(path string) (*MarketplaceConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config MarketplaceConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func createSkillsZip(zipPath string, marketplace *MarketplaceConfig, verbose bool, usePrefix bool, stats *PackageStats) error {
	// Create zip file
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Process each plugin
	for _, plugin := range marketplace.Plugins {
		if err := packagePlugin(plugin, zipWriter, verbose, usePrefix, stats); err != nil {
			fmt.Printf("%s[ERROR]%s Failed to package plugin '%s': %v\n", colorRed, colorReset, plugin.Name, err)
			return err
		}
	}

	return nil
}

func validatePlugin(plugin Plugin, verbose bool, usePrefix bool, stats *PackageStats) {
	if len(plugin.Skills) == 0 {
		if verbose {
			fmt.Printf("%s[SKIP]%s Plugin '%s' has no skills\n", colorYellow, colorReset, plugin.Name)
		}
		return
	}

	fmt.Printf("\n%s=== Validating plugin: %s ===%s\n", colorBlue, plugin.Name, colorReset)

	for _, skillPath := range plugin.Skills {
		skillName := filepath.Base(skillPath)
		var packagedName string
		if usePrefix {
			packagedName = fmt.Sprintf("%s-%s", plugin.Name, skillName)
		} else {
			packagedName = skillName
		}

		srcDir, err := filepath.Abs(skillPath)
		if err != nil {
			fmt.Printf("%s[ERROR]%s Failed to resolve %s: %v\n", colorRed, colorReset, skillPath, err)
			stats.SkillsFailed++
			continue
		}

		if _, err := os.Stat(srcDir); os.IsNotExist(err) {
			fmt.Printf("%s[ERROR]%s Source directory does not exist: %s\n", colorRed, colorReset, srcDir)
			stats.SkillsFailed++
			continue
		}

		skillFile := filepath.Join(srcDir, "SKILL.md")
		if _, err := os.Stat(skillFile); os.IsNotExist(err) {
			fmt.Printf("%s[ERROR]%s SKILL.md not found in %s\n", colorRed, colorReset, srcDir)
			stats.SkillsFailed++
			continue
		}

		fmt.Printf("%s[DRY RUN]%s Would package: %s\n", colorYellow, colorReset, packagedName)
		stats.SkillsPackaged++
	}
}

func packagePlugin(plugin Plugin, zipWriter *zip.Writer, verbose bool, usePrefix bool, stats *PackageStats) error {
	if len(plugin.Skills) == 0 {
		if verbose {
			fmt.Printf("%s[SKIP]%s Plugin '%s' has no skills\n", colorYellow, colorReset, plugin.Name)
		}
		return nil
	}

	fmt.Printf("\n%s=== Packaging plugin: %s ===%s\n", colorBlue, plugin.Name, colorReset)

	for _, skillPath := range plugin.Skills {
		if err := packageSkill(plugin.Name, skillPath, zipWriter, verbose, usePrefix, stats); err != nil {
			fmt.Printf("%s[ERROR]%s Failed to package %s: %v\n", colorRed, colorReset, skillPath, err)
			stats.SkillsFailed++
		} else {
			stats.SkillsPackaged++
		}
	}

	return nil
}

func packageSkill(pluginName, skillPath string, zipWriter *zip.Writer, verbose bool, usePrefix bool, stats *PackageStats) error {
	// Extract skill name from path
	skillName := filepath.Base(skillPath)

	// Create packaged skill name (with optional plugin prefix)
	var packagedName string
	if usePrefix {
		packagedName = fmt.Sprintf("%s-%s", pluginName, skillName)
	} else {
		packagedName = skillName
	}

	// Source path
	srcDir, err := filepath.Abs(skillPath)
	if err != nil {
		return fmt.Errorf("failed to resolve source path: %w", err)
	}

	// Check if source exists
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return fmt.Errorf("source directory does not exist: %s", srcDir)
	}

	// Check if SKILL.md exists
	skillFile := filepath.Join(srcDir, "SKILL.md")
	if _, err := os.Stat(skillFile); os.IsNotExist(err) {
		return fmt.Errorf("SKILL.md not found in %s", srcDir)
	}

	if verbose {
		fmt.Printf("  Adding %s to zip...\n", packagedName)
	}

	// Add all files from skill directory to zip
	fileCount := 0
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Get relative path from source directory
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		// Create path in zip with skill name as root
		zipPath := filepath.Join(packagedName, relPath)

		// Add file to zip
		if err := addFileToZip(zipWriter, path, zipPath); err != nil {
			return fmt.Errorf("failed to add %s: %w", relPath, err)
		}

		fileCount++
		if verbose {
			fmt.Printf("    %s✓%s Added: %s\n", colorGreen, colorReset, zipPath)
		}

		return nil
	})

	if err != nil {
		return err
	}

	stats.FilesAdded += fileCount
	fmt.Printf("%s[PACKAGED]%s %s (%d files added)\n", colorGreen, colorReset, packagedName, fileCount)

	return nil
}

func addFileToZip(zipWriter *zip.Writer, srcPath, zipPath string) error {
	// Open source file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Get file info for permissions
	info, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// Create zip file header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Use forward slashes for zip paths (platform independent)
	header.Name = filepath.ToSlash(zipPath)
	header.Method = zip.Deflate

	// Create writer for this file in zip
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Copy file contents to zip
	if _, err := io.Copy(writer, srcFile); err != nil {
		return err
	}

	return nil
}

func printHeader(title string) {
	fmt.Println()
	fmt.Printf("%s╔═══════════════════════════════════════════════════════╗%s\n", colorBlue, colorReset)
	fmt.Printf("%s║%s  %-50s %s║%s\n", colorBlue, colorReset, title, colorBlue, colorReset)
	fmt.Printf("%s╚═══════════════════════════════════════════════════════╝%s\n", colorBlue, colorReset)
	fmt.Println()
}

func printSummary(stats *PackageStats, zipPath string, dryRun bool) {
	fmt.Println()
	fmt.Printf("%s╔═══════════════════════════════════════════════════════╗%s\n", colorGreen, colorReset)
	fmt.Printf("%s║%s  %-50s %s║%s\n", colorGreen, colorReset, "Summary", colorGreen, colorReset)
	fmt.Printf("%s╚═══════════════════════════════════════════════════════╝%s\n", colorGreen, colorReset)

	if dryRun {
		fmt.Printf("\n%sDry run completed - no files were created%s\n", colorYellow, colorReset)
	}

	fmt.Printf("\n%sSkills packaged:%s   %d\n", colorBlue, colorReset, stats.SkillsPackaged)
	if stats.SkillsFailed > 0 {
		fmt.Printf("%sSkills failed:%s     %d\n", colorRed, colorReset, stats.SkillsFailed)
	}
	if !dryRun {
		fmt.Printf("%sFiles added:%s       %d\n", colorBlue, colorReset, stats.FilesAdded)
	}
	fmt.Println()

	if stats.SkillsPackaged > 0 && !dryRun {
		// Get file size
		info, err := os.Stat(zipPath)
		if err == nil {
			size := float64(info.Size()) / 1024 / 1024 // Convert to MB
			fmt.Printf("%s✓ Successfully created zip file!%s\n", colorGreen, colorReset)
			fmt.Printf("  Location: %s\n", zipPath)
			fmt.Printf("  Size: %.2f MB\n\n", size)
		}
	}
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%sERROR: %s%s\n", colorRed, fmt.Sprintf(format, args...), colorReset)
	os.Exit(1)
}
