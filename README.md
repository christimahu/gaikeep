# Gai Keep

![Gai Keep](web/media/gaikeep-avatar.jpg)

Gai Keep (Generated Artificial Intelligence Keeper) is a tool designed to save and organize content from Large Language Models (LLMs).

## Table of Contents

- [Project Status](#project-status)
- [Vision](#vision)
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Project Status

⚠️ **IMPORTANT**: This repository is currently in the early planning and scaffolding phase. It contains only a skeleton structure and does not yet implement the planned functionality. This is not an MVP or even a PoC - it's a foundation for future development.

## Vision

![Gai Keep at Desk](web/media/gaikeep-desk.jpg)

Gai Keep aims to be a companion tool for working with Large Language Models (LLMs). The planned functionality includes:

- Saving AI-generated content as timestamped JSON files
- Organizing content with tags
- Analyzing saved content to discover patterns and relationships
- Providing a simple interface to browse and search your AI knowledge base

## Technology Stack

- **Backend**: Go with gRPC
- **Storage**: JSON files
- **Frontend**: (Planned) Svelte
- **Future Machine Learning**: neural network-based natural language processing (NLP) in Go
- **Documentation**: pkgsite for Go documentation

## Getting Started

### Prerequisites

- **Go** 1.22 or higher
- **Protocol Buffer compiler** (protoc)
- **Protocol Buffer plugins for Go**
- **Make**

See [developer.md](docs/developer.md) for detailed installation instructions for each prerequisite.

### Quick Start

```bash
# Clone the repository
git clone https://github.com/ChristiMahu/gaikeep.git
cd gaikeep

# Install dependencies and build
make setup

# Run the server
make run

# In another terminal, test the API
make demo "This is a test entry"
```

For more detailed setup instructions, see [developer.md](docs/developer.md).

## Documentation

Gai Keep includes comprehensive documentation in the `docs/` directory:

- [Architecture Overview](docs/architecture.md) - System design and component interactions
- [Developer Guide](docs/developer.md) - Development environment setup and coding standards
- [Workflow Guide](docs/workflow.md) - Build system and development workflow
- [Testing Strategy](docs/testing.md) - Test organization and best practices

To view the interactive API documentation:

```bash
make docs
```

This starts a local documentation server at http://localhost:8080.

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

This project is governed by the [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Contact

For questions or suggestions, please open an issue on GitHub or contact the maintainer through the contact form at [https://christimahu.com/contact](https://christimahu.com/contact).
