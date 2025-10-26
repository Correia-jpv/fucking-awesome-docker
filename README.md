# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)][sindresorhus] [![Netlify Status](https://api.netlify.com/api/v1/badges/8ca86717-11ba-46d4-9d0a-700d8527f13b/deploy-status)](https://app.netlify.com/sites/awesome-docker/deploys)[![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/veggiemonk/awesome-docker/)[![Last Commit](https://img.shields.io/github/last-commit/veggiemonk/awesome-docker)](https://github.com/correia-jpv/fucking-awesome-docker/commits/main)<!-- omit in toc -->

> A curated list of Docker resources and projects

If you would like to contribute, please read [CONTRIBUTING.md][contributing] first.
It contains a lot of tips and guidelines to help keep things organized.
Just click [README.md][editreadme] to submit a [pull request][editreadme].
If this list is not complete, you can [contribute][editreadme] to make it so. Here is a great video tutorial to learn how to 🌎 [contribute on Github](egghead.io/lessons/javascript-identifying-how-to-contribute-to-an-open-source-project-on-github).

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for newcomers. See how to **[Contribute][contributing]** for tips!

**_If you see a link here that is not (any longer) a good fit, you can fix it by submitting a [pull request][editreadme] to improve this file. Thank you!_**

The creators and maintainers of this list do not receive any form of payment to accept a change made by any contributor. This page is not an official Docker product in any way. It is a list of links to projects and is maintained by volunteers. Everybody is welcome to contribute. The goal of this repo is to index open-source projects, not to advertise for profit.

