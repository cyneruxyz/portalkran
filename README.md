# portalkran
### Project Layout Generator for Go

Project Layout Generator is a simple yet powerful tool designed to help Go developers quickly set up the standard folder structure for their new projects. This utility automatically creates directories and files that are commonly used in well-organized Go applications. It follows community best practices for structuring Go projects to enhance maintainability and scalability.

#### Features
- Automated Generation: Generate the basic Go project structure with a single command.
- Customizable Templates: Modify existing templates or create new ones to suit your specific project needs.
- Community Standards: Adheres to widely accepted project layout conventions in the Go community.

#### Generated Directory Structure Includes:
- cmd/: Main applications for this project.
- pkg/: Library code that can be used by external applications.
- internal/: Private application and library code.
- api/: OpenAPI/Swagger specs, JSON schema files, protocol definition files.
- configs/: Configuration file templates or default configs.
- deploy/: system and container orchestration deployment configurations and templates.
- scripts/: Scripts to perform build, install, analysis, etc.
- build/: Packaging and Continuous Integration.
- test/: Additional external test apps and test data.
- docs/: Project documentation.

#### How to Use
1. Install the utility via go get or download the binary from the release page.
2. Navigate to your project directory.
3. Run the command `portalkran init` to generate the project structure.
4. Modify as needed based on the specific needs of your project.
5. Start developing your Go application with a well-structured project layout.

This tool is aimed to free up your time from setting up the project environment, allowing you to focus more on coding and less on organizing. Whether you're starting a new project or re-organizing an existing one, Project Layout Generator provides a solid starting point that adheres to industry standards.
