# Pipeline Parser

[![Test Pipeline Parser](https://github.com/khulnasoft-lab/pipeline-parser/actions/workflows/pr.yml/badge.svg)](https://github.com/khulnasoft-lab/pipeline-parser/actions/workflows/pr.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/khulnasoft-lab/pipeline-parser/blob/main/LICENSE)
[![go-report-card][go-report-card]](https://goreportcard.com/report/github.com/khulnasoft-lab/pipeline-parser)
![coverage report](https://img.shields.io/codecov/c/github/khulnasoft-lab/pipeline-parser)

[go-report-card]: https://goreportcard.com/badge/github.com/khulnasoft-lab/pipeline-parser

## Description

Pipeline Parser is KhulnaSoft's solution for parsing and analyzing pipeline files of popular CI yaml files in order to create a generic pipeline entity that can be used across platforms.

#### Supported Platforms:

| Platform
| :---:
| GitHub Workflows
| GitLab CI
| Azure Pipelines
| Bitbucket Pipelines

## Usage

### Package Usage

```golang
import (
    "os"

    "github.com/khulnasoft-lab/pipeline-parser/pkg/handler"
    "github.com/khulnasoft-lab/pipeline-parser/pkg/consts"
)

// Read the pipeline data as bytes array
buf, err := os.ReadFile("/path/to/workflow.yml")
if err != nil {
    return nil
}

// Parse the pipeline from the specific platform to the common pipeline object
pipeline, err := handler.Handle(buf, consts.GitHubPlatform, scmCredentials, organization, baseProviderUrl)
```

### CLI Usage

#### CLI flags

|      Flag       | Value  |                                       Description                                       | Default  |
| :-------------: | :----: | :-------------------------------------------------------------------------------------: | :------: |
|  platform (-p)  | string |                                  CI platform to parse                                   | `github` |
|   output (-o)   | string |                                      Output target                                      | `stdout` |
|   file-suffix   | string | File suffix for output file. This flag is useless if 'output' flag is not set to 'file' | `parsed` |
|      token      | string |                 SCM token to use for fetching remote files if necessary                 |          |
|  organization   | string |      The target organization when fetching remote files (used for Azure Pipelines)      |          |
| baseProviderUrl | string |       base api url for the pipeline provider (used for parsing remote templates)        |          |

#### Parse GitHub Workflow yaml

```bash
pipeline-parser -p github workflow.yml
```

#### Parse GitLab CI yaml

```bash
pipeline-parser -p gitlab .gitlab-ci.yml
```

#### Parse Azure Pipelines yaml

```bash
pipeline-parser -p azure .azure-pipelines.yml
```

#### Parse Bitbucket Pipelines yaml

```bash
pipeline-parser -p bitbucket .bitbucket-pipelines.yml
```

#### Parse multiple files in one execution

```bash
pipeline-parser -p github workflow-1.yml workflow-2.yml workflow-3.yml
```

## Local Development

First, execute the following command to enable the client's git hooks:

```
git config core.hooksPath .githooks
```
