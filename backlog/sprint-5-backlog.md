# Sprint 5 Backlog

**Sprint Goal:** Enhanced sub-issue integration - self-contained split, progress tracking, and cross-repo support

**Sprint Duration:** TBD
**Total Story Points:** 18

---

## Selected Stories

### Story 3.1: Native Sub-Issue Handling in Split (#49)
**Points:** 5 | **Priority:** High

**As a** developer using split command
**I want** split to work without external extensions
**So that** I don't need separate gh-sub-issue installed

**Acceptance Criteria:**
- [ ] `gh pmu split` uses internal sub-issue API code
- [ ] No dependency on gh-sub-issue extension
- [ ] Same functionality as current split + sub-issue combo
- [ ] Maintains backward compatibility

**Technical Notes:**
- Split command already exists and creates issues
- Need to verify it uses internal `AddSubIssue` API, not external extension
- Test that split works without gh-sub-issue installed

---

### Story 3.3: Sub-Issue Progress Tracking (#51)
**Points:** 5 | **Priority:** Medium

**As a** project manager reviewing epics
**I want** to see sub-issue completion percentages
**So that** I can track progress on large work items

**Acceptance Criteria:**
- [ ] `gh pmu view <issue>` shows progress bar for parents
- [ ] Shows "3 of 5 sub-issues complete (60%)"
- [ ] `gh pmu list --has-sub-issues` filters to parent issues
- [ ] Progress based on closed/total sub-issue count

**Technical Notes:**
- Use `subIssuesSummary` field from GraphQL API
- Update view command output format
- Add `--has-sub-issues` filter to list command

---

### Story 3.2: Cross-Repository Sub-Issues (#50)
**Points:** 8 | **Priority:** Medium

**As a** developer with multi-repo projects
**I want** to create sub-issues in different repositories
**So that** I can organize work across my codebase

**Acceptance Criteria:**
- [ ] `gh pmu sub add` works across repositories
- [ ] `gh pmu sub create --repo <owner/repo>` creates in specified repo
- [ ] Parent can be in different repo than child
- [ ] Validates repos are in same project
- [ ] Shows repo info in sub list output

**Technical Notes:**
- GraphQL API supports cross-repo sub-issues natively
- Need to update `sub create` to accept `--repo` flag
- Update `sub list` output to show repository for each sub-issue
- Consider validation: should repos be in same project?

---

## Deferred Stories

### Story 3.4: Recursive Operations on Issue Trees (#52)
**Points:** 8 | **Priority:** Low
**Reason:** Lower priority, more complex, defer to Sprint 6

---

## Sprint Metrics

| Metric | Value |
|--------|-------|
| Total Points | 18 |
| Stories | 3 |
| High Priority | 1 |
| Medium Priority | 2 |

## Definition of Done

- [ ] All acceptance criteria met
- [ ] Unit tests written and passing
- [ ] Integration tests for API calls
- [ ] Code follows Go conventions
- [ ] Command help text documented
- [ ] README updated if needed

## Dependencies

- GitHub API `sub_issues` feature header support (already implemented)
- `subIssuesSummary` field access (may need API update)

## Risks

1. **subIssuesSummary field access** - May require additional API permissions or headers
2. **Cross-repo validation** - Need to determine if/how to validate repos are related
