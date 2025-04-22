# Naked Time

Naked Time is a time management productivity application for use by individuals, teams, businesses, and enterprises. Naked Time provides tools to help users to:

1. Plan and track activities that they need to perform in their job (or life).
2. Prioritize the activities to ensure that the most important activities are completed first.
3. Determine whether to handle activities themselves or delegate the activities to others.
4. Prioritize and schedule work time during work hours over meetings and other non-value-adding activities.
5. Hyperfocus on the work that they need to perform when working on activities.
6. Review, analyze, and improve their use of time.

## Getting Started

Before cloning the Naked Time repository, please review the [software requirements](docs/software_requirements.md) to ensure that your development environment is configured correctly and the necessary tools are installed. When your development environment is ready, you can begin by cloning the Git repository locally on your computer from GitHub. Open a terminal (Apple macOS or Linux) or Command Prompt window (Microsoft Windows) and run:

    gh repo clone nakedsoftware/time

Before you can begin working with the Naked Time source code, you need to prepare your local repository for development. Naked Time requires the installation and configuration of external tools and libraries that are managed through package and dependency managers. In addition, Naked Time does make use of generated source code and project files for certain components. The steps required to prepare your local environment for development have been automated.

- On Apple macOS or Linux, in the terminal run:

```shell
cd time
./go setup
```

- On Microsoft Windows, in the Command Prompt window, run:

```batch
cd time
go.bat setup
```

Once the command is finished, you are ready to run, develop for, and debug the Naked Time product.