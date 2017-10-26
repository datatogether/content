# Content

[![GitHub](https://img.shields.io/badge/project-Data_Together-487b57.svg?style=flat-square)](http://github.com/datatogether)
[![Slack](https://img.shields.io/badge/slack-Archivers-b44e88.svg?style=flat-square)](https://archivers-slack.herokuapp.com/)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](./LICENSE) 
[![Codecov](https://img.shields.io/codecov/c/github/datatogether/task_mgmt.svg?style=flat-square)](https://codecov.io/gh/datatogether/task_mgmt)

Content is a service for serving archived content stored on amazon S3.

## License & Copyright

[Modelled on [project guidelines template](https://github.com/datatogether/roadmap/blob/master/PROJECT.md#license--copyright-readme-block) ]

## Getting Involved

We would love involvement from more people! If you notice any errors or would like to submit changes, please see our [Contributing Guidelines](./.github/CONTRIBUTING.md). 

We use GitHub issues for [tracking bugs and feature requests](https://github.com/datatogether/REPONAME/issues) and Pull Requests (PRs) for [submitting changes](https://github.com/datatogether/REPONAME/pulls)

## Usage

We're working on a `docker-compose` file for this guy. In the meantime, to run this service, you'll have to connect to an amazon S3 bucket and postgres server by setting the folling environment variables:
```shell
  AWS_REGION
  AWS_S3_BUCKET_NAME
  AWS_S3_BUCKET_PATH
  POSTGRES_DB_URL
```
once those are properly set:
```shell
  # cd to project directory, then:
  go install

  content
```

## Development

Coming soon!

## Deployment

Coming soon!