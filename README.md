# Sample WebComponent WebAPI

This repository contains a sample web api for a microfrontend that is deployed as a containerized application to a Kubernetes cluster. The microfrontend is managed by the microfrontends-controller repository and displayed by the microfrontends-webui repository. The UI is built using Stencil, TypeScript.

To build and run the WebAPI, use the go build and ./<executable_name> commands respectively. To deploy the WebAPI to a Kubernetes cluster, use the kubectl apply command with the yaml file provided in the repository https://github.com/SevcikMichal/microfrontends-webcomponent-sample (more details there).

The repository also includes a Dockerfile for building a Docker image of the API, which can be pushed to a Docker registry and then deployed to a Kubernetes cluster.