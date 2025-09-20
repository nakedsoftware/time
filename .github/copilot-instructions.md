# GitHub Copilot Instructions

## About Naked Time

Naked Time is a personal productivity tool designed to help users plan and manage their time effectively. Naked Time is based around the [Pomodoro Technique](https://en.wikipedia.org/wiki/Pomodoro_Technique). The main feature of the Pomodoro Technique is breaking work into intervals, called pomodoros. Each pomodoro is 25 minutes of focused work, followed by a 5-minute break. After four pomodoros, users take a longer break of 15-30 minutes. Once a pomodoro is started, the pomodoro must run to completion. Pomodoros are not divisible.

Naked Time and the Pomodoro Technique can help users to plan their day and track their progress. The first pomodoro of the day is used for planning the day. Naked Time maintains a list of activities that the user wants to accomplish in the future. This list of activities is called the Activity Inventory. At the start of each day, the user selects activities from the Activity Inventory to work on that day. The activities are added to the day plan and then the user can prioritize the activities. After the planning pomodoro is complete, the user can start iterating on the first activity by performing one or more pomodoros while focusing on that activity. When the activity is complete, the user can mark the activity as completed and move onto the next activity in the day plan. If the user finishes all the activities in the day plan, they can select more activities from the Activity Inventory to work on.

The last pomodoro of the day is used for recording and reflection. During this pomodoro, the user can record what they accomplished during the day and reflect on what went well and what could be improved. This helps the user to learn from their experiences and improve their productivity over time. The user can review their Activity Inventory and add any new activities that they want to work on in the future.

## Technology Stack

Naked Time is built as a text-based application running in the terminal. Naked Time will expose multiple commands that the user can run to interact with the application. Some commands perform an action using command-line flags and positional arguments to provide input data and output a response. Other commands will open an interactive text-based user interface (TUI) that the user can interact with using their keyboard and mouse. The TUI will provide a more user-friendly way to interact with the application and perform complex tasks.

- Naked Time should run on Windows, macOS, and Linux.
- Naked Time is built using the [Go programming language](https://go.dev/). 
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
    - The Docker container should be multi-architecture and support amd64, arm64.
- Naked Time is also distributed as a standalone binary that can be downloaded and run on Windows, macOS, and Linux.
- Naked Time is distributed via GitHub Releases.
- Naked Time is also distributed via Homebrew for macOS and Linux users.
- Naked Time is also distributed via Windows Package Manager (winget) for Windows users.
- Naked Time is also distributed as a `.deb` package for Debian-based Linux distributions and as an `rpm` package for Red Hat-based Linux distributions.

## Project Structure

- `cmd/`: Contains the main entry point for executable programs. Each executable program should be put in its own subdirectory with a `main.go` file. The subdirectory name will be the executable name when the program is built.
- `internal/`: Contains the core application code. This code is not intended to be used by other applications or libraries.
- `pkg/`: Contains code that is intended to be used by other applications or libraries.

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
