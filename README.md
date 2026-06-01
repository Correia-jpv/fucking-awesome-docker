# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)][sindresorhus] [![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/veggiemonk/awesome-docker/)[![Last Commit](https://img.shields.io/github/last-commit/veggiemonk/awesome-docker)](https://github.com/correia-jpv/fucking-awesome-docker/commits/main)<!-- omit in toc -->

> A curated list of projects for Docker.

If you would like to contribute, please read [CONTRIBUTING.md](https://github.com/correia-jpv/fucking-awesome-docker/blob/master/.github/CONTRIBUTING.md) first.
If this list is not complete, you can contribute to make it so.
If you see a link here that is not (any longer) a good fit, you can fix it by submitting a [pull request][editreadme] to improve this file. Thank you!

**The project has to be for Docker, not just using Docker.**

> Rule of thumb: if removing the Docker integration wouldn't kill the project's value proposition, it doesn't belong in the this list.

The creators and maintainers of this list do not receive any form of payment to accept a change made by any contributor.
This page is not an official Docker product in any way.
It is a list of links to projects and is maintained by volunteers.
Everybody is welcome to contribute.
The goal of this repo is to index open-source projects, not to advertise for profit.

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

_Source:_ 🌎 [What is Docker](www.docker.com/why-docker/)

# Contents <!-- omit in toc -->

<!-- TOC -->

- [Projects](#projects)
    - [Engine \& Runtime](#engine--runtime)
    - [Building Images](#building-images)
        - [Builder](#builder)
        - [Base Images](#base-images)
        - [Dockerfile](#dockerfile)
        - [Linter](#linter)
    - [Image Lifecycle](#image-lifecycle)
        - [Registry](#registry)
        - [Registry CLI](#registry-cli)
        - [Image Scanning \& SBOM](#image-scanning--sbom)
        - [Supply Chain](#supply-chain)
    - [Running Containers](#running-containers)
        - [Composition](#composition)
        - [Orchestration](#orchestration)
        - [Deployment \& Platforms](#deployment--platforms)
        - [Garbage Collection](#garbage-collection)
    - [Networking \& Proxies](#networking--proxies)
        - [Networking](#networking)
        - [Reverse Proxy](#reverse-proxy)
    - [Storage \& Data](#storage--data)
    - [Observability](#observability)
    - [Security](#security)
    - [User Interfaces](#user-interfaces)
        - [Desktop](#desktop)
        - [Terminal](#terminal)
        - [Web](#web)
        - [IDE Integrations](#ide-integrations)
    - [Developer Workflow](#developer-workflow)
        - [API Client](#api-client)
        - [CI/CD](#cicd)
        - [Development Environment](#development-environment)
        - [Serverless](#serverless)
        - [Testing](#testing)
        - [Wrappers](#wrappers)
    - [In-Container Tooling](#in-container-tooling)
- [Learning Resources](#learning-resources)
    - [Where to Start](#where-to-start)
    - [Where to Start (Windows)](#where-to-start-windows)
    - [Books \& Tutorials](#books--tutorials)
    - [Awesome Lists](#awesome-lists)
    - [Demos and Examples](#demos-and-examples)
    - [Good Tips](#good-tips)
    - [Raspberry Pi \& ARM](#raspberry-pi--arm)
    - [Security Articles](#security-articles)
    - [Videos](#videos)
    - [Communities and Meetups](#communities-and-meetups)
        - [Brazilian](#brazilian)
        - [English](#english)
        - [Russian](#russian)
        - [Spanish](#spanish)
- [Stargazers over time](#stargazers-over-time)

<!-- /TOC -->

# Projects

## Official Projects

- <b><code>&nbsp;71612⭐</code></b> <b><code>&nbsp;18961🍴</code></b> [Moby](https://github.com/moby/moby))
- 🌎 [Docker Hub](hub.docker.com)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker Compose](https://github.com/docker/compose/)) - Define and run multi-container applications with Docker.
- [Docker Registry][distribution] - The Docker toolset to pack, ship, store, and deliver content

## Engine & Runtime

- <b><code>&nbsp;29066⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;577🍴</code></b> [colima](https://github.com/abiosoft/colima)) - Container runtimes on macOS (and Linux) with minimal setup.
- <b><code>&nbsp;20791⭐</code></b> <b><code>&nbsp;&nbsp;3951🍴</code></b> [containerd](https://github.com/containerd/containerd)) - An open and reliable container runtime.
- <b><code>&nbsp;&nbsp;5617⭐</code></b> <b><code>&nbsp;&nbsp;1171🍴</code></b> [cri-o](https://github.com/cri-o/cri-o)) - Open Container Initiative-based implementation of Kubernetes Container Runtime Interface.
- <b><code>&nbsp;18428⭐</code></b> <b><code>&nbsp;&nbsp;1623🍴</code></b> [gVisor](https://github.com/google/gvisor)) - Application Kernel for Containers.
- <b><code>&nbsp;&nbsp;5191⭐</code></b> <b><code>&nbsp;&nbsp;1171🍴</code></b> [lxc](https://github.com/lxc/lxc)) - LXC - Linux Containers.
- <b><code>&nbsp;&nbsp;&nbsp;191⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Mocker](https://github.com/us/mocker)) - Docker-compatible container CLI for macOS, built on Apple's Containerization framework.
- <b><code>&nbsp;31852⭐</code></b> <b><code>&nbsp;&nbsp;3135🍴</code></b> [podman](https://github.com/containers/libpod)) - Libpod is a library used to create container pods. Home of Podman.
- <b><code>&nbsp;13250⭐</code></b> <b><code>&nbsp;&nbsp;2285🍴</code></b> [runc](https://github.com/opencontainers/runc)) - CLI tool for spawning and running containers according to the OCI specification.
- <b><code>&nbsp;&nbsp;&nbsp;483⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;167🍴</code></b> [runtime-tools](https://github.com/opencontainers/runtime-tools)) - Oci-runtime-tool is a collection of tools for working with the OCI runtime specification.
- <b><code>&nbsp;&nbsp;7427⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;438🍴</code></b> [youki](https://github.com/youki-dev/youki)) - Container runtime written in Rust, implementing the OCI runtime specification.

## Building Images

### Builder

Applications designed to help or simplify building **new** images

- <b><code>&nbsp;&nbsp;&nbsp;696⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;73🍴</code></b> [ansible-bender](https://github.com/ansible-community/ansible-bender)) - A tool utilising `ansible` and `buildah`.
- <b><code>&nbsp;&nbsp;1623⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;215🍴</code></b> [apko](https://github.com/chainguard-dev/apko)) - Declarative OCI image builder from apk packages; reproducible by design.
- <b><code>&nbsp;&nbsp;8812⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;899🍴</code></b> [buildah](https://github.com/containers/buildah)) - A tool that facilitates building OCI images.
- <b><code>&nbsp;10002⭐</code></b> <b><code>&nbsp;&nbsp;1430🍴</code></b> [BuildKit](https://github.com/moby/buildkit)) - Concurrent, cache-efficient, and Dockerfile-agnostic builder toolkit.
- <b><code>&nbsp;&nbsp;4396⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;642🍴</code></b> [buildx](https://github.com/docker/buildx)) - Official Docker CLI plugin for multi-platform builds backed by BuildKit.
- <b><code>&nbsp;&nbsp;&nbsp;112⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;44🍴</code></b> [cekit](https://github.com/cekit/cekit)) - A tool used by openshift to build base images using different build engines.
- <b><code>&nbsp;&nbsp;&nbsp;444⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [dlayer](https://github.com/orisano/dlayer)) - Docker layer analyzer.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;47⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [docker-companion](https://github.com/mudler/docker-companion)) - A command line tool written in Golang to squash and unpack docker images.
- <b><code>&nbsp;&nbsp;&nbsp;154⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker-repack](https://github.com/orf/docker-repack)) - Repacks a Docker image into a smaller, more efficient version that makes it significantly faster to pull.
- <b><code>&nbsp;23289⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;832🍴</code></b> [DockerSlim](https://github.com/docker-slim/docker-slim)) shrinks fat Docker images creating the smallest possible images.
- <b><code>&nbsp;12038⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;454🍴</code></b> [earthly](https://github.com/earthly/earthly)) - Containerized build automation with Dockerfile-meets-Makefile syntax.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;38⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [essex](https://github.com/utensils/essex)) - Boilerplate for Docker Based Projects: Essex is a CLI utility written in bash to quickly setup clean and consistent Docker projects with Makefile driven workflows.
- <b><code>&nbsp;&nbsp;&nbsp;513⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;102🍴</code></b> [HPC Container Maker](https://github.com/NVIDIA/hpc-container-maker)) - Generates Dockerfiles from a high level Python recipe, including building blocks for High-Performance Computing components.
- <b><code>&nbsp;&nbsp;3984⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;233🍴</code></b> [img](https://github.com/genuinetools/img)) - Standalone, daemon-less, unprivileged Dockerfile and OCI compatible container image builder.
- <b><code>&nbsp;&nbsp;8443⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;440🍴</code></b> [ko](https://github.com/ko-build/ko)) - Build and deploy Go applications as container images without a Dockerfile.
- <b><code>&nbsp;&nbsp;&nbsp;848⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;77🍴</code></b> [nix2container](https://github.com/nlewo/nix2container)) - Build OCI images with Nix without `docker load` round-trips.
- 🌎 [packer](developer.hashicorp.com/packer/integrations/hashicorp/docker/latest/components/builder/docker) - Hashicorp tool to build machine images including docker image integrated with configuration management tools like chef, puppet, ansible.
- 🌎 [Production-Ready Python Containers](pythonspeed.com/products/pythoncontainer/) - :yen: A template for creating production-ready Docker images for Python applications.
- <b><code>&nbsp;&nbsp;&nbsp;559⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34🍴</code></b> [RAUDI](https://github.com/cybersecsi/RAUDI)) - A tool to automatically update (and optionally push to Docker Hub) Docker Images for 3rd party software whenever theres is a new release/update/commit.
- <b><code>&nbsp;&nbsp;2929⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;167🍴</code></b> [runlike](https://github.com/lavie/runlike)) - Generate `docker run`command and options from running containers.
- <b><code>&nbsp;&nbsp;1184⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;104🍴</code></b> [Whaler](https://github.com/P3GLEG/Whaler)) - Program to reverse Docker images into Dockerfiles.


### Base Images

Minimal, hardened, or purpose-built container base images.

- <b><code>&nbsp;&nbsp;&nbsp;671⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;175🍴</code></b> [Chainguard Images](https://github.com/chainguard-images/images)) - Minimal, signed, SBOM-attested container images built on Wolfi.
- <b><code>&nbsp;22685⭐</code></b> <b><code>&nbsp;&nbsp;1397🍴</code></b> [distroless](https://github.com/GoogleContainerTools/distroless)) - Language focused docker images, minus the operating system.
- <b><code>&nbsp;&nbsp;&nbsp;599⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;172🍴</code></b> [melange](https://github.com/chainguard-dev/melange)) - Build apk packages from declarative YAML for use with apko.
- <b><code>&nbsp;&nbsp;1225⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;433🍴</code></b> [Wolfi](https://github.com/wolfi-dev/os)) - Undistro Linux designed for containers; glibc-based, signed, daily SBOMs.


### Dockerfile

- <b><code>&nbsp;&nbsp;&nbsp;187⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21🍴</code></b> [Dockerfile Generator](https://github.com/ozankasikci/dockerfile-generator)) `dfg` is both a Go library and an executable that produces valid Dockerfiles using various input channels.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;97⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [Dockershelf](https://github.com/Dockershelf/dockershelf)) - A repository that serves as a collector for docker recipes that are universal, efficient and slim. Images are updated, tested and published daily via a Travis cron job.
- 🌎 [Trsuted Builds](dockerfile.github.io/) - Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.

### Linter

- <b><code>&nbsp;&nbsp;&nbsp;203⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [Dockadvisor](https://github.com/deckrun/dockadvisor)) - Lightweight Dockerfile linter with 60+ rules, quality scoring, and security checks.
- <b><code>&nbsp;&nbsp;&nbsp;130⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [docker-image-size-limit](https://github.com/wemake-services/docker-image-size-limit)) - A tool to keep an eye on your docker images size.
- <b><code>&nbsp;12177⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;493🍴</code></b> [Hadolint](https://github.com/hadolint/hadolint)) - A Dockerfile linter that checks for best practices, common mistakes, and is also able to lint any bash written in `RUN` instructions;.

## Image Lifecycle

### Registry

Services to securely store your Docker images.

- 🌎 [Amazon Elastic Container Registry](aws.amazon.com/ecr/) - :yen: Amazon Elastic Container Registry (ECR) is a fully-managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images.
- 🌎 [Azure Container Registry](azure.microsoft.com/en-us/products/container-registry/#overview) - :yen: Manage a Docker private registry as a first-class Azure resource.
- 🌎 [Cloudsmith](cloudsmith.com/product/formats/docker-registry) - :yen: A fully managed package management SaaS, with first-class support for public and private Docker registries (and many others, incl. Helm charts for the Kubernetes ecosystem). Has a generous free-tier and is also completely free for open-source.
- 🌎 [Container Registry Service](container-registry.com/) - :yen: Harbor based Container Management Solution as a Service for teams and organizations. Free tier offers 1 GB storage for private repositories.
- 🌎 [Cycle.io](cycle.io/) - :yen: Bare-metal container hosting.
- 🌎 [DigitalOcean](www.digitalocean.com/products/container-registry) - :yen: DigitalOcean Container Registry.
- 🌎 [Docker Hub](hub.docker.com/) provided by Docker Inc.
- [Docker Registry v2][distribution] - The Docker toolset to pack, ship, store, and deliver content
- <b><code>&nbsp;&nbsp;3182⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;400🍴</code></b> [Dragonfly](https://github.com/dragonflyoss/Dragonfly2)) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology.
- 🌎 [GCP Artifact Registry](docs.cloud.google.com/artifact-registry/docs) - :yen: Fast, private Docker image storage on Google Cloud Platform.
- 🌎 [Gitea Container Registry](docs.gitea.com/usage/packages/container) - Integrated Docker registry in Gitea, ideal for private, small-scale image hosting.
- 🌎 [GitHub Container Registry](docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) - GitHub's solution for storing and managing Docker images, with tight integration into GitHub Actions.
- 🌎 [GitLab Container Registry](docs.gitlab.com/user/packages/container_registry/) - Registry focused on using its images in GitLab CI.
- <b><code>&nbsp;28600⭐</code></b> <b><code>&nbsp;&nbsp;5239🍴</code></b> [Harbor](https://github.com/goharbor/harbor)) An open source trusted cloud native registry project that stores, signs, and scans content. Supports replication, user management, access control and activity auditing.
- 🌎 [JFrog Artifactory](jfrog.com/artifactory/) - :yen: Artifact Repository Manager, can be used as private Docker Registry as well.
- <b><code>&nbsp;&nbsp;&nbsp;241⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [kontain.me](https://github.com/imjasonh/kontain.me)) - On-demand container image registry that builds and serves images when they are pulled.
- <b><code>&nbsp;&nbsp;6694⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;474🍴</code></b> [Kraken](https://github.com/uber/kraken)) - Uber's Highly scalable P2P docker registry, capable of distributing TBs of data in seconds.
- <b><code>&nbsp;&nbsp;&nbsp;183⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [NORA](https://github.com/getnora-io/nora)) - Lightweight multi-protocol artifact registry supporting Docker, Maven, npm, Cargo and PyPI in a single 32MB binary. Pull-through cache, Web UI, Prometheus metrics, RBAC auth.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [nscr](https://github.com/jhstatewide/nscr)) - A light-weight, self-contained container registry that's easy to run and maintain.
- 🌎 [Quay.io](quay.io/) - :yen: Secure hosting for private Docker repositories.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [Registryo](https://github.com/inmagik/registryo)) - UI and token based authentication server for onpremise docker registry.
- 🌎 [RepoFlow](www.repoflow.io) - A simple and easy-to-use package management platform with Docker support alongside other formats like PyPI, Maven, npm, and Helm. Includes smart search, built-in Docker image scanning, and a great free option for both self-hosted and cloud use.
- 🌎 [Sonatype Nexus Repository](www.sonatype.com/products/sonatype-nexus-repository) - Manage binaries and build artifacts across your software supply chain.

### Registry CLI

Daemonless command-line tools for inspecting, copying, and manipulating images in OCI/Docker registries.

- <b><code>&nbsp;&nbsp;3892⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;648🍴</code></b> [crane](https://github.com/google/go-containerregistry/tree/main/cmd/crane)) - Lightweight CLI to manipulate registry images, from `go-containerregistry`.
- <b><code>&nbsp;&nbsp;3892⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;648🍴</code></b> [go-containerregistry](https://github.com/google/go-containerregistry)) - Go library and CLI tools (`crane`, `gcrane`, `registry`) for working with container registries.
- <b><code>&nbsp;&nbsp;2284⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;237🍴</code></b> [oras](https://github.com/oras-project/oras)) - Push and pull arbitrary OCI artifacts to and from any OCI registry.
- <b><code>&nbsp;&nbsp;1845⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;123🍴</code></b> [regctl](https://github.com/regclient/regclient)) - Daemonless registry client; copy, inspect, modify, and sign OCI images.
- <b><code>&nbsp;10927⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;928🍴</code></b> [skopeo](https://github.com/containers/skopeo)) - Work with remote image registries: retrieve information, copy images, sign content.

### Image Scanning & SBOM

Image vulnerability scanners, SBOM generators, and digest pinning tools. Commercial entries marked `:yen:`.

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Anchor](https://github.com/SongStitch/anchor/)) - A tool to ensure reproducible builds by pinning dependencies inside your Dockerfiles.
- 🌎 [Anchor Enterprise](anchore.com/) - :yen: Analyze images for CVE vulnerabilities and against custom security policies.
- <b><code>&nbsp;10991⭐</code></b> <b><code>&nbsp;&nbsp;1208🍴</code></b> [Clair](https://github.com/quay/clair)) - Clair is an open source project for the static analysis of vulnerabilities in appc and docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;448⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;132🍴</code></b> [Docker Scout](https://github.com/docker/scout-cli)) - Official Docker CLI for SBOM generation, vulnerability analysis, and policy evaluation.
- <b><code>&nbsp;12315⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;802🍴</code></b> [Grype](https://github.com/anchore/grype)) - A vulnerability scanner for container images, filesystems and SBOMs.
- <b><code>&nbsp;&nbsp;1725⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;436🍴</code></b> [oscap-docker](https://github.com/OpenSCAP/openscap)) - OpenSCAP provides oscap-docker tool which is used to scan Docker containers and images.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [pindock](https://github.com/deadnews/pindock)) - Pin and update Docker image digests in Dockerfiles and compose files.
- <b><code>&nbsp;&nbsp;9041⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;862🍴</code></b> [Syft](https://github.com/anchore/syft)) - CLI tool and library for generating a Software Bill of Materials (SBOM) from container images and filesystems.
- <b><code>&nbsp;35262⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;404🍴</code></b> [Trivy](https://github.com/aquasecurity/trivy)) - Aqua Security's open source simple and comprehensive vulnerability scanner for containers (suitable for CI).

### Supply Chain

Signing, attestation, and provenance for container images.

- <b><code>&nbsp;&nbsp;5979⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;743🍴</code></b> [cosign](https://github.com/sigstore/cosign)) - Container signing, verification, and transparency log for OCI artifacts.
- <b><code>&nbsp;&nbsp;1003⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;155🍴</code></b> [in-toto](https://github.com/in-toto/in-toto)) - Framework for supply chain attestations; underpins SLSA and cosign provenance.
- <b><code>&nbsp;&nbsp;&nbsp;173⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;71🍴</code></b> [policy-controller](https://github.com/sigstore/policy-controller)) - Kubernetes admission controller enforcing cosign signatures on container images.
- <b><code>&nbsp;&nbsp;&nbsp;534⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;77🍴</code></b> [witness](https://github.com/in-toto/witness)) - Generate and verify in-toto attestations across the build pipeline.

## Running Containers

### Composition

- <b><code>&nbsp;&nbsp;3741⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;254🍴</code></b> [Composerize](https://github.com/magicmark/composerize)) - Convert docker run commands into docker-compose files.
- <b><code>&nbsp;&nbsp;&nbsp;301⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27🍴</code></b> [ctk](https://github.com/ctk-hq/ctk)) - Visual composer for container based workloads.
- <b><code>&nbsp;10533⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;798🍴</code></b> [kompose](https://github.com/kubernetes/kompose)) - Go from Docker Compose to Kubernetes.
- <b><code>&nbsp;&nbsp;&nbsp;382⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [plash](https://github.com/ihucos/plash)) - A container run and build engine - runs inside docker.
- <b><code>&nbsp;&nbsp;6092⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;596🍴</code></b> [podman-compose](https://github.com/containers/podman-compose)) - A script to run docker-compose.yml using podman.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;36⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [Smalte](https://github.com/roquie/smalte)) – Dynamically configure applications that require static configuration in docker container.

### Orchestration

- <b><code>&nbsp;&nbsp;&nbsp;241⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;84🍴</code></b> [CloudSlang](https://github.com/CloudSlang/cloud-slang)) - CloudSlang is a workflow engine to create Docker process automation.
- <b><code>&nbsp;&nbsp;3225⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;99🍴</code></b> [docker rollout](https://github.com/Wowu/docker-rollout)) - Zero downtime deployment for Docker Compose services.
- <b><code>122590⭐</code></b> <b><code>&nbsp;43225🍴</code></b> [Kubernetes](https://github.com/kubernetes/kubernetes)) - Open source orchestration system for Docker containers by Google.
- <b><code>&nbsp;&nbsp;5368⭐</code></b> <b><code>&nbsp;&nbsp;1662🍴</code></b> [Mesos](https://github.com/apache/mesos)) - Resource/Job scheduler for containers, VM's and physical hosts.
- [Nebula](https://github.com/nebula-orchestrator) - A Docker orchestration tool designed to manage massive scale distributed clusters.
- <b><code>&nbsp;16531⭐</code></b> <b><code>&nbsp;&nbsp;2088🍴</code></b> [Nomad](https://github.com/hashicorp/nomad)) - Easily deploy applications at any scale. A Distributed, Highly Available, Datacenter-Aware Scheduler.
- <b><code>&nbsp;25626⭐</code></b> <b><code>&nbsp;&nbsp;3199🍴</code></b> [Rancher](https://github.com/rancher/rancher)) - An open source project that provides a complete platform for operating Docker in production.
- <b><code>&nbsp;&nbsp;&nbsp;873⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;75🍴</code></b> [Swarm-cronjob](https://github.com/crazy-max/swarm-cronjob)) - Create jobs on a time-based schedule on Swarm.

### Deployment & Platforms

Self-hosted and managed cloud platforms (PaaS/CaaS, deployment automation). Commercial entries marked `:yen:`.

- 🌎 [Amazon ECS](aws.amazon.com/ecs/) - :yen: A management service on EC2 that supports Docker containers.
- 🌎 [Appfleet](appfleet.com/) - :yen: Edge platform to deploy and manage containerized services globally; routes traffic to the closest location for low latency.
- 🌎 [Azure AKS](azure.microsoft.com/en-us/products/kubernetes-service/) - :yen: Fully managed Kubernetes container orchestration service.
- 🌎 [blackfish](gitlab.com/blackfish/blackfish) - A CoreOS VM to build swarm clusters for Dev & Production.
- 🌎 [BosnD](gitlab.com/n0r1sk/bosnd) - BosnD, the boatswain daemon - A dynamic configuration file writer & service reloader for dynamically changing container environments.
- <b><code>&nbsp;15047⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;976🍴</code></b> [caprover](https://github.com/caprover/caprover)) - [Previously known as CaptainDuckDuck] Automated Scalable Webserver Package (automated Docker+nginx) - Heroku on Steroids.
- 🌎 [Cloud 66](www.cloud66.com) - :yen: Full-stack hosted container management as a service.
- 🌎 [Cloud Run Compose](docs.cloud.google.com/run/docs/deploy-run-compose) - :yen: Deploy `docker-compose.yaml` files directly to Google Cloud Run as a managed service.
- <b><code>&nbsp;&nbsp;1891⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;177🍴</code></b> [Convox Rack](https://github.com/convox/rack)) - Convox Rack is open source PaaS built on top of expert infrastructure automation and devops best practices.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [docker-to-iac](https://github.com/deploystackio/docker-to-iac)) - Translate docker run and commit into Infrastructure as Code templates for AWS, Render.com and DigitalOcean.
- <b><code>&nbsp;31909⭐</code></b> <b><code>&nbsp;&nbsp;2040🍴</code></b> [Dokku](https://github.com/dokku/dokku)) - Docker powered mini-Heroku that helps you build and manage the lifecycle of applications.
- <b><code>&nbsp;&nbsp;1152⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;59🍴</code></b> [Exoframe](https://github.com/exoframejs/exoframe)) - A self-hosted tool that allows simple one-command deployments using Docker.
- 🌎 [Giant Swarm](www.giantswarm.io/) - :yen: Simple microservice infrastructure. Deploy your containers in seconds.
- 🌎 [Google Container Engine](docs.cloud.google.com/kubernetes-engine/docs) - :yen: Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
- <b><code>&nbsp;&nbsp;1564⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;307🍴</code></b> [Grafeas](https://github.com/grafeas/grafeas)) - A common API for metadata about containers, from image and build details to security vulnerabilities.
- 🌎 [Mesosphere DC/OS Platform](d2iq.com/products/dcos) - :yen: Integrated platform for data and containers built on Apache Mesos.
- [OpenShift][openshift] - An open source PaaS built on [Kubernetes][kubernetes] and optimized for Dockerized app development and deployment by 🌎 [Red Hat](www.redhat.com/en)
- 🌎 [Red Hat OpenShift Dedicated](www.redhat.com/en/technologies/cloud-computing/openshift/dedicated) - :yen: Fully-managed Red Hat® OpenShift® service on Amazon Web Services and Google Cloud.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;57⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [swarm-ansible](https://github.com/LombardiDaniel/swarm-ansible?tab=readme-ov-file)) - Swarm-Ansible bootstraps a production-ready swarm cluster using ansible. Comes with tools to automate CI, help monitoring and traefik pre-configured for SSL certificates and simple-auth. Comes with a private registry and more!.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [SwarmManagement](https://github.com/hansehe/SwarmManagement)) - Swarm Management is a python application, installed with pip. The application makes it easy to manage a Docker Swarm by configuring a single yaml file describing which stacks to deploy, and which networks, configs or secrets to create.
- 🌎 [Triton](www.joyent.com/) - :yen: Elastic container-native infrastructure.
- <b><code>&nbsp;&nbsp;5279⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;553🍴</code></b> [Tsuru](https://github.com/tsuru/tsuru)) - Tsuru is an extensible and open source Platform as a Service software.
- <b><code>&nbsp;&nbsp;4688⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;234🍴</code></b> [werf](https://github.com/werf/werf)) - Werf is a CI/CD tool for building Docker images efficiently and deploying them to Kubernetes using GitOps.

### Garbage Collection

- <b><code>&nbsp;&nbsp;&nbsp;372⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;49🍴</code></b> [docker-custodian](https://github.com/Yelp/docker-custodian)) - Keep docker hosts tidy.
- <b><code>&nbsp;&nbsp;&nbsp;695⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;42🍴</code></b> [Docuum](https://github.com/stepchowfun/docuum)) - Least recently used (LRU) eviction of Docker images.

## Networking & Proxies

### Networking

Container networking, overlay networks, DNS/service-discovery bridges.

- [Calico][calico] - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [docker-dns](https://github.com/bytesharky/docker-dns)) - Lightweight DNS forwarder for Docker containers, resolves container names with custom suffixes (e.g. `.docker`) on the host to simplify service discovery.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Flannel](https://github.com/coreos/flannel/)) - Flannel is a virtual network that gives a subnet to each host for use with container runtimes.
- <b><code>&nbsp;10725⭐</code></b> <b><code>&nbsp;&nbsp;1090🍴</code></b> [netshoot](https://github.com/nicolaka/netshoot)) - The netshoot container has a powerful set of networking tools to help troubleshoot Docker networking issues.
- <b><code>&nbsp;&nbsp;4251⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;725🍴</code></b> [Pipework](https://github.com/jpetazzo/pipework)) - Software-Defined Networking for Linux Containers, Pipework works with "plain" LXC containers, and with the awesome Docker.
- <b><code>&nbsp;&nbsp;4676⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;904🍴</code></b> [registrator](https://github.com/gliderlabs/registrator)) - Service registry bridge for Docker.

### Reverse Proxy

Container-aware reverse proxies, ingress, and TLS-terminating front-ends with auto-discovery.

- <b><code>&nbsp;10561⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;611🍴</code></b> [BunkerWeb](https://github.com/bunkerity/bunkerweb)) - Open-source and next-gen Web Application Firewall (WAF).
- <b><code>&nbsp;&nbsp;4500⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;216🍴</code></b> [caddy-docker-proxy](https://github.com/lucaslorentz/caddy-docker-proxy)) - Caddy-based reverse proxy, configured with service or container labels.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;35⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [caddy-docker-upstreams](https://github.com/invzhi/caddy-docker-upstreams)) - Docker upstreams module for Caddy, configured with container labels.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;34⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Docker Dnsmasq Updater](https://github.com/moonbuggy/docker-dnsmasq-updater)) - Update a remote dnsmasq server with Docker container hostnames.
- <b><code>&nbsp;&nbsp;&nbsp;318⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;186🍴</code></b> [docker-flow-proxy](https://github.com/docker-flow/docker-flow-proxy)) - Reconfigures proxy every time a new service is deployed, or when a service is scaled.
- <b><code>&nbsp;&nbsp;7712⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;830🍴</code></b> [Let's Encrypt Nginx-proxy Companion](https://github.com/nginx-proxy/docker-letsencrypt-nginx-proxy-companion)) - A lightweight companion container for the nginx-proxy. It allow the creation/renewal of Let's Encrypt certificates automatically.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [mesh-router](https://github.com/Yundera/mesh-router)) - Free domain(nsl.sh) provider for Docker containers with automatic HTTPS routing. Uses Wireguard VPN to securely route subdomain requests across networks. Ideal for self-hosted NAS and cloud deployments.
- <b><code>&nbsp;33008⭐</code></b> <b><code>&nbsp;&nbsp;3753🍴</code></b> [Nginx Proxy Manager](https://github.com/jc21/nginx-proxy-manager)) - A beautiful web interface for proxying web based services with SSL.
- [nginx-proxy][nginxproxy] - Automated nginx proxy for Docker containers using docker-gen.
- <b><code>&nbsp;&nbsp;1400⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;108🍴</code></b> [OpenResty Manager](https://github.com/Safe3/openresty-manager)) - The easiest using, powerful and beautiful OpenResty Manager(Nginx Enhanced Version), open source alternative to OpenResty Edge.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;75⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;12🍴</code></b> [Swarm Router](https://github.com/flavioaiello/swarm-router)) - A zero config service name based router for docker swarm mode with a fresh and more secure approach.
- <b><code>&nbsp;63387⭐</code></b> <b><code>&nbsp;&nbsp;6025🍴</code></b> [Træfɪk](https://github.com/containous/traefik)) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd.

## Storage & Data

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Label Backup](https://github.com/resulgg/label-backup)) - A lightweight, Docker-aware backup agent that automatically discovers and backs up containerized databases (PostgreSQL, MySQL, MongoDB, Redis) based on Docker labels. Supports local storage and S3-compatible destinations with flexible scheduling via cron expressions.
- <b><code>&nbsp;&nbsp;3609⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;135🍴</code></b> [Docker Volume Backup](https://github.com/offen/docker-volume-backup)) Backup Docker volumes locally or to any S3 compatible storage.
- <b><code>&nbsp;&nbsp;1142⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;165🍴</code></b> [Netshare](https://github.com/ContainX/docker-volume-netshare)) Docker NFS, AWS EFS, Ceph & Samba/CIFS Volume Plugin.
- 🌎 [portworx](portworx.com) - :yen: Decentralized storage solution for persistent, shared and replicated volumes.
- 🌎 [quobyte](www.quobyte.com/) - :yen: Fully fault-tolerant distributed file system with a docker volume driver.
- <b><code>&nbsp;&nbsp;2222⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;331🍴</code></b> [REX-Ray](https://github.com/rexray/rexray)) provides a vendor agnostic storage orchestration engine. The primary design goal is to provide persistent storage for Docker, Kubernetes, and Mesos.

## Observability

Monitor Docker hosts, containers, and the services running inside them. Self-hosted and SaaS together; commercial entries marked `:yen:`.

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [ADRG](https://github.com/jaldertech/adrg)) - Dynamic Docker resource governor using cgroups v2 to manage system load.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [AppDynamics](https://github.com/Appdynamics/docker-monitoring-extension)) - :yen: Docker Monitoring extension gathers metrics from the Docker Remote API, either using Unix Socket or TCP.
- <b><code>&nbsp;&nbsp;1925⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;260🍴</code></b> [Autoheal](https://github.com/willfarrell/docker-autoheal)) - Monitor and restart unhealthy docker containers automatically.
- 🌎 [Better Stack](betterstack.com/community/guides/scaling-docker/) - :yen: A Docker-compatible observability stack that delivers log aggregation and uptime monitoring for containerized apps.
- <b><code>&nbsp;19166⭐</code></b> <b><code>&nbsp;&nbsp;2472🍴</code></b> [cAdvisor](https://github.com/google/cadvisor)) - Analyzes resource usage and performance characteristics of running containers.
- 🌎 [Datadog](www.datadoghq.com/) - :yen: Full-stack monitoring service with first-class Docker, Kubernetes, and Mesos support.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [DLIA](https://github.com/zorak1103/dlia)) - DLIA is an AI-powered Docker log monitoring agent that uses Large Language Models (LLMs) to intelligently analyze container logs, detect anomalies, and provide contextual insights over time.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [docker-exporter](https://github.com/dlepaux/docker-exporter)) - Lightweight Prometheus exporter for Docker container metrics written in Rust. Correct cgroup v2 memory working set on ARM64 (Raspberry Pi 5), runs non-root with a read-only socket, ~7 MiB idle RAM.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [Docker-Sentinel](https://github.com/Will-Luck/Docker-Sentinel)) - Automated container updates with per-container policies, rollback safety, and a real-time web dashboard.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [DockProbe](https://github.com/deep-on/dockprobe)) - Lightweight Docker monitoring dashboard in a single container. Real-time metrics, 6 anomaly detection rules, Telegram alerts, and 16 automated security scans. Zero config, ~50MB RAM.
- 🌎 [DockProc](gitlab.com/n0r1sk/dockproc) - I/O monitoring for containers on processlevel.
- <b><code>&nbsp;&nbsp;6515⭐</code></b> <b><code>&nbsp;&nbsp;1757🍴</code></b> [dockprom](https://github.com/stefanprodan/dockprom)) - Docker hosts and containers monitoring with Prometheus, Grafana, cAdvisor, NodeExporter and AlertManager.
- <b><code>&nbsp;&nbsp;&nbsp;416⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;18🍴</code></b> [Doku](https://github.com/amerkurev/doku)) - Doku is a simple web-based application that allows you to monitor Docker disk usage.
- [Dozzle](dozzle) - Monitor container logs in real-time with a browser or mobile device.
- <b><code>&nbsp;&nbsp;&nbsp;198⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [Drydock](https://github.com/CodesWhat/drydock)) - Container update monitoring with web dashboard, 23 registry providers, 20 notification triggers, and distributed agent architecture.
- 🌎 [Dynatrace](docs.dynatrace.com/docs/observe/infrastructure-observability/container-platform-monitoring) - :yen: Monitor containerized applications without installing agents or modifying your Run commands.
- 🌎 [Grafana Docker Dashboard Template](grafana.com/grafana/dashboards/179-docker-prometheus-monitoring/) - A template for your Docker, Grafana and Prometheus stack.
- <b><code>&nbsp;&nbsp;&nbsp;324⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14🍴</code></b> [Maintenant](https://github.com/kolapsis/maintenant)) - Self-discovering infrastructure monitoring for Docker and Kubernetes. Auto-detects containers via labels, with endpoint monitoring, heartbeats, TLS certificates, resource metrics, update intelligence, and a built-in status page. Single binary with embedded SPA.
- 🌎 [Middleware](middleware.io/) - :yen: Monitor Docker hosts, containers, logs, and application performance from a unified observability platform.
- 🌎 [Site24x7](www.site24x7.com/docker-monitoring.html) - :yen: Docker Monitoring for DevOps and IT, SaaS Pay-per-Host model.
- 🌎 [Sysdig Monitor](www.sysdig.com/products/monitor) - :yen: Software or SaaS service that monitors, alerts, and troubleshoots containers using system calls; container-specific features for Docker and Kubernetes.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Wiremap](https://github.com/codeofmario/wiremap)) - A self-hosted visual Docker network topology explorer with real-time log streaming, live stats, embedded terminal, and container inspection.

## Security

Container hardening, runtime security, policy, compliance, and forensics. Self-hosted and commercial together; commercial entries marked `:yen:`.

- 🌎 [Aqua Security](www.aquasec.com) - :yen: Securing container-based applications from Dev to Production on any platform.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [buildcage](https://github.com/dash14/buildcage)) - Restricts outbound network access during Docker builds to prevent supply chain attacks, working as a drop-in BuildKit remote driver for Docker Buildx, with ready-to-use GitHub Actions.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;86⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [CetusGuard](https://github.com/hectorm/cetusguard)) - CetusGuard is a tool that protects the Docker daemon socket by filtering calls to its API endpoints.
- <b><code>&nbsp;&nbsp;8758⭐</code></b> <b><code>&nbsp;&nbsp;1342🍴</code></b> [Checkov](https://github.com/bridgecrewio/checkov)) - Static analysis for infrastructure as code manifests (Terraform, Kubernetes, Cloudformation, Helm, Dockerfile, Kustomize) find security misconfiguration and fix them.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;97⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10🍴</code></b> [container-explorer](https://github.com/google/container-explorer)) - Forensic utility to explore Docker and containerd container details from mounted disk images.
- <b><code>&nbsp;&nbsp;5274⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;635🍴</code></b> [Deepfence Threat Mapper](https://github.com/deepfence/ThreatMapper)) - Powerful runtime vulnerability scanner for kubernetes, virtual machines and serverless.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [Den](https://github.com/us/den)) - Self-hosted sandbox runtime for AI agents with Docker containers, security hardening, REST API and WebSocket support.
- <b><code>&nbsp;&nbsp;9642⭐</code></b> <b><code>&nbsp;&nbsp;1039🍴</code></b> [docker-bench-security](https://github.com/docker/docker-bench-security)) - Script that checks for dozens of common best-practices around deploying Docker containers in production.
- <b><code>&nbsp;&nbsp;2540⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;200🍴</code></b> [docker-socket-proxy](https://github.com/Tecnativa/docker-socket-proxy)) - HAProxy-based fine-grained filter for the Docker API socket; widely used to expose a restricted socket to reverse proxies and homelab stacks.
- <b><code>&nbsp;&nbsp;2644⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;368🍴</code></b> [KICS](https://github.com/checkmarx/kics)) - An infrastructure-as-code scanning tool, find security vulnerabilities, compliance issues, and infrastructure misconfigurations early in the development cycle. Can be extended for additional policies.
- 🌎 [Prisma Cloud](www.paloaltonetworks.com/prisma/cloud) - :yen: (Previously Twistlock Security Suite) detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [segspec](https://github.com/dormstern/segspec)) - Extracts network dependencies from Docker Compose, Kubernetes manifests, Helm charts, and other config files to generate Kubernetes NetworkPolicies with evidence tracing.
- <b><code>&nbsp;&nbsp;9007⭐</code></b> <b><code>&nbsp;&nbsp;1021🍴</code></b> [Sysdig Falco](https://github.com/falcosecurity/falco)) - Sysdig Falco is an open source container security monitor. It can monitor application, container, host, and network activity and alert on unauthorized activity.
- 🌎 [Sysdig Secure](www.sysdig.com/solutions/cloud-detection-and-response-cdr) - :yen: Sysdig Secure addresses run-time security through behavioral monitoring and defense, and provides deep forensics based on open source Sysdig for incident response.
- 🌎 [Trend Micro DeepSecurity](www.trendmicro.com/en_us/business/products/hybrid-cloud/deep-security.html) - :yen: Trend Micro DeepSecurity offers runtime protection for container workloads and hosts as well as preruntime scanning of images to identify vulnerabilities, malware and content such as hardcoded secrets.

## User Interfaces

### Desktop

Native desktop applications for managing and monitoring docker hosts and clusters

- <b><code>&nbsp;&nbsp;&nbsp;162⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [Docker DB Manager](https://github.com/AbianS/docker-db-manager)) - Desktop app for managing Docker database containers with visual interface and one-click operations.
- 🌎 [Docker Desktop](www.docker.com/products/docker-desktop/) - Official native app. Only for Windows and MacOS.
- <b><code>&nbsp;&nbsp;&nbsp;603⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;96🍴</code></b> [Simple Docker UI](https://github.com/felixgborrego/simple-docker-ui)) - Built on Electron.
- <b><code>&nbsp;&nbsp;&nbsp;373⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;13🍴</code></b> [Stevedore](https://github.com/slonopotamus/stevedore)) - Good Docker Desktop replacement for Windows. Both Linux and Windows Containers are supported. [slonopotamus](https://github.com/slonopotamus).

### Terminal

TUIs, CLI tools, and shell integrations for Docker.

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;91⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [d4s](https://github.com/jr-k/d4s)) - A fast, keyboard-driven terminal UI to manage Docker containers, Compose stacks, and Swarm services with the ergonomics of K9s.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;14⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [dcinja](https://github.com/Falldog/dcinja)) - The powerful and smallest binary size of template engine for docker command line environment.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;23⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [dctl](https://github.com/FabienD/docker-stack)) - Dctl is a Cli tool that helps developers by allowing them to execute all docker compose commands anywhere in the terminal, and more.
- <b><code>&nbsp;&nbsp;&nbsp;133⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [decompose](https://github.com/s0rg/decompose)) - Reverse-engineering tool for docker environments.
- <b><code>&nbsp;54067⭐</code></b> <b><code>&nbsp;&nbsp;1993🍴</code></b> [dive](https://github.com/wagoodman/dive)) - A tool for exploring each layer in a docker image.
- <b><code>&nbsp;&nbsp;&nbsp;152⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [docker pushrm](https://github.com/christian-korneck/docker-pushrm)) - A Docker CLI plugin that lets you push the README.md file from the current directory to Docker Hub. Also supports Quay and Harbor.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [docker-captain](https://github.com/lucabello/docker-captain)) - A friendly CLI to manage multiple Docker Compose deployments with style — powered by Typer, Rich, questionary, and sh.
- <b><code>&nbsp;&nbsp;&nbsp;564⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;85🍴</code></b> [dockerfile-mode](https://github.com/spotify/dockerfile-mode)) - An Emacs mode for handling Dockerfiles.
- <b><code>&nbsp;&nbsp;&nbsp;264⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [dockerfilegraph](https://github.com/patrickhoefler/dockerfilegraph)) - Visualize your multi-stage Dockerfiles.
- <b><code>&nbsp;&nbsp;4014⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;163🍴</code></b> [dockly](https://github.com/lirantal/dockly)) - An interactive shell UI for managing Docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;318⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [DockMate](https://github.com/shubh-io/dockmate)) - Lightweight terminal-based Docker and Podman manager with a text-based user interface,.
- <b><code>&nbsp;&nbsp;2560⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;292🍴</code></b> [DockSTARTer](https://github.com/GhostWriters/DockSTARTer)) - DockSTARTer helps you get started with home server apps running in Docker.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;39⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [dprs](https://github.com/durableprogramming/dprs)) - A developer-focused TUI for managing Docker containers with real-time log streaming and container management.
- <b><code>&nbsp;&nbsp;3257⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;102🍴</code></b> [dry](https://github.com/moncho/dry)) - An interactive CLI for Docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;114⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [easydocker](https://github.com/joao-zanutto/easydocker)) - A Terminal UI highly inpired by k9s levaraging beatiful BubbleTea graphics.
- <b><code>&nbsp;&nbsp;&nbsp;639⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [goManageDocker](https://github.com/ajayd-san/gomanagedocker)) - TUI tool to view and manage your docker objects blazingly fast with sensible keybindings, also supports VIM navigation out of the box.
- <b><code>&nbsp;51228⭐</code></b> <b><code>&nbsp;&nbsp;1617🍴</code></b> [lazydocker](https://github.com/jesseduffield/lazydocker)) - The lazier way to manage everything docker. A simple terminal UI for both docker and docker-compose, written in Go with the gocui library.
- <b><code>&nbsp;&nbsp;1288⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;31🍴</code></b> [lazyjournal](https://github.com/Lifailon/lazyjournal)) - A interface for reading and filtering the logs output of Docker and Podman containers like [Dozzle](dozzle) but for the terminal with support for fuzzy find, regex and output coloring.
- <b><code>&nbsp;&nbsp;1693⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;45🍴</code></b> [oxker](https://github.com/mrjackwills/oxker)) - A simple tui to view & control docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;112⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [proco](https://github.com/shiwaforce/poco)) - Proco will help you to organise and manage Docker, Docker-Compose, Kubernetes projects of any complexity using simple YAML config files to shorten the route from finding your project to initialising it in your local environment.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;97⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [scuba](https://github.com/JonathonReinhart/scuba)) - Transparently use Docker containers to encapsulate software build environments,.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;86⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [supdock](https://github.com/segersniels/supdock)) - Allows for slightly more visual usage of Docker with an interactive prompt.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [swarmcli](https://github.com/Eldara-Tech/swarmcli)) - Swarm Management at the speed of thought — with real-time log streaming, instant shell access to containers, seamless port forwarding, and on-demand secret reveal capabilities, giving you full control over your Docker Swarm without breaking your flow.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;85⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3🍴</code></b> [tdocker](https://github.com/pivovarit/tdocker)) - A `docker ps` replacement for everyday container operations.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [wharf](https://github.com/idesyatov/wharf)) - A k9s-inspired TUI for Docker Compose with vim-style navigation, real-time CPU/MEM monitoring with braille charts, container file browser, SSH remote host support, and command mode.

### Web

- <b><code>&nbsp;&nbsp;5636⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;201🍴</code></b> [Arcane](https://github.com/getarcaneapp/arcane)) - An easy and modern Docker management platform, built with everybody in mind.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;85⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [CASA](https://github.com/knrdl/casa)) - Outsource the administration of a handful of containers to your co-workers,.
- <b><code>&nbsp;&nbsp;&nbsp;257⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;46🍴</code></b> [Container Web TTY](https://github.com/wrfly/container-web-tty)) - Connect your containers via a web-tty.
- <b><code>&nbsp;&nbsp;&nbsp;686⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;62🍴</code></b> [Docker Registry Browser](https://github.com/klausmeyer/docker-registry-browser)) - Web Interface for the Docker Registry HTTP API v2.
- <b><code>&nbsp;&nbsp;3345⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;594🍴</code></b> [docker-swarm-visualizer](https://github.com/dockersamples/docker-swarm-visualizer)) - Visualizes Docker services on a Docker Swarm (for running demos).
- <b><code>&nbsp;23358⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;749🍴</code></b> [dockge](https://github.com/louislam/dockge)) - Easy-to-use and reactive self-hosted docker compose.yaml stack-oriented manager.
- <b><code>&nbsp;11289⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;323🍴</code></b> [Komodo](https://github.com/mbecker20/komodo)) - A tool to build and deploy software on many servers.
- <b><code>&nbsp;37615⭐</code></b> <b><code>&nbsp;&nbsp;2831🍴</code></b> [Portainer](https://github.com/portainer/portainer)) - A lightweight management UI for managing your Docker hosts or Docker Swarm clusters.
- <b><code>&nbsp;&nbsp;3444⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;311🍴</code></b> [Swarmpit](https://github.com/swarmpit/swarmpit)) - Swarmpit provides simple and easy to use interface for your Docker Swarm cluster. You can manage your stacks, services, secrets, volumes, networks etc.
- <b><code>&nbsp;&nbsp;&nbsp;115⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5🍴</code></b> [usulnet](https://github.com/fr4nsys/usulnet)) - A complete and modern Docker management platform designed for sysadmin, devops with enterprise grade tools, cve scanner, ssh, rdp on web and much more.

### IDE Integrations

- JetBrains IDEs (IntelliJ IDEA, GoLand, WebStorm, CLion etc.) has 🌎 [built-in Docker plugin](www.jetbrains.com/help/idea/docker.html#managing-images)
- Eclipse 🌎 [Docker Tooling plugin](www.eclipse.org/community/eclipse_newsletter/2016/july/article2.php)
- <b><code>&nbsp;&nbsp;&nbsp;820⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;80🍴</code></b> [docker.el](https://github.com/Silex/docker.el)) Manage docker from Emacs.

## Developer Workflow

### API Client

- <b><code>&nbsp;&nbsp;&nbsp;148⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8🍴</code></b> [contajners](https://github.com/lispyclouds/contajners)) - An idiomatic, data-driven, REPL friendly Clojure client for OCI container engines.
- <b><code>&nbsp;&nbsp;&nbsp;120⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;28🍴</code></b> [Docker Client for JVM](https://github.com/gesellix/docker-client)) - A Docker remote api client library for the JVM, written in Groovy.
- 🌎 [Docker Client TypeScript](gitlab.com/masaeedu/docker-client) - Docker API client for JavaScript, automatically generated from Swagger API definition from moby repository.
- <b><code>&nbsp;&nbsp;&nbsp;247⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;33🍴</code></b> [docker-controller-bot](https://github.com/dgongut/docker-controller-bot)) - Telegram bot to control docker containers.
- <b><code>&nbsp;&nbsp;1930⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;652🍴</code></b> [docker-maven-plugin](https://github.com/fabric8io/docker-maven-plugin)) - A Maven plugin for running and creating Docker images.
- <b><code>&nbsp;&nbsp;2414⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;412🍴</code></b> [Docker.DotNet](https://github.com/Microsoft/Docker.DotNet)) - C#/.NET HTTP client for the Docker remote API.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;42⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22🍴</code></b> [Docker.Registry.DotNet](https://github.com/ChangemakerStudios/Docker.Registry.DotNet)) - .NET (C#) Client Library for interacting with a Docker Registry API (v2).
- <b><code>&nbsp;&nbsp;4897⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;489🍴</code></b> [dockerode](https://github.com/apocas/dockerode)) - Docker Remote API node.js module.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [go-dockerclient](https://github.com/fsouza/go-dockerclient/)) - Go HTTP client for the Docker remote API.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;81⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin)) - A Docker remote api plugin for Gradle.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;74⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;16🍴</code></b> [Portainer stack utils](https://github.com/greenled/portainer-stack-utils)) - Bash script to deploy/update/undeploy Docker stacks in a Portainer instance from a docker-compose yaml file.
- <b><code>&nbsp;&nbsp;&nbsp;732⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;113🍴</code></b> [sbt-docker](https://github.com/marcuslonnberg/sbt-docker)) - Create Docker images directly from sbt.

### CI/CD

Self-hosted CI engines, build accelerators, and hosted services that target Docker workflows. Commercial entries marked `:yen:`.

- 🌎 [Buddy](buddy.works) - :yen: The best of Git, build & deployment tools combined into one powerful tool that supercharged our development.
- <b><code>&nbsp;&nbsp;&nbsp;776⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;49🍴</code></b> [Captain](https://github.com/harbur/captain)) - Convert your Git workflow to Docker containers ready for Continuous Delivery.
- 🌎 [CircleCI](circleci.com/) - :yen: Push or pull Docker images from your build environment, or build and run containers right on CircleCI.
- 🌎 [CodeFresh](octopus.com/codefresh) - :yen: End-to-end build, test, and share for Docker applications, with automated testing.
- 🌎 [ConcourseCI](concourse-ci.org) - :yen: Pipeline-oriented CI SaaS platform for DevOps teams.
- <b><code>&nbsp;&nbsp;&nbsp;153⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;26🍴</code></b> [Defang](https://github.com/DefangLabs/defang)) - Deploy Docker Compose to your favorite cloud in minutes.
- 🌎 [Depot](depot.dev) - :yen: Build Docker images fast, in the cloud. Blazing fast compute, automatic intelligent caching, and zero configuration.
- <b><code>&nbsp;&nbsp;4686⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;150🍴</code></b> [Diun](https://github.com/crazy-max/diun)) - Receive notifications when an image or repository is updated on a Docker registry.
- <b><code>&nbsp;&nbsp;2342⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;86🍴</code></b> [dockcheck](https://github.com/mag37/dockcheck)) - A script checking updates for docker images without pulling then auto-update selected/all containers. With notifications, pruning and more.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/)) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
- <b><code>&nbsp;36267⭐</code></b> <b><code>&nbsp;&nbsp;3179🍴</code></b> [Drone](https://github.com/drone/drone)) - Continuous integration server built on Docker and configured using YAML files.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Gantry](https://github.com/shizunge/gantry)) - Automatically update selected Docker swarm services.
- 🌎 [GitLab Runner](gitlab.com/gitlab-org/gitlab-runner) - GitLab has integrated CI to test, build and deploy your code with the use of GitLab runners.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [Jaypore CI](https://github.com/theSage21/jaypore_ci)) - Simple, very flexible, powerful CI / CD / automation system configured in Python. Offline and local first.
- <b><code>&nbsp;&nbsp;&nbsp;159⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;20🍴</code></b> [Kraken CI](https://github.com/Kraken-CI/kraken)) - Modern CI/CD, open-source, on-premise system that is highly scalable and focused on testing. One of its executors is Docker. Developed.
- 🌎 [Screwdriver](screwdriver.cd/) - :yen: Yahoo's OpenSource buildplatform designed for Continous Delivery.
- 🌎 [Semaphore CI](semaphore.io/) - :yen: High-performance cloud CI that builds, tests and ships containers to production.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;49⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22🍴</code></b> [Skipper](https://github.com/Stratoscale/skipper)) - Easily dockerize your Git repository.
- 🌎 [Tekton CD](tekton.dev/) - A cloud-native pipeline resource.
- 🌎 [TravisCI](www.travis-ci.com/) - :yen: Hosted CI for GitHub projects with Docker support.

### Development Environment

- <b><code>&nbsp;13342⭐</code></b> <b><code>&nbsp;&nbsp;1306🍴</code></b> [coder](https://github.com/coder/coder)) - Remote development machines powered by Terraform or Docker.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;45⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9🍴</code></b> [dde](https://github.com/whatwedo/dde)) - Local development environment toolset based on Docker.
- <b><code>&nbsp;&nbsp;1335⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;50🍴</code></b> [DIP](https://github.com/bibendi/dip)) - CLI utility for straightforward provisioning and interacting with an application configured by docker-compose.
- <b><code>&nbsp;&nbsp;&nbsp;115⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [EnvCLI](https://github.com/EnvCLI/EnvCLI)) - Replace your local installation of Node, Go, ... with project-specific docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;632⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;17🍴</code></b> [Gebug](https://github.com/moshebe/gebug)) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [HarborPilot](https://github.com/potterwhite/HarborPilot)) - Automated multi-platform Docker image builder for embedded Linux development (RK3588, RV1126, RK3568). Features three-layer config inheritance, PORT_SLOT-based port allocation, and cross-version Ubuntu support (20.04/22.04/24.04).
- <b><code>&nbsp;&nbsp;4234⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;520🍴</code></b> [Lando](https://github.com/lando/lando)) - Lando is for developers who want to quickly specify and painlessly spin up the services and tools needed to develop their projects.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;21⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;0🍴</code></b> [uniget](https://github.com/uniget-org/cli)) - Uni(versal)get, the installer and updater for container tools and beyond (formerly docker-setup).
- <b><code>&nbsp;&nbsp;1108⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;122🍴</code></b> [Zsh-in-Docker](https://github.com/deluan/zsh-in-docker)) - Install Zsh, Oh-My-Zsh and plugins inside a Docker container with one line!.

### Serverless

- <b><code>&nbsp;&nbsp;6779⭐</code></b> <b><code>&nbsp;&nbsp;1179🍴</code></b> [Apache OpenWhisk](https://github.com/apache/openwhisk)) - A serverless, open source cloud platform that executes functions in response to events at any scale.
- 🌎 [Koyeb](www.koyeb.com/) - :yen: Koyeb is a developer-friendly serverless platform to deploy apps globally. Seamlessly run Docker containers, web apps, and APIs with git-based deployment, native autoscaling, a global edge network, and built-in service mesh and discovery.
- <b><code>&nbsp;26163⭐</code></b> <b><code>&nbsp;&nbsp;1967🍴</code></b> [OpenFaaS](https://github.com/openfaas/faas)) - A complete serverless functions framework for Docker and Kubernetes.

### Testing

- <b><code>&nbsp;&nbsp;2482⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;206🍴</code></b> [Container Structure Test](https://github.com/GoogleContainerTools/container-structure-test)) - A framework to validate the structure of an image by checking the outputs of commands or the contents of the filesystem.
- <b><code>&nbsp;&nbsp;5893⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;490🍴</code></b> [dgoss](https://github.com/goss-org/goss/tree/master/extras/dgoss)) - A fast YAML based tool for validating docker containers.
- <b><code>&nbsp;&nbsp;&nbsp;541⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;97🍴</code></b> [Kurtosis](https://github.com/kurtosis-tech/kurtosis)) - A composable build system for multi-container test environments that provides developers with: a powerful Python-like SDK for environment configuration, a compile-time validator to verify environment behavior & setup, and a runtime for environment execution, monitoring, & debugging capabilities.
- <b><code>&nbsp;&nbsp;3031⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;206🍴</code></b> [Pumba](https://github.com/alexei-led/pumba)) - Chaos testing tool for Docker. Can be deployed on kubernetes and CoreOS cluster.

### Wrappers

- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;97⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;25🍴</code></b> [Hokusai](https://github.com/artsy/hokusai)) - A Docker + Kubernetes CLI for application developers; used to containerize an application and to manage its lifecycle throughout development, testing, and release cycles. From [artsy](https://github.com/artsy).
- <b><code>&nbsp;&nbsp;2209⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;89🍴</code></b> [Preevy](https://github.com/livecycle/preevy)) - Preview environments for Docker and Docker Compose projects. Test your changes and get feedback from devs and non-devs (Product/Design) by deploying pull requests to the your cloud provider as part of your CI pipeline.
- <b><code>&nbsp;&nbsp;&nbsp;894⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;63🍴</code></b> [subuser](https://github.com/subuser-security/subuser)) - Makes it easy to securely and portably run graphical desktop applications in Docker.
- <b><code>&nbsp;&nbsp;1747⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;167🍴</code></b> [udocker](https://github.com/indigo-dc/udocker)) - A tool to execute simple docker containers in batch or interactive systems without root privileges.
- 🌎 [Vagrant - Docker provider](developer.hashicorp.com/vagrant/docs/providers/docker/basics) - Good starting point is <b><code>&nbsp;&nbsp;&nbsp;113⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;43🍴</code></b> [vagrant-docker-example](https://github.com/bubenkoff/vagrant-docker-example)).

## In-Container Tooling

Tools and applications that are either installed inside containers or designed to be run as a 🌎 [sidecar](learn.microsoft.com/en-us/azure/architecture/patterns/sidecar)

- <b><code>&nbsp;&nbsp;1648⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;67🍴</code></b> [cdebug](https://github.com/iximiuz/cdebug)) - Swiss-army knife for debugging running containers via ephemeral sidecars; works with Docker, containerd, and Kubernetes.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;56⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6🍴</code></b> [ckron](https://github.com/nicomt/ckron)) - A cron-style job scheduler for docker,.
- [CoreOS][coreos] - Linux for Massive Server Deployments
- <b><code>&nbsp;&nbsp;4627⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;612🍴</code></b> [docker-gen](https://github.com/jwilder/docker-gen)) - Generate files from docker container meta-data.
- <b><code>&nbsp;&nbsp;&nbsp;193⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;18🍴</code></b> [dockerize](https://github.com/powerman/dockerize)) - Utility to simplify running applications in docker containers.
- <b><code>&nbsp;&nbsp;4973⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;359🍴</code></b> [GoSu](https://github.com/tianon/gosu)) - Run this specific application as this specific user and get out of the pipeline (entrypoint script tool).
- <b><code>&nbsp;&nbsp;&nbsp;233⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;15🍴</code></b> [is-docker](https://github.com/sindresorhus/is-docker)) - Check if the process is running inside a Docker container.
- <b><code>&nbsp;&nbsp;&nbsp;142⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [microcheck](https://github.com/tarampampam/microcheck)) - Lightweight health check utilities for Docker containers (75 KB instead of 9.3 MB for httpcheck versus cURL) in pure C - http(s), port checks, and parallel execution are included.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Ofelia](https://github.com/mcuadros/ofelia/)) - Ofelia is a modern and low footprint job scheduler for docker environments, built on Go. Ofelia aims to be a replacement for the old fashioned cron. Supports configuration from container labels and/or configuration files.
- <b><code>&nbsp;&nbsp;1020⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;103🍴</code></b> [su-exec](https://github.com/ncopa/su-exec)) - This is a simple tool that will simply execute a program with different privileges. The program will be executed directly and not run as a child, like su and sudo does, which avoids TTY and signal issues. Why reinvent gosu? This does more or less exactly the same thing as gosu but it is only 10kb instead of 1.8MB.
- <b><code>&nbsp;&nbsp;2518⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;141🍴</code></b> [supercronic](https://github.com/aptible/supercronic)) - Crontab-compatible job runner, designed specifically to run in containers.

# Learning Resources

## Where to Start

- 🌎 [Benefits of using Docker](semaphore.io/blog/docker-benefits) for development and delivery, with a practical roadmap for adoption.
- 🌎 [Bootstrapping Microservices](www.manning.com/books/bootstrapping-microservices-with-docker-kubernetes-and-terraform) - A practical and project-based guide to building applications with microservices, starts by building a Docker image for a single microservice and publishing it to a private container registry, finishes by deploying a complete microservices application to a production Kubernetes cluster.
- <b><code>&nbsp;&nbsp;6044⭐</code></b> <b><code>&nbsp;&nbsp;2225🍴</code></b> [Docker Curriculum](https://github.com/prakhar1989/docker-curriculum)): A comprehensive tutorial for getting started with Docker. Teaches how to use Docker and deploy dockerized apps on AWS with Elastic Beanstalk and Elastic Container Service.
- 🌎 [Docker Documentation](docs.docker.com/): the official documentation.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;84⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;27🍴</code></b> [Docker for beginners](https://github.com/groda/big_data/blob/master/docker_for_beginners.md)): A tutorial for beginners who need to learn the basics of Docker—from "Hello world!" to basic interactions with containers, with simple explanations of the underlying concepts.
- 🌎 [Docker for novices](www.youtube.com/watch?v=xsjSadjKXns) An introduction to Docker for developers and testers who have never used it. (Video 1h40, recorded linux.conf.au 2019 — Christchurch, New Zealand)
- <b><code>&nbsp;&nbsp;&nbsp;289⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;252🍴</code></b> [Docker katas](https://github.com/eficode-academy/docker-katas)) A series of labs that will take you from "Hello Docker" to deploying a containerized web application to a server.
- 🌎 [Docker simplified in 55 seconds](www.youtube.com/watch?v=vP_4DlOH1G4): An animated high-level introduction to Docker. Think of it as a visual tl;dr that makes it easier to dive into more complex learning materials.
- 🌎 [Docker Training](training.mirantis.com) - :yen:
- <b><code>&nbsp;&nbsp;&nbsp;890⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;66🍴</code></b> [Dockerlings](https://github.com/furkan/dockerlings)): Learn docker from inside your terminal, with a modern TUI and bite sized exercises.
- 🌎 [Introduction à Docker](blog.stephane-robert.info/docs/conteneurs/moteurs-conteneurs/docker/) A dedicated section to master Docker on a French site about DevSecOps: From the basics to best practices, including optimizing, securing your containers...
- <b><code>&nbsp;&nbsp;&nbsp;243⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;32🍴</code></b> [Learn Docker](https://github.com/dwyl/learn-docker)): step-by-step tutorial and more resources (video, articles, cheat sheets)
- 🌎 [Learn Docker (Visually)](pagertree.com/learn/docker/overview) - A beginner-focused high-level overview of all the major components of Docker and how they fit together. Lots of high-quality images, examples, and resources.
- 🌎 [Play With Docker](training.play-with-docker.com/): PWD is a great way to get started with Docker from beginner to advanced users. Docker runs directly in your browser.
- <b><code>&nbsp;&nbsp;&nbsp;258⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;78🍴</code></b> [Practical Guide about Docker Commands in Spanish](https://github.com/brunocascio/docker-espanol)) This Spanish guide contains the use of basic docker commands with real life examples.
- <b><code>&nbsp;&nbsp;&nbsp;951⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;99🍴</code></b> [Setting Python Development Environment with VScode and Docker](https://github.com/RamiKrispin/vscode-python)): A step-by-step tutorial for setting up a dockerized Python development environment with VScode, Docker, and the Dev Container extension.
- 🌎 [The Docker Handbook](docker-handbook.farhan.dev/) An open-source book that teaches you the fundamentals, best practices and some intermediate Docker functionalities. The book is hosted on <b><code>&nbsp;&nbsp;&nbsp;868⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;126🍴</code></b> [fhsinchy/the-docker-handbook](https://github.com/fhsinchy/the-docker-handbook)) and the projects are hosted on <b><code>&nbsp;&nbsp;1392⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;327🍴</code></b> [fhsinchy/docker-handbook-projects](https://github.com/fhsinchy/docker-handbook-projects)) repository.

**Cheatsheets**

- <b><code>&nbsp;&nbsp;3937⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;485🍴</code></b> [eon01](https://github.com/eon01/DockerCheatSheet))
- <b><code>&nbsp;&nbsp;&nbsp;198⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19🍴</code></b> [dimonomid](https://github.com/dimonomid/docker-quick-ref)) (PDF)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;22⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [JensPiegsa](https://github.com/JensPiegsa/docker-cheat-sheet))
- <b><code>&nbsp;22529⭐</code></b> <b><code>&nbsp;&nbsp;4615🍴</code></b> [wsargent](https://github.com/wsargent/docker-cheat-sheet)) (Most popular)

## Where to Start (Windows)

- 🌎 [Docker on Windows behind a firewall](toedter.com/2015/05/11/docker-on-windows-behind-a-firewall/)
- 🌎 [Docker Reference Architecture: Modernizing Traditional .NET Framework Applications](docs.mirantis.com/containers/v3.0/dockeree-ref-arch/app-dev/modernize-dotnet-apps.html) - You will learn to identify the types of .NET Framework applications that are good candidates for containerization, the "lift-and-shift" approach to containerization.
- 🌎 [Docker with Microsoft SQL 2016 + ASP.NET](blog.alexellis.io/docker-does-sql2016-aspnet/) Demonstration running ASP.NET and SQL Server workloads in Docker
- 🌎 [Exploring ASP.NET Core with Docker in both Linux and Windows Containers](www.hanselman.com/blog/exploring-aspnet-core-with-docker-in-both-linux-and-windows-containers) Running ASP.NET Core apps in Linux and Windows containers, using [Docker for Windows][docker-for-windows]
- 🌎 [Running a Legacy ASP.NET App in a Windows Container](blog.sixeyed.com/dockerizing-nerd-dinner-part-1-running-a-legacy-asp-net-app-in-a-windows-container/) Steps for Dockerizing a legacy ASP.NET app and running as a Windows container
- 🌎 [Windows Containers and Docker: The 101](www.youtube.com/watch?v=N7SG2wEyQtM) - A 20-minute overview, using Docker to run PowerShell, ASP.NET Core and ASP.NET apps.
- 🌎 [Windows Containers Quick Start](learn.microsoft.com/en-us/virtualization/windowscontainers/about/) Overview of Windows containers, drilling down to Quick Starts for Windows 10 and Windows Server 2016

---

## Books & Tutorials

- <b><code>&nbsp;&nbsp;9904⭐</code></b> <b><code>&nbsp;&nbsp;2160🍴</code></b> [Cloud Native Landscape](https://github.com/cncf/landscape))
- 🌎 [Docker Blog](www.docker.com/blog/) - Regular updates about Docker, the community and tools.
- 🌎 [Docker Certification](intellipaat.com/docker-training-course/?US) - :yen: Will help you to will Learn Docker containerization, running Docker containers, Image creation, Dockerfile, Docker orchestration, security best practices, and more through hands-on projects and case studies and helps to clear Docker Certified Associate.
- 🌎 [Docker dev bookmarks](www.codever.dev/search?q=docker) - Use the tag 🌎 [docker](www.codever.dev/bookmarks/t/docker).
- 🌎 [Docker in Action, Second Edition](www.manning.com/books/docker-in-action-second-edition)
- 🌎 [Docker in Practice, Second Edition](www.manning.com/books/docker-in-practice-second-edition)
- 🌎 [Docker packaging guide for Python](pythonspeed.com/docker/) - A series of detailed articles on the specifics of Docker packaging for Python.
- 🌎 [Learn Docker in a Month of Lunches](www.manning.com/books/learn-docker-in-a-month-of-lunches)
- 🌎 [Learn Docker](coursesity.com/blog/best-docker-tutorials/) - Learn Docker - curated list of the top online docker tutorials and courses.
- 🌎 [Programming Community Curated Resources for learning Docker](hackr.io/tutorials/learn-docker)

## Awesome Lists

- <b><code>&nbsp;45447⭐</code></b> <b><code>&nbsp;&nbsp;8180🍴</code></b> [Awesome Compose](https://github.com/docker/awesome-compose)) - Docker Compose samples.
- <b><code>&nbsp;15977⭐</code></b> <b><code>&nbsp;&nbsp;2451🍴</code></b> [Awesome Kubernetes](https://github.com/ramitsurana/awesome-kubernetes))
- <b><code>&nbsp;&nbsp;2063⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;179🍴</code></b> [Awesome Linux Container](https://github.com/Friz-zy/awesome-linux-containers)) more general about container than this repo.
- <b><code>296528⭐</code></b> <b><code>&nbsp;13790🍴</code></b> [Awesome Selfhosted](https://github.com/awesome-selfhosted/awesome-selfhosted)) list of Free Software network services and web applications which can be hosted locally by running in a classical way (setup local web server and run applications from there) or in a Docker container.
- <b><code>&nbsp;34107⭐</code></b> <b><code>&nbsp;&nbsp;2034🍴</code></b> [Awesome Sysadmin](https://github.com/n1trux/awesome-sysadmin))
- <b><code>&nbsp;17053⭐</code></b> <b><code>&nbsp;&nbsp;1292🍴</code></b> [ToolsOfTheTrade](https://github.com/cjbarber/ToolsOfTheTrade)) a list of SaaS and On premise applications

## Demos and Examples

- 🌎 [An Annotated Docker Config for Frontend Web Development](nystudio107.com/blog/an-annotated-docker-config-for-frontend-web-development) A local development environment with Docker allows you to shrink-wrap the devops your project needs as config, making onboarding frictionless.
- <b><code>&nbsp;&nbsp;&nbsp;299⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;37🍴</code></b> [Local Docker DB](https://github.com/alexmacarthur/local-docker-db)) a list of docker-compose samples for a lot of databases
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;18🍴</code></b> [Webstack-micro](https://github.com/ferbs/webstack-micro)) Demo web app showing how Docker Compose might be used to set up an API Gateway, centralized authentication, background workers, and WebSockets as containerized services.

## Good Tips

- [Docker Caveats](http://docker-saigon.github.io/post/Docker-Caveats/) What You Should Know About Running Docker In Production (written 11 APRIL 2016).
- 🌎 [Docker Containers on the Desktop](blog.jessfraz.com/post/docker-containers-on-the-desktop/)
- 🌎 [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](www.flexera.com/blog/finops/)
- 🌎 [Don't Repeat Yourself with Anchors, Aliases and Extensions in Docker Compose Files](medium.com/@kinghuang/docker-compose-anchors-aliases-extensions-a1e4105d70bd)
- [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/)

## Raspberry Pi & ARM

- 🌎 [Docker Pirates ARMed with explosive stuff](blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Get Docker up and running on the RaspberryPi in three steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-three-steps))
- 🌎 [git push docker containers to linux devices](www.balena.io) Modern DevOps for IoT, leveraging git and Docker.
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-%28ARMv7%29-devices))

## Security Articles

- 🌎 [Bringing new security features to Docker](opensource.com/business/14/9/security-for-docker)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1🍴</code></b> [CVE Scanning Alpine images with Multi-stage builds in Docker 17.05](https://github.com/tomwillfixit/alpine-cvecheck))
- <b><code>&nbsp;&nbsp;&nbsp;606⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;79🍴</code></b> [Docker Secure Deployment Guidelines](https://github.com/AonCyberLabs/Docker-Secure-Deployment-Guidelines))
- 🌎 [Docker Security - Quick Reference](binarymist.io/publication/docker-security/)
- 🌎 [Docker Security: Are Your Containers Tightly Secured to the Ship? SlideShare](www.slideshare.net/slideshow/docker-security-are-your-containers-tightly-secured-to-the-ship/43834790)
- <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;?🍴</code></b> [How CVE's are handled on Offical Docker Images](https://github.com/docker-library/official-images/issues/1448))
- 🌎 [Lynis is an open source security auditing tool including Docker auditing](cisofy.com/lynis/)
- 🌎 [Security Best Practices for Building Docker Images](linux-audit.com/tags/docker/)
- 🌎 [Software Engineering Radio interview of Docker Security Team Lead (Diogo Mónica)](www.se-radio.net/2017/05/se-radio-episode-290-diogo-monica-on-docker-security/)
- 🌎 [Ten Docker Image Security Best Practices Cheat Sheet](snyk.io/blog/10-docker-image-security-best-practices/)
- 🌎 [Top ten most popular docker images each contain at least 30 vulnerabilities](snyk.io/blog/top-ten-most-popular-docker-images-each-contain-at-least-30-vulnerabilities/)
- 🌎 [Tuning Docker with the newest security enhancements](opensource.com/business/15/3/docker-security-tuning)
- 🌎 [10 best practices to containerize Node.js web applications with Docker](snyk.io/blog/10-best-practices-to-containerize-nodejs-web-applications-with-docker/)

## Videos

- 🌎 [Contributing to Docker by Andrew "Tianon" Page (InfoSiftr)](www.youtube.com/watch?v=1jwo8-1HYYg) (34:31)
- 🌎 [Deploying and scaling applications with Docker, Swarm, and a tiny bit of Python magic](www.youtube.com/watch?v=GpHMTR7P2Ms) (3:11:06)
- 🌎 [Docker and SELinux by Daniel Walsh from Red Hat](www.youtube.com/watch?v=zWGFqMuEHdw) (40:23)
- 🌎 [Docker Course](www.youtube.com/watch?v=UZpyvK6UGFo) (Spanish)
- 🌎 [Docker for Developers](www.youtube.com/watch?v=FdkNAjjO5yQ) (54:26)
- 🌎 [Docker from scratch](www.youtube.com/playlist?list=PLLhEJK7fQIxD-btrjrqdEfQHbkZnQrmqE) (1:22:01)
- 🌎 [Docker: How to Use Your Own Private Registry](www.youtube.com/watch?v=CAewZCBT4PI) (15:01)
- 🌎 [Docker in Production](www.youtube.com/watch?v=Glk5d5WP6MI) (36:05)
- 🌎 [Docker Primer to Docker Compose](www.youtube.com/watch?v=G-s2GXGAjTk) (1:56:45)
- 🌎 [Docker Registry from scratch](www.youtube.com/playlist?list=PLLhEJK7fQIxAz3d4Fj3edq7UcxEhdTCBm) (44:40)
- 🌎 [Docker Swarm from scratch](www.youtube.com/playlist?list=PLLhEJK7fQIxAY4gZd1Wl-GsLvg-e9Ap1e) (1:41:28)
- 🌎 [Extending Docker with Plugins](vimeo.com/110835013) (15:21)
- 🌎 [From Local Docker Development to Production Deployments](www.youtube.com/watch?v=7CZFpHUPqXw)
- 🌎 [Immutable Infrastructure with Docker and EC2 by Michael Bryzek (Gilt)](www.youtube.com/watch?v=GaHzdqFithc) (42:04)
- 🌎 [Introduction to Docker and containers](www.youtube.com/watch?v=ZVaRK10HBjo) (3:09:00)
- 🌎 [Logging on Docker: What You Need to Know](vimeo.com/123341629) (51:27)
- 🌎 [Performance Analysis of Docker - Jeremy Eder](www.youtube.com/watch?v=6f2E6PKYb0w) (1:36:58)
- 🌎 [Scalable Microservices with Kubernetes](www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) Free Udacity course
- 🌎 [State of containers: a debate with CoreOS, VMware and Google](www.youtube.com/watch?v=IiITP3yIRd8) (27:38)

## Communities and Meetups

### Brazilian

- 🌎 [Docker BR on Telegram](telegram.me/dockerbr)

### English

- 🌎 [Docker Community](www.docker.com/community/)
- 🌎 [Docker Events](www.docker.com/events/)
- 🌎 [Docker Online Meetup](www.meetup.com/en-AU/Docker-Online-Meetup/)
- 🌎 [Docker Reddit Community](www.reddit.com/r/docker/)

### Russian

- 🌎 [Docker Russian-speaking Community](t.me/docker_ru)

### Spanish

- 🌎 [Docker Tips](dockertips.com/)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/veggiemonk/awesome-docker.svg?variant=adaptive)](https://starchart.cc/veggiemonk/awesome-docker)

[calico]: https://github.com/projectcalico/calico
[coreos]: https://github.com/coreos
[distribution]: https://github.com/docker/distribution
[docker-for-windows]: https://docs.docker.com/desktop/setup/install/windows-install/
[editreadme]: https://github.com/correia-jpv/fucking-awesome-docker/edit/master/README.md
[kubernetes]: https://kubernetes.io
[nginxproxy]: https://github.com/nginx-proxy/nginx-proxy
[openshift]: https://okd.io/
[sindresorhus]: https://github.com/sindresorhus/awesome


## Source
<b><code>&nbsp;36124⭐</code></b> <b><code>&nbsp;&nbsp;3312🍴</code></b> [veggiemonk/awesome-docker](https://github.com/veggiemonk/awesome-docker))