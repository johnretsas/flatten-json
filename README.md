
# JSON Flattening Tool

## Overview

This tool is designed to flatten nested JSON structures into a simple key-value pair format. The flattened data can be easily exported to a CSV file, making it suitable for analysis and reporting.

## Features

- Recursively flattens nested JSON objects and arrays.
- Maintains the order of keys as they are encountered.
- Outputs the flattened data in CSV format.
- Simple and easy-to-use command-line interface.

## Requirements

- Go (1.14 or later)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/json-flattener.git
   cd json-flattener
   ```

2. Ensure you have Go installed. If you don't, you can download it from [golang.org](https://golang.org/dl/).

3. Build the application:
   ```bash
   go build
   ```

## Usage

To run the JSON flattening tool, use the following command:

```bash
go run main.go
```

### Input

1. Create a JSON file (e.g., `data.json`) in the root directory of the project with the desired nested JSON structure. Here’s an example:

   ```json
   {
     "project": {
       "name": "Project A",
       "tasks": [
         {
           "title": "Task 1",
           "status": "in progress"
         },
         {
           "title": "Task 2",
           "status": "completed"
         }
       ],
       "teamMembers": ["Alice", "Bob"]
     }
   }
   ```

### Output

After running the tool, it will create a CSV file named `flattened_data.csv` in the root directory, containing the flattened data.

## Example Output

Given the example JSON above, the flattened CSV will look like:

| project.name  | project.tasks[0].title | project.tasks[0].status | project.tasks[1].title | project.tasks[1].status | project.teamMembers[0] | project.teamMembers[1] |
|----------------|-------------------------|--------------------------|-------------------------|--------------------------|------------------------|------------------------|
| Project A      | Task 1                 | in progress              | Task 2                 | completed                | Alice                  | Bob                    |

## Code Structure

- `main.go`: The main application file that contains the logic for reading JSON, flattening it, and writing to a CSV file.
- `data.json`: Sample input file containing the nested JSON structure.

## How It Works

1. **Reading JSON**: The tool reads the JSON data from a specified file.
2. **Flattening**: It processes the JSON recursively, maintaining the order of keys and converting nested structures into dot-separated keys.
3. **CSV Export**: The flattened data is then written to a CSV file, with the keys as the header and their corresponding values in the subsequent rows.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, feel free to reach out to [your-email@example.com](mailto:your-email@example.com).

---

Feel free to customize the README to better suit your project's specifics, including adding installation instructions if you are using any external libraries or frameworks.