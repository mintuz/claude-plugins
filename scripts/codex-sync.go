package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

type SyncStats struct {
	SkillsSynced int
	SkillsFailed int
	FilesCreated int
}

func main() {
	// Parse command-line flags
	outputDir := flag.String("output", "", "Output directory for Codex skills (default: ~/.codex/skills)")
	pluginsDir := flag.String("plugins", "./plugins", "Directory containing Claude plugins")
	marketplaceFile := flag.String("marketplace", "./.claude-plugin/marketplace.json", "Path to marketplace.json")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	dryRun := flag.Bool("dry-run", false, "Perform a dry run without copying files")
	projectLevel := flag.Bool("project", false, "Install to .codex/skills in current directory instead of ~/.codex/skills")
	usePrefix := flag.Bool("prefix", false, "Prefix skill names with plugin name (e.g., core-commit-messages)")
	flag.Parse()

	// Determine output directory
	var targetDir string
	if *outputDir != "" {
		targetDir = *outputDir
	} else if *projectLevel {
		targetDir = ".codex/skills"
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fatal("Failed to get home directory: %v", err)
		}
		targetDir = filepath.Join(home, ".codex", "skills")
	}

	// Convert to absolute path
	absTargetDir, err := filepath.Abs(targetDir)
	if err != nil {
		fatal("Failed to resolve target directory: %v", err)
	}

	// Print configuration
	printHeader("Codex Skills Sync")
	fmt.Printf("%sTarget directory:%s %s\n", colorBlue, colorReset, absTargetDir)
	fmt.Printf("%sPlugins directory:%s %s\n", colorBlue, colorReset, *pluginsDir)
	if *dryRun {
		fmt.Printf("%sDry run mode: No files will be modified%s\n", colorYellow, colorReset)
	}
	fmt.Println()

	// Read marketplace.json
	marketplace, err := readMarketplace(*marketplaceFile)
	if err != nil {
		fatal("Failed to read marketplace.json: %v", err)
	}

	// Sync skills
	stats := &SyncStats{}
	for _, plugin := range marketplace.Plugins {
		syncPlugin(plugin, absTargetDir, *verbose, *dryRun, *usePrefix, stats)
	}

	// Print summary
	printSummary(stats, *dryRun)
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

func syncPlugin(plugin Plugin, targetDir string, verbose bool, dryRun bool, usePrefix bool, stats *SyncStats) {
	if len(plugin.Skills) == 0 {
		if verbose {
			fmt.Printf("%s[SKIP]%s Plugin '%s' has no skills\n", colorYellow, colorReset, plugin.Name)
		}
		return
	}

	fmt.Printf("\n%s=== Syncing plugin: %s ===%s\n", colorBlue, plugin.Name, colorReset)

	for _, skillPath := range plugin.Skills {
		if err := syncSkill(plugin.Name, skillPath, targetDir, verbose, dryRun, usePrefix, stats); err != nil {
			fmt.Printf("%s[ERROR]%s Failed to sync %s: %v\n", colorRed, colorReset, skillPath, err)
			stats.SkillsFailed++
		} else {
			stats.SkillsSynced++
		}
	}
}

func syncSkill(pluginName, skillPath, targetDir string, verbose bool, dryRun bool, usePrefix bool, stats *SyncStats) error {
	// Extract skill name from path (e.g., "./plugins/core/skills/commit-messages" -> "commit-messages")
	skillName := filepath.Base(skillPath)

	// Create Codex skill name (with optional plugin prefix)
	var codexSkillName string
	if usePrefix {
		codexSkillName = fmt.Sprintf("%s-%s", pluginName, skillName)
	} else {
		codexSkillName = skillName
	}

	// Source and destination paths
	srcDir, err := filepath.Abs(skillPath)
	if err != nil {
		return fmt.Errorf("failed to resolve source path: %w", err)
	}

	dstDir := filepath.Join(targetDir, codexSkillName)

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
		fmt.Printf("  %s → %s\n", srcDir, dstDir)
	}

	if dryRun {
		fmt.Printf("%s[DRY RUN]%s Would copy: %s\n", colorYellow, colorReset, codexSkillName)
		return nil
	}

	// Remove existing destination if it exists
	if _, err := os.Lstat(dstDir); err == nil {
		if err := os.RemoveAll(dstDir); err != nil {
			return fmt.Errorf("failed to remove existing destination: %w", err)
		}
	}

	// Ensure parent directory exists
	parentDir := filepath.Dir(dstDir)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	// Create destination directory
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Recursively copy all files
	fileCount := 0
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from source directory
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		// Destination path
		destPath := filepath.Join(dstDir, relPath)

		// If it's a directory, create it
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		// Copy file
		if err := copyFile(path, destPath); err != nil {
			return fmt.Errorf("failed to copy %s: %w", relPath, err)
		}

		fileCount++
		if verbose {
			fmt.Printf("    %s✓%s Copied: %s\n", colorGreen, colorReset, relPath)
		}

		return nil
	})

	if err != nil {
		return err
	}

	stats.FilesCreated += fileCount
	fmt.Printf("%s[SYNCED]%s %s (%d files copied)\n", colorGreen, colorReset, codexSkillName, fileCount)

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	// Copy file permissions
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, sourceInfo.Mode())
}

func printHeader(title string) {
	fmt.Println()
	fmt.Printf("%s╔═══════════════════════════════════════════════════════╗%s\n", colorBlue, colorReset)
	fmt.Printf("%s║%s  %-50s %s║%s\n", colorBlue, colorReset, title, colorBlue, colorReset)
	fmt.Printf("%s╚═══════════════════════════════════════════════════════╝%s\n", colorBlue, colorReset)
	fmt.Println()
}

func printSummary(stats *SyncStats, dryRun bool) {
	fmt.Println()
	fmt.Printf("%s╔═══════════════════════════════════════════════════════╗%s\n", colorGreen, colorReset)
	fmt.Printf("%s║%s  %-50s %s║%s\n", colorGreen, colorReset, "Summary", colorGreen, colorReset)
	fmt.Printf("%s╚═══════════════════════════════════════════════════════╝%s\n", colorGreen, colorReset)

	if dryRun {
		fmt.Printf("\n%sDry run completed - no files were modified%s\n", colorYellow, colorReset)
	}

	fmt.Printf("\n%sSkills synced:%s     %d\n", colorBlue, colorReset, stats.SkillsSynced)
	if stats.SkillsFailed > 0 {
		fmt.Printf("%sSkills failed:%s     %d\n", colorRed, colorReset, stats.SkillsFailed)
	}
	if !dryRun {
		fmt.Printf("%sFiles created:%s     %d\n", colorBlue, colorReset, stats.FilesCreated)
	}
	fmt.Println()

	if stats.SkillsSynced > 0 && !dryRun {
		fmt.Printf("%s✓ Successfully synced skills to Codex!%s\n\n", colorGreen, colorReset)
		fmt.Printf("You can now use these skills in Codex by typing $<skill-name>\n")
		fmt.Printf("Example: $commit-messages or $react\n\n")
	}
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%sERROR: %s%s\n", colorRed, fmt.Sprintf(format, args...), colorReset)
	os.Exit(1)
}
