# Copilot / AI Assistant Instructions for euler-go

Short, practical guidance so an AI coding agent can be productive immediately.

**Repository Layout:**
- `problemNNN/` — each Problem Euler solution lives in its own directory (e.g. `problem001/`).
- `common/` — shared helper utilities (e.g. `DigitsFromInt`, `GCD`, `PrimeFactors`).
- `run_solutions.sh` — convenience script that iterates `problem*/` and runs `go run solution.go`.
- `go.mod` — module root; external deps include `github.com/deckarep/golang-set/v2` and `github.com/etnz/permute`.

**Key patterns and conventions (do not change lightly):**
- Problem folders are named `problemNNN` and typically contain `solution.go` and optionally `solution_test.go` and data files (e.g. `0022_names.txt`).
- `solution.go` files use `package main` and expose a `main()` that prints `Problem N: <answer>`.
- Many problems define a `solve()` function that returns the computed value. Tests in the same folder call `solve()` directly. Expect function signatures to vary (some take args).
- Shared helpers are in `common` and imported with `github.com/ghonzo/euler-go/common`.

**Build / test / run workflows:**
- Fetch deps: `go mod tidy` (run from repo root).
- Run a single problem: `cd problem054 && go run solution.go` (or `go run ./problem054` from repo root).
- Run all solution programs with timing: `./run_solutions.sh` (expects each `problem*/solution.go` to be runnable).
- Run tests for a single folder: `go test ./problem054`.
- Run all tests: `go test ./...`.

**What to look for when editing/adding problems:**
- Keep `package main` for problem folders so `go run` and tests behave consistently.
- If adding helper functions put them under `common/` and import them using the module path. Reuse existing utilities (`DigitsFromInt`, `ProperDivisors`, `IsPrime`) rather than duplicating.
- If a problem uses a data file, place it inside the same `problemNNN/` folder and reference it with a relative path (e.g. `open("0054_poker.txt")`).

**Tests style / expectations:**
- Tests live alongside solutions in `solution_test.go` and use `package main` so they can call `solve()` directly.
- When adding tests, follow existing table-driven test style (see `problem001/solution_test.go`).

**External libraries & their uses:**
- `github.com/deckarep/golang-set/v2` — used for set operations (e.g. digit uniqueness checks).
- `github.com/etnz/permute` — used for permutation generation in permutation-heavy problems.

**Repository-specific gotchas & heuristics for an AI:**
- Some problems rely on small local data assets in the same folder (check for `*.txt` in each `problemNNN/`).
- Solutions were developed iteratively; occasionally a folder may contain a copied solution that doesn't match the problem number (e.g. copied from another problem). Verify correctness by reading the problem text or running tests.
- Function names and signatures vary — search for `solve(` in a folder to learn the expected signature for tests.

**Quick checklist for PRs or edits:**
- Run `go mod tidy` and `go test ./...` before proposing changes.
- Keep `solution.go` runnable with `go run` (no hidden input prompts). If you add CLI flags, ensure they don't break `run_solutions.sh`.
- Add/modify tests in `solution_test.go` using the existing table-driven pattern.

If anything here is unclear or you want more examples from specific problems (e.g. file IO in `problem022` or permutations in `problem024`), tell me which area and I'll expand the instructions.
