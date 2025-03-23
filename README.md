# Naked Time

## Licensing

Before using the source code for Naked Time, please review the [software license agreement](LICENSE.md). Naked Time is a commercial software product and is not open source. Please review the license agreement and [FAQ](LICENSE.md#FAQ) before proceeding.

## Press Release

### Naked Software, LLC, Announces the Launch of Naked Time: A Revolutionary Time Management Product

__TBD &mdash; [Surprise, AZ]__ &mdash; Naked Software, LLC, a new startup dedicated to creating innovative software solutions, is excited to announce the general availability of its first product, Naked Time. This groundbreaking time management tool is designed to help users plan their daily schedules, track time spent on various activities, and review their time allocation to make meaningful improvements in their productivity.

Naked Time offers a comprehensive suite of features that empower users to take control of their time:

- __Daily Schedule Planning__: Users can easily plan their day, ensuring that important tasks and meetings are accounted for.
- __Time Tracking__: Track how much time is spent in meetings or working on critical activities, providing valuable insights into daily productivity.
- __Review and Adjust__: Review how time was spent to make adjustments and improvements in future work processes and time allocation.

One of the standout features of Naked Time is its implementation of the Pomodoro Technique. This popular time management method encourages users to work in focused intervals of 25 minutes, known as Pomodoros, followed by a 5-minute break. This approach helps prevent burnout and maintains high levels of productivity throughout the day.

"Naked Time is designed to help users make the most of their time by provising a clear and structured approach to daily planning and time tracking," said Michael F. Collins, III, Founder and CEO of Naked Software, LLC. "We believe that by using the Pomodoro Technique, users can achieve a better work-life balance and improve their overall productivity."

Naked Time is not available for download and use. For more information about Naked Time and to start your journey towards better time management, visit https://www.nakedtime.app."

### About Naked Software, LLC

Naked Software, LLC is a software company focused on developing innovative software applications and services. Founded by Michael F. Collins, III, the company aims to create products to help users improve their productivity and achieve their goals. Naked Software is committed to providing high-quality software solutions that are both effective and user-friendly.

For media inquiries, please contact:

Michael F. Collins, III\
Founder & CEO\
Naked Software, LLC\
[michael.collins@naked.software](mailto:michael.collins@naked.software)\
623-297-5498\
https://www.naked.software

## About Naked Time

Naked Time is a time management and productivity product. Naked Time helps users to plan, track, and evaluate their use of time and find ways to improve how they use time to achieve valuable outcomes. Naked Time brings multiple time planning and management tools to users. Naked Time implements [the Pomodoro Technique](https://en.wikipedia.org/wiki/Pomodoro_Technique) to help users to hyperfocus on a single task at a time, while introducing frequent breaks to prevent burnout. Naked Time helps users to plan their days and weeks by aggregating events from multiple calendars, prioritizing meetings and work tasks, and helping the user to visualize their complete schedule to ensure that they are allocating sufficient time for important, value-driving activities. Naked Time helps users to assess the effectiveness of how they spend their time by presenting progress reports and allowing users to do individual retrospectives on how they can better utilize the time that they have available to achieve the goals that are important to them.

## Getting Started

Naked Time has dependencies on third-party software that must be installed in your development environment. Before attempting to work with the source code, please review the [software requirements](docs/software_requirements.md) and ensure that your development environment is properly configured.

After setting up your development environment, you are now able to clone the [Naked Time Git repository](https://github.com/nakedsoftware/time) into your development environment. To do this, open a terminal (Apple macOS or Linux) or Command Prompt window (Microsoft Windows), navigate to the location in your file system where you want to store the Naked Time source code, and run:

    gh repo clone nakedsoftware/time

If you are working with a personal or company fork of the repository, be sure to substitute `nakedsoftware/time` for the name of the forked repository that you are cloning from.

After cloning the repository, you need to prepare the repository for local development. Naked Time requires third-party development tools and libraries be installed before the source code will build correctly. The steps required to prepare the repository for local development have been automated as the [`setup.sh`](setup.sh) (Apple macOS or Linux) or [`Setup.bat`](Setup.bat) (Microsoft Windows) scripts in the root directory of the repository.

- __Apple macOS or Linux__: in the same terminal that you used to clone the repository, run:

```shell
cd time
./setup.sh
```

- __Microsoft Windows__: in the same Command Prompt window, run:

```batch
cd time
Setup.bat
```

After `setup.sh` (or `Setup.bat`) have completed successfully, your repository is ready to build, run, and develop for Naked Time.
