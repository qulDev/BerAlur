# Graph Report - C:\DEV\BerAlur  (2026-09-27)

## Corpus Check
- Corpus is ~16,121 words - fits in a single context window. You may not need a graph.

## Summary
- 224 nodes · 216 edges · 35 communities (22 shown, 13 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 21 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- Runtime And Tenant Config
- Web Test Dependencies
- API Bootstrap And Database
- Web App Package
- TypeScript Compiler Settings
- Workspace Build Scripts
- Deployment And API Foundation
- Root Tooling Package
- Graphify Project Tooling
- HTTP Router Contract
- TypeScript Project Scope
- Organization Data Model
- Production Data Safeguards
- Application Layout Metadata
- Homepage Component Tests
- Next App Configuration
- Next Type Definitions
- Initialization Test Plan
- Prettier Formatting Settings
- Database Documentation
- Environment Documentation
- Deferred Infrastructure Scope
- OpenAPI Service Contract
- Tenant Security Guidance
- API Module Workspace
- Monorepo Package Boundaries

## God Nodes (most connected - your core abstractions)
1. `scripts` - 19 edges
2. `compilerOptions` - 16 edges
3. `Graphify Full Pipeline` - 10 edges
4. `New()` - 8 edges
5. `NewRouter()` - 7 edges
6. `scripts` - 7 edges
7. `run()` - 6 edges
8. `include` - 6 edges
9. `Load()` - 5 edges
10. `CI Verification Workflow` - 5 edges

## Surprising Connections (you probably didn't know these)
- `Development Compose Overlay` --shares_data_with--> `Container Stack`  [INFERRED]
  compose.dev.yaml → compose.yaml
- `Local Compose Overlay` --shares_data_with--> `Container Stack`  [INFERRED]
  compose.local.yaml → compose.yaml
- `Project Graph Rules` --references--> `Graphify Full Pipeline`  [EXTRACTED]
  AGENTS.md → .codex/skills/graphify/SKILL.md
- `CI Verification Workflow` --references--> `sqlc Generation`  [EXTRACTED]
  .github/workflows/ci.yml → apps/api/sqlc.yaml
- `CI Verification Workflow` --references--> `Development Compose Overlay`  [EXTRACTED]
  .github/workflows/ci.yml → compose.dev.yaml

## Import Cycles
- None detected.

## Communities (35 total, 13 thin omitted)

### Community 0 - "Runtime And Tenant Config"
Cohesion: 0.10
Nodes (24): DATABASE_URL, Organization identity migration, sqlc and goose migrations, APP_ENV, Compose environment overrides, Environment credential isolation, Verifiable local foundation, Initialization verification boundaries (+16 more)

### Community 1 - "Web Test Dependencies"
Cohesion: 0.09
Nodes (23): devDependencies, eslint, eslint-config-next, jsdom, @testing-library/jest-dom, @testing-library/react, @types/node, @types/react (+15 more)

### Community 2 - "API Bootstrap And Database"
Cohesion: 0.13
Nodes (14): Logger, main(), run(), Queries, New(), T, TestOrganizationLookupOnPostgres(), Load() (+6 more)

### Community 3 - "Web App Package"
Cohesion: 0.11
Nodes (18): dependencies, next, react, react-dom, name, private, scripts, build (+10 more)

### Community 4 - "TypeScript Compiler Settings"
Cohesion: 0.11
Nodes (19): compilerOptions, allowJs, esModuleInterop, incremental, isolatedModules, jsx, lib, module (+11 more)

### Community 5 - "Workspace Build Scripts"
Cohesion: 0.11
Nodes (19): scripts, build, build:api, db:generate, db:migrate, db:up, dev:api, dev:web (+11 more)

### Community 6 - "Deployment And API Foundation"
Cohesion: 0.23
Nodes (13): sqlc Generation, Container Stack, Development Compose Overlay, Local Compose Overlay, Private Production Ports, Production Compose Overlay, API Contract, Membership-Derived Tenant Scope (+5 more)

### Community 7 - "Root Tooling Package"
Cohesion: 0.17
Nodes (11): devDependencies, @playwright/test, prettier, engines, node, name, packageManager, private (+3 more)

### Community 8 - "Graphify Project Tooling"
Cohesion: 0.25
Nodes (11): Project Graph Rules, Folder Watcher, URL Ingestion, Portable Graph Exports, Semantic Extraction Schema, Cross-Repository Graph Merge, Commit Graph Rebuild, Graph Query Traversal (+3 more)

### Community 9 - "HTTP Router Contract"
Cohesion: 0.22
Nodes (8): Context, Logger, NewRouter(), T, TestHTTPContract(), writeError(), Mux, ResponseWriter

### Community 10 - "TypeScript Project Scope"
Cohesion: 0.22
Nodes (8): exclude, include, .next/dev/types/**/*.ts, next-env.d.ts, .next/types/**/*.ts, node_modules, **/*.ts, **/*.tsx

### Community 11 - "Organization Data Model"
Cohesion: 0.25
Nodes (6): UUID, Context, Queries, UUID, Organization, Timestamptz

### Community 12 - "Production Data Safeguards"
Cohesion: 0.33
Nodes (6): Backup and restore safeguards, Forward-fix migration policy, Single deployment migration job, Production deployment sequence, Network exposure and TLS, Production customer-data readiness

## Knowledge Gaps
- **97 isolated node(s):** `trailingComma`, `Queries`, `beralur/apps/api`, `config`, `name` (+92 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **13 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `devDependencies` connect `Web Test Dependencies` to `Web App Package`?**
  _High betweenness centrality (0.026) - this node is a cross-community bridge._
- **Why does `scripts` connect `Workspace Build Scripts` to `Root Tooling Package`?**
  _High betweenness centrality (0.015) - this node is a cross-community bridge._
- **Why does `compilerOptions` connect `TypeScript Compiler Settings` to `TypeScript Project Scope`?**
  _High betweenness centrality (0.012) - this node is a cross-community bridge._
- **Are the 5 inferred relationships involving `New()` (e.g. with `main()` and `run()`) actually correct?**
  _`New()` has 5 INFERRED edges - model-reasoned connections that need verification._
- **What connects `trailingComma`, `Queries`, `beralur/apps/api` to the rest of the system?**
  _97 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Runtime And Tenant Config` be split into smaller, more focused modules?**
  _Cohesion score 0.09782608695652174 - nodes in this community are weakly interconnected._
- **Should `Web Test Dependencies` be split into smaller, more focused modules?**
  _Cohesion score 0.08695652173913043 - nodes in this community are weakly interconnected._