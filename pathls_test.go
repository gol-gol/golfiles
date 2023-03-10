package golfiles

import (
	"os"
	"path"
	"testing"
)

func TestPathLs(t *testing.T) {
	result, err := PathLs(".")
	if err != nil {
		t.Error("FAILED calling PathLs.")
	}
	if len(result) < 1 {
		t.Error("FAILED to list files in '.'.")
	}

	hasGitConfig := false
	cwd, _ := os.Getwd()
	gitConfigPath := path.Join(cwd, ".git", "config")
	for _, f := range result {
		if f == gitConfigPath {
			hasGitConfig = true
			break
		}
	}
	if !hasGitConfig {
		t.Errorf("FAILED to list file: %s", gitConfigPath)
	}
}

func TestPathLsN(t *testing.T) {
	result, err := PathLsN(".", 0)
	if err != nil {
		t.Error("FAILED calling PathLsN.")
	}
	if len(result) < 1 {
		t.Error("FAILED to list files in '.'.")
	}

	hasGitConfig := false
	cwd, _ := os.Getwd()
	gitConfigPath := path.Join(cwd, ".git", "config")
	for _, f := range result {
		if f == gitConfigPath {
			hasGitConfig = true
			break
		}
	}
	if hasGitConfig {
		t.Errorf("FAILED to list file: %s", gitConfigPath)
	}
}

func TestPathLsType(t *testing.T) {
	result, err := PathLsType(".", []string{".md"})
	if err != nil {
		t.Error("FAILED calling PathLsType.")
	}

	hasReadme := false
	cwd, _ := os.Getwd()
	readmePath := path.Join(cwd, "README.md")
	for _, f := range result {
		if f == readmePath {
			hasReadme = true
			break
		}
	}
	if !hasReadme {
		t.Errorf("FAILED to list file: %s", readmePath)
	}
}

func TestPathLsNType(t *testing.T) {
	result, err := PathLsNType(".", 0, []string{".md"})
	if err != nil {
		t.Error("FAILED calling PathLsNType.")
	}

	hasReadme := false
	cwd, _ := os.Getwd()
	readmePath := path.Join(cwd, "README.md")
	for _, f := range result {
		if f == readmePath {
			hasReadme = true
			break
		}
	}
	if !hasReadme {
		t.Errorf("FAILED to list file: %s", readmePath)
	}
}

func TestPathLsTypeNot(t *testing.T) {
	result, err := PathLsTypeNot(".", []string{".md"})
	if err != nil {
		t.Error("FAILED calling PathLsTypeNot.")
	}

	hasReadme := false
	hasGitConfig := false
	cwd, _ := os.Getwd()
	readmePath := path.Join(cwd, "README.md")
	gitConfigPath := path.Join(cwd, ".git", "config")
	for _, f := range result {
		if f == readmePath {
			hasReadme = true
		}
		if f == gitConfigPath {
			hasGitConfig = true
		}
	}
	if hasReadme {
		t.Errorf("FAILED to not list file: %s", readmePath)
	}
	if !hasGitConfig {
		t.Errorf("FAILED to list file: %s", gitConfigPath)
	}
}

func TestPathLsNTypeNot(t *testing.T) {
	result, err := PathLsNTypeNot(".", 0, []string{".md"})
	if err != nil {
		t.Error("FAILED calling PathLsTypeNot.")
	}

	hasReadme := false
	hasGitConfig := false
	hasLicense := false
	cwd, _ := os.Getwd()
	readmePath := path.Join(cwd, "README.md")
	gitConfigPath := path.Join(cwd, ".git", "config")
	licensePath := path.Join(cwd, "LICENSE")
	for _, f := range result {
		if f == readmePath {
			hasReadme = true
		}
		if f == gitConfigPath {
			hasGitConfig = true
		}
		if f == licensePath {
			hasLicense = true
		}
	}
	if hasReadme {
		t.Errorf("FAILED to not list file: %s", readmePath)
	}
	if hasGitConfig {
		t.Errorf("FAILED to not list file: %s", gitConfigPath)
	}
	if !hasLicense {
		t.Errorf("FAILED to list file: %s", licensePath)
	}
}
