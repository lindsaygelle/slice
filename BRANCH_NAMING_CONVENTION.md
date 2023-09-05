# Branch Naming Convention

This document outlines the branch naming convention to be followed for our project/repository. Consistent and descriptive branch names help us maintain a clear and organized codebase.

## Types of Branches

### 1. Feature Branches

Feature branches are used to develop new features or enhancements. They should be created from the `main` branch.

Naming Convention: `feature/descriptive-feature-name`

Example: `feature/user-authentication`

### 2. Bugfix Branches

Bugfix branches are used to address and fix bugs or issues in the code. They should be created from the `main` branch.

Naming Convention: `bugfix/descriptive-bug-description`

Example: `bugfix/navigation-bar-overflow`

### 3. Hotfix Branches

Hotfix branches are used to address critical issues that need immediate attention and cannot wait for the next regular release. They should be created from the `main` branch.

Naming Convention: `hotfix/descriptive-hotfix-description`

Example: `hotfix/critical-security-vulnerability`

## Special Cases

### 1. Release Branches

Release branches are used to prepare the codebase for a new release. They are created from the `main` branch when the code is ready for deployment. They might be prefixed with the release version.

Naming Convention: `release/version`

Example: `release/1.0.0`

### 2. Main Branch

The main branch is the primary branch of the repository and represents the latest stable version of the project.

Branch Name: `main`

## General Guidelines

- Use lowercase letters for branch names.
- Use hyphens to separate words in the branch names (e.g., `feature/new-feature`).
- Keep branch names concise and descriptive.
- Avoid using abbreviations or acronyms unless they are widely understood in the context of the project.

## Examples

### Good Examples

- `feature/user-profile`
- `bugfix/login-page-styling`
- `hotfix/404-page-error`

### Bad Examples

- `f/newfeature` (too short and lacks descriptive information)
- `fix` (not descriptive enough)

Remember to always follow this branch naming convention to ensure clarity and consistency in our code repository.

If you have any questions or need further clarification, feel free to reach out to the team.
