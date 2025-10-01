---
mode: agent
model: Claude Sonnet 4.5 (Preview) (copilot)
---
Create a draft GitHub release for the Naked Time product. The workflow should
be triggered automatically when a release branch is created. A release branch
will be named in the format `release/major.minor.path`, for example
`release/1.2.3`. The workflow should generate the draft release and keep it
up to date with the latest changes in the release branch. The workflow should
let GitHub automatically generate and update the release notes based on the
commits in the release branch. When the release branch is created, the workflow 
should build and include the executables for Windows, macOS, and Linux as assets
in the draft release.