All the links are monitored and tested with a home baked [Node.js script](https://github.com/correia-jpv/fucking-awesome-docker/blob/master/.github/workflows/pull_request.yml)

# Contents <!-- omit in toc -->

<!-- TOC -->

- [Legend](#legend)
- [What is Docker](#what-is-docker)
- [Where to start](#where-to-start)
- [Where to start (Windows)](#where-to-start-windows)
- [Projects](#projects)
  - [Container Operations](#container-operations)
    - [Container Composition](#container-composition)
    - [Deployment and Infrastructure](#deployment-and-infrastructure)
    - [Monitoring](#monitoring)
    - [Networking](#networking)
    - [Orchestration](#orchestration)
    - [PaaS](#paas)
    - [Reverse Proxy](#reverse-proxy)
    - [Runtime](#runtime)
    - [Security](#security)
    - [Service Discovery](#service-discovery)
    - [Volume Management / Data](#volume-management--data)
    - [User Interface](#user-interface)
      - [IDE integrations](#ide-integrations)
      - [Desktop](#desktop)
      - [Terminal](#terminal)
        - [Terminal UI](#terminal-ui)
        - [CLI tools](#cli-tools)
        - [Other](#other)
      - [Web](#web)
  - [Docker Images](#docker-images)
    - [Base Tools](#base-tools)
    - [Builder](#builder)
    - [Dockerfile](#dockerfile)
    - [Linter](#linter)
    - [Metadata](#metadata)
    - [Registry](#registry)
  - [Development with Docker](#development-with-docker)
    - [API Client](#api-client)
    - [CI/CD](#cicd)
    - [Development Environment](#development-environment)
    - [Garbage Collection](#garbage-collection)
    - [Serverless](#serverless)
    - [Testing](#testing)
    - [Wrappers](#wrappers)
  - [Services based on Docker (mostly :heavy\_dollar\_sign:)](#services-based-on-docker-mostly-heavy_dollar_sign)
    - [CI Services](#ci-services)
    - [CaaS](#caas)
    - [Monitoring Services](#monitoring-services)
- [Useful Resources](#useful-resources)
  - [Awesome Lists](#awesome-lists)
  - [Demos and Examples](#demos-and-examples)
  - [Good Tips](#good-tips)
  - [Raspberry Pi \& ARM](#raspberry-pi--arm)
  - [Security](#security-1)
  - [Videos](#videos)
- [Communities and Meetups](#communities-and-meetups)
  - [Brazilian](#brazilian)
  - [Chinese](#chinese)
  - [English](#english)
  - [Russian](#russian)
  - [Spanish](#spanish)
  - [Stargazers over time](#stargazers-over-time)

<!-- /TOC -->

# Legend

-   Abandoned :skull:
-   Beta :construction:
-   Monetized :heavy_dollar_sign:

# What is Docker

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

_Source:_ 🌎 [What is Docker](www.docker.com/why-docker/)

# Where to start

-   🌎 [Benefits of using Docker](semaphore.io/blog/docker-benefits) for development and delivery, with a practical roadmap for adoption.
-   🌎 [Bootstrapping Microservices](www.manning.com/books/bootstrapping-microservices-with-docker-kubernetes-and-terraform) by 🌎 [Ashley Davis](twitter.com/ashleydavis75) - A practical and project-based guide to building applications with microservices, starts by building a Docker image for a single microservice and publishing it to a private container registry, finishes by deploying a complete microservices application to a production Kubernetes cluster.
-   <b><code>&nbsp;&nbsp;5952⭐</code></b> <b><code>&nbsp;&nbsp;2214🍴</code></b> [Docker Curriculum](https://github.com/prakhar1989/docker-curriculum)): A comprehensive tutorial for getting started with Docker. Teaches how to use Docker and deploy dockerized apps on AWS with Elastic Beanstalk and Elastic Container Service.
-   🌎 [Docker Documentation](docs.docker.com/): the official documentation.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;84⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27🍴</code></b> [Docker for beginners](https://github.com/groda/big_data/blob/master/docker_for_beginners.md)): A tutorial for beginners who need to learn the basics of Docker—from "Hello world!" to basic interactions with containers, with simple explanations of the underlying concepts.
-   🌎 [Docker for novices](www.youtube.com/watch?v=xsjSadjKXns) An introduction to Docker for developers and testers who have never used it. (Video 1h40, recorded linux.conf.au 2019 — Christchurch, New Zealand) by Alex Clews.

-   <b><code>&nbsp;&nbsp;&nbsp;272⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;256🍴</code></b> [Docker katas](https://github.com/eficode-academy/docker-katas)) A series of labs that will take you from "Hello Docker" to deploying a containerized web application to a server.
-   🌎 [Docker simplified in 55 seconds](www.youtube.com/watch?v=vP_4DlOH1G4): An animated high-level introduction to Docker. Think of it as a visual tl;dr that makes it easier to dive into more complex learning materials.
-   🌎 [Docker Training](training.mirantis.com) :heavy_dollar_sign:

-   🌎 [Introduction à Docker](blog.stephane-robert.info/docs/conteneurs/moteurs-conteneurs/docker/) A dedicated section to master Docker on a French site about DevSecOps: From the basics to best practices, including optimizing, securing your containers... 
-   <b><code>&nbsp;&nbsp;&nbsp;242⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;32🍴</code></b> [Learn Docker](https://github.com/dwyl/learn-docker)): step-by-step tutorial and more resources (video, articles, cheat sheets) by [@dwyl](https://github.com/dwyl)
-   🌎 [Learn Docker (Visually)](pagertree.com/learn/docker/overview) - A beginner-focused high-level overview of all the major components of Docker and how they fit together. Lots of high-quality images, examples, and resources.
-   🌎 [Play With Docker](training.play-with-docker.com/): PWD is a great way to get started with Docker from beginner to advanced users. Docker runs directly in your browser.
-   <b><code>&nbsp;&nbsp;&nbsp;249⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;77🍴</code></b> [Practical Guide about Docker Commands in Spanish](https://github.com/brunocascio/docker-espanol)) This Spanish guide contains the use of basic docker commands with real life examples.
-   <b><code>&nbsp;&nbsp;&nbsp;939⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;98🍴</code></b> [Setting Python Development Environment with VScode and Docker](https://github.com/RamiKrispin/vscode-python)): A step-by-step tutorial for setting up a dockerized Python development environment with VScode, Docker, and the Dev Container extension.
-   🌎 [The Docker Handbook](docker-handbook.farhan.dev/) An open-source book that teaches you the fundamentals, best practices and some intermediate Docker functionalities. The book is hosted on <b><code>&nbsp;&nbsp;&nbsp;847⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;127🍴</code></b> [fhsinchy/the-docker-handbook](https://github.com/fhsinchy/the-docker-handbook)) and the projects are hosted on <b><code>&nbsp;&nbsp;1387⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;322🍴</code></b> [fhsinchy/docker-handbook-projects](https://github.com/fhsinchy/docker-handbook-projects)) repository.


**Cheatsheets** by

-   <b><code>&nbsp;&nbsp;3864⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;485🍴</code></b> [@eon01](https://github.com/eon01/DockerCheatSheet))
-   <b><code>&nbsp;&nbsp;&nbsp;196⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [@dimonomid](https://github.com/dimonomid/docker-quick-ref)) (PDF)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [@JensPiegsa](https://github.com/JensPiegsa/docker-cheat-sheet))
-   <b><code>&nbsp;22437⭐</code></b> <b><code>&nbsp;&nbsp;4657🍴</code></b> [@wsargent](https://github.com/wsargent/docker-cheat-sheet)) (Most popular)

# Where to start (Windows)

-   🌎 [Docker on Windows behind a firewall](toedter.com/2015/05/11/docker-on-windows-behind-a-firewall/) by 🌎 [@kaitoedter](twitter.com/kaitoedter)
-   🌎 [Docker Reference Architecture: Modernizing Traditional .NET Framework Applications](docs.mirantis.com/containers/v3.0/dockeree-ref-arch/app-dev/modernize-dotnet-apps.html) - You will learn to identify the types of .NET Framework applications that are good candidates for containerization, the "lift-and-shift" approach to containerization.
-   🌎 [Docker with Microsoft SQL 2016 + ASP.NET](blog.alexellis.io/docker-does-sql2016-aspnet/) Demonstration running ASP.NET and SQL Server workloads in Docker
-   🌎 [Exploring ASP.NET Core with Docker in both Linux and Windows Containers](www.hanselman.com/blog/exploring-aspnet-core-with-docker-in-both-linux-and-windows-containers) Running ASP.NET Core apps in Linux and Windows containers, using [Docker for Windows][docker-for-windows]
-   🌎 [Running a Legacy ASP.NET App in a Windows Container](blog.sixeyed.com/dockerizing-nerd-dinner-part-1-running-a-legacy-asp-net-app-in-a-windows-container/) Steps for Dockerizing a legacy ASP.NET app and running as a Windows container
-   🌎 [Windows Containers and Docker: The 101](www.youtube.com/watch?v=N7SG2wEyQtM) :movie_camera: - A 20-minute overview, using Docker to run PowerShell, ASP.NET Core and ASP.NET apps
-   🌎 [Windows Containers Quick Start](learn.microsoft.com/en-us/virtualization/windowscontainers/about/) Overview of Windows containers, drilling down to Quick Starts for Windows 10 and Windows Server 2016

---

# Projects

-   Moby = open source development
-   Docker CE = free product release based on Moby
-   Docker EE = commercial product release based on Docker CE.

> Docker EE is on the same code base as Docker CE, so also built from Moby, with commercial components added, such as "docker data center / universal control plane"

-   <b><code>&nbsp;70957⭐</code></b> <b><code>&nbsp;18821🍴</code></b> [Moby](https://github.com/moby/moby))
-   🌎 [Docker Images](hub.docker.com)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker Compose](https://github.com/docker/compose/)) (Define and run multi-container applications with Docker)
-   <b><code>&nbsp;&nbsp;6638⭐</code></b> <b><code>&nbsp;&nbsp;1966🍴</code></b> [Docker Machine](https://github.com/docker/machine)) :skull: (Machine management for a container-centric world)
-   [Docker Registry][distribution] (The Docker toolset to pack, ship, store, and deliver content)
-   <b><code>&nbsp;&nbsp;5747⭐</code></b> <b><code>&nbsp;&nbsp;1073🍴</code></b> [Docker Swarm](https://github.com/docker/swarm)) :skull: (Swarm: a Docker-native clustering system)

## Container Operations

### Container Composition

-   <b><code>&nbsp;&nbsp;&nbsp;145⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [bocker](https://github.com/icy/bocker)) (2) :skull: - Write Dockerfile completely in Bash. Extensible and simple. --> Reusable by [@icy](https://github.com/icy)
-   <b><code>&nbsp;12529⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;750🍴</code></b> [bocker](https://github.com/p8952/bocker)) (1) :skull: - Docker implemented in 100 lines of bash by [p8952](https://github.com/p8952)
-   <b><code>&nbsp;&nbsp;&nbsp;242⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [box](https://github.com/box-builder/box)) :skull: - Build Dockerfile images with a mruby DSL, includes flattening and layer manipulation
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [Capitan](https://github.com/byrnedo/capitan)) - Composable docker orchestration with added scripting support by [@byrnedo].
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [compose_plantuml](https://github.com/funkwerk/compose_plantuml)) :skull: - Generate Plantuml graphs from docker-compose files by [@funkwerk](https://github.com/funkwerk)
-   <b><code>&nbsp;&nbsp;3611⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;241🍴</code></b> [Composerize](https://github.com/magicmark/composerize)) - Convert docker run commands into docker-compose files
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;98⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [crowdr](https://github.com/polonskiy/crowdr)) - Tool for managing multiple Docker containers (`docker-compose` alternative) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@polonskiy](https://github.com/polonskiy/))
-   <b><code>&nbsp;&nbsp;&nbsp;299⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27🍴</code></b> [ctk](https://github.com/ctk-hq/ctk)) :construction: - Visual composer for container based workloads. By [@corpulent](https://github.com/corpulent)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;55⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [docker-compose-graphviz](https://github.com/abesto/docker-compose-graphviz)) :skull: - Turn a docker-compose.yml files into Graphviz .dot files by [@abesto](https://github.com/abesto)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;50⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [docker-config-update](https://github.com/sudo-bmitch/docker-config-update)) - Utility to update docker configs and secrets for deploying in a compose file by [@sudo-bmitch](https://github.com/sudo-bmitch)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [draw-compose](https://github.com/Alexis-benoist/draw-compose)) :skull: - Utility to draw a schema of a docker compose by [@Alexis-benoist](https://github.com/Alexis-benoist)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;80⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23🍴</code></b> [elsy](https://github.com/cisco/elsy)) - An opinionated, multi-language, build tool based on Docker and Docker Compose
-   <b><code>&nbsp;&nbsp;1383⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81🍴</code></b> [habitus](https://github.com/cloud66-oss/habitus)) - A Build Flow Tool for Docker by [@cloud66](https://github.com/cloud66)
-   <b><code>&nbsp;10283⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;791🍴</code></b> [kompose](https://github.com/kubernetes/kompose)) - Go from Docker Compose to Kubernetes
-   <b><code>&nbsp;&nbsp;&nbsp;615⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;40🍴</code></b> [Maestro](https://github.com/toscanini/maestro)) :skull: - Maestro provides the ability to easily launch, orchestrate and manage multiple Docker containers as single unit by [@tascanini](https://github.com/toscanini)
-   <b><code>&nbsp;&nbsp;&nbsp;157⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [percheron](https://github.com/ashmckenzie/percheron)) :skull: - Organise your Docker containers with muscle and intelligence by [@ashmckenzie](https://github.com/ashmckenzie)
-   <b><code>&nbsp;&nbsp;&nbsp;385⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [plash](https://github.com/ihucos/plash)) - A container run and build engine - runs inside docker.
-   <b><code>&nbsp;&nbsp;5791⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;543🍴</code></b> [podman-compose](https://github.com/containers/podman-compose)) - a script to run docker-compose.yml using podman by [@containers][containers]
-   <b><code>&nbsp;&nbsp;&nbsp;409⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [rocker-compose](https://github.com/grammarly/rocker-compose)) :skull: - Docker composition tool with idempotency features for deploying apps composed of multiple containers. By[@grammarly].
-   <b><code>&nbsp;&nbsp;1333⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;79🍴</code></b> [rocker](https://github.com/grammarly/rocker)) :skull: - Extended Dockerfile builder. Supports multiple FROMs, MOUNTS, templates, etc. by [@grammarly].
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [Smalte](https://github.com/roquie/smalte)) – Dynamically configure applications that require static configuration in docker container. By [@roquie](https://github.com/roquie)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;69⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Stacker](https://github.com/stacker/stacker-cli)) :skull: - Docker Compose Templates. Stacker provides an abstraction layer over Docker Compose and a better DX (developer experience).
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;30⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Stitchocker](https://github.com/alexaandrov/stitchocker)) - A lightweight and fast command line utility for conveniently grouping your docker-compose multiple container services. By [@alexaandrov](https://github.com/alexaandrov)
-   <b><code>&nbsp;&nbsp;&nbsp;200⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [Zodiac](https://github.com/CenturyLinkLabs/zodiac)) :skull: - A lightweight tool for easy deployment and rollback of dockerized applications. By [@CenturyLinkLabs][centurylinklabs]

### Deployment and Infrastructure

-   <b><code>&nbsp;&nbsp;1187⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;159🍴</code></b> [awesome-stacks](https://github.com/ethibox/awesome-stacks)) - Deploy 150+ open-source web apps with one Docker command
-   🌎 [blackfish](gitlab.com/blackfish/blackfish) - a CoreOS VM to build swarm clusters for Dev & Production by 🌎 [@blackfish](gitlab.com/blackfish/)
-   🌎 [BosnD](gitlab.com/n0r1sk/bosnd) - BosnD, the boatswain daemon - A dynamic configuration file writer & service reloader for dynamically changing container environments.
-   <b><code>&nbsp;&nbsp;1759⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;112🍴</code></b> [Centurion](https://github.com/newrelic/centurion)) - Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings. By [@newrelic](https://github.com/newrelic)
-   <b><code>&nbsp;&nbsp;&nbsp;429⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;65🍴</code></b> [Clocker](https://github.com/brooklyncentral/clocker)) - Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts, on both Docker and Marathon. It leverages [Calico][calico] and [Weave][weave] for networking and 🌎 [Brooklyn](brooklyn.apache.org/) for application blueprints. By [@brooklyncentral](https://github.com/brooklyncentral)
-   <b><code>&nbsp;&nbsp;&nbsp;107⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Conduit](https://github.com/ehazlett/conduit)) - Experimental deployment system for Docker by [@ehazlett](https://github.com/ehazlett)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;93⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [depcon](https://github.com/ContainX/depcon)) - Depcon is written in Go and allows you to easily deploy Docker containers to Apache Mesos/Marathon, Amazon ECS and Kubernetes. By [@ContainX][containx]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;56⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [deploy](https://github.com/ttiny/deploy)) :skull: - Git and Docker deployment tool. A middle ground between simple Docker composition tools and full blown cluster orchestration by [@ttiny](https://github.com/ttiny)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [docker-to-iac](https://github.com/deploystackio/docker-to-iac)) - Translate docker run and commit into Infrastructure as Code templates for AWS, Render.com and DigitalOcean by [@DeployStack](https://github.com/deploystackio)
-   <b><code>&nbsp;&nbsp;&nbsp;106⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [dockit](https://github.com/humblec/dockit)) :skull: - Do docker actions and Deploy gluster containers! By [@humblec](https://github.com/humblec)
-   <b><code>&nbsp;&nbsp;3839⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;206🍴</code></b> [gitkube](https://github.com/hasura/gitkube)) - Gitkube is a tool for building and deploying docker images on Kubernetes using `git push`. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Hasura](https://github.com/hasura/)).
-   <b><code>&nbsp;&nbsp;1549⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;305🍴</code></b> [Grafeas](https://github.com/grafeas/grafeas)) - A common API for metadata about containers, from image and build details to security vulnerabilities. By [grafeas](https://github.com/grafeas)
-   <b><code>&nbsp;&nbsp;&nbsp;420⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Longshoreman](https://github.com/longshoreman/longshoreman)) :skull: - Longshoreman automates application deployment using Docker. Just create a Docker repository (or use a service), configure the cluster using AWS or Digital Ocean (or whatever you like) and deploy applications using a Heroku-like CLI tool. By [longshoreman](https://github.com/longshoreman)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;57⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [swarm-ansible](https://github.com/LombardiDaniel/swarm-ansible?tab=readme-ov-file)) - Swarm-Ansible bootstraps a production-ready swarm cluster using ansible. Comes with tools to automate CI, help monitoring and traefik pre-configured for SSL certificates and simple-auth. Comes with a private registry and more!
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [SwarmManagement](https://github.com/hansehe/SwarmManagement)) - Swarm Management is a python application, installed with pip. The application makes it easy to manage a Docker Swarm by configuring a single yaml file describing which stacks to deploy, and which networks, configs or secrets to create.
-   <b><code>&nbsp;&nbsp;4546⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;218🍴</code></b> [werf](https://github.com/werf/werf)) - werf is a CI/CD tool for building Docker images efficiently and deploying them to Kubernetes using GitOps by [@flant](https://github.com/flant)

### Monitoring

- <b><code>&nbsp;&nbsp;1646⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;244🍴</code></b> [Autoheal](https://github.com/willfarrell/docker-autoheal)) - Monitor and restart unhealthy docker containers automatically.
- 🌎 [Axibase Collector](axibase.com/docs/axibase-collector/) - Axibase Collector streams performance counters, configuration changes and lifecycle events from the Docker engine(s) into Axibase Time Series Database for roll-up dashboards and integration with upstream monitoring systems.
- <b><code>&nbsp;18496⭐</code></b> <b><code>&nbsp;&nbsp;2408🍴</code></b> [cAdvisor](https://github.com/google/cadvisor)) - Analyzes resource usage and performance characteristics of running containers. Created by [@Google][google]
- <b><code>&nbsp;&nbsp;8457⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;588🍴</code></b> [Checkmate](https://github.com/bluewave-labs/checkmate)) - Checkmate is an open-source, self-hosted tool designed to track and monitor server hardware, uptime, response times, and incidents in real-time with beautiful visualizations.
- <b><code>&nbsp;&nbsp;&nbsp;107⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [Docker-Alertd](https://github.com/deltaskelta/docker-alertd)) - Monitor and send alerts based on docker container resource usage/statistics
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38🍴</code></b> [Docker-Flow-Monitor](https://github.com/docker-flow/docker-flow-monitor)) - Reconfigures Prometheus when a new service is updated or deployed automatically by [@docker-flow][docker-flow]
- <b><code>&nbsp;&nbsp;&nbsp;207⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27🍴</code></b> [Dockerana](https://github.com/dockerana/dockerana)) :skull: - packaged version of Graphite and Grafana, specifically targeted at metrics from Docker.
- 🌎 [DockProc](gitlab.com/n0r1sk/dockproc) - I/O monitoring for containers on processlevel.
- <b><code>&nbsp;&nbsp;6364⭐</code></b> <b><code>&nbsp;&nbsp;1774🍴</code></b> [dockprom](https://github.com/stefanprodan/dockprom)) - Docker hosts and containers monitoring with Prometheus, Grafana, cAdvisor, NodeExporter and AlertManager by [@stefanprodan](https://github.com/stefanprodan)
- <b><code>&nbsp;&nbsp;&nbsp;381⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [Doku](https://github.com/amerkurev/doku)) - Doku is a simple web-based application that allows you to monitor Docker disk usage. [@amerkurev](https://github.com/amerkurev)
- [Dozzle](dozzle) - Monitor container logs in real-time with a browser or mobile device. [@amir20](https://github.com/amir20)
- 🌎 [Dynatrace](www.dynatrace.com/solutions/container-monitoring/) :heavy_dollar_sign: - Monitor containerized applications without installing agents or modifying your Run commands
- <b><code>&nbsp;30329⭐</code></b> <b><code>&nbsp;&nbsp;1629🍴</code></b> [Glances](https://github.com/nicolargo/glances)) - A cross-platform curses-based system monitoring tool written in Python by [@nicolargo](https://github.com/nicolargo)
- 🌎 [Grafana Docker Dashboard Template](grafana.com/grafana/dashboards/179-docker-prometheus-monitoring/) - A template for your Docker, Grafana and Prometheus stack [@vegasbrianc][vegasbrianc]
- <b><code>&nbsp;&nbsp;6754⭐</code></b> <b><code>&nbsp;&nbsp;1196🍴</code></b> [HertzBeat](https://github.com/dromara/hertzbeat)) - An open-source real-time monitoring system with custom-monitor and agentless.
- <b><code>&nbsp;&nbsp;&nbsp;474⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;131🍴</code></b> [InfluxDB, cAdvisor, Grafana](https://github.com/vegasbrianc/docker-monitoring)) - InfluxDB Time series DB in combination with Grafana and cAdvisor by [@vegasbrianc][vegasbrianc]
- <b><code>&nbsp;&nbsp;&nbsp;135⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [LogJam](https://github.com/gocardless/logjam)) - :skull: Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward these messages on to a log collection server (such as logstash) by [@gocardless](https://github.com/gocardless)
- <b><code>&nbsp;&nbsp;4694⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;671🍴</code></b> [Logspout](https://github.com/gliderlabs/logspout)) - Log routing for Docker container logs by [@gliderlabs][gliderlabs]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [monit-docker](https://github.com/decryptus/monit-docker)) - Monitor docker containers resources usage or status and execute docker commands or inside containers. [@decryptus][decryptus]
- <b><code>&nbsp;&nbsp;&nbsp;567⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;72🍴</code></b> [NexClipper](https://github.com/NexClipper/NexClipper)) - NexClipper is the container monitoring and performance management solution specialized in Docker, Apache Mesos, Marathon, DC/OS, Mesosphere, Kubernetes by [@Nexclipper](https://github.com/NexClipper)
- <b><code>&nbsp;&nbsp;&nbsp;539⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;119🍴</code></b> [Out-of-the-box Host/Container Monitoring/Logging/Alerting Stack](https://github.com/uschtwill/docker_monitoring_logging_alerting)) - Docker host and container monitoring, logging and alerting out of the box using cAdvisor, Prometheus, Grafana for monitoring, Elasticsearch, Kibana and Logstash for logging and elastalert and Alertmanager for alerting. Set up in 5 Minutes. Secure mode for production use with built-in [Automated Nginx Reverse Proxy (jwilder's)][nginxproxy].
- <b><code>&nbsp;&nbsp;1618⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;69🍴</code></b> [Sidekick](https://github.com/runsidekick/sidekick)) 💲 - Open source live application debugger like Chrome DevTools for your backend. Collect traces and generate logs on-demand without stopping & redeploying your applications.
- 🌎 [SuperVisor CPM](t0xic0der.medium.com/simply-accessible-container-performance-monitoring-with-supervisor-7fb47f925f3b) <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Frontend Service](https://github.com/t0xic0der/supervisor-frontend-service/)) and <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Driver Service](https://github.com/t0xic0der/supervisor-driver-service/)) :construction: - A simple and accessible FOSS container performance monitoring service written in Python by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@t0xic0der](https://github.com/t0xic0der/))
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [SwarmAlert](https://github.com/gpulido/SwarmAlert)) - Monitors a Docker Swarm and sends Pushover alerts when it finds a container with no healthy service task running.
- <b><code>&nbsp;&nbsp;1197⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;264🍴</code></b> [Zabbix Docker module](https://github.com/monitoringartist/Zabbix-Docker-Monitoring)) - Zabbix module that provides discovery of running containers, CPU/memory/blk IO/net container metrics. Systemd Docker and LXC execution driver is also supported. It's a dynamically linked shared object library, so its performance is (~10x) better, than any script solution.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;53⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [Zabbix Docker](https://github.com/gomex/docker-zabbix)) - Monitor containers automatically using zabbix LLD feature.

### Networking

-   [Calico][calico] - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Flannel](https://github.com/coreos/flannel/)) - Flannel is a virtual network that gives a subnet to each host for use with container runtimes. By [@coreos][coreos]
-   <b><code>&nbsp;&nbsp;&nbsp;628⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94🍴</code></b> [Freeflow](https://github.com/Microsoft/Freeflow)) - High performance container overlay networks on Linux. Enabling RDMA (on both InfiniBand and RoCE) and accelerating TCP to bare metal performance. By [@Microsoft](https://github.com/Microsoft)
-   <b><code>&nbsp;&nbsp;9183⭐</code></b> <b><code>&nbsp;&nbsp;1007🍴</code></b> [MyIP](https://github.com/jason5ng32/MyIP)) - All in one IP Toolbox. Easy to check all your IPs, IP geolocation, check for DNS leaks, examine WebRTC connections, speed test, ping test, MTR test, check website availability, whois search and more. By [@jason5ng32](https://github.com/jason5ng32)
-   <b><code>&nbsp;&nbsp;9958⭐</code></b> <b><code>&nbsp;&nbsp;1068🍴</code></b> [netshoot](https://github.com/nicolaka/netshoot)) - The netshoot container has a powerful set of networking tools to help troubleshoot Docker networking issues by [@nicolaka](https://github.com/nicolaka)
-   <b><code>&nbsp;&nbsp;4254⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;727🍴</code></b> [Pipework](https://github.com/jpetazzo/pipework)) - Software-Defined Networking for Linux Containers, Pipework works with "plain" LXC containers, and with the awesome Docker. By [@jpetazzo][jpetazzo]
-   [Weave][weave] (The Docker network) - Weave creates a virtual network that connects Docker containers deployed across multiple hosts.

### Orchestration

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Ansible Linux Docker](https://github.com/Peco602/ansible-linux-docker)) - Run Ansible from a Linux container. By [@Peco602][peco602]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25🍴</code></b> [athena](https://github.com/athena-oss/athena)) - An automation platform with a plugin architecture that allows you to easily create and share services.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [blimp](https://github.com/tubesandlube/blimp)) :skull: - Uses Docker Machine to easily move a container from one Docker host to another, show containers running against all of your hosts, replicate a container across multiple hosts and more by [@defermat](https://github.com/defermat) and [@schvin](https://github.com/schvin)
-   <b><code>&nbsp;&nbsp;&nbsp;239⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;82🍴</code></b> [CloudSlang](https://github.com/CloudSlang/cloud-slang)) - CloudSlang is a workflow engine to create Docker process automation
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;29⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [clusterdock](https://github.com/clusterdock/clusterdock)) - Docker container orchestration to enable the testing of long-running cluster deployments
-   <b><code>&nbsp;&nbsp;&nbsp;750⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;170🍴</code></b> [Crane](https://github.com/Dataman-Cloud/crane)) - Control plane based on docker built-in swarm [@Dataman-Cloud](https://github.com/Dataman-Cloud)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;69⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;49🍴</code></b> [Docker Flow Swarm Listener](https://github.com/docker-flow/docker-flow-swarm-listener)) - Docker Flow Swarm Listener project is to listen to Docker Swarm events and send requests when a change occurs. By [@docker-flow][docker-flow]
-   <b><code>&nbsp;&nbsp;3000⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;87🍴</code></b> [docker rollout](https://github.com/Wowu/docker-rollout)) - Zero downtime deployment for Docker Compose services by [@Wowu](https://github.com/Wowu)
-   <b><code>&nbsp;&nbsp;&nbsp;267⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;28🍴</code></b> [gantryd](https://github.com/DevTable/gantryd)) :skull: - A framework for easy management of docker-based components across machines by [@DevTable](https://github.com/DevTable)
-   <b><code>&nbsp;&nbsp;&nbsp;296⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;41🍴</code></b> [Haven](https://github.com/codeabovelab/haven-platform)) - Haven is a simplified container management platform that integrates container, application, cluster, image, and registry managements. By [@codeabovelab](https://github.com/codeabovelab)
-   <b><code>&nbsp;&nbsp;2140⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;233🍴</code></b> [Helios](https://github.com/spotify/helios)) :skull: - A simple platform for deploying and managing containers across an entire fleet of servers by [@spotify][spotify]
-   <b><code>&nbsp;&nbsp;1463⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;125🍴</code></b> [Kontena](https://github.com/kontena/kontena)) :skull: - The developer friendly container and micro services platform. Works on any cloud, easy to setup, simple to use.
-   <b><code>118209⭐</code></b> <b><code>&nbsp;41584🍴</code></b> [Kubernetes](https://github.com/kubernetes/kubernetes)) - Open source orchestration system for Docker containers by Google
-   <b><code>&nbsp;&nbsp;1380⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;911🍴</code></b> [ManageIQ](https://github.com/ManageIQ/manageiq)) - Discover, optimize and control your hybrid IT. By [ManageIQ](https://github.com/ManageIQ)
-   <b><code>&nbsp;&nbsp;2985⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;412🍴</code></b> [Mantl](https://github.com/mantl/mantl)) - :skull: Mantl is a modern platform for rapidly deploying globally distributed services
-   <b><code>&nbsp;&nbsp;4052⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;836🍴</code></b> [Marathon](https://github.com/mesosphere/marathon)) - :skull: Marathon is a private PaaS built on Mesos. It automatically handles hardware or software failures and ensures that an app is "always on"
-   <b><code>&nbsp;&nbsp;5341⭐</code></b> <b><code>&nbsp;&nbsp;1673🍴</code></b> [Mesos](https://github.com/apache/mesos)) - Resource/Job scheduler for containers, VM's and physical hosts 🌎 [@apache](mesos.apache.org/)
-   [Nebula](https://github.com/nebula-orchestrator) - A Docker orchestration tool designed to manage massive scale distributed clusters.
-   <b><code>&nbsp;15889⭐</code></b> <b><code>&nbsp;&nbsp;2027🍴</code></b> [Nomad](https://github.com/hashicorp/nomad)) - Easily deploy applications at any scale. A Distributed, Highly Available, Datacenter-Aware Scheduler by [@hashicorp](https://github.com/hashicorp)
-   <b><code>&nbsp;&nbsp;1438⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;148🍴</code></b> [Panamax](https://github.com/CenturyLinkLabs/panamax-ui)) :skull: - An open-source project that makes deploying complex containerized apps as easy as Drag-and-Drop by [@CenturyLinkLabs][centurylinklabs].
-   <b><code>&nbsp;24800⭐</code></b> <b><code>&nbsp;&nbsp;3109🍴</code></b> [Rancher](https://github.com/rancher/rancher)) - An open source project that provides a complete platform for operating Docker in production by [@rancher][rancher].
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;73⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [RedHerd Framework](https://github.com/redherd-project/redherd-framework)) - RedHerd is a collaborative and serverless framework for orchestrating a geographically distributed group of assets capable of simulating complex offensive cyberspace operations. By [@RedHerdProject](https://github.com/redherd-project).
-   <b><code>&nbsp;&nbsp;&nbsp;845⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;73🍴</code></b> [Swarm-cronjob](https://github.com/crazy-max/swarm-cronjob)) - Create jobs on a time-based schedule on Swarm by [@crazy-max]

### PaaS

-   <b><code>&nbsp;&nbsp;&nbsp;390⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34🍴</code></b> [Atlantis](https://github.com/ooyala/atlantis)) :skull: - Atlantis is an Open Source PaaS for HTTP applications built on Docker and written in Go
-   <b><code>&nbsp;14561⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;938🍴</code></b> [caprover](https://github.com/caprover/caprover)) - [previously known as CaptainDuckDuck] Automated Scalable Webserver Package (automated Docker+nginx) - Heroku on Steroids
-   <b><code>&nbsp;&nbsp;1892⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;179🍴</code></b> [Convox Rack](https://github.com/convox/rack)) - Convox Rack is open source PaaS built on top of expert infrastructure automation and devops best practices.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Dcw](https://github.com/pbertera/dcw)) - Docker-compose SSH wrapper: a very poor man PaaS, exposing the docker-compose and custom-container commands defined in container labels.
-   <b><code>&nbsp;31411⭐</code></b> <b><code>&nbsp;&nbsp;2006🍴</code></b> [Dokku](https://github.com/dokku/dokku)) - Docker powered mini-Heroku that helps you build and manage the lifecycle of applications (originally by [@progrium][progrium])
-   <b><code>&nbsp;&nbsp;2679⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;156🍴</code></b> [Empire](https://github.com/remind101/empire)) - A PaaS built on top of Amazon EC2 Container Service (ECS)
-   <b><code>&nbsp;&nbsp;1146⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;59🍴</code></b> [Exoframe](https://github.com/exoframejs/exoframe)) - A self-hosted tool that allows simple one-command deployments using Docker
-   <b><code>&nbsp;&nbsp;7859⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;591🍴</code></b> [Flynn](https://github.com/flynn/flynn)) :skull: - A next generation open source platform as a service
-   <b><code>&nbsp;&nbsp;&nbsp;418⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36🍴</code></b> [Hephy Workflow](https://github.com/teamhephy/workflow)) - Open source PaaS for Kubernetes that adds a developer-friendly layer to any Kubernetes cluster, making it easy to deploy and manage applications. Fork of <b><code>&nbsp;&nbsp;1304⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;178🍴</code></b> [Deis Workflow](https://github.com/deis/workflow))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [Krane](https://github.com/krane/krane)) - Toolset for managing container workloads on remote servers
-   <b><code>&nbsp;&nbsp;1622⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;92🍴</code></b> [Nanobox](https://github.com/nanobox-io/nanobox)) :heavy_dollar_sign: - An application development platform that creates local environments that can then be deployed and scaled in the cloud.
-   [OpenShift][openshift] - An open source PaaS built on [Kubernetes][kubernetes] and optimized for Dockerized app development and deployment by 🌎 [Red Hat](www.redhat.com/en)
-   <b><code>&nbsp;&nbsp;5204⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;548🍴</code></b> [Tsuru](https://github.com/tsuru/tsuru)) - Tsuru is an extensible and open source Platform as a Service software

### Reverse Proxy

-   <b><code>&nbsp;&nbsp;9266⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;525🍴</code></b> [bunkerized-nginx](https://github.com/bunkerity/bunkerized-nginx)) - Web app hosting and reverse proxy secure by default. By [@bunkerity](https://github.com/bunkerity)
-   <b><code>&nbsp;&nbsp;3977⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;203🍴</code></b> [caddy-docker-proxy](https://github.com/lucaslorentz/caddy-docker-proxy)) - Caddy-based reverse proxy, configured with service or container labels. By [@lucaslorentz](https://github.com/lucaslorentz)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [caddy-docker-upstreams](https://github.com/invzhi/caddy-docker-upstreams)) - Docker upstreams module for Caddy, configured with container labels. By [@invzhi](https://github.com/invzhi)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Docker Dnsmasq Updater](https://github.com/moonbuggy/docker-dnsmasq-updater)) - Update a remote dnsmasq server with Docker container hostnames.
-   <b><code>&nbsp;&nbsp;&nbsp;320⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;188🍴</code></b> [docker-flow-proxy](https://github.com/docker-flow/docker-flow-proxy)) - Reconfigures proxy every time a new service is deployed, or when a service is scaled. By [@docker-flow][docker-flow]
-   <b><code>&nbsp;&nbsp;&nbsp;284⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;67🍴</code></b> [docker-proxy](https://github.com/silarsis/docker-proxy)) :skull: - Transparent proxy for docker containers, run in a docker container. By [@silarsis](https://github.com/silarsis)
-   <b><code>&nbsp;&nbsp;7315⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;624🍴</code></b> [fabio](https://github.com/fabiolb/fabio)) - A fast, modern, zero-conf load balancing HTTP(S) router for deploying microservices managed by consul. By [@magiconair](https://github.com/magiconair) (Frank Schroeder)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;46⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [h2o-proxy](https://github.com/zchee/h2o-proxy)) :skull: - Automated H2O reverse proxy for Docker containers. An alternative to [jwilder/nginx-proxy][nginxproxy] by [@zchee](https://github.com/zchee)
-   <b><code>&nbsp;&nbsp;7654⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;831🍴</code></b> [Let's Encrypt Nginx-proxy Companion](https://github.com/nginx-proxy/docker-letsencrypt-nginx-proxy-companion)) - A lightweight companion container for the nginx-proxy. It allow the creation/renewal of Let's Encrypt certificates automatically. By [@JrCs](https://github.com/JrCs)
-   <b><code>&nbsp;&nbsp;&nbsp;167⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [muguet](https://github.com/mattallty/muguet)) :skull: - DNS Server & Reverse proxy for Docker environments. By [@mattallty](https://github.com/mattallty)
-   <b><code>&nbsp;28956⭐</code></b> <b><code>&nbsp;&nbsp;3339🍴</code></b> [Nginx Proxy Manager](https://github.com/jc21/nginx-proxy-manager)) - A beautiful web interface for proxying web based services with SSL. By [@jc21](https://github.com/jc21)
-   [nginx-proxy][nginxproxy] - Automated nginx proxy for Docker containers using docker-gen by [@jwilder][jwilder]
-   <b><code>&nbsp;&nbsp;1140⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88🍴</code></b> [OpenResty Manager](https://github.com/Safe3/openresty-manager)) - The easiest using, powerful and beautiful OpenResty Manager(Nginx Enhanced Version), open source alternative to OpenResty Edge. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Safe3](https://github.com/Safe3/))
-   <b><code>&nbsp;&nbsp;&nbsp;167⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Swarm Ingress Router](https://github.com/tpbowden/swarm-ingress-router)) :skull: - Route DNS names to Swarm services based on labels. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tpbowden](https://github.com/tpbowden/))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;73⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [Swarm Router](https://github.com/flavioaiello/swarm-router)) - A «zero config» service name based router for docker swarm mode with a fresh and more secure approach. By [@flavioaiello](https://github.com/flavioaiello)
-   <b><code>&nbsp;57311⭐</code></b> <b><code>&nbsp;&nbsp;5536🍴</code></b> [Træfɪk](https://github.com/containous/traefik)) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd... By [@EmileVauge](https://github.com/emilevauge)

### Runtime

-   <b><code>&nbsp;&nbsp;1491⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94🍴</code></b> [aind](https://github.com/aind-containers/aind)) - :skull: AinD launches Android apps in Docker, by nesting Anbox containers inside Docker by [@aind-containers](https://github.com/aind-containers)
-   <b><code>&nbsp;&nbsp;5519⭐</code></b> <b><code>&nbsp;&nbsp;1123🍴</code></b> [cri-o](https://github.com/cri-o/cri-o)) - Open Container Initiative-based implementation of Kubernetes Container Runtime Interface by [cri-o](https://github.com/cri-o)
-   <b><code>&nbsp;&nbsp;5014⭐</code></b> <b><code>&nbsp;&nbsp;1150🍴</code></b> [lxc](https://github.com/lxc/lxc)) - LXC - Linux Containers
-   <b><code>&nbsp;29379⭐</code></b> <b><code>&nbsp;&nbsp;2832🍴</code></b> [podman](https://github.com/containers/libpod)) - libpod is a library used to create container pods. Home of Podman by [@containers][containers]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [rlxc](https://github.com/brauner/rlxc)) - LXC binary written in Rust by [@brauner](https://github.com/brauner)
-   <b><code>&nbsp;&nbsp;&nbsp;458⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;157🍴</code></b> [runtime-tools](https://github.com/opencontainers/runtime-tools)) - oci-runtime-tool is a collection of tools for working with the OCI runtime specification by [@opencontainers](https://github.com/opencontainers)

### Security

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Anchor](https://github.com/SongStitch/anchor/)) - A tool to ensure reproducible builds by pinning dependencies inside your Dockerfiles <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@SongStitch](https://github.com/songStitch/))
-   🌎 [Anchor Enterprise](anchore.com/) :heavy_dollar_sign: - Analyze images for CVE vulnerabilities and against custom security policies by [@Anchor](https://github.com/anchore)
-   🌎 [Aqua Security](www.aquasec.com) :heavy_dollar_sign: - Securing container-based applications from Dev to Production on any platform
-   <b><code>&nbsp;&nbsp;1219⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88🍴</code></b> [bane](https://github.com/genuinetools/bane)) - AppArmor profile generator for Docker containers by [@genuinetools][genuinetools]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;75⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [CetusGuard](https://github.com/hectorm/cetusguard)) - CetusGuard is a tool that protects the Docker daemon socket by filtering calls to its API endpoints
-   <b><code>&nbsp;&nbsp;&nbsp;521⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;119🍴</code></b> [CIS Docker Benchmark](https://github.com/dev-sec/cis-docker-benchmark)) - This [InSpec][inspec] compliance profile implement the CIS Docker 1.12.0 Benchmark in an automated way to provide security best-practice tests around Docker daemon and containers in a production environment. By [@dev-sec](https://github.com/dev-sec)
-   <b><code>&nbsp;&nbsp;8128⭐</code></b> <b><code>&nbsp;&nbsp;1268🍴</code></b> [Checkov](https://github.com/bridgecrewio/checkov)) - Static analysis for infrastructure as code manifests (Terraform, Kubernetes, Cloudformation, Helm, Dockerfile, Kustomize) find security misconfiguration and fix them. By [@bridgecrew](https://github.com/bridgecrewio)
-   <b><code>&nbsp;10836⭐</code></b> <b><code>&nbsp;&nbsp;1195🍴</code></b> [Clair](https://github.com/quay/clair)) - Clair is an open source project for the static analysis of vulnerabilities in appc and docker containers. By [@coreos][coreos]
-   <b><code>&nbsp;&nbsp;1214⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;171🍴</code></b> [Dagda](https://github.com/eliasgranderubio/dagda)) - Dagda is a tool to perform static analysis of known vulnerabilities, trojans, viruses, malware & other malicious threats in docker images/containers and to monitor the docker daemon and running docker containers for detecting anomalous activities. By [@eliasgranderubio](https://github.com/eliasgranderubio)
-   🌎 [Deepfence Enterprise](deepfence.io) :heavy_dollar_sign: - Full life cycle Cloud Native Workload Protection platform for kubernetes, virtual machines and serverless. By [@deepfence][deepfence]
-   <b><code>&nbsp;&nbsp;5184⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;634🍴</code></b> [Deepfence Threat Mapper](https://github.com/deepfence/ThreatMapper)) - Powerful runtime vulnerability scanner for kubernetes, virtual machines and serverless. By [@deepfence][deepfence]
-   <b><code>&nbsp;&nbsp;9510⭐</code></b> <b><code>&nbsp;&nbsp;1038🍴</code></b> [docker-bench-security](https://github.com/docker/docker-bench-security)) - script that checks for dozens of common best-practices around deploying Docker containers in production. By [@docker][docker]
-   <b><code>&nbsp;&nbsp;&nbsp;549⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;46🍴</code></b> [docker-explorer](https://github.com/google/docker-explorer)) - A tool to help forensicate offline docker acquisitions by [@Google][google]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [docker-lock](https://github.com/safe-waters/docker-lock)) - A cli-plugin for docker to automatically manage image digests by tracking them in a separate Lockfile. By [@safe-waters][safe-waters]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [dvwassl](https://github.com/Peco602/dvwassl)) - SSL-enabled Damn Vulnerable Web App to test Web Application Firewalls. By [@Peco602][peco602]
-   <b><code>&nbsp;&nbsp;2495⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;348🍴</code></b> [KICS](https://github.com/checkmarx/kics)) - an infrastructure-as-code scanning tool, find security vulnerabilities, compliance issues, and infrastructure misconfigurations early in the development cycle. Can be extended for additional policies. By [Checkmarx](https://github.com/Checkmarx)
-   <b><code>&nbsp;&nbsp;3284⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;520🍴</code></b> [notary](https://github.com/theupdateframework/notary)) - a server and a client for running and interacting with trusted collections. By [@TUF](https://github.com/theupdateframework)
-   <b><code>&nbsp;&nbsp;1591⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;409🍴</code></b> [oscap-docker](https://github.com/OpenSCAP/openscap)) - OpenSCAP provides oscap-docker tool which is used to scan Docker containers and images. By [OpenSCAP](https://github.com/OpenSCAP)
-   🌎 [Prisma Cloud](www.paloaltonetworks.com/prisma/cloud) :heavy_dollar_sign: - (previously Twistlock Security Suite) detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.
-   <b><code>&nbsp;&nbsp;7842⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;725🍴</code></b> [Syft](https://github.com/anchore/syft)) - CLI tool and library for generating a Software Bill of Materials (SBOM) from container images and filesystems.
-   <b><code>&nbsp;&nbsp;8338⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;946🍴</code></b> [Sysdig Falco](https://github.com/falcosecurity/falco)) - Sysdig Falco is an open source container security monitor. It can monitor application, container, host, and network activity and alert on unauthorized activity.
-   🌎 [Sysdig Secure](www.sysdig.com/solutions/cloud-detection-and-response-cdr) :heavy_dollar_sign: - Sysdig Secure addresses run-time security through behavioral monitoring and defense, and provides deep forensics based on open source Sysdig for incident response.
-   🌎 [Trend Micro DeepSecurity](www.trendmicro.com/en_us/business/products/hybrid-cloud/deep-security.html) :heavy_dollar_sign: - Trend Micro DeepSecurity offers runtime protection for container workloads and hosts as well as preruntime scanning of images to identify vulnerabilities, malware and content such as hardcoded secrets.
-   <b><code>&nbsp;29533⭐</code></b> <b><code>&nbsp;&nbsp;2792🍴</code></b> [Trivy](https://github.com/aquasecurity/trivy)) - Aqua Security's open source simple and comprehensive vulnerability scanner for containers (suitable for CI).

### Service Discovery

-   <b><code>&nbsp;&nbsp;1062⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;278🍴</code></b> [docker-consul](https://github.com/gliderlabs/docker-consul)) by [@progrium][progrium]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [docker-dns](https://github.com/bytesharky/docker-dns)) - Lightweight DNS forwarder for Docker containers, resolves container names with custom suffixes (e.g. `.docker`) on the host to simplify service discovery by [@bytesharky](https://github.com/bytesharky)
-   <b><code>&nbsp;50630⭐</code></b> <b><code>&nbsp;10197🍴</code></b> [etcd](https://github.com/etcd-io/etcd)) - Distributed reliable key-value store for the most critical data of a distributed system by [@etcd-io](https://github.com/etcd-io) (former part of CoreOS)
-   <b><code>&nbsp;37560⭐</code></b> <b><code>&nbsp;&nbsp;8107🍴</code></b> [istio](https://github.com/istio/istio)) - An open platform to connect, manage, and secure microservices by [@istio](https://github.com/istio)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;52⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [proxy](https://github.com/factorish/proxy)) :skull: - lightweight nginx based load balancer self using service discovery provided by registrator. by [@factorish](https://github.com/factorish)
-   <b><code>&nbsp;&nbsp;4674⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;911🍴</code></b> [registrator](https://github.com/gliderlabs/registrator)) - Service registry bridge for Docker by [@gliderlabs][gliderlabs] and [@progrium][progrium]

### Volume Management / Data

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;93⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [Blockbridge](https://github.com/blockbridge/blockbridge-docker-volume)) :heavy_dollar_sign:- The Blockbridge plugin is a volume plugin that provides access to an extensible set of container-based persistent storage options. It supports single and multi-host Docker environments with features that include tenant isolation, automated provisioning, encryption, secure deletion, snapshots and QoS. By [@blockbridge](https://github.com/blockbridge)
-   <b><code>&nbsp;&nbsp;1305⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;135🍴</code></b> [Convoy](https://github.com/rancher/convoy)) :skull: - an open-source Docker volume driver that can snapshot, backup and restore Docker volumes anywhere. By [@rancher][rancher]
-   <b><code>&nbsp;&nbsp;&nbsp;792⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;100🍴</code></b> [Docker Machine NFS](https://github.com/adlogix/docker-machine-nfs)) :skull: - Activates NFS for an existing boot2docker box created through Docker Machine on OS X.
-   <b><code>&nbsp;&nbsp;&nbsp;167⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [Docker Unison](https://github.com/leighmcculloch/docker-unison)) - :skull: A docker volume container using Unison for fast two-way folder sync. Created as an alternative to slow boot2docker volumes on OS X. By [@leighmcculloch](https://github.com/leighmcculloch)
-   - <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Label Backup](https://github.com/resulgg/label-backup)) - A lightweight, Docker-aware backup agent that automatically discovers and backs up containerized databases (PostgreSQL, MySQL, MongoDB, Redis) based on Docker labels. Supports local storage and S3-compatible destinations with flexible scheduling via cron expressions.
-   <b><code>&nbsp;&nbsp;3068⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;120🍴</code></b> [Docker Volume Backup](https://github.com/offen/docker-volume-backup)) Backup Docker volumes locally or to any S3 compatible storage. By [@offen](https://github.com/offen)
-   <b><code>&nbsp;&nbsp;&nbsp;876⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;126🍴</code></b> [Local Persist](https://github.com/MatchbookLab/local-persist)) Specify a mountpoint for your local volumes (created via `docker volume create`) so that files will always persist and so you can mount to different directories in different containers.
-   <b><code>&nbsp;57325⭐</code></b> <b><code>&nbsp;&nbsp;6370🍴</code></b> [Minio](https://github.com/minio/minio)) - S3 compatible object storage server in Docker containers
-   <b><code>&nbsp;&nbsp;1141⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;165🍴</code></b> [Netshare](https://github.com/ContainX/docker-volume-netshare)) Docker NFS, AWS EFS, Ceph & Samba/CIFS Volume Plugin. By [@ContainX][containx]
-   🌎 [portworx](portworx.com) :heavy_dollar_sign: - Decentralized storage solution for persistent, shared and replicated volumes.
-   🌎 [quobyte](www.quobyte.com/) :heavy_dollar_sign: - fully fault-tolerant distributed file system with a docker volume driver
-   <b><code>&nbsp;&nbsp;2213⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;332🍴</code></b> [REX-Ray](https://github.com/rexray/rexray)) provides a vendor agnostic storage orchestration engine. The primary design goal is to provide persistent storage for Docker, Kubernetes, and Mesos. By[@thecodeteam](https://github.com/thecodeteam) (DELL Technologies)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Storidge](https://github.com/Storidge/quick-start)) :heavy_dollar_sign: - Software-defined Persistent Storage for Kubernetes and Docker Swarm

### User Interface

#### IDE integrations

-   JetBrains IDEs (IntelliJ IDEA, GoLand, WebStorm, CLion etc.) has 🌎 [built-in Docker plugin](www.jetbrains.com/help/idea/docker.html#managing-images)
-   Eclipse 🌎 [Docker Tooling plugin](www.eclipse.org/community/eclipse_newsletter/2016/july/article2.php)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;95⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [denops-docker.vim](https://github.com/skanehira/denops-docker.vim)) - :skull: Manage docker containers and images in Vim. By [@skanehira]
-   <b><code>&nbsp;&nbsp;&nbsp;176⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [docker.vim](https://github.com/skanehira/docker.vim)) :skull: - Manage docker containers and images in Vim. By [@skanehira]
-   <b><code>&nbsp;&nbsp;&nbsp;796⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;80🍴</code></b> [docker.el](https://github.com/Silex/docker.el)) Manage docker from Emacs by [@Silex](https://github.com/Silex)

#### Desktop

Native desktop applications for managing and monitoring docker hosts and clusters

-   🌎 [Docker Desktop](www.docker.com/products/docker-desktop/) - Official native app. Only for Windows and MacOS
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;54⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Docker DB Manager](https://github.com/AbianS/docker-db-manager)) - Desktop app for managing Docker database containers with visual interface and one-click operations.
-   <b><code>&nbsp;&nbsp;&nbsp;793⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;110🍴</code></b> [Dockeron](https://github.com/dockeron/dockeron)) :skull: - A project built on Electron + Vue.js for Docker on desktop. [@fluency03](https://github.com/fluency03)
-   <b><code>&nbsp;&nbsp;2182⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;108🍴</code></b> [DockStation](https://github.com/DockStation/dockstation)) - A developer centric UI to configure, monitor, and manage services and containers 🌎 [@dock_station](twitter.com/dock_station)
-   <b><code>&nbsp;&nbsp;&nbsp;186⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Lifeboat](https://github.com/jplhomer/lifeboat)) - :skull: An easy way to launch Docker projects with a graphical interface on your Mac. [@jplhomer](https://github.com/jplhomer)
-   <b><code>&nbsp;&nbsp;&nbsp;603⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;96🍴</code></b> [Simple Docker UI](https://github.com/felixgborrego/simple-docker-ui)) - built on Electron. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@felixgborrego](https://github.com/felixgborrego/))
-   <b><code>&nbsp;&nbsp;&nbsp;359⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [Stevedore](https://github.com/slonopotamus/stevedore)) - Good Docker Desktop replacement for Windows. Both Linux and Windows Containers are supported. [@slonopotamus](https://github.com/slonopotamus)

#### Terminal

##### Terminal UI
-   <b><code>&nbsp;&nbsp;&nbsp;423⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;33🍴</code></b> [ctop (1)](https://github.com/yadutaf/ctop)) - :skull: A command line / text based Linux Containers monitoring tool that works just like you expect (Python) by [@yadutaf](https://github.com/yadutaf)
-   <b><code>&nbsp;17026⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;572🍴</code></b> [ctop (2)](https://github.com/bcicen/ctop)) - :skull: Top-like interface for container metrics (Golang) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@bcicen](https://github.com/bcicen/))
-   <b><code>&nbsp;52367⭐</code></b> <b><code>&nbsp;&nbsp;1916🍴</code></b> [dive](https://github.com/wagoodman/dive)) - A tool for exploring each layer in a docker image. By [wagoodman](https://github.com/wagoodman).
-   <b><code>&nbsp;&nbsp;&nbsp;123⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [dockdash](https://github.com/byrnedo/dockdash)) detailed stats. By [@byrnedo]
-   <b><code>&nbsp;&nbsp;&nbsp;782⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;33🍴</code></b> [Docker-mon](https://github.com/icecrime/docker-mon)) :skull: - Console-based Docker monitoring by [@icecrime](https://github.com/icecrime)
-   <b><code>&nbsp;&nbsp;3975⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;164🍴</code></b> [dockly](https://github.com/lirantal/dockly)) - An interactive shell UI for managing Docker containers by [@lirantal](https://github.com/lirantal)
-   <b><code>&nbsp;&nbsp;2483⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;285🍴</code></b> [DockSTARTer](https://github.com/GhostWriters/DockSTARTer)) - DockSTARTer helps you get started with home server apps running in Docker by [GhostWriters](https://github.com/GhostWriters)
-   <b><code>&nbsp;&nbsp;2324⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;103🍴</code></b> [docui](https://github.com/skanehira/docui)) - :skull: An interactive shell UI for managing Docker containers. Also works in Windows. By [@skanehira]
-   <b><code>&nbsp;&nbsp;3181⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;100🍴</code></b> [dry](https://github.com/moncho/dry)) - An interactive CLI for Docker containers by [@moncho](https://github.com/moncho)
-   <b><code>&nbsp;&nbsp;&nbsp;614⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [goManageDocker](https://github.com/ajayd-san/gomanagedocker)) - TUI tool to view and manage your docker objects blazingly fast with sensible keybindings, also supports VIM navigation out of the box by [@ajay-dsan](https://github.com/ajayd-san)
-   <b><code>&nbsp;47132⭐</code></b> <b><code>&nbsp;&nbsp;1481🍴</code></b> [lazydocker](https://github.com/jesseduffield/lazydocker)) - The lazier way to manage everything docker. A simple terminal UI for both docker and docker-compose, written in Go with the gocui library. By [@jesseduffield](https://github.com/jesseduffield)
-   <b><code>&nbsp;&nbsp;&nbsp;740⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [lazyjournal](https://github.com/Lifailon/lazyjournal)) - A interface for reading and filtering the logs output of Docker and Podman containers like [Dozzle](dozzle) but for the terminal with support for fuzzy find, regex and output coloring
-   <b><code>&nbsp;&nbsp;1210⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37🍴</code></b> [oxker](https://github.com/mrjackwills/oxker)) - A simple tui to view & control docker containers. Written in 🌎 [Rust](rust-lang.org/), making heavy use of <b><code>&nbsp;15586⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;483🍴</code></b> [ratatui](https://github.com/tui-rs-revival/ratatui)) & <b><code>&nbsp;&nbsp;1138⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;154🍴</code></b> [Bollard](https://github.com/fussybeaver/bollard)), by [@mrjackwills](https://github.com/mrjackwills)
-   <b><code>&nbsp;&nbsp;1034⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62🍴</code></b> [sen](https://github.com/TomasTomecek/sen)) - :skull: Terminal user interface for docker engine, by [@TomasTomecek][tomastomecek]

##### CLI tools

-   <b><code>&nbsp;&nbsp;&nbsp;246⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [captain](https://github.com/jenssegers/captain)) - Easily start and stop docker compose projects from any directory. By [@jenssegers](https://github.com/jenssegers)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [dcinja](https://github.com/Falldog/dcinja)) - The powerful and smallest binary size of template engine for docker command line environment. By [@Falldog](https://github.com/Falldog)
-   <b><code>&nbsp;&nbsp;&nbsp;110⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [dcp](https://github.com/exdx/dcp)) - A simple tool for copying files from container filesystems. By [@exdx](https://github.com/exdx)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [dctl](https://github.com/FabienD/docker-stack)) - dctl is a Cli tool that helps developers by allowing them to execute all docker compose commands anywhere in the terminal, and more. By [FabienD](https://github.com/FabienD)
-   <b><code>&nbsp;&nbsp;&nbsp;114⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [decompose](https://github.com/s0rg/decompose)) - Reverse-engineering tool for docker environments. By [@s0rg](https://github.com/s0rg)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [docker-captain](https://github.com/lucabello/docker-captain)) - A friendly CLI to manage multiple Docker Compose deployments with style — powered by Typer, Rich, questionary, and sh.
-   <b><code>&nbsp;&nbsp;&nbsp;457⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;67🍴</code></b> [docker-ls](https://github.com/mayflower/docker-ls)) - CLI tools for browsing and manipulating docker registries by [@mayflower](https://github.com/mayflower)
-   <b><code>&nbsp;&nbsp;&nbsp;147⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker pushrm](https://github.com/christian-korneck/docker-pushrm)) - A Docker CLI plugin that lets you push the README.md file from the current directory to Docker Hub. Also supports Quay and Harbor. By [@christian-korneck](https://github.com/christian-korneck)
-   <b><code>&nbsp;&nbsp;&nbsp;122⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [dockersql](https://github.com/crosbymichael/dockersql)) - :skull: A command line interface to query Docker using SQL by [@crosbymichael](https://github.com/crosbymichael)
-   <b><code>&nbsp;&nbsp;&nbsp;524⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;50🍴</code></b> [DVM](https://github.com/howtowhale/dvm)) - Docker version manager by [@howtowhale](https://github.com/howtowhale)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;30⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [goinside](https://github.com/iamsoorena/goinside)) - Get inside a running docker container easily. by [@iamsoorena](https://github.com/iamsoorena)
-   <b><code>&nbsp;&nbsp;2596⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;270🍴</code></b> [ns-enter](https://github.com/jpetazzo/nsenter)) - :skull: no more ssh, enter name spaces of container by [@jpetazzo][jpetazzo]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Pdocker](https://github.com/g31s/Pdocker)) - A simple tool to manage and maintain Docker for personal projects by [@g31s](https://github.com/g31s)
-   <b><code>&nbsp;&nbsp;&nbsp;109⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [proco](https://github.com/shiwaforce/poco)) - Proco will help you to organise and manage Docker, Docker-Compose, Kubernetes projects of any complexity using simple YAML config files to shorten the route from finding your project to initialising it in your local environment. by [@shiwaforce](https://github.com/shiwaforce)
-   <b><code>&nbsp;&nbsp;1697⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;173🍴</code></b> [reg](https://github.com/genuinetools/reg)) - :skull: Docker registry v2 command line client by [@genuinetools][genuinetools]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;91⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [scuba](https://github.com/JonathonReinhart/scuba)) - Transparently use Docker containers to encapsulate software build environments, by [@JonathonReinhart](https://github.com/JonathonReinhart)
-   <b><code>&nbsp;&nbsp;9921⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;870🍴</code></b> [skopeo](https://github.com/containers/skopeo)) - Work with remote images registries - retrieving information, images, signing content by [@containers][containers]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;83⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [supdock](https://github.com/segersniels/supdock)) - Allows for slightly more visual usage of Docker with an interactive prompt. By [@segersniels](https://github.com/segersniels)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;54⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [tsaotun](https://github.com/qazbnm456/tsaotun)) - Python based Assistance for Docker by [@qazbnm456](https://github.com/qazbnm456)
-   <b><code>&nbsp;&nbsp;&nbsp;655⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;41🍴</code></b> [wharfee](https://github.com/j-bennet/wharfee)) - :skull: Autocompletion and syntax highlighting for Docker commands. by [@j-bennet](https://github.com/j-bennet)

##### Other

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [dext-docker-registry-plugin](https://github.com/vutran/dext-docker-registry-plugin)) - Search the Docker Registry with the Dext smart launcher. By [@vutran](https://github.com/vutran)
-   <b><code>&nbsp;&nbsp;&nbsp;656⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;91🍴</code></b> [docker-ssh](https://github.com/jeroenpeeters/docker-ssh)) - SSH Server for Docker containers ~ Because every container should be accessible. By [@jeroenpeeters](https://github.com/jeroenpeeters)
-   <b><code>&nbsp;&nbsp;7369⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;396🍴</code></b> [dockercraft](https://github.com/docker/dockercraft)) - :skull: Docker + Minecraft = Dockercraft by [@docker][docker]
-   <b><code>&nbsp;&nbsp;&nbsp;551⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81🍴</code></b> [dockerfile-mode](https://github.com/spotify/dockerfile-mode)) An emacs mode for handling Dockerfiles by [@spotify][spotify]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;55⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [MultiDocker](https://github.com/marty90/multidocker)) - Create a secure multi-user Docker machine, where each user is segregated into an indepentent container.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;61⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [Powerline-Docker](https://github.com/adrianmo/powerline-docker)) - A Powerline segment for showing the status of Docker containers by [@adrianmo](https://github.com/adrianmo)

#### Web

-   <b><code>&nbsp;&nbsp;&nbsp;257⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;65🍴</code></b> [Admiral](https://github.com/vmware/admiral)) - :skull: Admiral™ is a highly scalable and very lightweight Container Management platform for deploying and managing container based applications. By [VMWare][vmware]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [CASA](https://github.com/knrdl/casa)) - Outsource the administration of a handful of containers to your co-workers, by [@knrdl](https://github.com/knrdl)
-   <b><code>&nbsp;&nbsp;&nbsp;257⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;46🍴</code></b> [Container Web TTY](https://github.com/wrfly/container-web-tty)) - Connect your containers via a web-tty [@wrfly](https://github.com/wrfly)
-   <b><code>&nbsp;&nbsp;&nbsp;747⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35🍴</code></b> [dockemon](https://github.com/ProductiveOps/dokemon)) - Docker Container Management GUI by [@productiveops](https://github.com/ProductiveOps)
-   <b><code>&nbsp;&nbsp;1551⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;229🍴</code></b> [Docker Compose UI](https://github.com/francescou/docker-compose-ui)) - :skull: Manage docker-compose via HTTP. docker-compose-ui runs in a Docker container, mounts the hosts docker socket and exposes a RESTful API and AngularJS GUI
-   <b><code>&nbsp;&nbsp;&nbsp;627⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;58🍴</code></b> [Docker Registry Browser](https://github.com/klausmeyer/docker-registry-browser)) - Web Interface for the Docker Registry HTTP API v2 by [@klausmeyer](https://github.com/klausmeyer)
-   <b><code>&nbsp;&nbsp;3145⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;351🍴</code></b> [Docker Registry UI (Joxit)](https://github.com/Joxit/docker-registry-ui)) - :skull: The simplest and cleanest UI for private registries by [@Joxit](https://github.com/Joxit)
-   <b><code>&nbsp;&nbsp;&nbsp;892⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;147🍴</code></b> [Docker Registry UI](https://github.com/atcol/docker-registry-ui)) - A web UI for easy private/local Docker Registry integration by [@atcol](https://github.com/atcol)
-   <b><code>&nbsp;&nbsp;&nbsp;543⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;134🍴</code></b> [docker-registry-web](https://github.com/mkuchin/docker-registry-web)) - Web UI, authentication service and event recorder for private docker registry v2 by [@mkuchin](https://github.com/mkuchin)
-   <b><code>&nbsp;&nbsp;3337⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;595🍴</code></b> [docker-swarm-visualizer](https://github.com/dockersamples/docker-swarm-visualizer)) - Visualizes Docker services on a Docker Swarm (for running demos).
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [dockering-on-rails](https://github.com/Electrofenster/dockerding-on-rails)) :skull: - Simple Web-Interface for Docker with a lot of features by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Electrofenster](https://github.com/Electrofenster/))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [DockerSurfer](https://github.com/Simone-Erba/DockerSurfer)) :skull: - A web service for analyze and browse dependencies between Docker images in the Docker registry, by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Simone-Erba](https://github.com/Simone-Erba/))
-   <b><code>&nbsp;20011⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;619🍴</code></b> [dockge](https://github.com/louislam/dockge)) - easy-to-use and reactive self-hosted docker compose.yaml stack-oriented manager by [@louislam](https://github.com/louislam).
-   <b><code>&nbsp;&nbsp;8122⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;210🍴</code></b> [Komodo](https://github.com/mbecker20/komodo)) - A tool to build and deploy software on many servers
-   <b><code>&nbsp;&nbsp;1683⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;95🍴</code></b> [Kubevious](https://github.com/kubevious/kubevious)) - A highly visual web UI for Kubernetes which renders configuration and state in an application centric way by [@rubenhak](https://github.com/rubenhak).
-   <b><code>&nbsp;&nbsp;&nbsp;602⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;44🍴</code></b> [Mafl](https://github.com/hywax/mafl)) - Minimalistic flexible homepage by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@hywax](https://github.com/hywax/))
-   <b><code>&nbsp;76489⭐</code></b> <b><code>&nbsp;&nbsp;6206🍴</code></b> [netdata](https://github.com/netdata/netdata)) - Real-time performance monitoring
-   <b><code>&nbsp;&nbsp;5328⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;280🍴</code></b> [OctoLinker](https://github.com/OctoLinker/OctoLinker)) - A browser extension for GitHub that makes the image name in a `Dockerfile` clickable and redirect you to the related Docker Hub page.
-   <b><code>&nbsp;34916⭐</code></b> <b><code>&nbsp;&nbsp;2702🍴</code></b> [Portainer](https://github.com/portainer/portainer)) - A lightweight management UI for managing your Docker hosts or Docker Swarm clusters by [@portainer](https://github.com/portainer)
-   <b><code>&nbsp;&nbsp;&nbsp;145⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24🍴</code></b> [Rapid Dashboard](https://github.com/ozlerhakan/rapid)) - A simple query dashboard to use Docker Remote API by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@ozlerhakan](https://github.com/ozlerhakan/))
-   <b><code>&nbsp;&nbsp;1935⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;266🍴</code></b> [Seagull](https://github.com/tobegit3hub/seagull)) - Friendly Web UI to monitor docker daemon. by [@tobegit3hub](https://github.com/tobegit3hub)
-   <b><code>&nbsp;&nbsp;3328⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;300🍴</code></b> [Swarmpit](https://github.com/swarmpit/swarmpit)) - Swarmpit provides simple and easy to use interface for your Docker Swarm cluster. You can manage your stacks, services, secrets, volumes, networks etc.
-   <b><code>&nbsp;&nbsp;&nbsp;659⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;82🍴</code></b> [Swirl](https://github.com/cuigh/swirl)) - Swirl is a web management tool for Docker, focused on swarm cluster By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@cuigh](https://github.com/cuigh/))
-   <b><code>&nbsp;21090⭐</code></b> <b><code>&nbsp;&nbsp;2713🍴</code></b> [Theia](https://github.com/eclipse-theia/theia)) - Extensible platform to develop full-fledged multi-language Cloud & Desktop IDE-like products with state-of-the-art web technologies.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Yacht](https://github.com/SelfhostedPro/Yacht)) :construction: - A Web UI for docker that focuses on templates and ease of use in order to make deployments as easy as possible. By [@SelfhostedPro](https://github.com/SelfhostedPro)

## Docker Images

### Base Tools

Tools and applications that are either installed inside containers or designed to be run as a 🌎 [sidecar](learn.microsoft.com/en-us/azure/architecture/patterns/sidecar)

-   <b><code>&nbsp;&nbsp;1054⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;72🍴</code></b> [amicontained](https://github.com/genuinetools/amicontained)) - Container introspection tool. Find out what container runtime is being used as well as features available by [@genuinetools][genuinetools]
-   <b><code>&nbsp;&nbsp;&nbsp;177⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31🍴</code></b> [Chaperone](https://github.com/garywiz/chaperone)) - A single PID1 process designed for docker containers. Does user management, log management, startup, zombie reaping, all in one small package. by [@garywiz](https://github.com/garywiz)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;56⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [ckron](https://github.com/nicomt/ckron)) - A cron-style job scheduler for docker, by [@nicomt](https://github.com/nicomt)
-   [CoreOS][coreos] - Linux for Massive Server Deployments
-   <b><code>&nbsp;21484⭐</code></b> <b><code>&nbsp;&nbsp;1319🍴</code></b> [distroless](https://github.com/GoogleContainerTools/distroless)) - Language focused docker images, minus the operating system, by [@GoogleContainerTools][googlecontainertools]
-   <b><code>&nbsp;&nbsp;5709⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;531🍴</code></b> [docker-alpine](https://github.com/gliderlabs/docker-alpine)) - A super small Docker base image _(5MB)_ using Alpine Linux by [@gliderlabs][gliderlabs]
-   <b><code>&nbsp;&nbsp;4589⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;616🍴</code></b> [docker-gen](https://github.com/jwilder/docker-gen)) - Generate files from docker container meta-data by [@jwilder][jwilder]
-   <b><code>&nbsp;&nbsp;&nbsp;191⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [dockerize](https://github.com/powerman/dockerize)) - Utility to simplify running applications in docker containers by [@jwilder][jwilder], [@powerman][powerman]
-   <b><code>&nbsp;&nbsp;4888⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;347🍴</code></b> [GoSu](https://github.com/tianon/gosu)) - Run this specific application as this specific user and get out of the pipeline (entrypoint script tool) by [@tianon](https://github.com/tianon)
-   <b><code>&nbsp;&nbsp;&nbsp;227⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [is-docker](https://github.com/sindresorhus/is-docker)) - Check if the process is running inside a Docker container by [@sindresorhus][sindresorhus]
-   <b><code>&nbsp;&nbsp;&nbsp;338⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [lstags](https://github.com/ivanilves/lstags)) - sync Docker images across registries by [@ivanilves](https://github.com/ivanilves)
-   <b><code>&nbsp;17430⭐</code></b> <b><code>&nbsp;&nbsp;2042🍴</code></b> [NVIDIA-Docker](https://github.com/NVIDIA/nvidia-docker)) - :skull: The NVIDIA Container Runtime for Docker by [@NVIDIA][nvidia]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Ofelia](https://github.com/mcuadros/ofelia/)) - Ofelia is a modern and low footprint job scheduler for docker environments, built on Go. Ofelia aims to be a replacement for the old fashioned cron. Supports configuration from container labels and/or configuration files.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;18⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [SparkView](https://github.com/beyondssl/sparkview-container)) - Access VMs, desktops, servers or applications anytime and from anywhere, without complex and costly client roll-outs or user management.
-   <b><code>&nbsp;&nbsp;1002⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;103🍴</code></b> [su-exec](https://github.com/ncopa/su-exec)) - This is a simple tool that will simply execute a program with different privileges. The program will be executed directly and not run as a child, like su and sudo does, which avoids TTY and signal issues. Why reinvent gosu? This does more or less exactly the same thing as gosu but it is only 10kb instead of 1.8MB. By [ncopa](https://github.com/ncopa)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [sue](https://github.com/theAkito/sue)) - Executes a program as a user different from the user running sue. This is a maintainable alternative to ncopa/su-exec, which is the better tianon/gosu. This one is far better (higher performance, smaller size), than the original gosu, however it is far easier to maintain, than su-exec, which is written in plain C. Made by [Akito][akito]
-   <b><code>&nbsp;&nbsp;2237⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;135🍴</code></b> [supercronic](https://github.com/aptible/supercronic)) - crontab-compatible job runner, designed specifically to run in containers by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@aptible](https://github.com/aptible/))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;30⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [TrivialRC](https://github.com/vorakl/TrivialRC)) - A minimalistic Runtime Configuration system and process manager for containers [@vorakl](https://github.com/vorakl)

### Builder

Applications designed to help or simplify building **new** images

-   <b><code>&nbsp;&nbsp;&nbsp;690⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;74🍴</code></b> [ansible-bender](https://github.com/ansible-community/ansible-bender)) - A tool utilising `ansible` and `buildah` by [@TomasTomecek][tomastomecek]
-   <b><code>&nbsp;&nbsp;8368⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;849🍴</code></b> [buildah](https://github.com/containers/buildah)) - A tool that facilitates building OCI images by [@containers][containers]
-   <b><code>&nbsp;&nbsp;9461⭐</code></b> <b><code>&nbsp;&nbsp;1315🍴</code></b> [BuildKit](https://github.com/moby/buildkit)) - Concurrent, cache-efficient, and Dockerfile-agnostic builder toolkit by [@moby project](https://github.com/moby)
-   <b><code>&nbsp;&nbsp;&nbsp;110⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;43🍴</code></b> [cekit](https://github.com/cekit/cekit)) - A tool used by openshift to build base images using different build engines by [@cekit](https://github.com/cekit).
-   <b><code>&nbsp;&nbsp;3792⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;236🍴</code></b> [container-diff](https://github.com/GoogleContainerTools/container-diff)) - :skull: An image tool for comparing and analyzing container images by [@GoogleContainerTools][googlecontainertools]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [container-factory](https://github.com/mutable/container-factory)) - Produces Docker images from tarballs of application source code by [@mutable](https://github.com/mutable)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [copy-docker-image](https://github.com/mdlavin/copy-docker-image)) - Copy a Docker image between registries without a full Docker installation by [@mdlavin](https://github.com/mdlavin)
-   <b><code>&nbsp;&nbsp;&nbsp;693⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;110🍴</code></b> [Derrick](https://github.com/alibaba/derrick)) - A tool help you to automate the generation of Dockerfile and dockerize application by scanning the code. By [@alibaba](https://github.com/alibaba).
-   <b><code>&nbsp;&nbsp;&nbsp;442⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [dlayer](https://github.com/orisano/dlayer)) - docker layer analyzer by [@orisano](https://github.com/orisano)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;47⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [docker-companion](https://github.com/mudler/docker-companion)) - A command line tool written in Golang to squash and unpack docker images by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mudler](https://github.com/mudler/))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;98⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [docker-make](https://github.com/CtripCloud/docker-make)) - Build, tag,and push a bunch of related docker images via a single command.
-   <b><code>&nbsp;&nbsp;&nbsp;202⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [docker-replay](https://github.com/bcicen/docker-replay)) - Generate `docker run`command and options from running containers. By [bcicen](https://github.com/bcicen)
-   <b><code>&nbsp;&nbsp;&nbsp;147⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker-repack](https://github.com/orf/docker-repack)) - Repacks a Docker image into a smaller, more efficient version that makes it significantly faster to pull. By [orf](https://github.com/orf)
-   <b><code>&nbsp;&nbsp;&nbsp;139⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [DockerMake](https://github.com/avirshup/DockerMake)) - A reproducible Docker image build system for complex software stacks. By [@avirshup](https://github.com/avirshup)
-   <b><code>&nbsp;22425⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;792🍴</code></b> [DockerSlim](https://github.com/docker-slim/docker-slim)) shrinks fat Docker images creating the smallest possible images.
-   <b><code>&nbsp;&nbsp;&nbsp;226⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Dockly](https://github.com/swipely/dockly)) - Dockly is a gem made to ease the pain of packaging an application in Docker by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@swipely](https://github.com/swipely/))
-   <b><code>&nbsp;&nbsp;&nbsp;258⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [dockramp](https://github.com/jlhawn/dockramp)) :skull: - Proof of Concept: A Client Driven Docker Image Builder by [@jlhawn](https://github.com/jlhawn)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [essex](https://github.com/utensils/essex)) - Boilerplate for Docker Based Projects: Essex is a CLI utility written in bash to quickly setup clean and consistent Docker projects with Makefile driven workflows. [@jamesbrink](https://github.com/jamesbrink)
-   <b><code>&nbsp;&nbsp;&nbsp;496⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;101🍴</code></b> [HPC Container Maker](https://github.com/NVIDIA/hpc-container-maker)) - Generates Dockerfiles from a high level Python recipe, including building blocks for High-Performance Computing components by [@NVIDIA][nvidia]
-   <b><code>&nbsp;&nbsp;3975⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;234🍴</code></b> [img](https://github.com/genuinetools/img)) - Standalone, daemon-less, unprivileged Dockerfile and OCI compatible container image builder by [@genuinetools][genuinetools]
-   <b><code>&nbsp;15700⭐</code></b> <b><code>&nbsp;&nbsp;1523🍴</code></b> [kaniko](https://github.com/GoogleContainerTools/kaniko)) - Build Container Images In Kubernetes. By [@GoogleContainerTools][googlecontainertools]
-   <b><code>&nbsp;&nbsp;2399⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;155🍴</code></b> [makisu](https://github.com/uber/makisu)) - :skull: Uber's fast and flexible unprivileged image builder for Mesos and Kubernetes, with distributed cache support. By [@uber](https://github.com/uber)
-   🌎 [packer](developer.hashicorp.com/packer/integrations/hashicorp/docker/latest/components/builder/docker) - Hashicorp tool to build machine images including docker image integrated with configuration management tools like chef, puppet, ansible
-   <b><code>&nbsp;&nbsp;&nbsp;132⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [portainer](https://github.com/duedil-ltd/portainer)) - Apache Mesos framework for building Docker images by [@duedil-ltd](https://github.com/duedil-ltd)
-   🌎 [Production-Ready Python Containers :heavy_dollar_sign:](pythonspeed.com/products/pythoncontainer/) - A template for creating production-ready Docker images for Python applications.
-   <b><code>&nbsp;&nbsp;&nbsp;557⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34🍴</code></b> [RAUDI](https://github.com/cybersecsi/RAUDI)) - A tool to automatically update (and optionally push to Docker Hub) Docker Images for 3rd party software whenever theres is a new release/update/commit. By [@SecSI](https://github.com/cybersecsi)
-   <b><code>&nbsp;&nbsp;2858⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;159🍴</code></b> [runlike](https://github.com/lavie/runlike)) - Generate `docker run`command and options from running containers by [@lavie](https://github.com/lavie)
-   <b><code>&nbsp;&nbsp;&nbsp;183⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [SkinnyWhale](https://github.com/djosephsen/skinnywhale)) :skull: - Skinnywhale helps you make smaller (as in megabytes) Docker containers.
-   <b><code>&nbsp;&nbsp;&nbsp;601⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36🍴</code></b> [Smith](https://github.com/oracle/smith)) - :skull: A Micocontainer Builder and can perform multi-stage builds after the image is built [Oracle][oracle]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [userdef](https://github.com/theAkito/userdef)) - An advanced `adduser` for your Alpine based Docker images. Made by [Akito][akito]
-   <b><code>&nbsp;&nbsp;1145⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;98🍴</code></b> [Whaler](https://github.com/P3GLEG/Whaler)) - Program to reverse Docker images into Dockerfiles by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@P3GLEG](https://github.com/P3GLEG/)).
-   <b><code>&nbsp;&nbsp;&nbsp;392⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [Whales](https://github.com/Gueils/whales)) - A tool to automatically dockerize your applications by [@icalialabs](https://github.com/IcaliaLabs).

### Dockerfile

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;65⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [chaperone-docker](https://github.com/garywiz/chaperone-docker)) - A set of images using the Chaperone process manager, including a lean Alpine image, LAMP, LEMP, and bare-bones base kits.
-   <b><code>&nbsp;&nbsp;&nbsp;185⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Dockerfile Generator](https://github.com/ozankasikci/dockerfile-generator)) `dfg` is both a Go library and an executable that produces valid Dockerfiles using various input channels.
-   🌎 [Dockerfile Project](dockerfile.github.io/) - Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.
-   <b><code>&nbsp;&nbsp;&nbsp;237⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [dockerfilegraph](https://github.com/patrickhoefler/dockerfilegraph)) - Visualize your multi-stage Dockerfiles. By [@PatrickHoefler](https://github.com/patrickhoefler)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [Dockershelf](https://github.com/Dockershelf/dockershelf)) - A repository that serves as a collector for docker recipes that are universal, efficient and slim. Images are updated, tested and published daily via a Travis cron job. Maintained by [@CollageLabs](https://github.com/CollageLabs).
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [dockmoor](https://github.com/MeneDev/dockmoor)) :skull: - Manage docker image references and help to create reproducible builds with Docker. By [@MeneDev](https://github.com/MeneDev)
-   [Vektorcloud](https://github.com/vektorcloud) - A collection of minimal, Alpine-based Docker images

Examples by:

-   🌎 [@0xy](gitlab.com/0xy/dockerfiles)
-   <b><code>&nbsp;&nbsp;&nbsp;250⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;217🍴</code></b> [@arun-gupta](https://github.com/arun-gupta/docker-images))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;64⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [@awesome-startup](https://github.com/awesome-startup/docker-compose))
-   <b><code>&nbsp;&nbsp;&nbsp;299⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62🍴</code></b> [@crosbymichael](https://github.com/crosbymichael/Dockerfiles))
-   <b><code>&nbsp;13915⭐</code></b> <b><code>&nbsp;&nbsp;2544🍴</code></b> [@jessfraz](https://github.com/jessfraz/dockerfiles))
-   <b><code>&nbsp;&nbsp;&nbsp;580⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;487🍴</code></b> [@komljen](https://github.com/komljen/dockerfile-examples))
-   <b><code>&nbsp;&nbsp;&nbsp;825⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;721🍴</code></b> [@kstaken](https://github.com/kstaken/dockerfile-examples))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [@ondrejmo](https://github.com/ondrejmo/Dockerfiles))
-   <b><code>&nbsp;&nbsp;3197⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;796🍴</code></b> [@vimagick](https://github.com/vimagick/dockerfiles))

### Linter

-   <b><code>&nbsp;&nbsp;&nbsp;129⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [docker-image-size-limit](https://github.com/wemake-services/docker-image-size-limit)) - A tool to keep an eye on your docker images size.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;44⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Dockerfile Linter action](https://github.com/buddy-works/dockerfile-linter)) - The linter lets you verify Dockerfile syntax to make sure it follows the best practices for building efficient Docker images.
-   <b><code>&nbsp;&nbsp;&nbsp;439⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [dockfmt](https://github.com/jessfraz/dockfmt)) :construction: - Dockerfile formatter and parser by [@jessfraz][jessfraz]
-   <b><code>&nbsp;&nbsp;1022⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;83🍴</code></b> [FROM:latest](https://github.com/replicatedhq/dockerfilelint)) - An opinionated Dockerfile linter by [@replicatedhq](https://github.com/replicatedhq)
-   <b><code>&nbsp;11607⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;473🍴</code></b> [Hadolint](https://github.com/hadolint/hadolint)) - A Dockerfile linter that checks for best practices, common mistakes, and is also able to lint any bash written in `RUN` instructions; by [@lukasmartinelli](https://github.com/lukasmartinelli)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [Whale-linter](https://github.com/jeromepin/whale-linter)) - :skull: A simple and small Dockerfile linter written in Python3+ without dependencies by [@jeromepin](https://github.com/jeromepin)

### Metadata

-   <b><code>&nbsp;&nbsp;3967⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;767🍴</code></b> [opencontainer](https://github.com/opencontainers/image-spec/blob/master/annotations.md)) - A convention and shared namespace for Docker labels defined by OCI Image Spec.

### Registry

Services to securely store your Docker images.

-   🌎 [Amazon Elastic Container Registry :heavy_dollar_sign:](aws.amazon.com/ecr/) - Amazon Elastic Container Registry (ECR) is a fully-managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images.
-   🌎 [Azure Container Registry :heavy_dollar_sign:](azure.microsoft.com/en-us/products/container-registry/#overview) - Manage a Docker private registry as a first-class Azure resource
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [CargoOS](https://github.com/RedCoolBeans/cargos-buildroot)) - A bare essential OS for running the Docker Engine on bare metal or Cloud. By [@RedCoolBeans](https://github.com/RedCoolBeans)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;59⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;18🍴</code></b> [cleanreg](https://github.com/hcguersoy/cleanreg)) - A small tool to delete image manifests from a Docker Registry implementing the API v2, dereferencing them for the GC by [@hcguersoy](https://github.com/hcguersoy)
-   🌎 [Cloudsmith :heavy_dollar_sign:](cloudsmith.com/product/formats/docker-registry) - A fully managed package management SaaS, with first-class support for public and private Docker registries (and many others, incl. Helm charts for the Kubernetes ecosystem). Has a generous free-tier and is also completely free for open-source.
-   🌎 [Container Registry Service :heavy_dollar_sign:](container-registry.com/) - Harbor based Container Management Solution as a Service for teams and organizations. Free tier offers 1 GB storage for private repositories.
-   🌎 [Cycle.io :heavy_dollar_sign:](cycle.io/) - Bare-metal container hosting.
-   🌎 [DigitalOcean :heavy_dollar_sign:](www.digitalocean.com/products/container-registry) - DigitalOcean Container Registry.
-   🌎 [Docker Hub](hub.docker.com/) provided by Docker Inc.
-   [Docker Registry v2][distribution] - The Docker toolset to pack, ship, store, and deliver content
-   <b><code>&nbsp;&nbsp;&nbsp;706⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34🍴</code></b> [Docket](https://github.com/netvarun/docket)) - Custom docker registry that allows for lightning fast deploys through bittorrent by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@netvarun](https://github.com/netvarun/))
-   <b><code>&nbsp;&nbsp;2831⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;348🍴</code></b> [Dragonfly](https://github.com/dragonflyoss/Dragonfly2)) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology.
-   🌎 [GCP Artifact Registry :heavy_dollar_sign:](cloud.google.com/artifact-registry/docs) Fast, private Docker image storage on Google Cloud Platform.
-   🌎 [Gitea Container Registry](docs.gitea.com/usage/packages/container) - Integrated Docker registry in Gitea, ideal for private, small-scale image hosting.
-   🌎 [GitHub Container Registry](docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) - GitHub's solution for storing and managing Docker images, with tight integration into GitHub Actions.
-   🌎 [GitLab Container Registry](docs.gitlab.com/user/packages/container_registry/) - Registry focused on using its images in GitLab CI
-   <b><code>&nbsp;26718⭐</code></b> <b><code>&nbsp;&nbsp;4988🍴</code></b> [Harbor](https://github.com/goharbor/harbor)) An open source trusted cloud native registry project that stores, signs, and scans content. Supports replication, user management, access control and activity auditing. By 🌎 [CNCF](www.cncf.io) formerly [VMWare][vmware]
-   🌎 [JFrog Artifactory :heavy_dollar_sign:](jfrog.com/artifactory/) - Artifact Repository Manager, can be used as private Docker Registry as well
-   <b><code>&nbsp;&nbsp;6570⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;459🍴</code></b> [Kraken](https://github.com/uber/kraken)) - Uber's Highly scalable P2P docker registry, capable of distributing TBs of data in seconds.
-   🌎 [Quay.io :heavy_dollar_sign:](quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [Registryo](https://github.com/inmagik/registryo)) - UI and token based authentication server for onpremise docker registry
-   🌎 [RepoFlow](www.repoflow.io) - A simple and easy-to-use package management platform with Docker support alongside other formats like PyPI, Maven, npm, and Helm. Includes smart search, built-in Docker image scanning, and a great free option for both self-hosted and cloud use.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Rescoyl](https://github.com/noteed/rescoyl)) - Private Docker registry (free and open source) by [@noteed](https://github.com/noteed)
-   🌎 [Sonatype Nexus Repository](www.sonatype.com/products/sonatype-nexus-repository) - Manage binaries and build artifacts across your software supply chain.

## Development with Docker

### API Client

-   <b><code>&nbsp;&nbsp;&nbsp;137⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [ahab](https://github.com/instacart/ahab)) - Docker event handling with Python by [@instacart](https://github.com/instacart)
-   <b><code>&nbsp;&nbsp;&nbsp;178⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [clj-docker-client](https://github.com/into-docker/clj-docker-client)) :skull: - Idiomatic Clojure client for the Docker remote API. By [@lispyclouds][lispyclouds]
-   <b><code>&nbsp;&nbsp;&nbsp;143⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [contajners](https://github.com/lispyclouds/contajners)) - An idiomatic, data-driven, REPL friendly Clojure client for OCI container engines. By [@lispyclouds][lispyclouds]
-   <b><code>&nbsp;&nbsp;&nbsp;118⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;28🍴</code></b> [Docker Client for JVM](https://github.com/gesellix/docker-client)) - A Docker remote api client library for the JVM, written in Groovy by [@gesellix][gesellix]
-   🌎 [Docker Client TypeScript](gitlab.com/masaeedu/docker-client) - Docker API client for JavaScript, automatically generated from Swagger API definition from moby repository. By [@masaeedu](https://github.com/masaeedu)
-   <b><code>&nbsp;&nbsp;1440⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;543🍴</code></b> [docker-client](https://github.com/spotify/docker-client)) :skull: - Java client for the Docker remote API. By [@spotify][spotify]
-   <b><code>&nbsp;&nbsp;&nbsp;213⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [docker-controller-bot](https://github.com/dgongut/docker-controller-bot)) - Telegram bot to control docker containers. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@dgongut](https://github.com/dgongut/))
-   <b><code>&nbsp;&nbsp;&nbsp;432⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;90🍴</code></b> [docker-it-scala](https://github.com/whisklabs/docker-it-scala)) - Docker integration testing kit with Scala by [@whisklabs](https://github.com/whisklabs)
-   <b><code>&nbsp;&nbsp;&nbsp;272⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;55🍴</code></b> [docker-java-api](https://github.com/amihaiemil/docker-java-api)) - Lightweight, truly object-oriented, Java client for Docker's API. By [@amihaiemil](https://github.com/amihaiemil)
-   <b><code>&nbsp;&nbsp;1920⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;650🍴</code></b> [docker-maven-plugin](https://github.com/fabric8io/docker-maven-plugin)) - A Maven plugin for running and creating Docker images by [@fabric8io](https://github.com/fabric8io)
-   <b><code>&nbsp;&nbsp;&nbsp;286⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;87🍴</code></b> [Docker-PowerShell](https://github.com/Microsoft/Docker-PowerShell)) - :skull: PowerShell Module for Docker
-   <b><code>&nbsp;&nbsp;2374⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;402🍴</code></b> [Docker.DotNet](https://github.com/Microsoft/Docker.DotNet)) - C#/.NET HTTP client for the Docker remote API by [@ahmetb](https://github.com/ahmetb)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;42⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Docker.Registry.DotNet](https://github.com/ChangemakerStudios/Docker.Registry.DotNet)) - .NET (C#) Client Library for interacting with a Docker Registry API (v2) [@rquackenbush](https://github.com/rquackenbush)
-   <b><code>&nbsp;&nbsp;2750⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;485🍴</code></b> [dockerfile-maven](https://github.com/spotify/dockerfile-maven)) - :skull: A Maven plugin for building and pushing Docker images by [@spotify][spotify]
-   <b><code>&nbsp;&nbsp;4690⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;480🍴</code></b> [dockerode](https://github.com/apocas/dockerode)) - Docker Remote API node.js module by [@apocas](https://github.com/apocas)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;76⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [DoMonit](https://github.com/eon01/DoMonit)) - A simple Docker Monitoring wrapper For Docker API
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [go-dockerclient](https://github.com/fsouza/go-dockerclient/)) - Go HTTP client for the Docker remote API by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@fsouza](https://github.com/fsouza/))
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin)) - A Docker remote api plugin for Gradle by [@gesellix][gesellix]
-   <b><code>&nbsp;&nbsp;&nbsp;584⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;189🍴</code></b> [libcompose](https://github.com/docker/libcompose)) - :skull: Go library for Docker Compose.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;75⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [Portainer stack utils](https://github.com/greenled/portainer-stack-utils)) :construction: - Bash script to deploy/update/undeploy Docker stacks in a Portainer instance from a docker-compose yaml file. By [@greenled](https://github.com/greenled).
-   <b><code>&nbsp;&nbsp;&nbsp;178⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36🍴</code></b> [sbt-docker-compose](https://github.com/Tapad/sbt-docker-compose)) - :skull: Integrates Docker Compose functionality into sbt by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@kurtkopchik](https://github.com/kurtkopchik/))
-   <b><code>&nbsp;&nbsp;&nbsp;735⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;111🍴</code></b> [sbt-docker](https://github.com/marcuslonnberg/sbt-docker)) - Create Docker images directly from sbt by [@marcuslonnberg](https://github.com/marcuslonnberg)

### CI/CD

-   🌎 [Buddy :heavy_dollar_sign:](buddy.works) - The best of Git, build & deployment tools combined into one powerful tool that supercharged our development.
-   <b><code>&nbsp;&nbsp;&nbsp;774⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;48🍴</code></b> [Captain](https://github.com/harbur/captain)) - Convert your Git workflow to Docker containers ready for Continuous Delivery by [@harbur](https://github.com/harbur).
-   <b><code>&nbsp;&nbsp;1072⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;170🍴</code></b> [Cyclone](https://github.com/caicloud/cyclone)) - Powerful workflow engine and end-to-end pipeline solutions implemented with native Kubernetes resources by [@caicloud](https://github.com/caicloud).
-   <b><code>&nbsp;&nbsp;&nbsp;143⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [Defang](https://github.com/DefangLabs/defang)) - Deploy Docker Compose to your favorite cloud in minutes by [@DefangLabs](https://github.com/DefangLabs)
-   🌎 [Depot :heavy_dollar_sign:](depot.dev) - Build Docker images fast, in the cloud. Blazing fast compute, automatic intelligent caching, and zero configuration. 🌎 [Done in seconds](depot.dev/#benchmarks).
-   <b><code>&nbsp;&nbsp;4067⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;140🍴</code></b> [Diun](https://github.com/crazy-max/diun)) - Receive notifications when an image or repository is updated on a Docker registry by [@crazy-max].
-   <b><code>&nbsp;&nbsp;1859⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;65🍴</code></b> [dockcheck](https://github.com/mag37/dockcheck)) - A script checking updates for docker images without pulling then auto-update selected/all containers. With notifications, pruning and more.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/)) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
-   <b><code>&nbsp;33525⭐</code></b> <b><code>&nbsp;&nbsp;2875🍴</code></b> [Drone](https://github.com/drone/drone)) - Continuous integration server built on Docker and configured using YAML files.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;80⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [Gantry](https://github.com/shizunge/gantry)) - Automatically update selected Docker swarm services.
-   🌎 [GitLab Runner](gitlab.com/gitlab-org/gitlab-runner) - GitLab has integrated CI to test, build and deploy your code with the use of GitLab runners.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;71🍴</code></b> [GOCD-Docker](https://github.com/gocd/gocd-docker)) :skull: - Go Server and Agent in docker containers to provision.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Jaypore CI](https://github.com/theSage21/jaypore_ci)) - Simple, very flexible, powerful CI / CD / automation system configured in Python. Offline and local first.
-   <b><code>&nbsp;&nbsp;&nbsp;159⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [Kraken CI](https://github.com/Kraken-CI/kraken)) - Modern CI/CD, open-source, on-premise system that is highly scalable and focused on testing. One of its executors is Docker. Developed by [@Kraken-CI](https://github.com/Kraken-CI).
-   <b><code>&nbsp;&nbsp;&nbsp;144⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38🍴</code></b> [Microservices Continuous Deployment](https://github.com/francescou/docker-continuous-deployment)) - Continuous deployment of a microservices application.
-   <b><code>&nbsp;&nbsp;&nbsp;968⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;132🍴</code></b> [mu](https://github.com/stelligent/mu)) - Tool to configure CI/CD of your container applications via AWS CodePipeline, CodeBuild and ECS [@Stelligent](https://github.com/stelligent)
-   <b><code>&nbsp;&nbsp;1611⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;157🍴</code></b> [Ouroboros](https://github.com/pyouroboros/ouroboros)) :skull: - Automatically update running Docker containers with notifications
-   <b><code>&nbsp;&nbsp;&nbsp;303⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62🍴</code></b> [Popper](https://github.com/systemslab/popper)) - Github actions workflow (HCL syntax) execution engine.
-   🌎 [Screwdriver :heavy_dollar_sign:](screwdriver.cd/) - Yahoo's OpenSource buildplatform designed for Continous Delivery.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;49⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23🍴</code></b> [Skipper](https://github.com/Stratoscale/skipper)) - Easily dockerize your Git repository by [@Stratoscale](https://github.com/Stratoscale)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;57⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [SwarmCI](https://github.com/ghostsquad/swarmci)) - Create a distributed, isolated task pipeline in your Docker Swarm.
-   🌎 [Tekton CD](tekton.dev/) - A cloud-native pipeline resource.
-   <b><code>&nbsp;23903⭐</code></b> <b><code>&nbsp;&nbsp;1029🍴</code></b> [Watchtower](https://github.com/containrrr/watchtower)) - Automatically update running Docker containers

### Development Environment

-   <b><code>&nbsp;&nbsp;&nbsp;680⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;50🍴</code></b> [batect](https://github.com/batect/batect)) - :skull: build and testing environments as code tool: Dockerised build and testing environments made easy by [@charleskorn](https://github.com/charleskorn)
-   <b><code>&nbsp;&nbsp;&nbsp;672⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [Binci](https://github.com/binci/binci)) - Containerize your development workflow. (formerly DevLab by [@TechnologyAdvice](https://github.com/TechnologyAdvice))
-   <b><code>&nbsp;&nbsp;8306⭐</code></b> <b><code>&nbsp;&nbsp;1281🍴</code></b> [Boot2Docker](https://github.com/boot2docker/boot2docker)) :skull: - Docker for OSX and Windows
-   <b><code>&nbsp;11299⭐</code></b> <b><code>&nbsp;&nbsp;1050🍴</code></b> [coder](https://github.com/coder/coder)) - remote development machines powered by Terraform or Docker by [@coder](https://github.com/coder)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [construi](https://github.com/lstephen/construi)) - Run your builds inside a Docker defined environment by [@lstephen](https://github.com/lstephen)
-   <b><code>&nbsp;&nbsp;&nbsp;272⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [Crashcart](https://github.com/oracle/crashcart)) - :skull: Sideload Linux binaries into a running container for troubleshooting by [@Oracle][oracle]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;44⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [dde](https://github.com/whatwedo/dde)) :construction: - Local development environment toolset based on Docker. By [@whatwedo](https://github.com/whatwedo)
-   <b><code>&nbsp;&nbsp;&nbsp;195⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [Devstep](https://github.com/fgrehm/devstep)) :skull: - Development environments powered by Docker and buildpacks by [@fgrehm][fgrehm]
-   <b><code>&nbsp;&nbsp;2113⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;107🍴</code></b> [Dinghy](https://github.com/codekitchen/dinghy)) - :skull: An alternative way to use Docker on Mac OS X using Docker Machine with virtualbox, vmware, xhyve or parallels
-   <b><code>&nbsp;&nbsp;1311⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;49🍴</code></b> [DIP](https://github.com/bibendi/dip)) - CLI utility for straightforward provisioning and interacting with an application configured by docker-compose. By [@bibendi](https://github.com/bibendi)
-   <b><code>&nbsp;&nbsp;2328⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;53🍴</code></b> [DLite](https://github.com/nlf/dlite)) :skull: - Simplest way to use Docker on OSX, no VM needed. By [@nlf](https://github.com/nlf)
-   <b><code>&nbsp;&nbsp;&nbsp;315⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35🍴</code></b> [dobi](https://github.com/dnephin/dobi)) - A build automation tool for Docker applications. By [@dnephin](https://github.com/dnephin)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;29⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [Docker Missing Tools](https://github.com/nandoquintana/docker-missing-tools)) - A set of bash commands to shortcut typical docker dev-ops. An alternative to creating typical helper scripts like "build.sh" and "deploy.sh" inside code repositories. By [@NandoQuintana](https://github.com/nandoquintana).
-   <b><code>&nbsp;&nbsp;1432⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;101🍴</code></b> [Docker osx dev](https://github.com/brikis98/docker-osx-dev)) :skull: - A productive development environment with Docker on OS X by [@brikis98](https://github.com/brikis98)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Docker-Arch](https://github.com/Ph3nol/Docker-Arch)) - Generate Web/CLI projects Dockerized development environments, from 1 simple YAML file. By [@Ph3nol](https://github.com/ph3nol)
-   <b><code>&nbsp;&nbsp;3563⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;285🍴</code></b> [Docker-sync](https://github.com/EugenMayer/docker-sync)) - Drastically improves performance (<b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [50-70x](https://github.com/EugenMayer/docker-sync/wiki/4.-Performance))) when using Docker for development on Mac OS X/Windows and Linux while sharing code to the container. By [@EugenMayer](https://github.com/EugenMayer)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;43⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [docker-vm](https://github.com/shyiko/docker-vm)) - Simple and transparent alternative to boot2docker (backed by Vagrant) by [@shyiko](https://github.com/shyiko)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [DockerBuildManagement](https://github.com/DIPSAS/DockerBuildManagement)) - :skull: Build Management is a python application, installed with pip. The application makes it easy to manage a build system based on Docker by configuring a single yaml file describing how to build, test, run or publish a containerized solution.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;82⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [DockerDL](https://github.com/matifali/dockerdl)) - Deep Learning Docker Images. Don't waste time setting up a deep learning env when you can get a deep learning environment with everything pre-installed.
-   <b><code>&nbsp;&nbsp;&nbsp;369⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Dusty](https://github.com/gamechanger/dusty)) - :skull: Managed Docker development environments on OS X
-   <b><code>&nbsp;&nbsp;7085⭐</code></b> <b><code>&nbsp;&nbsp;1197🍴</code></b> [Eclipse Che](https://github.com/eclipse/che)) - Developer workspace server with Docker runtimes, cloud IDE, next-generation Eclipse IDE
-   <b><code>&nbsp;&nbsp;&nbsp;112⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [EnvCLI](https://github.com/EnvCLI/EnvCLI)) - Replace your local installation of Node, Go, ... with project-specific docker containers. By [@EnvCLI](https://github.com/EnvCLI)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;67⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [ESP32 Linux - Docker builder](https://github.com/hpsaturn/esp32s3-linux)) - Container solution to compile Linux and develop it for ESP32 microcontrollers - By [@Hpsaturn](https://github.com/hpsaturn)
-   <b><code>&nbsp;&nbsp;1591⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;120🍴</code></b> [footloose](https://github.com/weaveworks/footloose)) - :skull: Spin containers that look like Virtual Machines - By [@dlespiau](https://github.com/dlespiau)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;77⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [forward2docker](https://github.com/bsideup/forward2docker)) :skull: - Utility to auto forward a port from localhost into ports on Docker containers running in a boot2docker VM by [@bsideup](https://github.com/bsideup)
-   <b><code>&nbsp;&nbsp;&nbsp;635⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [Gebug](https://github.com/moshebe/gebug)) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Kitt](https://github.com/senges/kitt)) - A portable and disposable Shell environment, based on Docker and Nix. By [@senges](https://github.com/senges)
-   <b><code>&nbsp;&nbsp;4209⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;531🍴</code></b> [Lando](https://github.com/lando/lando)) - Lando is for developers who want to quickly specify and painlessly spin up the services and tools needed to develop their projects. By 🌎 [Tandem](www.thinktandem.io/)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;33⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Rust Universal Compiler](https://github.com/Peco602/rust-universal-compiler)) - Container solution to compile Rust projects for Linux, macOS and Windows. By [@Peco602][peco602]
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;18⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [uniget](https://github.com/uniget-org/cli)) - uni(versal)get, the installer and updater for container tools and beyond (formerly docker-setup). By [@nicholasdille](https://github.com/nicholasdille)
-   <b><code>&nbsp;&nbsp;1890⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;96🍴</code></b> [Vagga](https://github.com/tailhook/vagga)) - Vagga is a containerisation tool without daemons. It is a fully-userspace container engine inspired by Vagrant and Docker, specialized for development environments by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tailhook](https://github.com/tailhook/))
-   <b><code>&nbsp;&nbsp;1060⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;121🍴</code></b> [Zsh-in-Docker](https://github.com/deluan/zsh-in-docker)) - Install Zsh, Oh-My-Zsh and plugins inside a Docker container with one line! By 🌎 [Deluan](www.deluan.com)

### Garbage Collection

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [caduc](https://github.com/tjamet/caduc)) - A docker garbage collector cleaning stuff you did not use recently
-   <b><code>&nbsp;&nbsp;1299⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;89🍴</code></b> [Docker Clean](https://github.com/ZZROTDesign/docker-clean)) - A script that cleans Docker containers, images and volumes by [@zzrotdesign](https://github.com/ZZROTDesign)
-   <b><code>&nbsp;&nbsp;&nbsp;121⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22🍴</code></b> [docker_gc](https://github.com/pdacity/docker_gc)) - Image for automatic removing unused Docker Swarm objects. Also works just as Docker Service by [@pdacity](https://github.com/pdacity)
-   <b><code>&nbsp;&nbsp;&nbsp;589⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;117🍴</code></b> [Docker-cleanup](https://github.com/meltwater/docker-cleanup)) :skull: - Automatic Docker image, container and volume cleanup by [@meltwater](https://github.com/meltwater)
-   <b><code>&nbsp;&nbsp;&nbsp;374⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;50🍴</code></b> [docker-custodian](https://github.com/Yelp/docker-custodian)) - Keep docker hosts tidy. By [@Yelp](https://github.com/Yelp)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker-garby](https://github.com/konstruktoid/docker-garby)) - :skull: Docker garbage collection script by [@konstruktoid](https://github.com/konstruktoid).
-   <b><code>&nbsp;&nbsp;5036⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;443🍴</code></b> [docker-gc](https://github.com/spotify/docker-gc)) :skull: - A cron job that will delete old stopped containers and unused images by [@spotify][spotify]
-   <b><code>&nbsp;&nbsp;&nbsp;659⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;41🍴</code></b> [Docuum](https://github.com/stepchowfun/docuum)) - Least recently used (LRU) eviction of Docker images by [@stepchowfun](https://github.com/stepchowfun)
-   <b><code>&nbsp;&nbsp;&nbsp;109⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [sherdock](https://github.com/rancher/sherdock)) :skull: - Automatic GC of images based on regexp by [@rancher][rancher]

### Serverless

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;28🍴</code></b> [AMP](https://github.com/appcelerator-archive/amp)) :skull: - The open source unified CaaS/FaaS platform for Docker, batteries included. By [@Appcelerator](https://github.com/appcelerator-archive)
-   <b><code>&nbsp;&nbsp;6708⭐</code></b> <b><code>&nbsp;&nbsp;1172🍴</code></b> [Apache OpenWhisk](https://github.com/apache/openwhisk)) - a serverless, open source cloud platform that executes functions in response to events at any scale. By [@apache](https://github.com/apache)
-   <b><code>&nbsp;&nbsp;5800⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;419🍴</code></b> [Docker-Lambda](https://github.com/lambci/docker-lambda)) - :skull: Docker images and test runners that replicate the live AWS Lambda environment. By [@lamb-ci](https://github.com/lambci)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [Funker](https://github.com/bfirsh/funker-example-voting-app)) - Functions as Docker containers example voting app. By [@bfirsh](https://github.com/bfirsh)
-   <b><code>&nbsp;&nbsp;3216⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;227🍴</code></b> [IronFunctions](https://github.com/iron-io/functions)) - The serverless microservices platform FaaS (Functions as a Service) which uses Docker containers to run Any language or AWS Lambda functions
-   🌎 [Koyeb](www.koyeb.com/) :heavy_dollar_sign: - Koyeb is a developer-friendly serverless platform to deploy apps globally. Seamlessly run Docker containers, web apps, and APIs with git-based deployment, native autoscaling, a global edge network, and built-in service mesh and discovery.
-   <b><code>&nbsp;25948⭐</code></b> <b><code>&nbsp;&nbsp;1970🍴</code></b> [OpenFaaS](https://github.com/openfaas/faas)) - A complete serverless functions framework for Docker and Kubernetes. By [OpenFaaS](https://github.com/openfaas)
-   <b><code>&nbsp;&nbsp;&nbsp;602⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;47🍴</code></b> [SCAR](https://github.com/grycap/scar)) - Serverless Container-aware Architectures (SCAR) is a serverless framework that allows easy deployment and execution of containers (e.g. Docker) in Serverless environments (e.g. Lambda) by [@grycap](https://github.com/grycap)

### Testing

-   <b><code>&nbsp;&nbsp;2433⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;202🍴</code></b> [Container Structure Test](https://github.com/GoogleContainerTools/container-structure-test)) - A framework to validate the structure of an image by checking the outputs of commands or the contents of the filesystem. By [@GoogleContainerTools][googlecontainertools]
-   <b><code>&nbsp;&nbsp;5814⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;486🍴</code></b> [dgoss](https://github.com/aelsabbahy/goss/tree/master/extras/dgoss)) - A fast YAML based tool for validating docker containers.
-   <b><code>&nbsp;&nbsp;&nbsp;181⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [DockerSpec](https://github.com/zuazo/dockerspec)) - A small Ruby Gem to run RSpec and Serverspec, Infrataster and Capybara tests against Dockerfiles or Docker images easily. By [@zuazo](https://github.com/zuazo)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Dockunit](https://github.com/dockunit/platform)) :skull: - Docker based integration tests. A simple Node based utility for running Docker based unit tests. By [@dockunit](https://github.com/dockunit)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [EZDC](https://github.com/lynchborg/ezdc)) - Golang test harness for easily setting up tests that rely on services in a docker-compose.yml. By [@byrnedo]
-   [InSpec][inspec] - InSpec is an open-source testing framework for infrastructure with a human- and machine-readable language for specifying compliance, security and policy requirements. By [@chef](https://github.com/chef)
-   <b><code>&nbsp;&nbsp;&nbsp;491⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;83🍴</code></b> [Kurtosis](https://github.com/kurtosis-tech/kurtosis)) - A composable build system for multi-container test environments that provides developers with: a powerful Python-like SDK for environment configuration, a compile-time validator to verify environment behavior & setup, and a runtime for environment execution, monitoring, & debugging capabilities. By 🌎 [Kurtosis](www.kurtosis.com/)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Pull Dog](https://github.com/apps/pull-dog)) - A GitHub app that automatically creates Docker-based test environments for your pull requests, from your docker-compose files. Not open source.
-   <b><code>&nbsp;&nbsp;2936⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;199🍴</code></b> [Pumba](https://github.com/alexei-led/pumba)) - Chaos testing tool for Docker. Can be deployed on kubernetes and CoreOS cluster. By [@alexei-led](https://github.com/alexei-led)

### Wrappers

-   🌎 [Ansible](docs.ansible.com/ansible/latest/collections/community/general/docker_container_module.html) - Manage the life cycle of Docker containers. By RedHat
-   <b><code>&nbsp;&nbsp;&nbsp;894⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;55🍴</code></b> [Azk](https://github.com/azukiapp/azk)) - :skull: Orchestrate development environments on your local machine by [@azukiapp](https://github.com/azukiapp)
-   <b><code>&nbsp;&nbsp;&nbsp;166⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [Beluga](https://github.com/cortexmedia/Beluga)) :skull: - CLI to deploy docker containers on a single server or low amount of servers. By [@cortextmedia](https://github.com/cortexmedia)
-   <b><code>&nbsp;&nbsp;&nbsp;332⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [dexec](https://github.com/docker-exec/dexec)) - Command line interface written in Go for running code with Docker Exec images.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;63⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [dockerized](https://github.com/benzaita/dockerized-cli)) - Seamlessly execute commands in a container.
-   <b><code>&nbsp;&nbsp;&nbsp;385⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38🍴</code></b> [Dray](https://github.com/CenturyLinkLabs/dray)) - An engine for managing the execution of container-based workflows by [@CenturyLinkLabs][centurylinklabs]
-   <b><code>&nbsp;&nbsp;&nbsp;140⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [FuGu](https://github.com/mattes/fugu)) :skull: - Docker run wrapper without orchestration by [@mattes](https://github.com/mattes)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;95⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25🍴</code></b> [Hokusai](https://github.com/artsy/hokusai)) - A Docker + Kubernetes CLI for application developers; used to containerize an application and to manage its lifecycle throughout development, testing, and release cycles. From [@artsy](https://github.com/artsy)
-   <b><code>&nbsp;&nbsp;2162⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;87🍴</code></b> [Preevy](https://github.com/livecycle/preevy)) - Preview environments for Docker and Docker Compose projects. Test your changes and get feedback from devs and non-devs (Product/Design) by deploying pull requests to the your cloud provider as part of your CI pipeline.
-   <b><code>&nbsp;&nbsp;2143⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;110🍴</code></b> [Shutit](https://github.com/ianmiell/shutit)) - Tool for building and maintaining complex Docker deployments by [@ianmiell](https://github.com/ianmiell)
-   <b><code>&nbsp;&nbsp;&nbsp;893⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;63🍴</code></b> [subuser](https://github.com/subuser-security/subuser)) - Makes it easy to securely and portably run graphical desktop applications in Docker
-   <b><code>&nbsp;&nbsp;&nbsp;782⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;118🍴</code></b> [T.A.D.S. boilerplate](https://github.com/Thomvaill/tads-boilerplate)) - :skull: The power of Ansible and Terraform + the simplicity of Docker Swarm = Infrastructure as Code and DevOps best practices. By [@Thomvaill](https://github.com/Thomvaill)
-   <b><code>&nbsp;&nbsp;&nbsp;117⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;30🍴</code></b> [Terraform cloud-init config](https://github.com/christippett/terraform-cloudinit-container-server)) - Terraform module for deploying a single Docker image or `docker-compose.yaml` file to any Cloud™ VM
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [Turbo](https://github.com/ramitsurana/turbo)) - Simple and Powerful utility for docker. By [@ramitsurana][ramitsurana]
-   <b><code>&nbsp;&nbsp;1620⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;154🍴</code></b> [udocker](https://github.com/indigo-dc/udocker)) - A tool to execute simple docker containers in batch or interactive systems without root privileges by [@inidigo-dc](https://github.com/indigo-dc)
-   🌎 [Vagrant - Docker provider](developer.hashicorp.com/vagrant/docs/providers/docker/basics) - Good starting point is <b><code>&nbsp;&nbsp;&nbsp;114⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;45🍴</code></b> [vagrant-docker-example](https://github.com/bubenkoff/vagrant-docker-example)) by [@bubenkoff](https://github.com/bubenkoff)

## Services based on Docker (mostly :heavy_dollar_sign:)

### CI Services

-   🌎 [CircleCI](circleci.com/) :heavy_dollar_sign: - Push or pull Docker images from your build environment, or build and run containers right on CircleCI.
-   🌎 [CodeFresh](codefresh.io) :heavy_dollar_sign: - Everything you need to build, test, and share your Docker applications. Provides automated end to end testing.
-   🌎 [CodeShip](www.cloudbees.com/products/codeship) :heavy_dollar_sign: - Work with your established Docker workflows while automating your testing and deployment tasks with our hosted platform dedicated to speed and security.
-   🌎 [ConcourseCI](concourse-ci.org) :heavy_dollar_sign: - A CI SaaS platform for developers and DevOps teams pipeline oriented.
-   🌎 [Semaphore CI](semaphore.io/) :heavy_dollar_sign: — A high-performance cloud solution that makes it easy to build, test and ship your containers to production.
-   🌎 [TravisCI](www.travis-ci.com/) :heavy_dollar_sign: - A Free github projects continuous integration Saas platform for developers and Devops.

### CaaS

-   🌎 [Amazon ECS](aws.amazon.com/ecs/) :heavy_dollar_sign: - A management service on EC2 that supports Docker containers.
-   🌎 [Appfleet](appfleet.com/) :heavy_dollar_sign: - Edge platform to deploy and manage containerized services globally. The system will route the traffic to the closest location for lower latency.
-   🌎 [Azure AKS](azure.microsoft.com/en-us/products/kubernetes-service/) :heavy_dollar_sign: - Simplify Kubernetes management, deployment, and operations. Use a fully managed Kubernetes container orchestration service.
-   🌎 [Cloud 66](www.cloud66.com) :heavy_dollar_sign: - Full-stack hosted container management as a service
-   🌎 [Giant Swarm](www.giantswarm.io/) :heavy_dollar_sign: - Simple microservice infrastructure. Deploy your containers in seconds.
-   🌎 [Google Container Engine](cloud.google.com/kubernetes-engine/docs/) :heavy_dollar_sign: - Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
-   🌎 [Mesosphere DC/OS Platform](d2iq.com/products/dcos) :heavy_dollar_sign: - Integrated platform for data and containers built on Apache Mesos by 🌎 [@mesosphere](d2iq.com)
-   🌎 [Red Hat OpenShift Dedicated](www.redhat.com/en/technologies/cloud-computing/openshift/dedicated) :heavy_dollar_sign: - Fully-managed Red Hat® OpenShift® service on Amazon Web Services and Google Cloud
-   🌎 [Triton](www.joyent.com/) :heavy_dollar_sign: - Elastic container-native infrastructure by Joyent.
-   🌎 [Virtuozzo Application Platform](www.virtuozzo.com/application-platform-partners/) :heavy_dollar_sign: - Deploy and manage your projects with turnkey PaaS across a wide network of reliable service providers

### Monitoring Services

-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [AppDynamics](https://github.com/Appdynamics/docker-monitoring-extension)) - Docker Monitoring extension gathers metrics from the Docker Remote API, either using Unix Socket or TCP.
-   🌎 [Better Stack](betterstack.com/community/guides/scaling-docker/) :heavy_dollar_sign: - A Docker-compatible observability stack that delivers robust log aggregation and uptime monitoring capabilities for various software application.
-   🌎 [Broadcom Docker Monitoring](www.broadcom.com/info/aiops/docker-monitoring) :heavy_dollar_sign: - Agile Operations solutions from Broadcom deliver the modern Docker monitoring businesses need to accelerate and optimize the performance of microservices and the dynamic Docker environments running them. Monitor both the Docker environment and apps that run inside them. (former CA Technologies)
-   🌎 [Collecting docker logs and stats with Splunk](www.splunk.com/en_us/blog/tips-and-tricks/collecting-docker-logs-and-stats-with-splunk.html)
-   🌎 [Datadog](www.datadoghq.com/) :heavy_dollar_sign: - Datadog is a full-stack monitoring service for large-scale cloud environments that aggregates metrics/events from servers, databases, and applications. It includes support for Docker, Kubernetes, and Mesos.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [DockStat](https://github.com/its4nik/dockstat)) :construction: - A full fletched (WIP) Docker management solution featuring plugin support and community integration by [its4nik](https://github.com/its4nik)
-   🌎 [Prometheus](prometheus.io/) :heavy_dollar_sign: - Open-source service monitoring system and time series database
-   🌎 [Site24x7](www.site24x7.com/docker-monitoring.html) :heavy_dollar_sign: - Docker Monitoring for DevOps and IT is a SaaS Pay per Host model
-   <b><code>&nbsp;&nbsp;&nbsp;210⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;32🍴</code></b> [SPM for Docker](https://github.com/sematext/sematext-agent-docker)) :heavy_dollar_sign: - Monitoring of host and container metrics, Docker events and logs. Automatic log parser. Anomaly Detection and alerting for metrics and logs. [@sematext](https://github.com/sematext)
-   🌎 [Sysdig Monitor](www.sysdig.com/products/monitor) :heavy_dollar_sign: - Sysdig Monitor can be used as either software or a SaaS service to monitor, alert, and troubleshoot containers using system calls. It has container-specific features for Docker and Kubernetes.

# Useful Resources

-   **[Valuable Docker Links](http://nane.kratzke.pages.mylab.th-luebeck.de/about/blog/2014/08/24/valuable-docker-links/)** High quality articles about docker! **MUST SEE**
-   <b><code>&nbsp;&nbsp;9713⭐</code></b> <b><code>&nbsp;&nbsp;2079🍴</code></b> [Cloud Native Landscape](https://github.com/cncf/landscape))
-   🌎 [Docker Blog](www.docker.com/blog/) - regular updates about Docker, the community and tools
-   🌎 [Docker Certification](intellipaat.com/docker-training-course/?US) :heavy_dollar_sign: will help you to will Learn Docker containerization, running Docker containers, Image creation, Dockerfile, Docker orchestration, security best practices, and more through hands-on projects and case studies and helps to clear Docker Certified Associate.

-   🌎 [Docker dev bookmarks](www.codever.dev/search?q=docker) - use the tag 🌎 [docker](www.codever.dev/bookmarks/t/docker)
-   🌎 [Docker in Action, Second Edition](www.manning.com/books/docker-in-action-second-edition)
-   🌎 [Docker in Practice, Second Edition](www.manning.com/books/docker-in-practice-second-edition)
-   🌎 [Docker packaging guide for Python](pythonspeed.com/docker/) - a series of detailed articles on the specifics of Docker packaging for Python.
-   🌎 [Learn Docker in a Month of Lunches](www.manning.com/books/learn-docker-in-a-month-of-lunches)
-   🌎 [Learn Docker](coursesity.com/blog/best-docker-tutorials/) - Learn Docker - curated list of the top online docker tutorials and courses.
-   🌎 [Programming Community Curated Resources for learning Docker](hackr.io/tutorials/learn-docker)

## Awesome Lists

-   <b><code>&nbsp;&nbsp;1967⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;217🍴</code></b> [Awesome CI/CD](https://github.com/cicdops/awesome-ciandcd)) - Not specific to docker but relevant.
-   <b><code>&nbsp;42038⭐</code></b> <b><code>&nbsp;&nbsp;7794🍴</code></b> [Awesome Compose](https://github.com/docker/awesome-compose)) - Docker Compose samples
-   <b><code>&nbsp;15656⭐</code></b> <b><code>&nbsp;&nbsp;2324🍴</code></b> [Awesome Kubernetes](https://github.com/ramitsurana/awesome-kubernetes)) by [@ramitsurana][ramitsurana]
-   <b><code>&nbsp;&nbsp;1973⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;172🍴</code></b> [Awesome Linux Container](https://github.com/Friz-zy/awesome-linux-containers)) more general about container than this repo, by [@Friz-zy](https://github.com/Friz-zy).
-   <b><code>254716⭐</code></b> <b><code>&nbsp;11789🍴</code></b> [Awesome Selfhosted](https://github.com/awesome-selfhosted/awesome-selfhosted)) list of Free Software network services and web applications which can be hosted locally by running in a classical way (setup local web server and run applications from there) or in a Docker container. By [@Kickball](https://github.com/Kickball)
-   <b><code>&nbsp;31494⭐</code></b> <b><code>&nbsp;&nbsp;1846🍴</code></b> [Awesome Sysadmin](https://github.com/n1trux/awesome-sysadmin)) by [@n1trux](https://github.com/n1trux)
-   <b><code>&nbsp;16869⭐</code></b> <b><code>&nbsp;&nbsp;1285🍴</code></b> [ToolsOfTheTrade](https://github.com/cjbarber/ToolsOfTheTrade)) a list of SaaS and On premise applications by [@cjbarber](https://github.com/cjbarber)

## Demos and Examples

-   🌎 [An Annotated Docker Config for Frontend Web Development](nystudio107.com/blog/an-annotated-docker-config-for-frontend-web-development) A local development environment with Docker allows you to shrink-wrap the devops your project needs as config, making onboarding frictionless.
-   <b><code>&nbsp;&nbsp;&nbsp;300⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37🍴</code></b> [Local Docker DB](https://github.com/alexmacarthur/local-docker-db)) a list of docker-compose samples for a lot of databases by [@alexmacarthur](https://github.com/alexmacarthur)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;89⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;18🍴</code></b> [Webstack-micro](https://github.com/ferbs/webstack-micro)) Demo web app showing how Docker Compose might be used to set up an API Gateway, centralized authentication, background workers, and WebSockets as containerized services.

## Good Tips

-   [Docker Caveats](http://docker-saigon.github.io/post/Docker-Caveats/) What You Should Know About Running Docker In Production (written 11 APRIL 2016) **MUST SEE**
-   🌎 [Docker Containers on the Desktop](blog.jessfraz.com/post/docker-containers-on-the-desktop/) - The **funniest way** to learn about docker by [@jessfraz][jessfraz] who also gave a 🌎 [presentation](www.youtube.com/watch?v=1qlLUf7KtAw) about it @ DockerCon 2015
-   🌎 [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](www.flexera.com/blog/finops/)
-   <b><code>&nbsp;&nbsp;4092⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;154🍴</code></b> [Dockerfile best practices](https://github.com/hexops/dockerfile)) - This repository has best-practices for writing Dockerfiles
-   🌎 [Don't Repeat Yourself with Anchors, Aliases and Extensions in Docker Compose Files](medium.com/@kinghuang/docker-compose-anchors-aliases-extensions-a1e4105d70bd) by [@King Chung Huang](https://github.com/kinghuang)
-   [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm][fgrehm]

## Raspberry Pi & ARM

-   🌎 [Docker Pirates ARMed with explosive stuff](blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Get Docker up and running on the RaspberryPi in three steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-three-steps))
-   🌎 [git push docker containers to linux devices](www.balena.io) Modern DevOps for IoT, leveraging git and Docker.
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-%28ARMv7%29-devices))

## Security

-   🌎 [Bringing new security features to Docker](opensource.com/business/14/9/security-for-docker)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [CVE Scanning Alpine images with Multi-stage builds in Docker 17.05](https://github.com/tomwillfixit/alpine-cvecheck)) by 🌎 [@tomwillfixit](twitter.com/tomwillfixit)
-   <b><code>&nbsp;&nbsp;&nbsp;607⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;79🍴</code></b> [Docker Secure Deployment Guidelines](https://github.com/AonCyberLabs/Docker-Secure-Deployment-Guidelines))
-   🌎 [Docker Security - Quick Reference](binarymist.io/publication/docker-security/)
-   🌎 [Docker Security: Are Your Containers Tightly Secured to the Ship? SlideShare](www.slideshare.net/slideshow/docker-security-are-your-containers-tightly-secured-to-the-ship/43834790)
-   <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [How CVE's are handled on Offical Docker Images](https://github.com/docker-library/official-images/issues/1448))
-   🌎 [Lynis is an open source security auditing tool including Docker auditing](cisofy.com/lynis/)
-   🌎 [Security Best Practices for Building Docker Images](linux-audit.com/tags/docker/)
-   🌎 [Software Engineering Radio interview of Docker Security Team Lead (Diogo Mónica)](www.se-radio.net/2017/05/se-radio-episode-290-diogo-monica-on-docker-security/)
-   🌎 [Ten Docker Image Security Best Practices Cheat Sheet](snyk.io/blog/10-docker-image-security-best-practices/)
-   🌎 [Top ten most popular docker images each contain at least 30 vulnerabilities](snyk.io/blog/top-ten-most-popular-docker-images-each-contain-at-least-30-vulnerabilities/)
-   🌎 [Tuning Docker with the newest security enhancements](opensource.com/business/15/3/docker-security-tuning)
-   🌎 [10 best practices to containerize Node.js web applications with Docker](snyk.io/blog/10-best-practices-to-containerize-nodejs-web-applications-with-docker/)

## Videos

-   🌎 [Contributing to Docker by Andrew "Tianon" Page (InfoSiftr)](www.youtube.com/watch?v=1jwo8-1HYYg) (34:31)
-   🌎 [Deploying and scaling applications with Docker, Swarm, and a tiny bit of Python magic](www.youtube.com/watch?v=GpHMTR7P2Ms) (3:11:06) by [@jpetazzo][jpetazzo]
-   🌎 [Docker and SELinux by Daniel Walsh from Red Hat](www.youtube.com/watch?v=zWGFqMuEHdw) (40:23)
-   🌎 [Docker Course](www.youtube.com/watch?v=UZpyvK6UGFo) (Spanish) by [@pablokbs](https://github.com/pablokbs)
-   🌎 [Docker for Developers](www.youtube.com/watch?v=FdkNAjjO5yQ) (54:26) by [@jpetazzo][jpetazzo] <== Good introduction, context, demo
-   🌎 [Docker from scratch](www.youtube.com/playlist?list=PLLhEJK7fQIxD-btrjrqdEfQHbkZnQrmqE) (1:22:01) on YouTube by Paris Nakita Kejser
-   🌎 [Docker: How to Use Your Own Private Registry](www.youtube.com/watch?v=CAewZCBT4PI) (15:01)
-   🌎 [Docker in Production](www.youtube.com/watch?v=Glk5d5WP6MI) by [@jpetazzo][jpetazzo] (36:05)
-   🌎 [Docker Primer to Docker Compose](www.youtube.com/watch?v=G-s2GXGAjTk) (1:56:45) on YouTube by LoginRadius
-   🌎 [Docker Registry from scratch](www.youtube.com/playlist?list=PLLhEJK7fQIxAz3d4Fj3edq7UcxEhdTCBm) (44:40) on YouTube by Paris Nakita Kejser
-   🌎 [Docker Swarm from scratch](www.youtube.com/playlist?list=PLLhEJK7fQIxAY4gZd1Wl-GsLvg-e9Ap1e) (1:41:28) on YouTube by Paris Nakita Kejser
-   🌎 [Extending Docker with Plugins](vimeo.com/110835013) (15:21)
-   🌎 [From Local Docker Development to Production Deployments](www.youtube.com/watch?v=7CZFpHUPqXw) by [@jpetazzo][jpetazzo] @ AWS re:Invent 2015
-   🌎 [Immutable Infrastructure with Docker and EC2 by Michael Bryzek (Gilt)](www.youtube.com/watch?v=GaHzdqFithc) (42:04)
-   🌎 [Introduction to Docker and containers](www.youtube.com/watch?v=ZVaRK10HBjo) (3:09:00) by [@jpetazzo][jpetazzo]
-   🌎 [Logging on Docker: What You Need to Know](vimeo.com/123341629) (51:27)
-   🌎 [Performance Analysis of Docker - Jeremy Eder](www.youtube.com/watch?v=6f2E6PKYb0w) (1:36:58)
-   🌎 [Scalable Microservices with Kubernetes](www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) Free Udacity course
-   🌎 [State of containers: a debate with CoreOS, VMware and Google](www.youtube.com/watch?v=IiITP3yIRd8) (27:38)

# Communities and Meetups

## Brazilian

-   🌎 [Docker BR on Telegram](telegram.me/dockerbr)

## Chinese

-   [DockerOne](http://dockone.io/) Docker Community (in Chinese) by [@LiYingJie](http://dockone.io/people/%E6%9D%8E%E9%A2%96%E6%9D%B0)

## English

-   🌎 [Docker Community](www.docker.com/community/)
-   🌎 [Docker Events](www.docker.com/events/)
-   🌎 [Docker Online Meetup](www.meetup.com/en-AU/Docker-Online-Meetup/)
-   🌎 [Docker Reddit Community](www.reddit.com/r/docker/)

## Russian

-   🌎 [Docker Russian-speaking Community](t.me/docker_ru)

## Spanish

-   🌎 [Docker Tips](dockertips.com/)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/veggiemonk/awesome-docker.svg)](https://starchart.cc/veggiemonk/awesome-docker)

[contributing]: https://github.com/correia-jpv/fucking-awesome-docker/blob/master/.github/CONTRIBUTING.md
[akito]: https://github.com/theAkito
[calico]: https://github.com/projectcalico/calico
[centurylinklabs]: https://github.com/CenturyLinkLabs
[containx]: https://github.com/ContainX
[containers]: https://github.com/containers
[coreos]: https://github.com/coreos
[deepfence]: https://github.com/deepfence
[distribution]: https://github.com/docker/distribution
[docker-flow]: https://github.com/docker-flow
[docker-for-windows]: https://docs.docker.com/desktop/setup/install/windows-install/
[docker]: https://github.com/docker
[dozzle]: https://github.com/amir20/dozzle
[editreadme]: https://github.com/correia-jpv/fucking-awesome-docker/edit/master/README.md
[fgrehm]: https://github.com/fgrehm
[gesellix]: https://github.com/gesellix
[genuinetools]: https://github.com/genuinetools
[gliderlabs]: https://github.com/gliderlabs
[google]: https://github.com/google
[googlecontainertools]: https://github.com/GoogleContainerTools
[inspec]: https://github.com/inspec/inspec
[jessfraz]: https://github.com/jessfraz
[jpetazzo]: https://github.com/jpetazzo
[jwilder]: https://github.com/jwilder
[kubernetes]: https://kubernetes.io
[lispyclouds]: https://github.com/lispyclouds
[nvidia]: https://github.com/nvidia
[nginxproxy]: https://github.com/nginx-proxy/nginx-proxy
[openshift]: https://okd.io/
[oracle]: https://github.com/oracle
[peco602]: https://github.com/Peco602
[powerman]: https://github.com/powerman
[progrium]: https://github.com/progrium
[ramitsurana]: https://github.com/ramitsurana
[rancher]: https://github.com/rancher
[safe-waters]: https://github.com/safe-waters
[sindresorhus]: https://github.com/sindresorhus/awesome
[spotify]: https://github.com/spotify
[tomastomecek]: https://github.com/TomasTomecek
[vegasbrianc]: https://github.com/vegasbrianc
[weave]: https://github.com/weaveworks/weave
[vmware]: https://github.com/vmware
[@byrnedo]: https://github.com/byrnedo
[@crazy-max]: https://github.com/crazy-max
[@grammarly]: https://github.com/grammarly
[@skanehira]: https://github.com/skanehira

## Source
<b><code>&nbsp;34269⭐</code></b> <b><code>&nbsp;&nbsp;3165🍴</code></b> [veggiemonk/awesome-docker](https://github.com/veggiemonk/awesome-docker))