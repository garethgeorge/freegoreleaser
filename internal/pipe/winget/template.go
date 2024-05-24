package winget

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/caarlos0/log"
	"github.com/garethgeorge/freegoreleaser/internal/artifact"
	"github.com/garethgeorge/freegoreleaser/internal/yaml"
	"github.com/garethgeorge/freegoreleaser/pkg/config"
	"github.com/garethgeorge/freegoreleaser/pkg/context"
)

func createYAML(ctx *context.Context, winget config.Winget, in any, tp artifact.Type) error {
	versionContent, err := yaml.Marshal(in)
	if err != nil {
		return err
	}

	filename := winget.PackageIdentifier + extFor(tp)
	path := path.Join(ctx.Config.Dist, "winget", winget.Path, filename)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	log.WithField("path", path).Info("writing")
	if err := os.WriteFile(path, []byte(strings.Join([]string{
		generatedHeader,
		langserverLineFor(tp),
		string(versionContent),
	}, "\n")), 0o644); err != nil { //nolint: gosec
		return fmt.Errorf("failed to write winget version: %w", err)
	}

	ctx.Artifacts.Add(&artifact.Artifact{
		Name: filename,
		Path: path,
		Type: tp,
		Extra: map[string]interface{}{
			wingetConfigExtra: winget,
			artifact.ExtraID:  winget.Name,
		},
	})

	return nil
}

const (
	manifestVersion         = "1.6.0"
	versionLangServer       = "# yaml-language-server: $schema=https://aka.ms/winget-manifest.version.1.6.0.schema.json"
	installerLangServer     = "# yaml-language-server: $schema=https://aka.ms/winget-manifest.installer.1.6.0.schema.json"
	defaultLocaleLangServer = "# yaml-language-server: $schema=https://aka.ms/winget-manifest.defaultLocale.1.6.0.schema.json"
	defaultLocale           = "en-US"
	generatedHeader         = `# This file was generated by GoReleaser. DO NOT EDIT.`
)

//nolint:tagliatelle
type Version struct {
	PackageIdentifier string `yaml:"PackageIdentifier,omitempty"`
	PackageVersion    string `yaml:"PackageVersion,omitempty"`
	DefaultLocale     string `yaml:"DefaultLocale,omitempty"`
	ManifestType      string `yaml:"ManifestType,omitempty"`
	ManifestVersion   string `yaml:"ManifestVersion,omitempty"`
}

//nolint:tagliatelle
type InstallerItemFile struct {
	RelativeFilePath     string `yaml:"RelativeFilePath,omitempty"`
	PortableCommandAlias string `yaml:"PortableCommandAlias,omitempty"`
}

//nolint:tagliatelle
type InstallerItem struct {
	Architecture         string              `yaml:"Architecture,omitempty"`
	NestedInstallerType  string              `yaml:"NestedInstallerType,omitempty"`
	NestedInstallerFiles []InstallerItemFile `yaml:"NestedInstallerFiles,omitempty"`
	InstallerURL         string              `yaml:"InstallerUrl,omitempty"`
	InstallerSha256      string              `yaml:"InstallerSha256,omitempty"`
	UpgradeBehavior      string              `yaml:"UpgradeBehavior,omitempty"`
}

//nolint:tagliatelle
type Installer struct {
	PackageIdentifier string          `yaml:"PackageIdentifier,omitempty"`
	PackageVersion    string          `yaml:"PackageVersion,omitempty"`
	InstallerLocale   string          `yaml:"InstallerLocale,omitempty"`
	InstallerType     string          `yaml:"InstallerType,omitempty"`
	Commands          []string        `yaml:"Commands,omitempty"`
	ReleaseDate       string          `yaml:"ReleaseDate,omitempty"`
	Installers        []InstallerItem `yaml:"Installers,omitempty"`
	ManifestType      string          `yaml:"ManifestType,omitempty"`
	ManifestVersion   string          `yaml:"ManifestVersion,omitempty"`
	Dependencies      Dependencies    `yaml:"Dependencies,omitempty"`
}

//nolint:tagliatelle
type Dependencies struct {
	PackageDependencies []PackageDependency `yaml:"PackageDependencies,omitempty"`
}

//nolint:tagliatelle
type PackageDependency struct {
	PackageIdentifier string `yaml:"PackageIdentifier"`
	MinimumVersion    string `yaml:"MinimumVersion,omitempty"`
}

//nolint:tagliatelle
type Locale struct {
	PackageIdentifier   string   `yaml:"PackageIdentifier,omitempty"`
	PackageVersion      string   `yaml:"PackageVersion,omitempty"`
	PackageLocale       string   `yaml:"PackageLocale,omitempty"`
	Publisher           string   `yaml:"Publisher,omitempty"`
	PublisherURL        string   `yaml:"PublisherUrl,omitempty"`
	PublisherSupportURL string   `yaml:"PublisherSupportUrl,omitempty"`
	Author              string   `yaml:"Author,omitempty"`
	PackageName         string   `yaml:"PackageName,omitempty"`
	PackageURL          string   `yaml:"PackageUrl,omitempty"`
	License             string   `yaml:"License,omitempty"`
	LicenseURL          string   `yaml:"LicenseUrl,omitempty"`
	Copyright           string   `yaml:"Copyright,omitempty"`
	CopyrightURL        string   `yaml:"CopyrightUrl,omitempty"`
	ShortDescription    string   `yaml:"ShortDescription,omitempty"`
	Description         string   `yaml:"Description,omitempty"`
	Moniker             string   `yaml:"Moniker,omitempty"`
	Tags                []string `yaml:"Tags,omitempty"`
	ReleaseNotes        string   `yaml:"ReleaseNotes,omitempty"`
	ReleaseNotesURL     string   `yaml:"ReleaseNotesUrl,omitempty"`
	ManifestType        string   `yaml:"ManifestType,omitempty"`
	ManifestVersion     string   `yaml:"ManifestVersion,omitempty"`
}

var fromGoArch = map[string]string{
	"amd64": "x64",
	"386":   "x86",
	"arm64": "arm64",
}
