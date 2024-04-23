# Dionysia

> A live streaming platform based on open source software

## What is this?

This project is a live streaming platform based on open source software. The main goal is to create a platform that can be used by anyone to stream their content, be it a game, a podcast, a talk show, etc.

It plugs many open source projects together to create a seamless experience for video developers.

Among the features that are planned for this project are:
* Ingest streams from RTMP, SRT, RTSP, and other protocols;
* Transcode/encode streaming using customizable profiles;
* ABR (Adaptive Bitrate) streaming for different devices and network conditions;
* Distribute transcoding/encoding to many workers as needed;
* Output streams in HLS (DASH coming soon);
* API to manage inputs, workers, and metrics.

If you are interested in reading my ideas and thoughts about this project, you can read the [dev log](docs/DEV_LOG.md)

## How does it work?

The platform is composed of several components:

![System design](docs/static/architecture.png)

* **Media Server**: The server that will receive the stream that will be pulled by a worker.
* **Management API**: An API used to manage streams, check status, watch metrics, etc.
* **Workers**: A worker that will pull the stream from the media server, encode and transcode the video to enable playback on various devices and under different network conditions, and distribute it to viewers through a CDN.
* **Routing API**: An API used to route CDN requests to the right origins.

Let's see how the components interact in detail:

The **Management API** is used to manage the streams, workers, and metrics. It can be used to defined how the stream will be ingested, and how it will be transcoded.

A **Media Server** is useful to receive the content from streamers. It supports many ingest protocols like RTMP, RTSP, SRT and so on. Each protocol has its own characteristics. This Media Server can be deployed behind a load balancer to scale horizontally.

**Workers** are responsible for pulling the stream from the Media Server, transcoding the video, and saving all the data locally in the nodes. The workers can be deployed in a cluster to scale horizontally.

**Routing API** is used to route the CDN requests to the right origins. It can be used to distribute the requests to the right worker. The Routing API is a module under the Management API.

## Tech Stack

The platform is built using the following technologies:

* [Go](https://go.dev/) for all the services and workers
    * Easy to write concurrent code, and easy to deploy
* [mediamtx](https://github.com/bluenviron/mediamtx) for the media server
    * Supports a wide range of ingest protocols
* [gpac](https://gpac.io/) for the encoding, transcoding, and packaging
    * Awesome tool to handle complex video tasks like transcoding and packaging
* [asynq](https://github.com/hibiken/asynq) for distributing tasks to workers
    * Scalable, easy to use task queue for Go

## Development

### Running

All tasks are managed by a Justfile. You can see all the available tasks by running:

```sh
just -l
```

To run the platform, you need to have Docker and docker compose installed. Then, you can run the following command:

```sh
docker compose up
```

### Code Quality

To run the test suite:

```sh
just test
```

To run code linting:

```sh
just lint
```

## Contributing

Feel free to contribute to this project by opening issues or pull requests.

How could you contribute?

* Opening an issue to report a bug or suggest a new feature;
* Writing new features;
* Covering the code with tests;
* Improving the code quality;
* Improving the documentation.

If you want to contribute, please read the [CONTRIBUTING.md](CONTRIBUTING.md) file.

For more **advanced** contributions, please, create an issue to discuss the feature you want to implement. If it is a complex decision or a big feature, it is better to create an ADR (Architecture Decision Record) to write down the decision and the reasons behind it. Check the [adr](docs/adr) folder for more information.

## Roadmap

- [ ] Video and audio transcoding (profiles)
- [ ] Ingest authentication
- [ ] Playback authentication
- [ ] API authentication

## Why the name?

> Dionysia is a festival in ancient Greece in honor of Dionysus, the god of wine, fertility, and theater. It was a time of celebration, where people would gather to watch plays, dance, and drink wine.

Thank you @josethz00 for the name suggestion!
