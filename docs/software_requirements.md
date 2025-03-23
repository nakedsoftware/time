# Software Requirements

Naked Time makes use of several third-party products, development tools, or libraries that must be installed in the development environment in order to clone the [Git repository](https://github.com/nakedsoftware/time) and work with the source code. Please review the following list of software packages and ensure that they are installed in your development environment before attempting to clone and work with the Naked Time repository.

Note that not all software packages are required for every platform. Platforms where each package is required are noted for each software package below.

1. [Homebrew](#homebrew)
1. [Git](#homebrew)
1. [GitHub CLI](#homebrew)

## Homebrew

- __Apple macOS__
- __Linux__

[Homebrew](https://brew.sh) is a package manager for Apple macOS and Linux developers. Homebrew is used to install many popular open source and commercial software products, development tools, and development libraries. The Naked Time team uses Homebrew to install third-party software packages and keep them up-to-date whenever possible.

To install Homebrew, open a terminal and run:

    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

After Homebrew is successfully installed, you may need to close your terminal and start a new terminal before the environmental changes for Homebrew take effect.

## Git

[Git](https://git-scm.com) is a distributed version control management solution for software developers and teams. Git differs from traditional version control systems in that instead of having a central server, each developer maintains their own complete copy of the version control repository. Developers can work on their own repository while online or offline, and can easily share their changes with other developers either directly or using a shared Git repository such as a repository hosted on [GitHub](https://github.com).

The Naked Time team uses Git for managing all of the source code for the Naked Time product. The Naked Time team uses Git's lightweight branches to provide a stable environment for the development of new features, fixing application defects, and preparing new product releases.

- __Apple macOS or Linux__: Git can be installed using [Homebrew](#homebrew). In a terminal, run:

```shell
brew install git
```

- __Microsoft Windows__: Git can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install --id Git.Git -e --source winget
```

## GitHub CLI

[GitHub CLI](https://cli.github.com) is a command line interface for [GitHub](https://www.github.com). GitHub CLI can be used for cloning Git repositories from GitHub, starting [GitHub Actions](https://github.com/features/actions) workflows, or automating tasks that operate on repositories or projects hosted in GitHub. The Naked Time team uses GitHub ClI for cloning the Naked Time project and for automating activities using GitHub repositories and projects.

- __Apple macOS or Linux__: GitHub CLI can be installed using [Homebrew](#homebrew). In a terminal, run:

```shell
brew install dh
```

- __Microsoft Windows__: GitHub CLI can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install --id GitHub.cli
```
