# Pelith Assessment

## Introduction

Pelith Assessment is a backend system designed to address specific engineering challenges. It focuses on building a
campaign management platform for Uniswap users, enabling them to complete tasks, earn points, and redeem rewards.

## Project Goals

- **Task System**: Implement a configurable multitask system, including onboarding and share pool tasks.
- **Campaign Modes**: Support real-time event handling and backtesting mode for analyzing historical data.
- **Data Integrity**: Ensure accuracy in task tracking and point distribution.
- **Scalability**: Allow dynamic addition of new tasks and trading pairs to meet diverse business requirements.

## Features

1. **Task System**
    - **Onboarding Task**: Users receive 100 points by completing a transaction of at least 1000u.
    - **Share Pool Task**: Distribute 10,000 points among users based on their transaction volume proportion in a
      specific pool.

2. **Data Modes**
    - **Real-Time Mode**: Process live data for immediate task updates.
    - **Backtest Mode**: Simulate past campaigns using historical data.

3. **APIs**
    - Fetch user task completion status.
    - Retrieve task point distribution history.

4. **Extensibility**
    - Dynamically add new trading pair tasks.
    - Provide WebSocket support for real-time updates.

## Technical Overview

- **Architecture**: Built with Go, PostgreSQL for persistent storage, and Redis for caching if needed.
- **Data Storage**: Campaign parameters and user data are stored in PostgreSQL, ensuring ACID compliance.
- **Implementation Details**:
    - **Computation Module**: Calculates user transaction volume proportions and point distribution based on Uniswap V2
      data.
    - **Campaign Status Management**: Configure campaign start time via environment variables or the database.
- **Testing**: Focus on testing core logic with a minimum of 50% unit test coverage.

## Future Enhancements

- Add a leaderboard API for user rankings based on points.
- Introduce CI/CD pipelines for automated testing and deployment.
- Optimize system performance for large-scale data processing.

## How to Use

### Prerequisites

1. **Docker** and **Docker Compose** installed.
2. **Go** environment set up (if required).

### Running the Application

#### Using Makefile

1. **Build the Application**:
   ```bash
   make build
   ```
   Compiles the application and outputs the binary to the `build` directory.

2. **Run the Application**:
   ```bash
   make run
   ```
   Starts the compiled application.

3. **Test the Application**:
   ```bash
   make test
   ```
   Runs all test cases.

4. **Lint the Code**:
   ```bash
   make lint
   ```
   Checks for linting issues using `golangci-lint`.

5. **Clean Up Build Artifacts**:
   ```bash
   make clean
   ```
   Removes all files in the `build` directory.

6. **Help**:
   To see all available Makefile commands:
   ```bash
   make help
   ```

## Documentation

For more detailed information, please refer to the documentation in the `docs/` directory:

- [System Architecture](docs/architecture.md): Explanation of the system's architecture and components.
- [Project Layout](docs/project-layout.md): Overview of the project's directory structure and organization.

---

Explore the `docs/` folder for additional details.

## References

- [Uniswap V2 Documentation](https://docs.uniswap.org/contracts/v2/reference/smart-contracts/pair)
- [Etherscan Pool Events](https://etherscan.io/address/0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc#events)

## Author

[blackhorseya](https://github.com/blackhorseya)

---

Feel free to raise issues or provide suggestions!
