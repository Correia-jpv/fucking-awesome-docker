# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)][sindresorhus] [![Netlify Status](https://api.netlify.com/api/v1/badges/8ca86717-11ba-46d4-9d0a-700d8527f13b/deploy-status)](https://app.netlify.com/sites/awesome-docker/deploys)<!-- omit in toc -->

> A curated list of Docker resources and projects

If you would like to contribute, please read [CONTRIBUTING.md][contributing] first.
It contains a lot of tips and guidelines to help keep things organized.
Just click [README.md][editreadme] to submit a [pull request][editreadme].
If this list is not complete, you can [contribute][editreadme] to make it so. Here is a great video tutorial to learn how to 🌎 [contribute on Github](https://egghead.io/lessons/javascript-identifying-how-to-contribute-to-an-open-source-project-on-github).

**_You can see the updates on 🌎 [TWITTER](https://twitter.com/awesome_docker)_**

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for newcomers. See how to **[Contribute][contributing]** for tips!

**_If you see a link here that is not (any longer) a good fit, you can fix it by submitting a [pull request][editreadme] to improve this file. Thank you!_**

The creators and maintainers of this list do not receive any form of payment to accept a change made by any contributor. This page is not an official Docker product in any way. It is a list of links to projects and is maintained by volunteers. Everybody is welcome to contribute. The goal of this repo is to index open-source projects, not to advertise for profit.

All the links are monitored and tested with a home baked <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Node.js script](https://github.com/veggiemonk/awesome-docker/blob/master/pull_request.mjs)

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
  - [Services based on Docker (mostly :heavy_dollar_sign:)](#services-based-on-docker-mostly-heavy_dollar_sign)
    - [CI Services](#ci-services)
    - [CaaS](#caas)
    - [Monitoring Services](#monitoring-services)
- [Useful Resources](#useful-resources)
  - [Awesome Lists](#awesome-lists)
  - [Demos and Examples](#demos-and-examples)
  - [Good Tips](#good-tips)
  - [Raspberry Pi & ARM](#raspberry-pi--arm)
  - [Security](#security-1)
  - [Videos](#videos)
- [Communities and Meetups](#communities-and-meetups)
  - [Brazilian](#brazilian)
  - [Chinese](#chinese)
  - [English](#english)
  - [Russian](#russian)
  - [Spanish](#spanish)
  - [Stargazers over time](#stargazers-over-time)
  - [Contributor over time](#contributor-over-time)

<!-- /TOC -->

# Legend

- Abandoned :skull:
- Beta :construction:
- Monetized :heavy_dollar_sign:

# What is Docker

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

_Source:_ 🌎 [What is Docker](https://www.docker.com/why-docker/)

# Where to start

- 🌎 [Benefits of using Docker](https://semaphoreci.com/blog/docker-benefits) for development and delivery, with a practical roadmap for adoption.
- 🌎 [Bootstrapping Microservices](https://www.manning.com/books/bootstrapping-microservices-with-docker-kubernetes-and-terraform) by 🌎 [Ashley Davis](https://twitter.com/ashleydavis75) - A practical and project-based guide to building applications with microservices, starts by building a Docker image for a single microservice and publishing it to a private container registry, finishes by deploying a complete microservices application to a production Kubernetes cluster.
- <b><code>&nbsp;&nbsp;4738⭐</code></b> <b><code>&nbsp;&nbsp;1896🍴</code></b> [Docker Curriculum](https://github.com/prakhar1989/docker-curriculum): A comprehensive tutorial for getting started with Docker. Teaches how to use Docker and deploy dockerized apps on AWS with Elastic Beanstalk and Elastic Container Service.
- 🌎 [Docker Documentation](https://docs.docker.com/): the official documentation.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker for beginners](https://github.com/groda/big_data/blob/master/docker_for_beginners.md): A tutorial for beginners who need to learn the basics of Docker—from "Hello world!" to basic interactions with containers, with simple explanations of the underlying concepts.
- 🌎 [Docker for novices](https://www.youtube.com/watch?v=xsjSadjKXns) An introduction to Docker for developers and testers who have never used it. (Video 1h40, recorded linux.conf.au 2019 — Christchurch, New Zealand) by Alex Clews.
- 🌎 [Docker simplified in 55 seconds](https://www.youtube.com/watch?v=vP_4DlOH1G4): An animated high-level introduction to Docker. Think of it as a visual tl;dr that makes it easier to dive into more complex learning materials.
- 🌎 [Docker Training](https://training.mirantis.com) :heavy_dollar_sign:
- 🌎 [Docker Tutorial for Beginners (Updated 2019 version)](https://hashnode.com/post/docker-tutorial-for-beginners-cjrj2hg5001s2ufs1nker9he2) — In this Docker tutorial, you'll learn all the basics and learn how you can containerize Node.js and Go applications. Even if you aren't familiar with these languages it should be easy for you to follow this tutorial and use any other language.
- 🌎 [Katacoda](https://www.katacoda.com/courses/docker): Learn Docker using Interactive Browser-Based Labs
- <b><code>&nbsp;&nbsp;&nbsp;199⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;28🍴</code></b> [Learn Docker](https://github.com/dwyl/learn-docker): step-by-step tutorial and more resources (video, articles, cheat sheets) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@dwyl](https://github.com/dwyl)
- 🌎 [Play With Docker](https://training.play-with-docker.com/): PWD is a great way to get started with Docker from beginner to advanced users. Docker runs directly in your browser.
- <b><code>&nbsp;&nbsp;&nbsp;179⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;64🍴</code></b> [Practical Guide about Docker Commands in Spanish](https://github.com/brunocascio/docker-espanol) This spanish guide contains the use of basic docker commands with real life examples.
- 🌎 [Practical Introduction to Container Terminology](https://developers.redhat.com/blog/2018/02/22/container-terminology-practical-introduction) The landscape for container technologies is larger than just docker. Without a good handle on the terminology, It can be difficult to grasp the key differences between docker and (pick your favorites, CRI-O, rkt, lxc/lxd) or understand what the Open Container Initiative is doing to standardize container technology.
- 🌎 [The Docker Handbook](https://docker-handbook.farhan.dev/) An open-source book that teaches you the fundamentals, best practices and some intermediate Docker functionalities. The book is hosted on <b><code>&nbsp;&nbsp;&nbsp;618⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;83🍴</code></b> [fhsinchy/the-docker-handbook](https://github.com/fhsinchy/the-docker-handbook) and the projects are hosted on <b><code>&nbsp;&nbsp;1081⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;248🍴</code></b> [fhsinchy/docker-handbook-projects](https://github.com/fhsinchy/docker-handbook-projects) repository.

**Cheatsheets** by

- <b><code>&nbsp;&nbsp;3359⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;404🍴</code></b> [@eon01](https://github.com/eon01/DockerCheatSheet)
- <b><code>&nbsp;&nbsp;&nbsp;174⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [@dimonomid](https://github.com/dimonomid/docker-quick-ref) (PDF)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [@JensPiegsa](https://github.com/JensPiegsa/docker-cheat-sheet)
- <b><code>&nbsp;20917⭐</code></b> <b><code>&nbsp;&nbsp;4430🍴</code></b> [@wsargent](https://github.com/wsargent/docker-cheat-sheet) (Most popular)

# Where to start (Windows)

- 🌎 [A Comparative Study of Docker Engine on Windows Server vs Linux Platform](https://collabnix.com/a-comparative-study-of-docker-engine-on-windows-server-vs-linux-platform/) Comparing the feature sets and implementations of Docker on Windows and Linux
- 🌎 [Build And Run Your First Docker Windows Server Container](https://www.docker.com/blog/build-your-first-docker-windows-server-container/) Walkthrough installing Docker on Windows 10, building a Docker image and running a Windows container
- 🌎 [Docker on Windows behind a firewall](https://toedter.com/2015/05/11/docker-on-windows-behind-a-firewall/) by 🌎 [@kaitoedter](https://twitter.com/kaitoedter)
- 🌎 [Docker Quick Tutorial](https://vegibit.com/docker-quick-tutorial/) Introduces you to the official Getting Started Tutorial using Windows 10 and Hyper-V.
- 🌎 [Docker Reference Architecture: Modernizing Traditional .NET Framework Applications](https://docs.mirantis.com/containers/v3.0/dockeree-ref-arch/app-dev/modernize-dotnet-apps.html) - You will learn to identify the types of .NET Framework applications that are good candidates for containerization, the "lift-and-shift" approach to containerization.
- 🌎 [Docker with Microsoft SQL 2016 + ASP.NET](https://blog.alexellis.io/docker-does-sql2016-aspnet/) Demonstration running ASP.NET and SQL Server workloads in Docker
- 🌎 [Exploring ASP.NET Core with Docker in both Linux and Windows Containers](https://www.hanselman.com/blog/exploring-aspnet-core-with-docker-in-both-linux-and-windows-containers) Running ASP.NET Core apps in Linux and Windows containers, using [Docker for Windows][docker-for-windows]
- 🌎 [Running a Legacy ASP.NET App in a Windows Container](https://blog.sixeyed.com/dockerizing-nerd-dinner-part-1-running-a-legacy-asp-net-app-in-a-windows-container/) Steps for Dockerizing a legacy ASP.NET app and running as a Windows container
- 🌎 [Windows Containers and Docker: The 101](https://www.youtube.com/watch?v=N7SG2wEyQtM) :movie_camera: - A 20-minute overview, using Docker to run PowerShell, ASP.NET Core and ASP.NET apps
- 🌎 [Windows Containers Quick Start](https://docs.microsoft.com/en-us/virtualization/windowscontainers/about/) Overview of Windows containers, drilling down to Quick Starts for Windows 10 and Windows Server 2016

---

# Projects

- Moby = open source development
- Docker CE = free product release based on Moby
- Docker EE = commercial product release based on Docker CE.

> Docker EE is on the same code base as Docker CE, so also built from Moby, with commercial components added, such as "docker data center / universal control plane"

- <b><code>&nbsp;63254⭐</code></b> <b><code>&nbsp;18138🍴</code></b> [Moby](https://github.com/moby/moby)
- 🌎 [Docker Images](https://hub.docker.com)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker Compose](https://github.com/docker/compose/) (Define and run multi-container applications with Docker)
- <b><code>&nbsp;&nbsp;6499⭐</code></b> <b><code>&nbsp;&nbsp;1968🍴</code></b> [Docker Machine](https://github.com/docker/machine) (Machine management for a container-centric world)
- [Docker Registry][distribution] (The Docker toolset to pack, ship, store, and deliver content)
- <b><code>&nbsp;&nbsp;5802⭐</code></b> <b><code>&nbsp;&nbsp;1112🍴</code></b> [Docker Swarm](https://github.com/docker/swarm) (Swarm: a Docker-native clustering system)

## Container Operations

### Container Composition

- <b><code>&nbsp;&nbsp;&nbsp;130⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [bocker](https://github.com/icy/bocker) (2) :skull: - Write Dockerfile completely in Bash. Extensible and simple. --> Reusable by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@icy](https://github.com/icy)
- <b><code>&nbsp;&nbsp;8669⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;633🍴</code></b> [bocker](https://github.com/p8952/bocker) (1) :skull: - Docker implemented in 100 lines of bash by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [p8952](https://github.com/p8952)
- <b><code>&nbsp;&nbsp;&nbsp;237⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [box](https://github.com/box-builder/box) - Build Dockerfile images with a mruby DSL, includes flattening and layer manipulation
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [Capitan](https://github.com/byrnedo/capitan) - Composable docker orchestration with added scripting support by [@byrnedo].
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;77⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [compose_plantuml](https://github.com/funkwerk/compose_plantuml) :skull: - Generate Plantuml graphs from docker-compose files by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@funkwerk](https://github.com/funkwerk)
- <b><code>&nbsp;&nbsp;1627⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;103🍴</code></b> [Composerize](https://github.com/magicmark/composerize) - Convert docker run commands into docker-compose files
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;87⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [crowdr](https://github.com/polonskiy/crowdr) - Tool for managing multiple Docker containers (`docker-compose` alternative) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@polonskiy](https://github.com/polonskiy/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;47⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [docker-compose-graphviz](https://github.com/abesto/docker-compose-graphviz) :skull: - Turn a docker-compose.yml files into Graphviz .dot files by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@abesto](https://github.com/abesto)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [docker-config-update](https://github.com/sudo-bmitch/docker-config-update) - Utility to update docker configs and secrets for deploying in a compose file by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@sudo-bmitch](https://github.com/sudo-bmitch)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;83⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [draw-compose](https://github.com/Alexis-benoist/draw-compose) :skull: - Utility to draw a schema of a docker compose by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Alexis-benoist](https://github.com/Alexis-benoist)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;73⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [elsy](https://github.com/cisco/elsy) - An opinionated, multi-language, build tool based on Docker and Docker Compose
- <b><code>&nbsp;&nbsp;1249⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;90🍴</code></b> [habitus](https://github.com/cloud66-oss/habitus) - A Build Flow Tool for Docker by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@cloud66](https://github.com/cloud66)
- <b><code>&nbsp;&nbsp;7717⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;640🍴</code></b> [kompose](https://github.com/kubernetes/kompose) - Go from Docker Compose to Kubernetes
- <b><code>&nbsp;&nbsp;&nbsp;616⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;40🍴</code></b> [Maestro](https://github.com/toscanini/maestro) :skull: - Maestro provides the ability to easily launch, orchestrate and manage multiple Docker containers as single unit by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tascanini](https://github.com/toscanini)
- <b><code>&nbsp;&nbsp;&nbsp;131⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [Nuxx](https://github.com/nuxxapp/nuxx) - Visually generate docker-compose.yaml code with drag-n-drop UI. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@corpulent](https://github.com/corpulent)
- <b><code>&nbsp;&nbsp;&nbsp;158⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [percheron](https://github.com/ashmckenzie/percheron) :skull: - Organise your Docker containers with muscle and intelligence by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@ashmckenzie](https://github.com/ashmckenzie)
- <b><code>&nbsp;&nbsp;&nbsp;360⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [plash](https://github.com/ihucos/plash) - A container run and build engine - runs inside docker.
- <b><code>&nbsp;&nbsp;3026⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;283🍴</code></b> [podman-compose](https://github.com/containers/podman-compose) - a script to run docker-compose.yml using podman by [@containers][containers]
- <b><code>&nbsp;&nbsp;&nbsp;409⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [rocker-compose](https://github.com/grammarly/rocker-compose) :skull: - Docker composition tool with idempotency features for deploying apps composed of multiple containers. B 🌎 [@grammarly](grammarly).
- <b><code>&nbsp;&nbsp;1342⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81🍴</code></b> [rocker](https://github.com/grammarly/rocker) :skull: - Extended Dockerfile builder. Supports multiple FROMs, MOUNTS, templates, etc. by 🌎 [@grammarly](grammarly).
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Smalte](https://github.com/roquie/smalte) – Dynamically configure applications that require static configuration in docker container. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@roquie](https://github.com/roquie)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;63⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Stacker](https://github.com/stacker/stacker-cli) :skull: - Docker Compose Templates. Stacker provides an abstraction layer over Docker Compose and a better DX (developer experience).
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Stitchocker](https://github.com/alexaandrov/stitchocker) - A lightweight and fast command line utility for conveniently grouping your docker-compose multiple container services. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@alexaandrov](https://github.com/alexaandrov)
- <b><code>&nbsp;&nbsp;&nbsp;194⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [Zodiac](https://github.com/CenturyLinkLabs/zodiac) :skull: - A lightweight tool for easy deployment and rollback of dockerized applications. By [@CenturyLinkLabs][centurylinklabs]

### Deployment and Infrastructure

- <b><code>&nbsp;&nbsp;&nbsp;350⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;72🍴</code></b> [awesome-stacks](https://github.com/ethibox/awesome-stacks) - Deploy 80+ open-source web apps with one Docker command
- 🌎 [blackfish](https://gitlab.com/blackfish/blackfish) - a CoreOS VM to build swarm clusters for Dev & Production by 🌎 [@blackfish](https://gitlab.com/blackfish/)
- 🌎 [BosnD](https://gitlab.com/n0r1sk/bosnd) - BosnD, the boatswain daemon - A dynamic configuration file writer & service reloader for dynamically changing container environments.
- <b><code>&nbsp;&nbsp;1717⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;111🍴</code></b> [Centurion](https://github.com/newrelic/centurion) - Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@newrelic](https://github.com/newrelic)
- <b><code>&nbsp;&nbsp;&nbsp;425⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;66🍴</code></b> [Clocker](https://github.com/brooklyncentral/clocker) - Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts, on both Docker and Marathon. It leverages [Calico][calico] and [Weave][weave] for networking and 🌎 [Brooklyn](https://brooklyn.apache.org/) for application blueprints. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@brooklyncentral](https://github.com/brooklyncentral)
- <b><code>&nbsp;&nbsp;&nbsp;105⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Conduit](https://github.com/ehazlett/conduit) - Experimental deployment system for Docker by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@ehazlett](https://github.com/ehazlett)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;91⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [depcon](https://github.com/ContainX/depcon) - Depcon is written in Go and allows you to easily deploy Docker containers to Apache Mesos/Marathon, Amazon ECS and Kubernetes. By [@ContainX][containx]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;54⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [deploy](https://github.com/ttiny/deploy) :skull: - Git and Docker deployment tool. A middle ground between simple Docker composition tools and full blown cluster orchestration by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@ttiny](https://github.com/ttiny)
- <b><code>&nbsp;&nbsp;&nbsp;106⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [dockit](https://github.com/humblec/dockit) :skull: - Do docker actions and Deploy gluster containers! By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@humblec](https://github.com/humblec)
- <b><code>&nbsp;&nbsp;3661⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;209🍴</code></b> [gitkube](https://github.com/hasura/gitkube) - Gitkube is a tool for building and deploying docker images on Kubernetes using `git push`. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Hasura](https://github.com/hasura/).
- <b><code>&nbsp;&nbsp;1351⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;282🍴</code></b> [Grafeas](https://github.com/grafeas/grafeas) - A common API for metadata about containers, from image and build details to security vulnerabilities. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [grafeas](https://github.com/grafeas)
- <b><code>&nbsp;&nbsp;&nbsp;424⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [Longshoreman](https://github.com/longshoreman/longshoreman) :skull: - Longshoreman automates application deployment using Docker. Just create a Docker repository (or use a service), configure the cluster using AWS or Digital Ocean (or whatever you like) and deploy applications using a Heroku-like CLI tool. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [longshoreman](https://github.com/longshoreman)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [SwarmManagement](https://github.com/hansehe/SwarmManagement) - Swarm Management is a python application, installed with pip. The application makes it easy to manage a Docker Swarm by configuring a single yaml file describing which stacks to deploy, and which networks, configs or secrets to create.
- <b><code>&nbsp;&nbsp;3157⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;154🍴</code></b> [werf](https://github.com/werf/werf) - werf is a CI/CD tool for building Docker images efficiently and deploying them to Kubernetes using GitOps by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@flant](https://github.com/flant)

### Monitoring

- <b><code>&nbsp;&nbsp;&nbsp;699⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;151🍴</code></b> [Autoheal](https://github.com/willfarrell/docker-autoheal) - Monitor and restart unhealthy docker containers automatically.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Axibase Collector](https://github.com/axibase/docker-axibase-collector) - Axibase Collector streams performance counters, configuration changes and lifecycle events from the Docker engine(s) into Axibase Time Series Database for roll-up dashboards and integration with upstream monitoring systems.
- <b><code>&nbsp;13492⭐</code></b> <b><code>&nbsp;&nbsp;2036🍴</code></b> [cAdvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers. Created by [@Google][google]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;92⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [Docker-Alertd](https://github.com/deltaskelta/docker-alertd) - Monitor and send alerts based on docker container resource usage/statistics
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;78⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37🍴</code></b> [Docker-Flow-Monitor](https://github.com/docker-flow/docker-flow-monitor) - Reconfigures Prometheus when a new service is updated or deployed automatically by [@docker-flow][docker-flow]
- <b><code>&nbsp;&nbsp;&nbsp;202⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;28🍴</code></b> [Dockerana](https://github.com/dockerana/dockerana) :skull: - packaged version of Graphite and Grafana, specifically targeted at metrics from Docker.
- 🌎 [DockProc](https://gitlab.com/n0r1sk/dockproc) - I/O monitoring for containers on processlevel.
- <b><code>&nbsp;&nbsp;4838⭐</code></b> <b><code>&nbsp;&nbsp;1446🍴</code></b> [dockprom](https://github.com/stefanprodan/dockprom) - Docker hosts and containers monitoring with Prometheus, Grafana, cAdvisor, NodeExporter and AlertManager by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@stefanprodan](https://github.com/stefanprodan)
- <b><code>&nbsp;&nbsp;1564⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;117🍴</code></b> [Dozzle](https://github.com/amir20/dozzle) - Monitor container logs in real-time with a browser or mobile device. <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@amir20](https://github.com/amir20)
- 🌎 [Dynatrace](https://www.dynatrace.com/solutions/container-monitoring/) :heavy_dollar_sign: - Monitor containerized applications without installing agents or modifying your Run commands
- <b><code>&nbsp;20605⭐</code></b> <b><code>&nbsp;&nbsp;1299🍴</code></b> [Glances](https://github.com/nicolargo/glances) - A cross-platform curses-based system monitoring tool written in Python by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@nicolargo](https://github.com/nicolargo)
- 🌎 [Grafana Docker Dashboard Template](https://grafana.com/grafana/dashboards/179) - A template for your Docker, Grafana and Prometheus stack [@vegasbrianc][vegasbrianc]
- <b><code>&nbsp;&nbsp;&nbsp;455⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;134🍴</code></b> [InfluxDB, cAdvisor, Grafana](https://github.com/vegasbrianc/docker-monitoring) - InfluxDB Time series DB in combination with Grafana and cAdvisor by [@vegasbrianc][vegasbrianc]
- <b><code>&nbsp;&nbsp;&nbsp;134⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [LogJam](https://github.com/gocardless/logjam) - Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward these messages on to a log collection server (such as logstash) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@gocardless](https://github.com/gocardless)
- <b><code>&nbsp;&nbsp;4435⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;688🍴</code></b> [Logspout](https://github.com/gliderlabs/logspout) - Log routing for Docker container logs by [@gliderlabs][gliderlabs]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [monit-docker](https://github.com/decryptus/monit-docker) - Monitor docker containers resources usage or status and execute docker commands or inside containers. [@decryptus][decryptus]
- <b><code>&nbsp;&nbsp;&nbsp;556⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;71🍴</code></b> [NexClipper](https://github.com/NexClipper/NexClipper) - NexClipper is the container monitoring and performance management solution specialized in Docker, Apache Mesos, Marathon, DC/OS, Mesosphere, Kubernetes by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Nexclipper](https://github.com/NexClipper)
- <b><code>&nbsp;&nbsp;&nbsp;497⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;119🍴</code></b> [Out-of-the-box Host/Container Monitoring/Logging/Alerting Stack](https://github.com/uschtwill/docker_monitoring_logging_alerting) - Docker host and container monitoring, logging and alerting out of the box using cAdvisor, Prometheus, Grafana for monitoring, Elasticsearch, Kibana and Logstash for logging and elastalert and Alertmanager for alerting. Set up in 5 Minutes. Secure mode for production use with built-in [Automated Nginx Reverse Proxy (jwilder's)][nginxproxy].
- 🌎 [SuperVisor CPM](https://t0xic0der.medium.com/simply-accessible-container-performance-monitoring-with-supervisor-7fb47f925f3b) <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Frontend Service](https://github.com/t0xic0der/supervisor-frontend-service/) and <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Driver Service](https://github.com/t0xic0der/supervisor-driver-service/) :construction: - A simple and accessible FOSS container performance monitoring service written in Python by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@t0xic0der](https://github.com/t0xic0der/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [SwarmAlert](https://github.com/gpulido/SwarmAlert) - Monitors a Docker Swarm and sends Pushover alerts when it finds a container with no healthy service task running.
- <b><code>&nbsp;&nbsp;1153⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;272🍴</code></b> [Zabbix Docker module](https://github.com/monitoringartist/Zabbix-Docker-Monitoring) - Zabbix module that provides discovery of running containers, CPU/memory/blk IO/net container metrics. Systemd Docker and LXC execution driver is also supported. It's a dynamically linked shared object library, so its performance is (~10x) better, than any script solution.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;47⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [Zabbix Docker](https://github.com/gomex/docker-zabbix) - Monitor containers automatically using zabbix LLD feature.

### Networking

- [Calico-Docker][calico] - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Flannel](https://github.com/coreos/flannel/) - Flannel is a virtual network that gives a subnet to each host for use with container runtimes. By [@coreos][coreos]
- <b><code>&nbsp;&nbsp;&nbsp;519⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;74🍴</code></b> [Freeflow](https://github.com/Microsoft/Freeflow) - High performance container overlay networks on Linux. Enabling RDMA (on both InfiniBand and RoCE) and accelerating TCP to bare metal performance. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Microsoft](https://github.com/Microsoft)
- <b><code>&nbsp;&nbsp;4690⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;651🍴</code></b> [netshoot](https://github.com/nicolaka/netshoot) - The netshoot container has a powerful set of networking tools to help troubleshoot Docker networking issues by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@nicolaka](https://github.com/nicolaka)
- <b><code>&nbsp;&nbsp;4001⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;725🍴</code></b> [Pipework](https://github.com/jpetazzo/pipework) - Software-Defined Networking for Linux Containers, Pipework works with "plain" LXC containers, and with the awesome Docker. By [@jpetazzo][jpetazzo]
- [Weave][weave] (The Docker network) - Weave creates a virtual network that connects Docker containers deployed across multiple hosts.

### Orchestration

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;87⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24🍴</code></b> [athena](https://github.com/athena-oss/athena) - An automation platform with a plugin architecture that allows you to easily create and share services.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [blimp](https://github.com/tubesandlube/blimp) :skull: - Uses Docker Machine to easily move a container from one Docker host to another, show containers running against all of your hosts, replicate a container across multiple hosts and more by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@defermat](https://github.com/defermat) and <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@schvin](https://github.com/schvin)
- <b><code>&nbsp;&nbsp;&nbsp;220⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;78🍴</code></b> [CloudSlang](https://github.com/CloudSlang/cloud-slang) - CloudSlang is a workflow engine to create Docker process automation
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [clusterdock](https://github.com/clusterdock/clusterdock) - Docker container orchestration to enable the testing of long-running cluster deployments
- <b><code>&nbsp;&nbsp;&nbsp;748⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;169🍴</code></b> [Crane](https://github.com/Dataman-Cloud/crane) - Control plane based on docker built-in swarm <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Dataman-Cloud](https://github.com/Dataman-Cloud)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;64⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;52🍴</code></b> [Docker Flow Swarm Listener](https://github.com/docker-flow/docker-flow-swarm-listener) - Docker Flow Swarm Listener project is to listen to Docker Swarm events and send requests when a change occurs. By [@docker-flow][docker-flow]
- <b><code>&nbsp;&nbsp;&nbsp;268⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31🍴</code></b> [gantryd](https://github.com/DevTable/gantryd) :skull: - A framework for easy management of docker-based components across machines by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@DevTable](https://github.com/DevTable)
- <b><code>&nbsp;&nbsp;&nbsp;246⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;40🍴</code></b> [Haven](https://github.com/codeabovelab/haven-platform) - Haven is a simplified container management platform that integrates container, application, cluster, image, and registry managements. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@codeabovelab](https://github.com/codeabovelab)
- <b><code>&nbsp;&nbsp;2084⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;254🍴</code></b> [Helios](https://github.com/spotify/helios) :skull: - A simple platform for deploying and managing containers across an entire fleet of servers by [@spotify][spotify]
- <b><code>&nbsp;&nbsp;1484⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;129🍴</code></b> [Kontena](https://github.com/kontena/kontena) :skull: - The developer friendly container and micro services platform. Works on any cloud, easy to setup, simple to use.
- <b><code>&nbsp;89187⭐</code></b> <b><code>&nbsp;32685🍴</code></b> [Kubernetes](https://github.com/kubernetes/kubernetes) - Open source orchestration system for Docker containers by Google
- <b><code>&nbsp;&nbsp;1194⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;836🍴</code></b> [ManageIQ](https://github.com/ManageIQ/manageiq) - Discover, optimize and control your hybrid IT. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [ManageIQ](https://github.com/ManageIQ)
- <b><code>&nbsp;&nbsp;3034⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;433🍴</code></b> [Mantl](https://github.com/mantl/mantl) - Mantl is a modern platform for rapidly deploying globally distributed services
- <b><code>&nbsp;&nbsp;4044⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;872🍴</code></b> [Marathon](https://github.com/mesosphere/marathon) - Marathon is a private PaaS built on Mesos. It automatically handles hardware or software failures and ensures that an app is "always on"
- <b><code>&nbsp;&nbsp;4908⭐</code></b> <b><code>&nbsp;&nbsp;1704🍴</code></b> [Mesos](https://github.com/apache/mesos) - Resource/Job scheduler for containers, VM's and physical hosts 🌎 [@apache](https://mesos.apache.org/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Nebula](https://github.com/nebula-orchestrator) - A Docker orchestration tool designed to manage massive scale distributed clusters.
- <b><code>&nbsp;11158⭐</code></b> <b><code>&nbsp;&nbsp;1559🍴</code></b> [Nomad](https://github.com/hashicorp/nomad) - Easily deploy applications at any scale. A Distributed, Highly Available, Datacenter-Aware Scheduler by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@hashicorp](https://github.com/hashicorp)
- <b><code>&nbsp;&nbsp;1448⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;168🍴</code></b> [Panamax](https://github.com/CenturyLinkLabs/panamax-ui) :skull: - An open-source project that makes deploying complex containerized apps as easy as Drag-and-Drop by [@CenturyLinkLabs][centurylinklabs].
- <b><code>&nbsp;19289⭐</code></b> <b><code>&nbsp;&nbsp;2582🍴</code></b> [Rancher](https://github.com/rancher/rancher) - An open source project that provides a complete platform for operating Docker in production by [@rancher][rancher].
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;46⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [RedHerd Framework](https://github.com/redherd-project/redherd-framework) - RedHerd is a collaborative and serverless framework for orchestrating a geographically distributed group of assets capable of simulating complex offensive cyberspace operations. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@RedHerdProject](https://github.com/redherd-project).
- <b><code>&nbsp;&nbsp;&nbsp;529⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;54🍴</code></b> [Swarm-cronjob](https://github.com/crazy-max/swarm-cronjob) - Create jobs on a time-based schedule on Swarm by 🌎 [@crazy-max](crazy-max)

### PaaS

- <b><code>&nbsp;&nbsp;&nbsp;387⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37🍴</code></b> [Atlantis](https://github.com/ooyala/atlantis) :skull: - Atlantis is an Open Source PaaS for HTTP applications built on Docker and written in Go
- <b><code>&nbsp;&nbsp;8702⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;601🍴</code></b> [caprover](https://github.com/caprover/caprover) - [previously known as CaptainDuckDuck] Automated Scalable Webserver Package (automated Docker+nginx) - Heroku on Steroids
- <b><code>&nbsp;&nbsp;1854⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;183🍴</code></b> [Convox Rack](https://github.com/convox/rack) - Convox Rack is open source PaaS built on top of expert infrastructure automation and devops best practices.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Dcw](https://github.com/pbertera/dcw) - Docker-compose SSH wrapper: a very poor man PaaS, exposing the docker-compose and custom-container commands defined in container labels.
- <b><code>&nbsp;23203⭐</code></b> <b><code>&nbsp;&nbsp;1745🍴</code></b> [Dokku](https://github.com/dokku/dokku) - Docker powered mini-Heroku that helps you build and manage the lifecycle of applications (originally by [@progrium][progrium])
- <b><code>&nbsp;&nbsp;2667⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;163🍴</code></b> [Empire](https://github.com/remind101/empire) - A PaaS built on top of Amazon EC2 Container Service (ECS)
- <b><code>&nbsp;&nbsp;1016⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;61🍴</code></b> [Exoframe](https://github.com/exoframejs/exoframe) - A self-hosted tool that allows simple one-command deployments using Docker
- <b><code>&nbsp;&nbsp;7885⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;612🍴</code></b> [Flynn](https://github.com/flynn/flynn) - A next generation open source platform as a service
- <b><code>&nbsp;&nbsp;&nbsp;364⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37🍴</code></b> [Hephy Workflow](https://github.com/teamhephy/workflow) - Open source PaaS for Kubernetes that adds a developer-friendly layer to any Kubernetes cluster, making it easy to deploy and manage applications. Fork of <b><code>&nbsp;&nbsp;1298⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;187🍴</code></b> [Deis Workflow](https://github.com/deis/workflow)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;69⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [Krane](https://github.com/krane/krane) - Toolset for managing container workloads on remote servers
- <b><code>&nbsp;&nbsp;1553⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;98🍴</code></b> [Nanobox](https://github.com/nanobox-io/nanobox) :heavy_dollar_sign: - An application development platform that creates local environments that can then be deployed and scaled in the cloud.
- [OpenShift][openshift] - An open source PaaS built on [Kubernetes][kubernetes] and optimized for Dockerized app development and deployment by 🌎 [Red Hat](https://www.redhat.com/en)
- <b><code>&nbsp;&nbsp;3997⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;501🍴</code></b> [Tsuru](https://github.com/tsuru/tsuru) - Tsuru is an extensible and open source Platform as a Service software

### Reverse Proxy

- <b><code>&nbsp;&nbsp;2540⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;163🍴</code></b> [bunkerized-nginx](https://github.com/bunkerity/bunkerized-nginx) - Web app hosting and reverse proxy secure by default. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@bunkerity](https://github.com/bunkerity)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Docker Dnsmasq Updater](https://github.com/moonbuggy/docker-dnsmasq-updater) - Update a remote dnsmasq server with Docker container hostnames.
- <b><code>&nbsp;&nbsp;&nbsp;287⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;192🍴</code></b> [docker-flow-proxy](https://github.com/docker-flow/docker-flow-proxy) - Reconfigures proxy every time a new service is deployed, or when a service is scaled. By [@docker-flow][docker-flow]
- <b><code>&nbsp;&nbsp;&nbsp;258⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;54🍴</code></b> [docker-proxy](https://github.com/silarsis/docker-proxy) :skull: - Transparent proxy for docker containers, run in a docker container. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@silarsis](https://github.com/silarsis)
- <b><code>&nbsp;&nbsp;6943⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;619🍴</code></b> [fabio](https://github.com/fabiolb/fabio) - A fast, modern, zero-conf load balancing HTTP(S) router for deploying microservices managed by consul. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@magiconair](https://github.com/magiconair) (Frank Schroeder)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;44⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [h2o-proxy](https://github.com/zchee/h2o-proxy) :skull: - Automated H2O reverse proxy for Docker containers. An alternative to [jwilder/nginx-proxy][nginxproxy] by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@zchee](https://github.com/zchee)
- <b><code>&nbsp;&nbsp;6568⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;771🍴</code></b> [Let's Encrypt Nginx-proxy Companion](https://github.com/nginx-proxy/docker-letsencrypt-nginx-proxy-companion) - A lightweight companion container for the nginx-proxy. It allow the creation/renewal of Let's Encrypt certificates automatically. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@JrCs](https://github.com/JrCs)
- <b><code>&nbsp;&nbsp;&nbsp;166⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [muguet](https://github.com/mattallty/muguet) :skull: - DNS Server & Reverse proxy for Docker environments. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mattallty](https://github.com/mattallty)
- <b><code>&nbsp;&nbsp;8104⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;992🍴</code></b> [Nginx Proxy Manager](https://github.com/jc21/nginx-proxy-manager) - A beautiful web interface for proxying web based services with SSL. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jc21](https://github.com/jc21)
- [nginx-proxy][nginxproxy] - Automated nginx proxy for Docker containers using docker-gen by [@jwilder][jwilder]
- <b><code>&nbsp;&nbsp;&nbsp;168⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Swarm Ingress Router](https://github.com/tpbowden/swarm-ingress-router) :skull: - Route DNS names to Swarm services based on labels. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tpbowden](https://github.com/tpbowden/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;57⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Swarm Router](https://github.com/flavioaiello/swarm-router) - A «zero config» service name based router for docker swarm mode with a fresh and more secure approach. By 🌎 [@flavioaiello](https://twitter.com/flavioaiello)
- <b><code>&nbsp;38438⭐</code></b> <b><code>&nbsp;&nbsp;4196🍴</code></b> [Træfɪk](https://github.com/containous/traefik) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd... By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@EmileVauge](https://github.com/emilevauge)

### Runtime

- <b><code>&nbsp;&nbsp;1457⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;87🍴</code></b> [aind](https://github.com/aind-containers/aind) - AinD launches Android apps in Docker, by nesting Anbox containers inside Docker by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@aind-containers](https://github.com/aind-containers)
- <b><code>&nbsp;&nbsp;4009⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;798🍴</code></b> [cri-o](https://github.com/cri-o/cri-o) - Open Container Initiative-based implementation of Kubernetes Container Runtime Interface by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [cri-o](https://github.com/cri-o)
- <b><code>&nbsp;&nbsp;3714⭐</code></b> <b><code>&nbsp;&nbsp;1017🍴</code></b> [lxc](https://github.com/lxc/lxc) - LXC - Linux Containers
- <b><code>&nbsp;14128⭐</code></b> <b><code>&nbsp;&nbsp;1463🍴</code></b> [podman](https://github.com/containers/libpod) - libpod is a library used to create container pods. Home of Podman by [@containers][containers]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [rlxc](https://github.com/brauner/rlxc) - LXC binary written in Rust by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@brauner](https://github.com/brauner)
- <b><code>&nbsp;&nbsp;&nbsp;321⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;124🍴</code></b> [runtime-tools](https://github.com/opencontainers/runtime-tools) - oci-runtime-tool is a collection of tools for working with the OCI runtime specification by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@opencontainers](https://github.com/opencontainers)

### Security

- <b><code>&nbsp;&nbsp;&nbsp;355⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;60🍴</code></b> [Anchor Engine](https://github.com/anchore/anchore) - Analyze images for CVE vulnerabilities and against custom security policies by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Anchor](https://github.com/anchore)
- 🌎 [Aqua Security](https://www.aquasec.com) :heavy_dollar_sign: - Securing container-based applications from Dev to Production on any platform
- <b><code>&nbsp;&nbsp;1035⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;80🍴</code></b> [bane](https://github.com/genuinetools/bane) - AppArmor profile generator for Docker containers by [@genuinetools][genuinetools]
- <b><code>&nbsp;&nbsp;&nbsp;396⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;98🍴</code></b> [CIS Docker Benchmark](https://github.com/dev-sec/cis-docker-benchmark) - This [InSpec][inspec] compliance profile implement the CIS Docker 1.12.0 Benchmark in an automated way to provide security best-practice tests around Docker daemon and containers in a production environment. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@dev-sec](https://github.com/dev-sec)
- <b><code>&nbsp;&nbsp;8786⭐</code></b> <b><code>&nbsp;&nbsp;1074🍴</code></b> [Clair](https://github.com/quay/clair) - Clair is an open source project for the static analysis of vulnerabilities in appc and docker containers. By [@coreos][coreos]
- <b><code>&nbsp;&nbsp;&nbsp;970⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;150🍴</code></b> [Dagda](https://github.com/eliasgranderubio/dagda) - Dagda is a tool to perform static analysis of known vulnerabilities, trojans, viruses, malware & other malicious threats in docker images/containers and to monitor the docker daemon and running docker containers for detecting anomalous activities. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@eliasgranderubio](https://github.com/eliasgranderubio)
- 🌎 [Deepfence Enterprise](https://deepfence.io) :heavy_dollar_sign: - Full life cycle Cloud Native Workload Protection platform for kubernetes, virtual machines and serverless. By 🌎 [@deepfence](deepfence)
- <b><code>&nbsp;&nbsp;1609⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;180🍴</code></b> [Deepfence Threat Mapper](https://github.com/deepfence/ThreatMapper) - Powerful runtime vulnerability scanner for kubernetes, virtual machines and serverless. By 🌎 [@deepfence](deepfence)
- <b><code>&nbsp;&nbsp;7808⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;890🍴</code></b> [docker-bench-security](https://github.com/docker/docker-bench-security) - script that checks for dozens of common best-practices around deploying Docker containers in production. By [@docker][docker]
- <b><code>&nbsp;&nbsp;&nbsp;403⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;39🍴</code></b> [docker-explorer](https://github.com/google/docker-explorer) - A tool to help forensicate offline docker acquisitions by [@Google][google]
- <b><code>&nbsp;&nbsp;&nbsp;364⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [docker-lock](https://github.com/safe-waters/docker-lock) - A cli-plugin for docker to automatically manage image digests by tracking them in a separate Lockfile. By [@safe-waters][safe-waters]
- <b><code>&nbsp;&nbsp;1061⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;160🍴</code></b> [KICS](https://github.com/checkmarx/kics) - an infrastructure-as-code scanning tool, find security vulnerabilities, compliance issues, and infrastructure misconfigurations early in the development cycle. Can be extended for additional policies. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Checkmarx](https://github.com/Checkmarx)
- <b><code>&nbsp;&nbsp;2833⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;509🍴</code></b> [notary](https://github.com/theupdateframework/notary) - a server and a client for running and interacting with trusted collections. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@TUF](https://github.com/theupdateframework)
- <b><code>&nbsp;&nbsp;&nbsp;941⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;307🍴</code></b> [oscap-docker](https://github.com/OpenSCAP/openscap) - OpenSCAP provides oscap-docker tool which is used to scan Docker containers and images. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [OpenSCAP](https://github.com/OpenSCAP)
- 🌎 [Prisma Cloud](https://www.paloaltonetworks.com/prisma/cloud) :heavy_dollar_sign: - (previously Twistlock Security Suite) detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.
- <b><code>&nbsp;&nbsp;4914⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;666🍴</code></b> [Sysdig Falco](https://github.com/falcosecurity/falco) - Sysdig Falco is an open source container security monitor. It can monitor application, container, host, and network activity and alert on unauthorized activity.
- 🌎 [Sysdig Secure](https://sysdig.com/products/secure/runtime-security/) :heavy_dollar_sign: - Sysdig Secure addresses run-time security through behavioral monitoring and defense, and provides deep forensics based on open source Sysdig for incident response.
- 🌎 [Trend Micro DeepSecurity](https://www.trendmicro.com/en_us/business/products/hybrid-cloud/deep-security.html) :heavy_dollar_sign: - Trend Micro DeepSecurity offers runtime protection for container workloads and hosts as well as preruntime scanning of images to identify vulnerabilities, malware and content such as hardcoded secrets.
- <b><code>&nbsp;12383⭐</code></b> <b><code>&nbsp;&nbsp;1156🍴</code></b> [Trivy](https://github.com/aquasecurity/trivy) - Aqua Security's open source simple and comprehensive vulnerability scanner for containers (suitable for CI).

### Service Discovery

- <b><code>&nbsp;&nbsp;1075⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;297🍴</code></b> [docker-consul](https://github.com/gliderlabs/docker-consul) by [@progrium][progrium]
- <b><code>&nbsp;40153⭐</code></b> <b><code>&nbsp;&nbsp;8621🍴</code></b> [etcd](https://github.com/etcd-io/etcd) - Distributed reliable key-value store for the most critical data of a distributed system by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@etcd-io](https://github.com/etcd-io) (former part of CoreOS)
- <b><code>&nbsp;30483⭐</code></b> <b><code>&nbsp;&nbsp;6400🍴</code></b> [istio](https://github.com/istio/istio) - An open platform to connect, manage, and secure microservices by 🌎 [@IstioMesh](istio)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;51⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [proxy](https://github.com/factorish/proxy) :skull: - lightweight nginx based load balancer self using service discovery provided by registrator. by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@factorish](https://github.com/factorish)
- <b><code>&nbsp;&nbsp;4574⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;921🍴</code></b> [registrator](https://github.com/gliderlabs/registrator) - Service registry bridge for Docker by [@gliderlabs][gliderlabs] and [@progrium][progrium]

### Volume Management / Data

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;92⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [Blockbridge](https://github.com/blockbridge/blockbridge-docker-volume) :heavy_dollar_sign:- The Blockbridge plugin is a volume plugin that provides access to an extensible set of container-based persistent storage options. It supports single and multi-host Docker environments with features that include tenant isolation, automated provisioning, encryption, secure deletion, snapshots and QoS. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@blockbridge](https://github.com/blockbridge)
- <b><code>&nbsp;&nbsp;1315⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;137🍴</code></b> [Convoy](https://github.com/rancher/convoy) - an open-source Docker volume driver that can snapshot, backup and restore Docker volumes anywhere. By [@rancher][rancher]
- <b><code>&nbsp;&nbsp;&nbsp;801⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;107🍴</code></b> [Docker Machine NFS](https://github.com/adlogix/docker-machine-nfs) Activates NFS for an existing boot2docker box created through Docker Machine on OS X.
- <b><code>&nbsp;&nbsp;&nbsp;166⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25🍴</code></b> [Docker Unison](https://github.com/leighmcculloch/docker-unison) A docker volume container using Unison for fast two-way folder sync. Created as an alternative to slow boot2docker volumes on OS X. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@leighmcculloch](https://github.com/leighmcculloch)
- <b><code>&nbsp;&nbsp;&nbsp;764⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;105🍴</code></b> [Local Persist](https://github.com/MatchbookLab/local-persist) Specify a mountpoint for your local volumes (created via `docker volume create`) so that files will always persist and so you can mount to different directories in different containers.
- <b><code>&nbsp;33550⭐</code></b> <b><code>&nbsp;&nbsp;3943🍴</code></b> [Minio](https://github.com/minio/minio) - S3 compatible object storage server in Docker containers
- <b><code>&nbsp;&nbsp;1061⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;162🍴</code></b> [Netshare](https://github.com/ContainX/docker-volume-netshare) Docker NFS, AWS EFS, Ceph & Samba/CIFS Volume Plugin. By [@ContainX][containx]
- 🌎 [portworx](https://portworx.com) :heavy_dollar_sign: - Decentralized storage solution for persistent, shared and replicated volumes.
- 🌎 [quobyte](https://www.quobyte.com/) :heavy_dollar_sign: - fully fault-tolerant distributed file system with a docker volume driver
- <b><code>&nbsp;&nbsp;2028⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;306🍴</code></b> [REX-Ray](https://github.com/rexray/rexray) provides a vendor agnostic storage orchestration engine. The primary design goal is to provide persistent storage for Docker, Kubernetes, and Mesos. B <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@thecodeteam](https://github.com/thecodeteam) (DELL Technologies)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Storidge](https://github.com/Storidge/quick-start) :heavy_dollar_sign: - Software-defined Persistent Storage for Kubernetes and Docker Swarm

### User Interface

#### IDE integrations

- JetBrains IDEs (IntelliJ IDEA, GoLand, WebStorm, CLion etc.) has 🌎 [built-in Docker plugin](https://www.jetbrains.com/help/idea/docker.html#managing-images)
- Eclipse 🌎 [Docker Tooling plugin](https://www.eclipse.org/community/eclipse_newsletter/2016/july/article2.php)
- <b><code>&nbsp;&nbsp;&nbsp;167⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [docker.vim](https://github.com/skanehira/docker.vim) - Manage docker containers and images in Vim. By [@skanehira]
- <b><code>&nbsp;&nbsp;&nbsp;612⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;66🍴</code></b> [docker.el](https://github.com/Silex/docker.el) Manage docker from Emacs by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Silex](https://github.com/Silex)

#### Desktop

Native desktop applications for managing and montoring docker hosts and clusters

- 🌎 [Docker Desktop](https://www.docker.com/products/docker-desktop/) - Official native app. Only for Windows and MacOS
- <b><code>&nbsp;&nbsp;&nbsp;779⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;114🍴</code></b> [Dockeron](https://github.com/dockeron/dockeron) - A project built on Electron + Vue.js for Docker on desktop. <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@fluency03](https://github.com/fluency03)
- 🌎 [DockerStacks](http://docker-stacks.com/) - Local LAMP/LEMP Development Studio <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@sfx101](https://github.com/sfx101)
- <b><code>&nbsp;&nbsp;1884⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;97🍴</code></b> [DockStation](https://github.com/DockStation/dockstation) - A developer centric UI to configure, monitor, and manage services and containers 🌎 [@dock_station](https://twitter.com/dock_station)
- <b><code>&nbsp;&nbsp;&nbsp;175⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23🍴</code></b> [Lifeboat](https://github.com/jplhomer/lifeboat) - An easy way to launch Docker projects with a graphical interface on your Mac. <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jplhomer](https://github.com/jplhomer)
- <b><code>&nbsp;&nbsp;&nbsp;580⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88🍴</code></b> [Simple Docker UI](https://github.com/felixgborrego/simple-docker-ui) - built on Electron. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@felixgborrego](https://github.com/felixgborrego/)
- <b><code>&nbsp;&nbsp;&nbsp;196⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [Stevedore](https://github.com/slonopotamus/stevedore) - Good Docker Desktop replacement for Windows. Both Linux and Windows Containers are supported. <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@slonopotamus](https://github.com/slonopotamus)

#### Terminal

##### Terminal UI

- <b><code>&nbsp;&nbsp;&nbsp;417⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;32🍴</code></b> [ctop (1)](https://github.com/yadutaf/ctop) - :skull: A command line / text based Linux Containers monitoring tool that works just like you expect (Python) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@yadutaf](https://github.com/yadutaf)
- <b><code>&nbsp;12922⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;494🍴</code></b> [ctop (2)](https://github.com/bcicen/ctop) - Top-like interface for container metrics (Golang) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@bcicen](https://github.com/bcicen/)
- <b><code>&nbsp;32213⭐</code></b> <b><code>&nbsp;&nbsp;1288🍴</code></b> [dive](https://github.com/wagoodman/dive) - A tool for exploring each layer in a docker image. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [wagoodman](https://github.com/wagoodman).
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;99⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [dockdash](https://github.com/byrnedo/dockdash) detailed stats. By [@byrnedo]
- <b><code>&nbsp;&nbsp;&nbsp;776⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38🍴</code></b> [Docker-mon](https://github.com/icecrime/docker-mon) :skull: - Console-based Docker monitoring by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@icecrime](https://github.com/icecrime)
- <b><code>&nbsp;&nbsp;3285⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;153🍴</code></b> [dockly](https://github.com/lirantal/dockly) - An interactive shell UI for managing Docker containers by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@lirantal](https://github.com/lirantal)
- <b><code>&nbsp;&nbsp;1681⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;187🍴</code></b> [DockSTARTer](https://github.com/GhostWriters/DockSTARTer) - DockSTARTer helps you get started with home server apps running in Docker by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [GhostWriters](https://github.com/GhostWriters)
- <b><code>&nbsp;&nbsp;2250⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;99🍴</code></b> [docui](https://github.com/skanehira/docui) - An interactive shell UI for managing Docker containers. Also works in Windows. By [@skanehira]
- <b><code>&nbsp;&nbsp;2542⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;92🍴</code></b> [dry](https://github.com/moncho/dry) - An interactive CLI for Docker containers by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@moncho](https://github.com/moncho)
- <b><code>&nbsp;23002⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;861🍴</code></b> [lazydocker](https://github.com/jesseduffield/lazydocker) - The lazier way to manage everything docker. A simple terminal UI for both docker and docker-compose, written in Go with the gocui library. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jesseduffield](https://github.com/jesseduffield)
- <b><code>&nbsp;&nbsp;&nbsp;898⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62🍴</code></b> [sen](https://github.com/TomasTomecek/sen) - :skull: Terminal user interface for docker engine, by [@TomasTomecek][tomastomecek]

##### CLI tools

- <b><code>&nbsp;&nbsp;&nbsp;192⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [captain](https://github.com/jenssegers/captain) - Easily start and stop docker compose projects from any directory. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jenssegers](https://github.com/jenssegers)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [dcinja](https://github.com/Falldog/dcinja) - The powerful and smallest binary size of template engine for docker command line environment. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Falldog](https://github.com/Falldog)
- <b><code>&nbsp;&nbsp;&nbsp;435⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;58🍴</code></b> [docker-ls](https://github.com/mayflower/docker-ls) - CLI tools for browsing and manipulating docker registries by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mayflower](https://github.com/mayflower)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;86⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker pushrm](https://github.com/christian-korneck/docker-pushrm) - A Docker CLI plugin that lets you push the README.md file from the current directory to Docker Hub. Also supports Quay and Harbor. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@christian-korneck](https://github.com/christian-korneck)
- <b><code>&nbsp;&nbsp;&nbsp;117⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [dockersql](https://github.com/crosbymichael/dockersql) - A command line interface to query Docker using SQL by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@crosbymichael](https://github.com/crosbymichael)
- <b><code>&nbsp;&nbsp;&nbsp;514⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;51🍴</code></b> [DVM](https://github.com/howtowhale/dvm) - Docker version manager by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@howtowhale](https://github.com/howtowhale)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [goinside](https://github.com/iamsoorena/goinside) - Get inside a running docker container easily. by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@iamsoorena](https://github.com/iamsoorena)
- <b><code>&nbsp;&nbsp;2552⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;280🍴</code></b> [ns-enter](https://github.com/jpetazzo/nsenter) - :skull: no more ssh, enter name spaces of container by [@jpetazzo][jpetazzo]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Pdocker](https://github.com/g31s/Pdocker) - A simple tool to manage and maintain Docker for personal projects by 🌎 [@g31s](g31s)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;80⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [proco](https://github.com/shiwaforce/poco) - Proco will help you to organise and manage Docker, Docker-Compose, Kubernetes projects of any complexity using simple YAML config files to shorten the route from finding your project to initialising it in your local environment. by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@shiwaforce](https://github.com/shiwaforce)
- <b><code>&nbsp;&nbsp;1535⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;158🍴</code></b> [reg](https://github.com/genuinetools/reg) - Docker registry v2 command line client by [@genuinetools][genuinetools]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;64⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [scuba](https://github.com/JonathonReinhart/scuba) - Transparently use Docker containers to encapsulate software build environments, by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@JonathonReinhart](https://github.com/JonathonReinhart)
- <b><code>&nbsp;&nbsp;4760⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;552🍴</code></b> [skopeo](https://github.com/containers/skopeo) - Work with remote images registries - retrieving information, images, signing content by [@containers][containers]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [supdock](https://github.com/segersniels/supdock) - Allows for slightly more visual usage of Docker with an interactive prompt. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@segersniels](https://github.com/segersniels)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;51⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [tsaotun](https://github.com/qazbnm456/tsaotun) - Python based Assistance for Docker by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@qazbnm456](https://github.com/qazbnm456)
- <b><code>&nbsp;&nbsp;&nbsp;650⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38🍴</code></b> [wharfee](https://github.com/j-bennet/wharfee) - Autocompletion and syntax highlighting for Docker commands. by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@j-bennet](https://github.com/j-bennet)

##### Other

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [dext-docker-registry-plugin](https://github.com/vutran/dext-docker-registry-plugin) - Search the Docker Registry with the Dext smart launcher. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@vutran](https://github.com/vutran)
- <b><code>&nbsp;&nbsp;&nbsp;602⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;90🍴</code></b> [docker-ssh](https://github.com/jeroenpeeters/docker-ssh) - SSH Server for Docker containers ~ Because every container should be accessible. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jeroenpeeters](https://github.com/jeroenpeeters)
- <b><code>&nbsp;&nbsp;6642⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;391🍴</code></b> [dockercraft](https://github.com/docker/dockercraft) - Docker + Minecraft = Dockercraft by [@docker][docker]
- <b><code>&nbsp;&nbsp;&nbsp;455⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;71🍴</code></b> [dockerfile-mode](https://github.com/spotify/dockerfile-mode) An emacs mode for handling Dockerfiles by [@spotify][spotify]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [MultiDocker](https://github.com/marty90/multidocker) - Create a secure multi-user Docker machine, where each user is segregated into an indepentent container.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;60⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Powerline-Docker](https://github.com/adrianmo/powerline-docker) - A Powerline segment for showing the status of Docker containers by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@adrianmo](https://github.com/adrianmo)

#### Web

- <b><code>&nbsp;&nbsp;&nbsp;254⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;67🍴</code></b> [Admiral](https://github.com/vmware/admiral) Admiral™ is a highly scalable and very lightweight Container Management platform for deploying and managing container based applications. By [VMWare]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [CASA](https://github.com/knrdl/casa) - Outsource the administration of a handful of containers to your co-workers, by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@knrdl](https://github.com/knrdl)
- <b><code>&nbsp;&nbsp;&nbsp;193⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35🍴</code></b> [Container Web TTY](https://github.com/wrfly/container-web-tty) - Connect your containers via a web-tty <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@wrfly](https://github.com/wrfly)
- <b><code>&nbsp;&nbsp;1467⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;219🍴</code></b> [Docker Compose UI](https://github.com/francescou/docker-compose-ui) - Manage docker-compose via HTTP. docker-compose-ui runs in a Docker container, mounts the hosts docker socket and exposes a RESTful API and AngularJS GUI
- <b><code>&nbsp;&nbsp;&nbsp;318⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38🍴</code></b> [Docker Registry Browser](https://github.com/klausmeyer/docker-registry-browser) - Web Interface for the Docker Registry HTTP API v2 by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@klausmeyer](https://github.com/klausmeyer)
- <b><code>&nbsp;&nbsp;1263⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;218🍴</code></b> [Docker Registry UI (Joxit)](https://github.com/Joxit/docker-registry-ui) - The simplest and cleanest UI for private registries by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Joxit](https://github.com/Joxit)
- <b><code>&nbsp;&nbsp;&nbsp;892⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;152🍴</code></b> [Docker Registry UI](https://github.com/atcol/docker-registry-ui) - A web UI for easy private/local Docker Registry integration by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@atcol](https://github.com/atcol)
- <b><code>&nbsp;&nbsp;&nbsp;505⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;129🍴</code></b> [docker-registry-web](https://github.com/mkuchin/docker-registry-web) - Web UI, authentication service and event recorder for private docker registry v2 by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mkuchin](https://github.com/mkuchin)
- <b><code>&nbsp;&nbsp;3050⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;527🍴</code></b> [docker-swarm-visualizer](https://github.com/dockersamples/docker-swarm-visualizer) - Visualizes Docker services on a Docker Swarm (for running demos).
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [dockering-on-rails](https://github.com/Electrofenster/dockerding-on-rails) :skull: - Simple Web-Interface for Docker with a lot of features by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Electrofenster](https://github.com/Electrofenster/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [DockerSurfer](https://github.com/Simone-Erba/DockerSurfer) :skull: - A web service for analyze and browse dependencies between Docker images in the Docker registry, by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Simone-Erba](https://github.com/Simone-Erba/)
- <b><code>&nbsp;&nbsp;1224⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81🍴</code></b> [Kubevious](https://github.com/kubevious/kubevious) - A highly visual web UI for Kubernetes which renders configuration and state in an application centric way by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@rubenhak](https://github.com/rubenhak).
- <b><code>&nbsp;59583⭐</code></b> <b><code>&nbsp;&nbsp;5307🍴</code></b> [netdata](https://github.com/netdata/netdata) - Real-time performance monitoring
- <b><code>&nbsp;&nbsp;4960⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;313🍴</code></b> [OctoLinker](https://github.com/OctoLinker/OctoLinker) - A browser extension for GitHub that makes the image name in a `Dockerfile` clickable and redirect you to the related Docker Hub page.
- <b><code>&nbsp;22025⭐</code></b> <b><code>&nbsp;&nbsp;1976🍴</code></b> [Portainer](https://github.com/portainer/portainer) - A lightweight management UI for managing your Docker hosts or Docker Swarm clusters by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@portainer](https://github.com/portainer)
- <b><code>&nbsp;&nbsp;&nbsp;125⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24🍴</code></b> [Rapid Dashboard](https://github.com/ozlerhakan/rapid) - A simple query dashboard to use Docker Remote API by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@ozlerhakan](https://github.com/ozlerhakan/)
- <b><code>&nbsp;&nbsp;1911⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;279🍴</code></b> [Seagull](https://github.com/tobegit3hub/seagull) - Friendly Web UI to monitor docker daemon. by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tobegit3hub](https://github.com/tobegit3hub)
- <b><code>&nbsp;&nbsp;2395⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;237🍴</code></b> [Swarmpit](https://github.com/swarmpit/swarmpit) - Swarmpit provides simple and easy to use interface for your Docker Swarm cluster. You can manage your stacks, services, secrets, volumes, networks etc.
- <b><code>&nbsp;&nbsp;&nbsp;404⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;67🍴</code></b> [Swirl](https://github.com/cuigh/swirl) - Swirl is a web management tool for Docker, focused on swarm cluster By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@cuigh](https://github.com/cuigh/)
- <b><code>&nbsp;16987⭐</code></b> <b><code>&nbsp;&nbsp;2171🍴</code></b> [Theia](https://github.com/eclipse-theia/theia) - Extensible platform to develop full-fledged multi-language Cloud & Desktop IDE-like products with state-of-the-art web technologies.
- <b><code>&nbsp;&nbsp;1733⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;99🍴</code></b> [Yacht](https://github.com/SelfhostedPro/Yacht) :construction: - A Web UI for docker that focuses on templates and ease of use in order to make deployments as easy as possible. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@SelfhostedPro](https://github.com/SelfhostedPro)

## Docker Images

### Base Tools

Tools and applications that are either installed inside containers or designed to be run as a 🌎 [sidecar](https://docs.microsoft.com/en-us/azure/architecture/patterns/sidecar)

- <b><code>&nbsp;&nbsp;&nbsp;802⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;56🍴</code></b> [amicontained](https://github.com/genuinetools/amicontained) - Container introspection tool. Find out what container runtime is being used as well as features available by [@genuinetools][genuinetools]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [autodock](https://github.com/prologic/autodock) - Daemon for Docker Automation by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@prologic](https://github.com/prologic)
- <b><code>&nbsp;&nbsp;&nbsp;174⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27🍴</code></b> [Chaperone](https://github.com/garywiz/chaperone) - A single PID1 process designed for docker containers. Does user management, log management, startup, zombie reaping, all in one small package. by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@garywiz](https://github.com/garywiz)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [ckron](https://github.com/nicomt/ckron) - A cron-style job scheduler for docker, by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@nicomt](https://github.com/nicomt)
- [CoreOS][coreos] - Linux for Massive Server Deployments
- <b><code>&nbsp;12970⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;741🍴</code></b> [distroless](https://github.com/GoogleContainerTools/distroless) - Language focused docker images, minus the operating system, by [@GoogleContainerTools][googlecontainertools]
- <b><code>&nbsp;&nbsp;5582⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;542🍴</code></b> [docker-alpine](https://github.com/gliderlabs/docker-alpine) - A super small Docker base image _(5MB)_ using Alpine Linux by [@gliderlabs][gliderlabs]
- <b><code>&nbsp;&nbsp;4032⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;579🍴</code></b> [docker-gen](https://github.com/jwilder/docker-gen) - Generate files from docker container meta-data by [@jwilder][jwilder]
- <b><code>&nbsp;&nbsp;4376⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;352🍴</code></b> [dockerize](https://github.com/jwilder/dockerize) - Utility to simplify running applications in docker containers by [@jwilder][jwilder]
- <b><code>&nbsp;&nbsp;4061⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;278🍴</code></b> [GoSu](https://github.com/tianon/gosu) - Run this specific application as this specific user and get out of the pipeline (entrypoint script tool) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tianon](https://github.com/tianon)
- <b><code>&nbsp;&nbsp;&nbsp;180⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [is-docker](https://github.com/sindresorhus/is-docker) - Check if the process is running inside a Docker container by [@sindresorhus][sindresorhus]
- <b><code>&nbsp;&nbsp;&nbsp;288⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25🍴</code></b> [lstags](https://github.com/ivanilves/lstags) - sync Docker images across registries by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@ivanilves](https://github.com/ivanilves)
- <b><code>&nbsp;14710⭐</code></b> <b><code>&nbsp;&nbsp;1899🍴</code></b> [NVIDIA-Docker](https://github.com/NVIDIA/nvidia-docker) - The NVIDIA Container Runtime for Docker by [@NVIDIA][nvidia]
- <b><code>&nbsp;&nbsp;&nbsp;848⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;84🍴</code></b> [su-exec](https://github.com/ncopa/su-exec) - This is a simple tool that will simply execute a program with different privileges. The program will be executed directly and not run as a child, like su and sudo does, which avoids TTY and signal issues. Why reinvent gosu? This does more or less exactly the same thing as gosu but it is only 10kb instead of 1.8MB. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [ncopa](https://github.com/ncopa)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [sue](https://github.com/theAkito/sue) - Executes a program as a user different from the user running sue. This is a maintainable alternative to ncopa/su-exec, which is the better tianon/gosu. This one is far better (higher performance, smaller size), than the original gosu, however it is far easier to maintain, than su-exec, which is written in plain C. Made by [Akito][Akito]
- <b><code>&nbsp;&nbsp;1083⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;71🍴</code></b> [supercronic](https://github.com/aptible/supercronic) - crontab-compatible job runner, designed specifically to run in containers by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@aptible](https://github.com/aptible/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [TrivialRC](https://github.com/vorakl/TrivialRC) - A minimalistic Runtime Configuration system and process manager for containers <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@vorakl](https://github.com/vorakl)

### Builder

Applications designed to help or simplify building **new** images

- <b><code>&nbsp;&nbsp;&nbsp;585⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;64🍴</code></b> [ansible-bender](https://github.com/ansible-community/ansible-bender) - A tool utilising `ansible` and `buildah` by [@TomasTomecek][tomastomecek]
- <b><code>&nbsp;&nbsp;5392⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;602🍴</code></b> [buildah](https://github.com/containers/buildah) - A tool that facilitates building OCI images by [@containers][containers]
- <b><code>&nbsp;&nbsp;5307⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;692🍴</code></b> [BuildKit](https://github.com/moby/buildkit) - Concurrent, cache-efficient, and Dockerfile-agnostic builder toolkit by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@moby project](https://github.com/moby)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;63⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34🍴</code></b> [cekit](https://github.com/cekit/cekit) - A tool used by openshift to build base images using different build engines by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@cekit](https://github.com/cekit).
- <b><code>&nbsp;&nbsp;3197⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;210🍴</code></b> [container-diff](https://github.com/GoogleContainerTools/container-diff) - An image tool for comparing and analyzing container images by [@GoogleContainerTools][googlecontainertools]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;58⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [container-factory](https://github.com/mutable/container-factory) - Produces Docker images from tarballs of application source code by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mutable](https://github.com/mutable)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [copy-docker-image](https://github.com/mdlavin/copy-docker-image) - Copy a Docker image between registries without a full Docker installation by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mdlavin](https://github.com/mdlavin)
- <b><code>&nbsp;&nbsp;&nbsp;635⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;111🍴</code></b> [Derrick](https://github.com/alibaba/derrick) - A tool help you to automate the generation of Dockerfile and dockerize application by scanning the code. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@alibaba](https://github.com/alibaba).
- <b><code>&nbsp;&nbsp;&nbsp;362⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [dlayer](https://github.com/orisano/dlayer) - docker layer analyzer by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@orisano](https://github.com/orisano)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [docker-companion](https://github.com/mudler/docker-companion) - A command line tool written in Golang to squash and unpack docker images by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mudler](https://github.com/mudler/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;89⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22🍴</code></b> [docker-make](https://github.com/CtripCloud/docker-make) - Build, tag,and push a bunch of related docker images via a single command.
- <b><code>&nbsp;&nbsp;&nbsp;178⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [docker-replay](https://github.com/bcicen/docker-replay) - Generate `docker run`command and options from running containers. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [bcicen](https://github.com/bcicen)
- <b><code>&nbsp;&nbsp;&nbsp;126⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [DockerMake](https://github.com/avirshup/DockerMake) - A reproducible Docker image build system for complex software stacks. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@avirshup](https://github.com/avirshup)
- <b><code>&nbsp;13726⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;503🍴</code></b> [DockerSlim](https://github.com/docker-slim/docker-slim) shrinks fat Docker images creating the smallest possible images.
- <b><code>&nbsp;&nbsp;&nbsp;218⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Dockly](https://github.com/swipely/dockly) - Dockly is a gem made to ease the pain of packaging an application in Docker by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@swipely](https://github.com/swipely/)
- <b><code>&nbsp;&nbsp;&nbsp;263⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [dockramp](https://github.com/jlhawn/dockramp) :skull: - Proof of Concept: A Client Driven Docker Image Builder by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jlhawn](https://github.com/jlhawn)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [essex](https://github.com/utensils/essex) - Boilerplate for Docker Based Projects: Essex is a CLI utility written in bash to quickly setup clean and consistent Docker projects with Makefile driven workflows. <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jamesbrink](https://github.com/jamesbrink)
- <b><code>&nbsp;&nbsp;&nbsp;352⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;73🍴</code></b> [HPC Container Maker](https://github.com/NVIDIA/hpc-container-maker) - Generates Dockerfiles from a high level Python recipe, including building blocks for High-Performance Computing components by [@NVIDIA][nvidia]
- <b><code>&nbsp;&nbsp;3640⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;220🍴</code></b> [img](https://github.com/genuinetools/img) - Standalone, daemon-less, unprivileged Dockerfile and OCI compatible container image builder by [@genuinetools][genuinetools]
- <b><code>&nbsp;10482⭐</code></b> <b><code>&nbsp;&nbsp;1069🍴</code></b> [kaniko](https://github.com/GoogleContainerTools/kaniko) - Build Container Images In Kubernetes. By [@GoogleContainerTools][googlecontainertools]
- <b><code>&nbsp;&nbsp;2409⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;156🍴</code></b> [makisu](https://github.com/uber/makisu) - Uber's fast and flexible unprivileged image builder for Mesos and Kubernetes, with distributed cache support. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@uber](https://github.com/uber)
- 🌎 [packer](https://www.packer.io/plugins/builders/docker) - Hashicorp tool to build machine images including docker image integrated with configuration management tools like chef, puppet, ansible
- <b><code>&nbsp;&nbsp;&nbsp;130⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [portainer](https://github.com/duedil-ltd/portainer) - Apache Mesos framework for building Docker images by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@duedil-ltd](https://github.com/duedil-ltd)
- 🌎 [Production-Ready Python Containers :heavy_dollar_sign:](https://pythonspeed.com/products/pythoncontainer/) - A template for creating production-ready Docker images for Python applications.
- <b><code>&nbsp;&nbsp;&nbsp;519⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;28🍴</code></b> [RAUDI](https://github.com/cybersecsi/RAUDI) - A tool to automatically update (and optionally push to Docker Hub) Docker Images for 3rd party software whenever theres is a new release/update/commit. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@SecSI](https://github.com/cybersecsi) 
- <b><code>&nbsp;&nbsp;&nbsp;766⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;87🍴</code></b> [runlike](https://github.com/lavie/runlike) - Generate `docker run`command and options from running containers by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@lavie](https://github.com/lavie)
- <b><code>&nbsp;&nbsp;&nbsp;181⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [SkinnyWhale](https://github.com/djosephsen/skinnywhale) :skull: - Skinnywhale helps you make smaller (as in megabytes) Docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;600⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;40🍴</code></b> [Smith](https://github.com/oracle/smith) - A Micocontainer Builder and can perform multi-stage builds after the image is built [Oracle][oracle]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [userdef](https://github.com/theAkito/userdef) - An advanced `adduser` for your Alpine based Docker images. Made by [Akito][Akito]
- <b><code>&nbsp;&nbsp;&nbsp;812⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;65🍴</code></b> [Whaler](https://github.com/P3GLEG/Whaler) - Program to reverse Docker images into Dockerfiles by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@P3GLEG](https://github.com/P3GLEG/).
- <b><code>&nbsp;&nbsp;&nbsp;338⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [Whales](https://github.com/Gueils/whales) - A tool to automatically dockerize your applications by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@icalialabs](https://github.com/IcaliaLabs).

### Dockerfile

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;64⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [chaperone-docker](https://github.com/garywiz/chaperone-docker) - A set of images using the Chaperone process manager, including a lean Alpine image, LAMP, LEMP, and bare-bones base kits.
- <b><code>&nbsp;&nbsp;&nbsp;130⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [Dockerfile Generator](https://github.com/ozankasikci/dockerfile-generator) `dfg` is both a Go library and an executable that produces valid Dockerfiles using various input channels.
- 🌎 [Dockerfile Project](https://dockerfile.github.io/) - Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [dockerfilegraph](https://github.com/patrickhoefler/dockerfilegraph) - Visualize your multi-stage Dockerfiles. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@PatrickHoefler](https://github.com/patrickhoefler)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;54⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Dockershelf](https://github.com/Dockershelf/dockershelf) - A repository that serves as a collector for docker recipes that are universal, efficient and slim. Images are updated, tested and published daily via a Travis cron job. Maintained by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@CollageLabs](https://github.com/CollageLabs).
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [dockmoor](https://github.com/MeneDev/dockmoor) :construction: - Manage docker image references and help to create reproducible builds with Docker. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@MeneDev](https://github.com/MeneDev)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Vektorcloud](https://github.com/vektorcloud) - A collection of minimal, Alpine-based Docker images

Examples by:

- <b><code>&nbsp;&nbsp;&nbsp;239⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;221🍴</code></b> [@arun-gupta](https://github.com/arun-gupta/docker-images)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;56⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [@awesome-startup](https://github.com/awesome-startup/docker-compose)
- <b><code>&nbsp;&nbsp;&nbsp;298⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;64🍴</code></b> [@crosbymichael](https://github.com/crosbymichael/Dockerfiles)
- <b><code>&nbsp;12575⭐</code></b> <b><code>&nbsp;&nbsp;2282🍴</code></b> [@jessfraz](https://github.com/jessfraz/dockerfiles)
- <b><code>&nbsp;&nbsp;&nbsp;457⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;266🍴</code></b> [@komljen](https://github.com/komljen/dockerfile-examples)
- <b><code>&nbsp;&nbsp;&nbsp;802⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;695🍴</code></b> [@kstaken](https://github.com/kstaken/dockerfile-examples)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [@ondrejmo](https://github.com/ondrejmo/Dockerfiles)
- 🌎 [@pandrew](https://gitlab.com/pandrew/dockerfiles)
- <b><code>&nbsp;&nbsp;2794⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;734🍴</code></b> [@vimagick](https://github.com/vimagick/dockerfiles)

### Linter

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;94⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [docker-image-size-limit](https://github.com/wemake-services/docker-image-size-limit) - A tool to keep an eye on your docker images size.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [Dockerfile Linter action](https://github.com/buddy-works/dockerfile-linter) - The linter lets you verify Dockerfile syntax to make sure it follows the best practices for building efficient Docker images.
- <b><code>&nbsp;&nbsp;&nbsp;393⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [dockfmt](https://github.com/jessfraz/dockfmt) :construction: - Dockerfile formatter and parser by [@jessfraz][jessfraz]
- <b><code>&nbsp;&nbsp;&nbsp;873⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;85🍴</code></b> [FROM:latest](https://github.com/replicatedhq/dockerfilelint) - An opinionated Dockerfile linter by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@replicatedhq](https://github.com/replicatedhq)
- <b><code>&nbsp;&nbsp;6988⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;312🍴</code></b> [Hadolint](https://github.com/hadolint/hadolint) - A Dockerfile linter that checks for best practices, common mistakes, and is also able to lint any bash written in `RUN` instructions; by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@lukasmartinelli](https://github.com/lukasmartinelli)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [Whale-linter](https://github.com/jeromepin/whale-linter) - A simple and small Dockerfile linter written in Python3+ without dependencies by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@jeromepin](https://github.com/jeromepin)

### Metadata

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [opencontainer](https://github.com/opencontainers/image-spec/blob/master/annotations.md) - A convention and shared namespace for Docker labels defined by OCI Image Spec.

### Registry

Services to securely store your Docker images.

- 🌎 [Amazon Elastic Container Registry :heavy_dollar_sign:](https://aws.amazon.com/ecr/) - Amazon Elastic Container Registry (ECR) is a fully-managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images.
- 🌎 [Azure Container Registry :heavy_dollar_sign:](https://azure.microsoft.com/en-us/services/container-registry/) - Manage a Docker private registry as a first-class Azure resource
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [CargoOS](https://github.com/RedCoolBeans/cargos-buildroot) - A bare essential OS for running the Docker Engine on bare metal or Cloud. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@RedCoolBeans](https://github.com/RedCoolBeans)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;48⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [cleanreg](https://github.com/hcguersoy/cleanreg) - A small tool to delete image manifests from a Docker Registry implementing the API v2, dereferencing them for the GC by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@hcguersoy](https://github.com/hcguersoy)
- 🌎 [Cloudsmith :heavy_dollar_sign:](https://cloudsmith.com/docker-registry/) - A fully managed package management SaaS, with first-class support for public and private Docker registries (and many others, incl. Helm charts for the Kubernetes ecosystem). Has a generous free-tier and is also completely free for open-source.
- 🌎 [Container Registry Service :heavy_dollar_sign:](https://container-registry.com/) - Harbor based Container Management Solution as a Service for teams and organizations. Free tier offers 1 GB storage for private repositories.
- 🌎 [Cycle.io :heavy_dollar_sign:](https://cycle.io/) - Bare-metal container hosting.
- 🌎 [DigitalOcean :heavy_dollar_sign:](https://www.digitalocean.com/products/container-registry) - DigitalOcean Container Registry.
- 🌎 [Docker Hub](https://hub.docker.com/) provided by Docker Inc.
- [Docker Registry v2][distribution] - The Docker toolset to pack, ship, store, and deliver content
- <b><code>&nbsp;&nbsp;&nbsp;693⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35🍴</code></b> [Docket](https://github.com/netvarun/docket) - Custom docker registry that allows for lightning fast deploys through bittorrent by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@netvarun](https://github.com/netvarun/)
- 🌎 [GCE Container Registry :heavy_dollar_sign:](https://cloud.google.com/container-registry/) Fast, private Docker image storage on Google Cloud Platform
- 🌎 [GitLab Container Registry](https://docs.gitlab.com/ee/user/packages/container_registry/) - Registry focused on using its images in GitLab CI
- <b><code>&nbsp;17565⭐</code></b> <b><code>&nbsp;&nbsp;4087🍴</code></b> [Harbor](https://github.com/goharbor/harbor) An open source trusted cloud native registry project that stores, signs, and scans content. Supports replication, user management, access control and activity auditing. By 🌎 [CNCF](Sandbox) formerly [VMWare]
- 🌎 [JFrog Artifactory :heavy_dollar_sign:](https://jfrog.com/artifactory/) - Artifact Repository Manager, can be used as private Docker Registry as well
- <b><code>&nbsp;&nbsp;5055⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;346🍴</code></b> [Kraken](https://github.com/uber/kraken) - Uber's Highly scalable P2P docker registry, capable of distributing TBs of data in seconds.
- <b><code>&nbsp;&nbsp;2935⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;485🍴</code></b> [Portus](https://github.com/SUSE/Portus) - Authorization service and frontend for Docker registry (v2) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@SUSE](https://github.com/SUSE)
- 🌎 [Quay.io :heavy_dollar_sign:](https://quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [Registryo](https://github.com/inmagik/registryo) - UI and token based authentication server for onpremise docker registry
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Rescoyl](https://github.com/noteed/rescoyl) - Private Docker registry (free and open source) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@noteed](https://github.com/noteed)
- 🌎 [Sonatype Nexus Repository](https://www.sonatype.com/products/nexus-repository) - Manage binaries and build artifacts across your software supply chain.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [TreeScale](https://github.com/treescale) - Build and Distribute container based applications. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tigranbs](https://github.com/tigranbs)

## Development with Docker

### API Client

- <b><code>&nbsp;&nbsp;&nbsp;137⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [ahab](https://github.com/instacart/ahab) - Docker event handling with Python by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@instacart](https://github.com/instacart)
- <b><code>&nbsp;&nbsp;&nbsp;176⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [clj-docker-client](https://github.com/into-docker/clj-docker-client) :skull: - Idiomatic Clojure client for the Docker remote API. By [@lispyclouds][lispyclouds]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;78⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [contajners](https://github.com/lispyclouds/contajners) - An idiomatic, data-driven, REPL friendly Clojure client for OCI container engines. By [@lispyclouds][lispyclouds]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;93⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27🍴</code></b> [Docker Client for JVM](https://github.com/gesellix/docker-client) - A Docker remote api client library for the JVM, written in Groovy by [@gesellix][gesellix]
- 🌎 [Docker Client TypeScript](https://gitlab.com/masaeedu/docker-client) - Docker API client for JavaScript, automatically generated from Swagger API definition from moby repository. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@masaeedu](https://github.com/masaeedu)
- <b><code>&nbsp;&nbsp;1425⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;551🍴</code></b> [docker-client](https://github.com/spotify/docker-client) :skull: - Java client for the Docker remote API. By [@spotify][spotify]
- <b><code>&nbsp;&nbsp;&nbsp;425⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;90🍴</code></b> [docker-it-scala](https://github.com/whisklabs/docker-it-scala) - Docker integration testing kit with Scala by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@whisklabs](https://github.com/whisklabs)
- <b><code>&nbsp;&nbsp;&nbsp;243⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;51🍴</code></b> [docker-java-api](https://github.com/amihaiemil/docker-java-api) - Lightweight, truly object-oriented, Java client for Docker's API. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@amihaiemil](https://github.com/amihaiemil)
- <b><code>&nbsp;&nbsp;1611⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;572🍴</code></b> [docker-maven-plugin](https://github.com/fabric8io/docker-maven-plugin) - A Maven plugin for running and creating Docker images by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@fabric8io](https://github.com/fabric8io)
- <b><code>&nbsp;&nbsp;&nbsp;279⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;96🍴</code></b> [Docker-PowerShell](https://github.com/Microsoft/Docker-PowerShell) - PowerShell Module for Docker
- <b><code>&nbsp;&nbsp;1709⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;333🍴</code></b> [Docker.DotNet](https://github.com/Microsoft/Docker.DotNet) - C#/.NET HTTP client for the Docker remote API by 🌎 [@ahmetalpbalkan](ahmetalpbalkan)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [Docker.Registry.DotNet](https://github.com/ChangemakerStudios/Docker.Registry.DotNet) - .NET (C#) Client Library for interacting with a Docker Registry API (v2) <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@rquackenbush](https://github.com/rquackenbush)
- <b><code>&nbsp;&nbsp;2636⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;452🍴</code></b> [dockerfile-maven](https://github.com/spotify/dockerfile-maven) - A Maven plugin for building and pushing Docker images by [@spotify][spotify]
- <b><code>&nbsp;&nbsp;3549⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;434🍴</code></b> [dockerode](https://github.com/apocas/dockerode) - Docker Remote API node.js module by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@apocas](https://github.com/apocas)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;70⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [DoMonit](https://github.com/eon01/DoMonit) - A simple Docker Monitoring wrapper For Docker API
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [go-dockerclient](https://github.com/fsouza/go-dockerclient/) - Go HTTP client for the Docker remote API by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@fsouza](https://github.com/fsouza/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;82⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin) - A Docker remote api plugin for Gradle by [@gesellix][gesellix]
- <b><code>&nbsp;&nbsp;&nbsp;591⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;198🍴</code></b> [libcompose](https://github.com/docker/libcompose) - Go library for Docker Compose.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;67⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [Portainer stack utils](https://github.com/greenled/portainer-stack-utils) :construction: - Bash script to deploy/update/undeploy Docker stacks in a Portainer instance from a docker-compose yaml file. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@greenled](https://github.com/greenled).
- <b><code>&nbsp;&nbsp;&nbsp;169⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;30🍴</code></b> [sbt-docker-compose](https://github.com/Tapad/sbt-docker-compose) - Integrates Docker Compose functionality into sbt by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@kurtkopchik](https://github.com/kurtkopchik/)
- <b><code>&nbsp;&nbsp;&nbsp;717⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;105🍴</code></b> [sbt-docker](https://github.com/marcuslonnberg/sbt-docker) - Create Docker images directly from sbt by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@marcuslonnberg](https://github.com/marcuslonnberg)

### CI/CD

- 🌎 [Buddy :heavy_dollar_sign:](https://buddy.works) - The best of Git, build & deployment tools combined into one powerful tool that supercharged our development.
- <b><code>&nbsp;&nbsp;&nbsp;749⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;52🍴</code></b> [Captain](https://github.com/harbur/captain) - Convert your Git workflow to Docker containers ready for Continuous Delivery by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@harbur](https://github.com/harbur).
- <b><code>&nbsp;&nbsp;1027⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;164🍴</code></b> [Cyclone](https://github.com/caicloud/cyclone) - Powerful workflow engine and end-to-end pipeline solutions implemented with native Kubernetes resources by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@caicloud](https://github.com/caicloud).
- <b><code>&nbsp;&nbsp;1353⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;65🍴</code></b> [Diun](https://github.com/crazy-max/diun) - Receive notifications when an image or repository is updated on a Docker registry by 🌎 [@crazy-max](crazy-max).
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;75⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [Dockupdater](https://github.com/dockupdater/dockupdater) - Automatically keep your docker services and your docker containers up-to-date
- <b><code>&nbsp;25135⭐</code></b> <b><code>&nbsp;&nbsp;2467🍴</code></b> [Drone](https://github.com/drone/drone) - Continuous integration server built on Docker and configured using YAML files.
- 🌎 [GitLab Runner](https://gitlab.com/gitlab-org/gitlab-runner) - GitLab has integrated CI to test, build and deploy your code with the use of GitLab runners.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;95⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;74🍴</code></b> [GOCD-Docker](https://github.com/gocd/gocd-docker) :skull: - Go Server and Agent in docker containers to provision.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;74⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [Kraken CI](https://github.com/Kraken-CI/kraken) - Modern CI/CD, open-source, on-premise system that is highly scalable and focused on testing. One of its executors is Docker. Developed by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Kraken-CI](https://github.com/Kraken-CI).
- <b><code>&nbsp;&nbsp;&nbsp;142⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;39🍴</code></b> [Microservices Continuous Deployment](https://github.com/francescou/docker-continuous-deployment) - Continuous deployment of a microservices application.
- <b><code>&nbsp;&nbsp;&nbsp;965⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;137🍴</code></b> [mu](https://github.com/stelligent/mu) - Tool to configure CI/CD of your container applications via AWS CodePipeline, CodeBuild and ECS <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Stelligent](https://github.com/stelligent)
- <b><code>&nbsp;&nbsp;1509⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;169🍴</code></b> [Ouroboros](https://github.com/pyouroboros/ouroboros) :skull: - Automatically update running Docker containers with notifications
- <b><code>&nbsp;&nbsp;&nbsp;258⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62🍴</code></b> [Popper](https://github.com/systemslab/popper) - Github actions workflow (HCL syntax) execution engine.
- 🌎 [Screwdriver :heavy_dollar_sign:](https://screwdriver.cd/) - Yahoo's OpenSource buildplatform designed for Continous Delivery.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [Skipper](https://github.com/Stratoscale/skipper) - Easily dockerize your Git repository by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Stratoscale](https://github.com/Stratoscale)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;48⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [SwarmCI](https://github.com/ghostsquad/swarmci) - Create a distributed, isolated task pipeline in your Docker Swarm.
- <b><code>&nbsp;10963⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;623🍴</code></b> [Watchtower](https://github.com/containrrr/watchtower) - Automatically update running Docker containers

### Development Environment

- <b><code>&nbsp;&nbsp;&nbsp;634⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;47🍴</code></b> [batect](https://github.com/batect/batect) - build and testing environments as code tool: Dockerised build and testing environments made easy by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@charleskorn](https://github.com/charleskorn)
- <b><code>&nbsp;&nbsp;&nbsp;672⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;29🍴</code></b> [Binci](https://github.com/binci/binci) - Containerize your development workflow. (formerly DevLab by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@TechnologyAdvice](https://github.com/TechnologyAdvice))
- <b><code>&nbsp;&nbsp;8399⭐</code></b> <b><code>&nbsp;&nbsp;1325🍴</code></b> [Boot2Docker](https://github.com/boot2docker/boot2docker) :skull: - Docker for OSX and Windows
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [construi](https://github.com/lstephen/construi) - Run your builds inside a Docker defined environment by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@lstephen](https://github.com/lstephen)
- <b><code>&nbsp;&nbsp;&nbsp;271⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24🍴</code></b> [Crashcart](https://github.com/oracle/crashcart) - Sideload Linux binaries into a running container for troubleshooting by [@Oracle][oracle]
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [dde](https://github.com/whatwedo/dde) :construction: - Local development environment toolset based on Docker. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@whatwedo](https://github.com/whatwedo)
- <b><code>&nbsp;&nbsp;&nbsp;198⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [Devstep](https://github.com/fgrehm/devstep) :skull: - Development environments powered by Docker and buildpacks by [@fgrehm][fgrehm]
- <b><code>&nbsp;&nbsp;2123⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;114🍴</code></b> [Dinghy](https://github.com/codekitchen/dinghy) - An alternative way to use Docker on Mac OS X using Docker Machine with virtualbox, vmware, xhyve or parallels
- <b><code>&nbsp;&nbsp;1003⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;29🍴</code></b> [DIP](https://github.com/bibendi/dip) - CLI utility for straightforward provisioning and interacting with an application configured by docker-compose. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@bibendi](https://github.com/bibendi)
- <b><code>&nbsp;&nbsp;2359⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62🍴</code></b> [DLite](https://github.com/nlf/dlite) :skull: - Simplest way to use Docker on OSX, no VM needed. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@nlf](https://github.com/nlf)
- <b><code>&nbsp;&nbsp;&nbsp;296⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36🍴</code></b> [dobi](https://github.com/dnephin/dobi) - A build automation tool for Docker applications. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@dnephin](https://github.com/dnephin)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [Docker Missing Tools](https://github.com/nandoquintana/docker-missing-tools) - A set of bash commands to shortcut typical docker dev-ops. An alternative to creating typical helper scripts like "build.sh" and "deploy.sh" inside code repositories. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@NandoQuintana](https://github.com/nandoquintana).
- <b><code>&nbsp;&nbsp;1435⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;110🍴</code></b> [Docker osx dev](https://github.com/brikis98/docker-osx-dev) :skull: - A productive development environment with Docker on OS X by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@brikis98](https://github.com/brikis98)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;29⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [Docker-Arch](https://github.com/Ph3nol/Docker-Arch) - Generate Web/CLI projects Dockerized development environments, from 1 simple YAML file. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Ph3nol](https://github.com/ph3nol)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker-setup](https://github.com/nicholasdille/docker-setup) - Install, configure and update container tools. By 🌎 [@nicholasdille](https://twitter.com/nicholasdille)
- <b><code>&nbsp;&nbsp;3364⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;215🍴</code></b> [Docker-sync](https://github.com/EugenMayer/docker-sync) - Drastically improves performance  <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [50-70x](https://github.com/EugenMayer/docker-sync/wiki/4.-Performance)) when using Docker for development on Mac OS X/Windows and Linux while sharing code to the container. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@EugenMayer](https://github.com/EugenMayer)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [docker-vm](https://github.com/shyiko/docker-vm) - Simple and transparent alternative to boot2docker (backed by Vagrant) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@shyiko](https://github.com/shyiko)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [DockerBuildManagement](https://github.com/DIPSAS/DockerBuildManagement) - Build Management is a python application, installed with pip. The application makes it easy to manage a build system based on Docker by configuring a single yaml file describing how to build, test, run or publish a containerized solution.
- <b><code>&nbsp;&nbsp;&nbsp;377⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [Dusty](https://github.com/gamechanger/dusty) - Managed Docker development environments on OS X
- <b><code>&nbsp;&nbsp;6678⭐</code></b> <b><code>&nbsp;&nbsp;1208🍴</code></b> [Eclipse Che](https://github.com/eclipse/che) - Developer workspace server with Docker runtimes, cloud IDE, next-generation Eclipse IDE
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;93⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [EnvCLI](https://github.com/EnvCLI/EnvCLI) - Replace your local installation of Node, Go, ... with project-specific docker containers. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@EnvCLI](https://github.com/EnvCLI)
- <b><code>&nbsp;&nbsp;1487⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;115🍴</code></b> [footloose](https://github.com/weaveworks/footloose) - Spin containers that look like Virtual Machines - By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@dlespiau](https://github.com/dlespiau)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;75⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [forward2docker](https://github.com/bsideup/forward2docker) :skull: - Utility to auto forward a port from localhost into ports on Docker containers running in a boot2docker VM by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@bsideup](https://github.com/bsideup)
- <b><code>&nbsp;&nbsp;&nbsp;572⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Gebug](https://github.com/moshebe/gebug) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.
- <b><code>&nbsp;&nbsp;3370⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;539🍴</code></b> [Lando](https://github.com/lando/lando) - Lando is for developers who want to quickly specify and painlessly spin up the services and tools needed to develop their projects. By 🌎 [Tandem](https://thinktandem.io/)
- <b><code>&nbsp;&nbsp;1783⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;98🍴</code></b> [Vagga](https://github.com/tailhook/vagga) - Vagga is a containerisation tool without daemons. It is a fully-userspace container engine inspired by Vagrant and Docker, specialized for development environments by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@tailhook](https://github.com/tailhook/)
- <b><code>&nbsp;&nbsp;&nbsp;443⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;71🍴</code></b> [Zsh-in-Docker](https://github.com/deluan/zsh-in-docker) - Install Zsh, Oh-My-Zsh and plugins inside a Docker container with one line! By 🌎 [Deluan](https://www.deluan.com)

### Garbage Collection

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [caduc](https://github.com/tjamet/caduc) - A docker garbage collector cleaning stuff you did not use recently
- <b><code>&nbsp;&nbsp;1264⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88🍴</code></b> [Docker Clean](https://github.com/ZZROTDesign/docker-clean) - A script that cleans Docker containers, images and volumes by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@zzrotdesign](https://github.com/ZZROTDesign)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;55⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [docker_gc](https://github.com/pdacity/docker_gc) - Image for automatic removing unused Docker Swarm objects. Also works just as Docker Service by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@pdacity](https://github.com/pdacity)
- <b><code>&nbsp;&nbsp;&nbsp;584⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;117🍴</code></b> [Docker-cleanup](https://github.com/meltwater/docker-cleanup) :skull: - Automatic Docker image, container and volume cleanup by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@meltwater](https://github.com/meltwater)
- <b><code>&nbsp;&nbsp;&nbsp;348⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;53🍴</code></b> [docker-custodian](https://github.com/Yelp/docker-custodian) - Keep docker hosts tidy. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Yelp](https://github.com/Yelp)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker-garby](https://github.com/konstruktoid/docker-garby) - Docker garbage collection script by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@konstruktoid](https://github.com/konstruktoid).
- <b><code>&nbsp;&nbsp;5091⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;459🍴</code></b> [docker-gc](https://github.com/spotify/docker-gc) :skull: - A cron job that will delete old stopped containers and unused images by [@spotify][spotify]
- <b><code>&nbsp;&nbsp;&nbsp;111⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [sherdock](https://github.com/rancher/sherdock) :skull: - Automatic GC of images based on regexp by [@rancher][rancher]

### Serverless

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;77⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;29🍴</code></b> [AMP](https://github.com/appcelerator-archive/amp) :skull: - The open source unified CaaS/FaaS platform for Docker, batteries included. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Appcelerator](https://github.com/appcelerator-archive)
- <b><code>&nbsp;&nbsp;5700⭐</code></b> <b><code>&nbsp;&nbsp;1095🍴</code></b> [Apache OpenWhisk](https://github.com/apache/openwhisk) - a serverless, open source cloud platform that executes functions in response to events at any scale. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@apache](https://github.com/apache)
- <b><code>&nbsp;&nbsp;5732⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;441🍴</code></b> [Docker-Lambda](https://github.com/lambci/docker-lambda) - Docker images and test runners that replicate the live AWS Lambda environment. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@lamb-ci](https://github.com/lambci)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [Funker](https://github.com/bfirsh/funker-example-voting-app) - Functions as Docker containers example voting app. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@bfirsh](https://github.com/bfirsh)
- <b><code>&nbsp;&nbsp;3011⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;231🍴</code></b> [IronFunctions](https://github.com/iron-io/functions) - The serverless microservices platform FaaS (Functions as a Service) which uses Docker containers to run Any language or AWS Lambda functions
- 🌎 [Koyeb](https://www.koyeb.com/) :heavy_dollar_sign: - Koyeb is a developer-friendly serverless platform to deploy apps globally. Seamlessly run Docker containers, web apps, and APIs with git-based deployment, native autoscaling, a global edge network, and built-in service mesh and discovery.
- <b><code>&nbsp;21667⭐</code></b> <b><code>&nbsp;&nbsp;1767🍴</code></b> [OpenFaaS](https://github.com/openfaas/faas) - A complete serverless functions framework for Docker and Kubernetes. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [OpenFaaS](https://github.com/openfaas)
- <b><code>&nbsp;&nbsp;&nbsp;575⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;44🍴</code></b> [SCAR](https://github.com/grycap/scar) - Serverless Container-aware Architectures (SCAR) is a serverless framework that allows easy deployment and execution of containers (e.g. Docker) in Serverless environments (e.g. Lambda) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@grycap](https://github.com/grycap)

### Testing

- <b><code>&nbsp;&nbsp;1846⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;158🍴</code></b> [Container Structure Test](https://github.com/GoogleContainerTools/container-structure-test) - A framework to validate the structure of an image by checking the outputs of commands or the contents of the filesystem. By [@GoogleContainerTools][googlecontainertools]
- <b><code>&nbsp;&nbsp;4791⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;428🍴</code></b> [dgoss](https://github.com/aelsabbahy/goss/tree/master/extras/dgoss) - A fast YAML based tool for validating docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;182⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [DockerSpec](https://github.com/zuazo/dockerspec) - A small Ruby Gem to run RSpec and Serverspec, Infrataster and Capybara tests against Dockerfiles or Docker images easily. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@zuazo](https://github.com/zuazo)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;44⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11🍴</code></b> [Dockunit](https://github.com/dockunit/platform) :skull: - Docker based integration tests. A simple Node based utility for running Docker based unit tests. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@dockunit](https://github.com/dockunit)
- [InSpec][inspec] - InSpec is an open-source testing framework for infrastructure with a human- and machine-readable language for specifying compliance, security and policy requirements. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@chef](https://github.com/chef)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Pull Dog](https://github.com/apps/pull-dog) - A GitHub app that automatically creates Docker-based test environments for your pull requests, from your docker-compose files. Not open source.
- <b><code>&nbsp;&nbsp;2285⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;173🍴</code></b> [Pumba](https://github.com/alexei-led/pumba) - Chaos testing tool for Docker. Can be deployed on kubernetes and CoreOS cluster. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@alexei-led](https://github.com/alexei-led)

### Wrappers

- 🌎 [Ansible](https://docs.ansible.com/ansible/latest/collections/community/general/docker_container_module.html) - Manage the life cycle of Docker containers. By RedHat
- <b><code>&nbsp;&nbsp;&nbsp;907⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;63🍴</code></b> [Azk](https://github.com/azukiapp/azk) - Orchestrate development environments on your local machine by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@azukiapp](https://github.com/azukiapp)
- <b><code>&nbsp;&nbsp;&nbsp;166⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [Beluga](https://github.com/cortexmedia/Beluga) :skull: - CLI to deploy docker containers on a single server or low amount of servers. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@cortextmedia](https://github.com/cortexmedia)
- <b><code>&nbsp;&nbsp;&nbsp;320⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [dexec](https://github.com/docker-exec/dexec) - Command line interface written in Go for running code with Docker Exec images.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;53⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [dockerized](https://github.com/benzaita/dockerized-cli) - Seamlessly execute commands in a container.
- <b><code>&nbsp;&nbsp;&nbsp;379⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;39🍴</code></b> [Dray](https://github.com/CenturyLinkLabs/dray) - An engine for managing the execution of container-based workflows by [@CenturyLinkLabs][centurylinklabs]
- <b><code>&nbsp;&nbsp;&nbsp;138⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [FuGu](https://github.com/mattes/fugu) :skull: - Docker run wrapper without orchestration by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@mattes](https://github.com/mattes)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;82⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23🍴</code></b> [Hokusai](https://github.com/artsy/hokusai) - A Docker + Kubernetes CLI for application developers; used to containerize an application and to manage its lifecycle throughout development, testing, and release cycles. From <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@artsy](https://github.com/artsy)
- <b><code>&nbsp;&nbsp;2105⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;132🍴</code></b> [Shutit](https://github.com/ianmiell/shutit) - Tool for building and maintaining complex Docker deployments by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@ianmiell](https://github.com/ianmiell)
- <b><code>&nbsp;&nbsp;&nbsp;866⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;66🍴</code></b> [subuser](https://github.com/subuser-security/subuser) - Makes it easy to securely and portably run graphical desktop applications in Docker
- <b><code>&nbsp;&nbsp;&nbsp;574⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;92🍴</code></b> [T.A.D.S. boilerplate](https://github.com/Thomvaill/tads-boilerplate) - The power of Ansible and Terraform + the simplicity of Docker Swarm = Infrastructure as Code and DevOps best practices. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Thomvaill](https://github.com/Thomvaill)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;63⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25🍴</code></b> [Terraform cloud-init config](https://github.com/christippett/terraform-cloudinit-container-server) - Terraform module for deploying a single Docker image or `docker-compose.yaml` file to any Cloud™ VM
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;24⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [Turbo](https://github.com/ramitsurana/turbo) - Simple and Powerful utility for docker. By [@ramitsurana][ramitsurana]
- <b><code>&nbsp;&nbsp;&nbsp;940⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;102🍴</code></b> [udocker](https://github.com/indigo-dc/udocker) - A tool to execute simple docker containers in batch or interactive systems without root privileges by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@inidigo-dc](https://github.com/indigo-dc)
- 🌎 [Vagrant - Docker provider](https://www.vagrantup.com/docs/providers/docker/basics) - Good starting point is <b><code>&nbsp;&nbsp;&nbsp;117⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;45🍴</code></b> [vagrant-docker-example](https://github.com/bubenkoff/vagrant-docker-example) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@bubenkoff](https://github.com/bubenkoff)

## Services based on Docker (mostly :heavy_dollar_sign:)

### CI Services

- 🌎 [CircleCI](https://circleci.com/) :heavy_dollar_sign: - Push or pull Docker images from your build environment, or build and run containers right on CircleCI.
- 🌎 [CodeFresh](https://codefresh.io) :heavy_dollar_sign: - Everything you need to build, test, and share your Docker applications. Provides automated end to end testing.
- 🌎 [CodeShip](https://www.cloudbees.com/products/codeship) :heavy_dollar_sign: - Work with your established Docker workflows while automating your testing and deployment tasks with our hosted platform dedicated to speed and security.
- 🌎 [ConcourseCI](https://concourse-ci.org) :heavy_dollar_sign: - A CI SaaS platform for developers and DevOps teams pipeline oriented.
- 🌎 [Semaphore CI](https://semaphoreci.com/) :heavy_dollar_sign: — A high-performance cloud solution that makes it easy to build, test and ship your containers to production.
- 🌎 [TravisCI](https://travis-ci.org/) :heavy_dollar_sign: - A Free github projects continuous integration Saas platform for developers and Devops.

### CaaS

- 🌎 [Amazon ECS](https://aws.amazon.com/ecs/) :heavy_dollar_sign: - A management service on EC2 that supports Docker containers.
- 🌎 [Appfleet](https://appfleet.com/) :heavy_dollar_sign: - Edge platform to deploy and manage containerized services globally. The system will route the traffic to the closest location for lower latency.
- 🌎 [Azure AKS](https://azure.microsoft.com/en-us/services/kubernetes-service/) :heavy_dollar_sign: - Simplify Kubernetes management, deployment, and operations. Use a fully managed Kubernetes container orchestration service.
- 🌎 [Cloud 66](https://www.cloud66.com) :heavy_dollar_sign: - Full-stack hosted container management as a service
- 🌎 [Dockhero](https://dockhero.io/) :heavy_dollar_sign: - Dockhero is a Heroku add-on which turns a Docker image into a microservice attached to the Heroku app. Currently in beta.
- 🌎 [Giant Swarm](https://www.giantswarm.io/) :heavy_dollar_sign: - Simple microservice infrastructure. Deploy your containers in seconds.
- 🌎 [Google Container Engine](https://cloud.google.com/kubernetes-engine/docs/) :heavy_dollar_sign: - Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
- 🌎 [Mesosphere DC/OS Platform](https://d2iq.com/products/dcos) :heavy_dollar_sign: - Integrated platform for data and containers built on Apache Mesos by 🌎 [@mesosphere](https://d2iq.com)
- 🌎 [Red Hat CodeReady Workspaces](https://developers.redhat.com/products/codeready-workspaces/overview) - A collaborative Kubernetes-native solution for rapid application development that delivers consistent developer environments on Red Hat OpenShift to allow anyone with a browser to contribute code in under two minutes.
- 🌎 [Red Hat OpenShift Dedicated](https://www.redhat.com/en/technologies/cloud-computing/openshift/dedicated) :heavy_dollar_sign: - Fully-managed Red Hat® OpenShift® service on Amazon Web Services and Google Cloud
- 🌎 [Sloppy.io](https://sloppy.io/en/) :heavy_dollar_sign: - all-in-one solution for container deployment and hosting – made and hosted in Germany
- 🌎 [Triton](https://www.joyent.com/) :heavy_dollar_sign: - Elastic container-native infrastructure by Joyent.
- 🌎 [Virtuozzo Application Platform](https://www.virtuozzo.com/application-platform-partners/) :heavy_dollar_sign: - Deploy and manage your projects with turnkey PaaS across a wide network of reliable service providers

### Monitoring Services

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [AppDynamics](https://github.com/Appdynamics/docker-monitoring-extension) - Docker Monitoring extension gathers metrics from the Docker Remote API, either using Unix Socket or TCP.
- 🌎 [Broadcom Docker Monitoring](https://www.broadcom.com/info/aiops/docker-monitoring) :heavy_dollar_sign: - Agile Operations solutions from Broadcom deliver the modern Docker monitoring businesses need to accelerate and optimize the performance of microservices and the dynamic Docker environments running them. Monitor both the Docker environment and apps that run inside them. (former CA Technologies)
- 🌎 [Collecting docker logs and stats with Splunk](https://www.splunk.com/en_us/blog/tips-and-tricks/collecting-docker-logs-and-stats-with-splunk.html)
- 🌎 [Datadog](https://www.datadoghq.com/) :heavy_dollar_sign: - Datadog is a full-stack monitoring service for large-scale cloud environments that aggregates metrics/events from servers, databases, and applications. It includes support for Docker, Kubernetes, and Mesos.
- 🌎 [Prometheus](https://prometheus.io/) :heavy_dollar_sign: - Open-source service monitoring system and time series database
- 🌎 [Site24x7](https://www.site24x7.com/docker-monitoring.html) :heavy_dollar_sign: - Docker Monitoring for DevOps and IT is a SaaS Pay per Host model
- <b><code>&nbsp;&nbsp;&nbsp;200⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35🍴</code></b> [SPM for Docker](https://github.com/sematext/sematext-agent-docker) :heavy_dollar_sign: - Monitoring of host and container metrics, Docker events and logs. Automatic log parser. Anomaly Detection and alerting for metrics and logs. <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@sematext](https://github.com/sematext)
- 🌎 [Sysdig Monitor](https://sysdig.com/products/monitor/) :heavy_dollar_sign: - Sysdig Monitor can be used as either software or a SaaS service to monitor, alert, and troubleshoot containers using system calls. It has container-specific features for Docker and Kubernetes.

# Useful Resources

- * 🌎 [Valuable Docker Links](http://nane.kratzke.pages.mylab.th-luebeck.de/about/blog/2014/08/24/valuable-docker-links/)** High quality articles about docker! **MUST SEE**
- 🌎 [Become a Docker Power User with Visual Studio Code](https://www.thebyte.io/become-docker-power-user-with-vs-code) - :heavy_dollar_sign: A training course to help you become a Docker Power user with Visual Studio Code
- <b><code>&nbsp;&nbsp;8337⭐</code></b> <b><code>&nbsp;&nbsp;1518🍴</code></b> [Cloud Native Landscape](https://github.com/cncf/landscape)
- 🌎 [Docker Certification](https://intellipaat.com/docker-training-course/) :heavy_dollar_sign: will help you to will Learn Docker containerization, running Docker containers, Image creation, Dockerfile, Docker orchestration, security best practices, and more through hands-on projects and case studies and helps to clear Docker Certified Associate.
- 🌎 [Docker Community on Hashnode](https://hashnode.com/n/docker)
- 🌎 [Docker dev bookmarks](https://www.codever.land/search?q=docker) - use the tag 🌎 [docker](https://www.codever.land/bookmarks/t/docker)
- 🌎 [Docker in Action, Second Edition](https://www.manning.com/books/docker-in-action-second-edition)
- 🌎 [Docker in Practice, Second Edition](https://www.manning.com/books/docker-in-practice-second-edition)
- 🌎 [Docker Newsletter](https://www.docker.com/newsletter-subscription/) Docker newsletter for regular updates
- 🌎 [Docker packaging guide for Python](https://pythonspeed.com/docker/) - a series of detailed articles on the specifics of Docker packaging for Python.
- 🌎 [Learn Docker in a Month of Lunches](https://www.manning.com/books/learn-docker-in-a-month-of-lunches)
- 🌎 [Learn Docker](https://coursesity.com/blog/best-docker-tutorials/) - Learn Docker - curated list of the top online docker tutorials and courses.
- 🌎 [Programming Community Curated Resources for learning Docker](https://hackr.io/tutorials/learn-docker)


## Awesome Lists

- <b><code>&nbsp;&nbsp;1445⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;159🍴</code></b> [Awesome CI/CD](https://github.com/cicdops/awesome-ciandcd) - Not specific to docker but relevant.
- <b><code>&nbsp;16023⭐</code></b> <b><code>&nbsp;&nbsp;2427🍴</code></b> [Awesome Compose](https://github.com/docker/awesome-compose) - Docker Compose samples
- <b><code>&nbsp;12809⭐</code></b> <b><code>&nbsp;&nbsp;2053🍴</code></b> [Awesome Kubernetes](https://github.com/ramitsurana/awesome-kubernetes) by [@ramitsurana][ramitsurana]
- <b><code>&nbsp;&nbsp;1314⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;158🍴</code></b> [Awesome Linux Container](https://github.com/Friz-zy/awesome-linux-containers) more general about container than this repo, by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Friz-zy](https://github.com/Friz-zy).
- <b><code>&nbsp;90522⭐</code></b> <b><code>&nbsp;&nbsp;5885🍴</code></b> [Awesome Selfhosted](https://github.com/awesome-selfhosted/awesome-selfhosted) list of Free Software network services and web applications which can be hosted locally by running in a classical way (setup local web server and run applications from there) or in a Docker container. By <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@Kickball](https://github.com/Kickball)
- <b><code>&nbsp;14040⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;936🍴</code></b> [Awesome Sysadmin](https://github.com/n1trux/awesome-sysadmin) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@n1trux](https://github.com/n1trux)
- <b><code>&nbsp;15655⭐</code></b> <b><code>&nbsp;&nbsp;1254🍴</code></b> [ToolsOfTheTrade](https://github.com/cjbarber/ToolsOfTheTrade) a list of SaaS and On premise applications by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@cjbarber](https://github.com/cjbarber)

## Demos and Examples

- 🌎 [An Annotated Docker Config for Frontend Web Development](https://nystudio107.com/blog/an-annotated-docker-config-for-frontend-web-development) A local development environment with Docker allows you to shrink-wrap the devops your project needs as config, making onboarding frictionless.
- <b><code>&nbsp;&nbsp;&nbsp;231⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [Local Docker DB](https://github.com/alexmacarthur/local-docker-db) a list of docker-compose samples for a lot of databases by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@alexmacarthur](https://github.com/alexmacarthur)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;72⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [Webstack-micro](https://github.com/ferbs/webstack-micro) Demo web app showing how Docker Compose might be used to set up an API Gateway, centralized authentication, background workers, and WebSockets as containerized services.

## Good Tips

- 🌎 [Dealing with linked containers dependency in docker-compose](http://brunorocha.org/python/dealing-with-linked-containers-dependency-in-docker-compose.html) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@rochacbruno](https://github.com/rochacbruno)
- 🌎 [Docker Caveats](http://docker-saigon.github.io/post/Docker-Caveats/) What You Should Know About Running Docker In Production (written 11 APRIL 2016) **MUST SEE**
- 🌎 [Docker Containers on the Desktop](https://blog.jessfraz.com/post/docker-containers-on-the-desktop/) - The **funniest way** to learn about docker by [@jessfraz][jessfraz] who also gave a 🌎 [presentation](https://www.youtube.com/watch?v=1qlLUf7KtAw) about it @ DockerCon 2015
- 🌎 [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](https://www.flexera.com/blog/cloud/docker-vs-vms-combining-both-for-cloud-portability-nirvana/)
- <b><code>&nbsp;&nbsp;3659⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;126🍴</code></b> [Dockerfile best practices](https://github.com/hexops/dockerfile) - This repository has best-practices for writing Dockerfiles
- 🌎 [Don't Repeat Yourself with Anchors, Aliases and Extensions in Docker Compose Files](https://medium.com/@kinghuang/docker-compose-anchors-aliases-extensions-a1e4105d70bd) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@King Chung Huang](https://github.com/kinghuang)
- 🌎 [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm][fgrehm]

## Raspberry Pi & ARM

- 🌎 [Docker Pirates ARMed with explosive stuff](https://blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Get Docker up and running on the RaspberryPi in three steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-three-steps)
- 🌎 [git push docker containers to linux devices](https://www.balena.io) Modern DevOps for IoT, leveraging git and Docker.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-%28ARMv7%29-devices)

## Security

- 🌎 [Bringing new security features to Docker](https://opensource.com/business/14/9/security-for-docker)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [CVE Scanning Alpine images with Multi-stage builds in Docker 17.05](https://github.com/tomwillfixit/alpine-cvecheck) by 🌎 [@tomwillfixit](https://twitter.com/tomwillfixit)
- <b><code>&nbsp;&nbsp;&nbsp;597⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;82🍴</code></b> [Docker Secure Deployment Guidelines](https://github.com/AonCyberLabs/Docker-Secure-Deployment-Guidelines)
- 🌎 [Docker Security - Quick Reference](https://binarymist.io/publication/docker-security/)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker Security Cheat Sheet](https://github.com/konstruktoid/Docker/blob/master/Security/CheatSheet.adoc)
- 🌎 [Docker Security: Are Your Containers Tightly Secured to the Ship? SlideShare](https://fr.slideshare.net/MichaelBoelen/docker-security-are-your-containers-tightly-secured-to-the-ship)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [How CVE's are handled on Offical Docker Images](https://github.com/docker-library/official-images/issues/1448)
- 🌎 [Lynis is an open source security auditing tool including Docker auditing](https://cisofy.com/lynis/)
- 🌎 [Security Best Practices for Building Docker Images](https://linux-audit.com/tag/docker/)
- 🌎 [Software Engineering Radio interview of Docker Security Team Lead (Diogo Mónica)](https://www.se-radio.net/2017/05/se-radio-episode-290-diogo-monica-on-docker-security/)
- 🌎 [Ten Docker Image Security Best Practices Cheat Sheet](https://snyk.io/blog/10-docker-image-security-best-practices/)
- 🌎 [Top ten most popular docker images each contain at least 30 vulnerabilities](https://snyk.io/blog/top-ten-most-popular-docker-images-each-contain-at-least-30-vulnerabilities/)
- 🌎 [Tuning Docker with the newest security enhancements](https://opensource.com/business/15/3/docker-security-tuning)

## Videos

- 🌎 [Contributing to Docker by Andrew "Tianon" Page (InfoSiftr)](https://www.youtube.com/watch?v=1jwo8-1HYYg) (34:31)
- 🌎 [Deploying and scaling applications with Docker, Swarm, and a tiny bit of Python magic](https://www.youtube.com/watch?v=GpHMTR7P2Ms) (3:11:06) by [@jpetazzo][jpetazzo]
- 🌎 [Docker and SELinux by Daniel Walsh from Red Hat](https://www.youtube.com/watch?v=zWGFqMuEHdw) (40:23)
- 🌎 [Docker Course](https://www.youtube.com/watch?v=UZpyvK6UGFo) (Spanish) by <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [@pablokbs](https://github.com/pablokbs)
- 🌎 [Docker for Developers](https://www.youtube.com/watch?v=FdkNAjjO5yQ) (54:26) by [@jpetazzo][jpetazzo] <== Good introduction, context, demo
- 🌎 [Docker from scratch](https://www.youtube.com/playlist?list=PLLhEJK7fQIxD-btrjrqdEfQHbkZnQrmqE) (1:22:01) on YouTube by Paris Nakita Kejser
- 🌎 [Docker: How to Use Your Own Private Registry](https://www.youtube.com/watch?v=CAewZCBT4PI) (15:01)
- 🌎 [Docker in Production](https://www.youtube.com/watch?v=Glk5d5WP6MI) by [@jpetazzo][jpetazzo] (36:05)
- 🌎 [Docker Primer to Docker Compose](https://www.youtube.com/watch?v=G-s2GXGAjTk) (1:56:45) on YouTube by LoginRadius
- 🌎 [Docker Registry from scratch](https://www.youtube.com/playlist?list=PLLhEJK7fQIxAz3d4Fj3edq7UcxEhdTCBm) (44:40) on YouTube by Paris Nakita Kejser
- 🌎 [Docker Swarm from scratch](https://www.youtube.com/playlist?list=PLLhEJK7fQIxAY4gZd1Wl-GsLvg-e9Ap1e) (1:41:28) on YouTube by Paris Nakita Kejser
- 🌎 [Extending Docker with Plugins](https://vimeo.com/110835013) (15:21)
- 🌎 [From Local Docker Development to Production Deployments](https://www.youtube.com/watch?v=7CZFpHUPqXw) by [@jpetazzo][jpetazzo] @ AWS re:Invent 2015
- 🌎 [Immutable Infrastructure with Docker and EC2 by Michael Bryzek (Gilt)](https://www.youtube.com/watch?v=GaHzdqFithc) (42:04)
- 🌎 [Introduction to Docker and containers](https://www.youtube.com/watch?v=ZVaRK10HBjo) (3:09:00) by [@jpetazzo][jpetazzo]
- 🌎 [Logging on Docker: What You Need to Know](https://vimeo.com/123341629) (51:27)
- 🌎 [Performance Analysis of Docker - Jeremy Eder](https://www.youtube.com/watch?v=6f2E6PKYb0w) (1:36:58)
- 🌎 [Scalable Microservices with Kubernetes](https://www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) Free Udacity course
- 🌎 [State of containers: a debate with CoreOS, VMware and Google](https://www.youtube.com/watch?v=IiITP3yIRd8) (27:38)
- 🌎 [List of online courses (free and paid) about Docker](https://skillcombo.com/topic/docker/)

# Communities and Meetups

## Brazilian

- 🌎 [Docker BR on Slack](http://docker-br.herokuapp.com) - Auto invite url
- 🌎 [Docker BR on Telegram](https://telegram.me/dockerbr)

## Chinese

- 🌎 [DockerOne](http://dockone.io/) Docker Community (in Chinese) by 🌎 [@LiYingJie](http://dockone.io/people/%E6%9D%8E%E9%A2%96%E6%9D%B0)

## English

- 🌎 [Docker Community](https://www.docker.com/community/)
- 🌎 [Docker Events](https://www.docker.com/events/)
- 🌎 [Docker Online Meetup](https://www.meetup.com/en-AU/Docker-Online-Meetup/)
- 🌎 [Docker Reddit Community](https://www.reddit.com/r/docker/)

## Russian

- 🌎 [Docker Russian-speaking Community](https://t.me/docker_ru)

## Spanish

- 🌎 [Docker Tips](https://dockertips.com/)


## Stargazers over time

[![Stargazers over time](https://starchart.cc/veggiemonk/awesome-docker.svg)](https://starchart.cc/veggiemonk/awesome-docker)

## Contributor over time

[![Contributor over time](https://api7.ai/contributor-graph?chart=contributorOverTime&repo=veggiemonk/awesome-docker)


[contributing]: https://github.com/veggiemonk/awesome-docker/blob/master/.github/CONTRIBUTING.md
[Akito]: https://github.com/theAkito
[calico]: https://github.com/projectcalico/calicoctl
[centurylinklabs]: https://github.com/CenturyLinkLabs
[containx]: https://github.com/ContainX
[containers]: https://github.com/containers
[coreos]: https://github.com/coreos
[cncf]: https://www.cncf.io
[crazy-max]: https://github.com/crazy-max
[deepfence]: https://github.com/deepfence
[distribution]: https://github.com/docker/distribution
[docker-flow]: https://github.com/docker-flow
[docker-for-windows]: https://docs.docker.com/docker-for-windows/
[docker]: https://github.com/docker
[editreadme]: https://github.com/veggiemonk/awesome-docker/edit/master/README.md
[fgrehm]: https://github.com/fgrehm
[gesellix]: https://github.com/gesellix
[genuinetools]: https://github.com/genuinetools
[gliderlabs]: https://github.com/gliderlabs
[grammarly]: https://github.com/grammarly
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
[openshift]: https://www.okd.io
[oracle]: https://github.com/oracle
[progrium]: https://github.com/progrium
[ramitsurana]: https://github.com/ramitsurana
[rancher]: https://github.com/rancher
[safe-waters]: https://github.com/safe-waters
[sindresorhus]: https://github.com/sindresorhus/awesome
[spotify]: https://github.com/spotify
[tomastomecek]: https://github.com/TomasTomecek
[vegasbrianc]: https://github.com/vegasbrianc
[weave]: https://github.com/weaveworks/weave
[VMWare]: https://github.com/vmware
[@byrnedo]: https://github.com/byrnedo
[@skanehira]: https://github.com/skanehira
