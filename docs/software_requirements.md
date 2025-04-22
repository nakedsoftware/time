# Software Requirements

Naked Time has requirements for specific software products or packages to be installed and configured in your development environment (local machine or virtual machine) to build, run, and develop for Naked Time. Please review the list below and ensure that all software packages are installed and configured correctly.

Note that not all software packages are required for every platform (Apple macOS, Linux, or Microsoft Windows). Each software package indicates which platform the package is required for.

1. [Homebrew](#homebrew)
1. [Git](#git)
1. [GitHub CLI](#github-cli)
1. [Docker Desktop](#docker-desktop)
1. [Visual Studio Code](#visual-studio-code)
1. [Remote Development Extension Pack for Visual Studio Code](#remote-development-extension-pack-for-visual-studio-code)
1. [Fast Node Manager](#fast-node-manager)

## Homebrew

- :white_check_mark: Apple macOS
- :white_check_mark: Linux
- :x: Microsoft Windows

[Homebrew](https://brew.sh) is a package manager for Apple macOS and Linux. Homebrew can be used to install and manage many popular open source or commercial software packages and libraries. The Naked Time team uses Homebrew whenever possible to install and manage external dependencies that need to be managed outside of our repository.

To install Homebrew, open a terminal and run:

    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

After Homebrew is installed, you will have to stop and restart your terminal for the environmental changes to take affect.

## Git

- :white_check_mark: Apple macOS
- :white_check_mark: Linux
- :white_check_mark: Microsoft Windows

[Git](https://git-scm.com) is a distributed version control system. Git is used to manage source code for software projects, facilitate software development workflows, and track changes that are made to the product source code over time. Git works differently than traditional version control systems in that each developer maintains their own local copy of the repository. No central server is needed. This model allows software developers to continue to work and make changes to source code whether online or offline. Git supports sharing changes by developers either directly with other developers or using a central Git repository in a service such as [GitHub](https://github.com).

- __Apple macOS or Linux__: Git can be installed using [Homebrew](#homebrew). In a terminal, run:

```shell
brew install git
```

- __Microsoft Windows__: Git can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install --id Git.Git -e --source winget
```

## GitHub CLI

- :white_check_mark: Apple macOS
- :white_check_mark: Linux
- :white_check_mark: Microsoft Windows

[GitHub CLI](https://cli.github.com) is a command line interface for [GitHub](https://github.com). GitHub CLI supports performing one-off commands or operations on Git repositories and projects, or automating complex tasks.

- __Apple macOS or Linux__: GitHub CLI can be installed using [Homebrew](#homebrew). In a terminal, run:

```shell
brew install gh
```

- __Microsoft Windows_: GitHub CLI can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install --id GitHub.cli
```

## Docker Desktop

- :white_check_mark: Apple macOS
- :x: Linux
- :white_check_mark: Microsoft Windows

[Docker Desktop](https://www.docker.com/products/docker-desktop/) supports the development of container images and execution of containers for Apple macOS and Microsoft Windows. Docker Desktop is used to run the Naked Time development environment within a container and for running services that are needed for development and execution of the Naked Time product.

Docker Desktop can be downloaded from the [website](https://www.docker.com/products/docker-desktop/).

## Visual Studio Code

- :white_check_mark: Apple macOS
- :white_check_mark: Linux
- :white_check_mark: Microsoft Windows

[Visual Studio Code](https://code.visualstudio.com) is a free open-source text editor and development environment created by [Microsoft](https://microsoft.com). Visual Studio Code is extremely extensible and has a large community-supported marketplace of available extensions for programming languages and supports integrating other development tools into the Visual Studio Code editor. The Naked Time team uses Visual Studio Code for development and also for executing the development environment in a development container using [Docker Desktop](#docker-desktop).

## Remote Development Extension Pack for Visual Studio Code

- :white_check_mark: Apple macOS
- :white_check_mark: Linux
- :white_check_mark: Microsoft Windows

The [Remote Development Extension Pack for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) is a collection of extensions for [Visual Studio Code](#visual-studio-code) that support developing using containers or connecting to remote developer machines. Naked Time uses the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) included in the extension pack to run and connect to development containers running on your local machine.

The Remote Development Extension Pack can be installed from the [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack).

# Fast Node Manager

- :white_check_mark: Apple macOS
- :white_check_mark: Linux
- :white_check_mark: Microsoft Windows

[Fast Node Manager](https://github.com/Schniz/fnm) is a version management tool for [Node.js](https://nodejs.org). Fast Node Manager can install and switch between different versions of Node.js for different projects. The version of Node.js supported by Naked Time is stored in the [`.node-version`](../.node-version) file in the root directory of the repository. When the developer switches into the Naked Time repository, Fast Node Manager will automatically select the supported version of Node.js to use with Naked Time.

- __Apple macOS or Linux__: Fast Node Manager can be installed using [Homebrew](#homebrew). In a terminal, run:

```bash
curl -fsSL https://fnm.vercel.app/install | bash
```

This command will download and run an installation script that will automatically use Homebrew to install Fast Node Manager.

- __Microsoft Windows__: Fast Node Manager can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install Schniz.fnm
```
