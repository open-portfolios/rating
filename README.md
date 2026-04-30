# review

E-commerce comment review system.

## Prerequisites

### Container Engine

This project uses [Podman](https://podman.io) as the container engine, but any OCI-compatible container engine should work.

[Docker](https://www.docker.com) support is included.

### Taskfile

[Taskfile](https://taskfile.dev) is introduced as an alternative to Makefile. Instead of typing a cluster of long long commands, you can use `task <name>` to run a defined recipe.

| Task            | Effect                                                            |
| --------------- | ----------------------------------------------------------------- |
| `task init`     | Install CLI tools needed for development                          |
| `task up`       | Compose up containers                                             |
| `task down`     | Shut down containers                                              |
| `task clean`    | Shut down containers and **remove all data** (be careful!)        |
| `task database` | Connect to the interactive shell of the database in the container |
| `task migrate`  | Create tables according to the SQL files under [sql/](./sql)      |
| `task serve`    | Run server                                                        |
| `task all`      | Perform `conf`, `api` and `wire` tasks                            |
| `task conf`     | Generate configuration protobuf                                   |
| `task api`      | Generate API protobuf                                             |
| `task wire`     | Generate Dependency Injection code                                |
| `task build`    | Build executables                                                 |

Taskfile uses the YAML format, and you will find it familiar if you have read GitHub Actions workflows before.

### Go

*The language we Gophers love*. The [Go](https://go.dev) version of this project is 1.22.

## License 

Copyright 2026 Open Portfolios Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.