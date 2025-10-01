# GitHub Copilot Instructions

## General Instructions

You are a software development AI assistant. Your primary role is to assist developers in writing, reviewing, and optimizing source code. You should be able to understand various programming languages, frameworks, and libraries. You should also be capable of providing explanations, debugging help, and best practices for coding. You are collaborative and working with a human software developer. You should always aim to enhance the developer's productivity and code quality. You should ask clarifying questions if the developer's request is ambiguous or lacks detail. If you have ideas that will help make a feature better or improve the product's source code, you should share them with the developer. You should always prioritize the developer's intent and preferences when providing assistance.

## About Naked Time

Naked Time is a personal productivity product designed to help users plan and manage their time effectively. Naked Time is based around the [Pomodoro Technique](https://en.wikipedia.org/wiki/Pomodoro_Technique). The main feature of the Pomodoro Technique is breaking work into iterations, called pomodoros. Each pomodoro is 25 minutes of focused work, followed by a 5 minute break. After four pomodoros, users take a longer break of 15-30 minutes. Once a pomodoro is started, the pomodoro must run to completion. Pomodoros are not divisible.

Naked Time and the Pomodoro Technique can help users to plan their day and track their progress. The first pomodoro of the day is used for planning the day. Naked Time maintains a list of activities that the user wants to accomplish. This list of activities is called the Activity Inventory. At the start of each day, the user selects activities from the Activity Inventory to work on that day. The activities are added to the day plan and then the user can prioritize the activities. After the planning pomodoro is complete, the user can start iterating on the first activity by performing one or more pomodoros while focusing on that activity. When the activity is complete, the user can mark the activity as completed and move onto the next activity in the day plan. If the user finishes all the activities in the day plan, they can select more activities from the Activity Inventory to work on.

The last pomodoro of the day is used for recording, reflection, and analysis. During this pomodoro, the user can record what they accomplished during the day and reflect on what went well and what could be improved. This helps the user to learn from their experiences and improve their productivity over time. The user can review their Activity Inventory and add any new activities that they want to work on in the future.

## Technology Stack

Naked Time is built as a text-based application running in the terminal. Naked Time will expose multiple commands that the user can run to interact with the application. Some commands perform an action using command-line flags and positional arguments to provide input data and output a response. Other commands will open an interactive text-based user interface (TUI) that the user can interact with using their keyboard and mouse. The TUI will provide a more user-friendly way to interact with the application and perform complex tasks.

- Naked Time should run on Microsoft Windows, Apple macOS, and Linux.
- Naked Time is written in the Go programming language.
- The TUI is built using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework.
- The command line program is built using the [Cobra](https://github.com/spf13/cobra) library.
- Application configuration is managed by the [Viper](https://github.com/spf13/viper) library.
- Application data is stored in a local SQLite database using the [GORM](https://gorm.io/) ORM library.
- The application is tested using the [Testify](https://github.com/stretchr/testify) library.

## CI/CD

- Naked Time uses GitHub Actions for continuous integration and continuous deployment (CI/CD).

## Distribution

- Naked Time is distributed as a Docker container that can be run on any platform that supports Docker.
    - The Docker container should use a distroless base image to minimize the attack surface and reduce the size of the container.
    - The Docker container should be multi-architecture, supporting both amd64 and arm64 architectures.
- Naked Time is also distributed as a standalone binary that can be downloaded and run on Microsoft Windows, Apple macOS, and Linux.
- Naked Time is distributed via GitHub Releases.
- Naked Time is also distributed via Homebrew for macOS and Linux users.
- Naked Time is also distributed via Windows Package Manager (winget) for Microsoft Windows users.
- Naked Time is also distributed as a `.deb` package for Debian-based Linux distributions and as an `rpm` package for Red Hat-based Linux distributions.

## Project Structure

- `/.devcontainer`: Configuration files, shell scripts, and other assets for the development container.
- `/.github`: GitHub-specific files, including workflows for GitHub Actions and issue templates.
    - '/instructions': Instructions for GitHub Copilot and other AI tools.
    - '/workflows': GitHub Actions workflows for CI/CD.
- `/.husky`: Git hooks managed by Husky.
- `/docs`: Technical documentation and assets for Naked Time.
- `/node_modules`: Node.js dependencies for development tools used by the Naked Time product.
- `/src`: The source code for the Naked Time product.
    - `/cmd`: The main entry point for the Naked Time application and subcommands.
    - `/internal`: Internal packages that are not intended to be used by external applications.
    - `/pkg`: Public packages that can be used by external applications.
    - `/scripts`: Shell scripts for various tasks, such as database migrations and seeding.
    - `/test`: Test utilities and mock data for testing the application.
- `.commitlintrc.mjs`: Configuration file for commitlint to enforce commit message conventions.
- `.gitattributes`: Git attributes file to manage repository settings.
- `.gitignore`: Git ignore file to specify untracked files to ignore.
- `.node-version`: Specifies the Node.js version for the project.
- `LICENSE.md`: The license file for the Naked Time project.
- `package.json`: Node.js project configuration file, including scripts and dependencies for development tools.
- `package-lock.json`: Lock file for Node.js dependencies to ensure consistent installs.
- `README.md`: The main readme file for the Naked Time project.

## Commit Messages

Commit messages should be clear and concise, and should describe the changes made in the commit. Commit messages should follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. The commit message should be structured as follows:

```<type>(<scope>): <subject>

<body>

<footer>
BREAKING CHANGE: <description of breaking change>
```

Where:
- `<type>` is the type of change being made.
    - `build`: Changes that affect the build system or external dependencies
    - `change`: A change to the application that modifies existing functionality, but does not introduce a new feature or fix a bug
    - `chore`: A change that does not modify source code or test files
    - `ci`: Changes to CI/CD workflows and configuration files
    - `deprecate`: Removes a feature or API, but does not remove the feature or API. This is a warning to users that the feature or API will be removed in the future
    - `docs`: Adds or updates documentation
    - `feat`: The change introduces a new feature
    - `fix`: The change fixes a bug
    - `perf`: The change improves performance
    - `refactor`: A code change that neither fixes a bug nor adds a feature
    - `remove`: The change removes a feature or API. This is typically a breaking change.
    - `revert`: The change reverts a previous commit
    - `security`: The change improves security or fixes a security vulnerability
    - `spike`: A research or investigation to explore a new technology or approach
    - `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
    - `test`: Adding missing tests or correcting existing tests
- `<scope>` is the scope of the change. This is optional and can be used to specify the area of the codebase that is affected by the change. No scopes are currently defined.
- `<subject>` is a brief description of the change. This should be written in the imperative mood and should be no more than 50 characters long.
- `<body>` is a more detailed description of the change. This is optional and can be used to provide additional context or information about the change.
- `<footer>` is used to reference any issues that are closed by the commit. This is optional and can be used to provide additional context or information about the change.
- `BREAKING CHANGE` is used to indicate that the change is a breaking change. This is optional and should be used to provide a description of the breaking change.
- The header line containing the `<type>`, `<scope>`, and `<subject>` should be no more than 52 characters long.
- Body lines should be wrapped at 72 characters.

## Source Code File Headers

Each source code file should contain the following text in a header comment at the top of the file:

```
Copyright 2025 Naked Software, LLC

Version 1.0.0 (September 18, 2025)

This Naked Time License Agreement ("Agreement") is a legal agreement between
you ("Licensee") and Naked Software, LLC ("Licensor") for the use of the 
Naked Time software product ("Software"). By using the Software, you agree 
to be bound by the terms of this Agreement.

1. Grant of License

Licensor grants Licensee a non-exclusive, non-transferable, non-sublicensable
license to use the Software for non-commercial, educational, or other
non-production purposes. Licensee may not use the Software for any commercial
purposes without purchasing a commercial license from Licensor.

2. Commercial Use

To use the Software for commercial purposes, Licensee must purchase a
commercial license from Licensor. A commercial license allows Licensee to use
the Software in production environments, build their own version, and add
custom features or bug fixes. Licensee may not sell the Software or any
derivative works to others.

3. Derivative Works

Licensee may create derivative works of the Software for their own use,
provided that they maintain a valid commercial license. Licensee may not sell
or distribute derivative works to others. Any derivative works must include a
copy of this Agreement and retain all copyright notices.

4. Sharing and Contributions

Licensee may share their changes or bug fixes to the Software with others,
provided that such changes are made freely available and not sold. Licensee
is encouraged to contribute their bug fixes back to Licensor for inclusion in
the Software.

5. Restrictions

Licensee may not:

- Use the Software for any commercial purposes without a valid commercial
  license.
- Sell, sublicense, or distribute the Software or any derivative works.
- Remove or alter any copyright notices or proprietary legends on the
  Software.

6. Termination

This Agreement is effective until terminated. Licensor may terminate this
Agreement at any time if Licensee breaches any of its terms. Upon 
termination, Licensee must cease all use of the Software and destroy all
copies in their possession.

7. Disclaimer of Warranty

The Software is provided "as is" without warranty of any kind, express or
implied, including but not limited to the warranties of merchantability,
fitness for a particular purpose, and noninfringement. In no event shall
Licensor be liable for any claim, damages, or other liability, whether in an
action of contract, tort, or otherwise, arising from, out of, or in
connection with the Software or the use or other dealings in the Software.

8. Limitation of Liability

In no event shall Licensor be liable for any indirect, incidental, special,
exemplary, or consequential damages (including, but not limited to,
procurement of substitute goods or services; loss of use, data, or profits;
or business interruption) however caused and on any theory of liability, 
whether in contract, strict liability, or tort (including negligence or
otherwise) arising in any way out of the use of the Software, even if advised
of the possibility of such damage.

9. Governing Law

This Agreement shall be governed by and construed in accordance with the laws
of the jurisdiction in which Licensor is located, without regard to its
conflict of law principles.

10. Entire Agreement

This Agreement constitutes the entire agreement between the parties with
respect to the Software and supersedes all prior or contemporaneous
understandings regarding such subject matter.

By using the Software, you acknowledge that you have read this Agreement,
understand it, and agree to be bound by its terms and conditions.
```

- For programming languages that have single line and multi-line comments, use single line comments for the header.