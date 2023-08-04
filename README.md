Base Project
==============

This repository is a monorepo containing various components managed using Helm for dependencies, Protobuf files for schemas, and microservices within the "services" folder. Below, you'll find more information on each section along with links to relevant nodes.

Services
--------

The "services" directory holds microservices that make up the core functionality of the application. Here are the services:

-   [Backend Service](/services/backend): This service contains the backend logic of the application. It is responsible for handling data processing and business logic.

-   [Frontend Service](/services/frontend): The frontend service is responsible for the user interface of the application. It handles rendering and user interactions.

Helm Charts
-----------

The "charts" directory is used to manage dependencies and Helm charts for the application.

### Dependency Charts

-   [Minio Chart](/charts/dependency/minio): The Minio chart sets up an object storage server, enabling easy storage and retrieval of files.

-   [Traefik Chart](/charts/dependency/traefik): Traefik is a popular reverse proxy and load balancer that simplifies routing and SSL configuration.

-   [Zitadel Chart](/charts/dependency/zitadel): The Zitadel chart helps integrate Zitadel, an open-source IAM (Identity and Access Management) solution, into the application.

### Services Chart

The "services" chart is responsible for deploying the core microservices of the application. It combines the backend and frontend services.

Protos
------

The "protos" directory contains Protobuf files that define the schema and structure of data used in the application.

-   [Main Protobuf](/protos/main.proto): The main Protobuf file serves as the entry point, defining the basic structure of the data used throughout the application.

-   [Exception Protobufs](/protos/exception): This subdirectory contains Protobuf files related to exceptions and error handling in the application.
    <details><summary>:arrow_down:</summary>

    -   [Code Proto](/protos/exception/code.proto): The file defines error codes used in the application.

    -   [Throw Proto](/protos/exception/throw.proto): The file contains definitions for various exceptions that can be thrown by the services.
    </details>
License
-------

This project is licensed under the terms of the [LICENSE](/LICENSE) file.

Makefile
--------

The [Makefile](/Makefile) contains useful build and development scripts to ease the development process.


Feel free to explore each section and dive into the respective folders to get more details about each component and how they work together. If you have any questions or feedback, please don't hesitate to reach out via GitHub issues or contact the project maintainers. Happy coding!