:warning: It is still under development and not production-ready.
---

Base Project
==============

This repository is a monorepo containing various components managed using Helm for dependencies, Protobuf files for schemas, and microservices within the "services" folder. Below, you'll find more information on each section along with links to relevant nodes.

Services
--------

The "services" directory holds microservices that make up the core functionality of the application. Here are the services:

-   [Todo Service](/services/todo): This service contains the Todo logic of the application. It is responsible for handling data processing and business logic.

-   [Frontend Service](/services/frontend): The frontend service is responsible for the user interface of the application. It handles rendering and user interactions.

Helm Charts
-----------

The "charts" directory is used to manage dependencies and Helm charts for the application.

### Dependency Charts

-   [Minio ](/charts/dependency/minio): The Minio chart sets up an object storage server, enabling easy storage and retrieval of files.

-   [Traefik ](/charts/dependency/traefik): Traefik is a popular reverse proxy and load balancer that simplifies routing and SSL configuration.

-   [Zitadel ](/charts/dependency/zitadel): The Zitadel chart helps integrate Zitadel, an open-source IAM (Identity and Access Management) solution, into the application.
  
-   [CockroachDB ](https://chat.openai.com/charts/dependency/cockroachdb): The CockroachDB chart sets up a distributed SQL database, providing scalable and reliable data storage.

Protos
------

The "protos" directory contains Protobuf files that define the schema and structure of data used in the application.

-   [Todo Protobuf](/protos/base/todo.proto): This protobuf schema defines a basic todo application's data structure with fields for tasks and their properties. It also includes service methods for adding, retrieving, updating, and deleting tasks


Makefile
--------

The [Makefile](/Makefile) contains useful build and development scripts to ease the development process.


Makefile Commands and Descriptions

| Command              | Description                                 |
|----------------------|---------------------------------------------|
| `make help`          | Display this help message.                 |
| `make up`            | Deploy Helm charts.                         |
| `make down`          | Uninstall Helm charts.                      |
| `make stop`          | Stop running Helm releases for services.    |
| `make start`         | Start previously stopped Helm releases.    |
| `make up-dependency` | Deploy dependency Helm charts only.         |
| `make up-service`    | Deploy service Helm charts only.            |
| `make down-dependency`| Uninstall dependency Helm charts only.      |
| `make down-service`  | Uninstall service Helm charts only.         |

## Note:

Before using these commands, ensure you have the required Helm charts and configurations in the appropriate directories specified in the Makefile. Also, make sure you have set the environment variables in the `.env` file appropriately.

-------

This project is licensed under the terms of the [LICENSE](/LICENSE) file.