# Agent Guidelines for awesome-docker

## Commands
- Build CLI: `make build` (or `go build -o awesome-docker ./cmd/awesome-docker`)
- Run tests: `make test` (runs `go test ./internal/... -v`)
- Lint README rules: `make lint` (runs `./awesome-docker lint`)
- Auto-fix lint issues: `./awesome-docker lint --fix`
- Check links: `make check` (runs `./awesome-docker check`; `GITHUB_TOKEN` enables GitHub repo checks)
- PR validation: `./awesome-docker validate` (lint + external link checks in PR mode)
- Build website: `./awesome-docker build` (generates `website/index.html` from `README.md`)
- Health scoring: `make health` (requires `GITHUB_TOKEN`, updates `config/health_cache.yaml`)
- Generate health report: `make report`

## Architecture
- **Main content**: `README.md` (curated Docker/container resources)
- **CLI entrypoint**: `cmd/awesome-docker/main.go` (Cobra commands)
- **Core packages**:
  - `internal/parser` - parse README sections and entries
  - `internal/linter` - alphabetical/order/format validation + autofix
  - `internal/checker` - HTTP and GitHub link checks
  - `internal/scorer` - repository health scoring and report generation
  - `internal/cache` - exclude list and health cache read/write
  - `internal/builder` - render README to website HTML from template
- **Config**:
  - `config/exclude.yaml` - known link-check exclusions
  - `config/website.tmpl.html` - HTML template for site generation
  - `config/health_cache.yaml` - persisted health scoring cache
- **Website output**: `website/index.html`

## Code Style
- **Language**: Go
- **Formatting**: Keep code `gofmt`-clean
- **Testing**: Add/adjust table-driven tests in `internal/*_test.go` for behavior changes
- **Error handling**: Return wrapped errors (`fmt.Errorf("context: %w", err)`) from command handlers
- **CLI conventions**: Keep command behavior consistent with existing Cobra commands (`lint`, `check`, `health`, `build`, `report`, `validate`)

## CI/Automation
- **PR + weekly validation**: `.github/workflows/pull_request.yml`
  - Triggers on pull requests to `master` and weekly schedule
  - Builds Go CLI and runs `./awesome-docker validate`
- **Weekly broken links issue**: `.github/workflows/broken_links.yml`
  - Runs `./awesome-docker check`
  - Opens/updates `broken-links` issue when failures are found
- **Weekly health report issue**: `.github/workflows/health_report.yml`
  - Runs `./awesome-docker health` then `./awesome-docker report`
  - Opens/updates `health-report` issue
- **GitHub Pages deploy**: `.github/workflows/deploy-pages.yml`
  - On push to `master`, builds CLI, runs `./awesome-docker build`, deploys `website/`

## Content Guidelines (from CONTRIBUTING.md)
- Use one link per entry
- Prefer project/repository URLs over marketing pages
- Keep entries alphabetically ordered within each section
- Keep descriptions concise and concrete
- Use `:heavy_dollar_sign:` only for paid/commercial services
- Remove archived/deprecated projects instead of tagging them
- Avoid duplicate links and redirect variants
