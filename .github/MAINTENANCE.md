# Maintenance Guide

This guide describes how maintainers keep the list and automation healthy.

## Automated Workflows

### Pull Requests / Weekly QA (`pull_request.yml`)

- Runs on pull requests and weekly on Saturday.
- Builds the Go CLI and runs `./awesome-docker validate`.

### Broken Links Report (`broken_links.yml`)

- Runs weekly on Saturday and on manual trigger.
- Executes `./awesome-docker check`.
- Opens/updates a `broken-links` issue when problems are found.

### Weekly Health Report (`health_report.yml`)

- Runs weekly on Monday and on manual trigger.
- Executes `./awesome-docker health` then `./awesome-docker report`.
- Opens/updates a `health-report` issue.

### Deploy to GitHub Pages (`deploy-pages.yml`)

- Runs on pushes to `master` and manual trigger.
- Builds website with `./awesome-docker build` and publishes `website/`.

## Day-to-Day Commands

```bash
# Build CLI
make build

# README lint/validation
make lint

# Auto-fix formatting issues
./awesome-docker lint --fix

# Link checks and health checks (requires GITHUB_TOKEN)
make check
make health
make report
```

## Content Maintenance Policy

- Remove archived/deprecated projects instead of tagging them.
- Remove broken links that cannot be fixed.
- Keep sections alphabetically sorted.
- Keep descriptions short and actionable.

## Suggested Review Cadence

### Weekly

- Triage open `broken-links` and `health-report` issues.
- Merge straightforward quality PRs.

### Monthly

- Review sections for stale/duplicate entries.
- Re-run `check` and `health` manually if needed.

### Quarterly

- Review `.github` docs and templates for drift.
- Confirm workflows still match repository tooling and policies.

## Contributor Support

When requesting PR changes, be explicit and actionable:

- point to section/order problems,
- explain why a link should be removed,
- suggest exact wording when description quality is the issue.

---

Last updated: 2026-02-27
