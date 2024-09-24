# GAIKeep

## Overview

**GAIKeep** (Generated Artificial Intelligence Keeper) is a companion tool designed to enhance your experience with Large Language Models (LLMs). It aggregates requests, listens for your cue, and keeps the important responses you tag in a local file or database.

### Features

- **Response Aggregation:** Automatically captures and stores responses from various LLM tools.
- **Tagging System:** Mark responses with custom tags for easy retrieval.
- **AI Hallucination Detection:** A future feature to compare responses across different tools to detect and alert on AI-generated inaccuracies.

## Getting Started

### Prerequisites

- Git
- Go 1.15 or higher

### Installation

1. **Clone the repository**
   ```
   git clone https://github.com/ChristiMahu/gaikeep.git
   cd gaikeep
   ```

2. **Run Setup**
   This will install all necessary dependencies and prepare the environment.
   ```
   make setup
   ```

3. **Build the application**
   ```
   make
   ```

4. **Run GAIKeep**
   ```
   ./bin/gaikeep
   ```

## Usage

Use GAIKeep to interact with supported LLMs like ChatGPT, Gemini, and MidJourney. Here’s how to invoke a command:
```
./bin/gaikeep testgpt
```
This command will send a predefined prompt to ChatGPT and save the response in the `output` directory.

## Contributing

We welcome contributions from the community! If you're interested in helping make GAIKeep better, feel free to fork the repository and submit pull requests. You can also suggest features, report bugs, or share your user experience for improvement.

## Support

If you encounter any issues or have questions, please file a bug report here: [GitHub Issues](https://github.com/ChristiMahu/gaikeep/issues).

## License

GAIKeep is licensed under the GNU General Public License Version 3 (GPLv3). See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- ChatGPT for generative text capabilities.
- Gemini for versatile AI interactions.
- MidJourney for creative image generation.

## About the Project

GAIKeep is developed and maintained by Christi Mahu, a product leadership professional with a focus on generative AI and Large Language Models (LLMs). 
