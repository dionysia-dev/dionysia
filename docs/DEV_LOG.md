# Dev Log

## 2024-04-03

I've started working on the Management API code, which will be used to manage streams, check their status, monitor metrics, and more.

Think of a stream as a TV channel. Users can ingest video using tools like OBS and then send it to a Media Server.

Today, I've pushed some initial code and will continue to develop the documentation and the API itself.

In the coming days, I'll write a notification system. This system will alert workers to pull a stream from the Media Server for encoding or transcoding. We'll leverage a queued messaging system to distribute these tasks across a cluster of machines.

## 2024-03-31

I've been trying mediamtx while working on the design of our API.

This API will be used to manage the streams:

* A user creates a new stream and receives a stream key (UUID);
* This UUID will be used to push the stream to the server, for example rtmp://streaming-platform/UUID;
* We can use *onReady* mediamtx event to POST our API a user has just started streaming;
* And then we will enqueue a task do encode the stream and generate the ABR and HLS streams.

## 2024-03-29

I've been thinking about the system design of the platform.

We need a few components to make it work:

* **Streaming server**: A server we can use to push our streaming. [mediamtx](https://github.com/bluenviron/mediamtx) seems to be a good option because it supports a wide range of protocols.
* **Encoding**: We need to encode the video to different bitrates and resolutions.[GPAC](https://gpac.io/) seems to be a good option because it supports a wide range of codecs and formats.
* **Packaging**: We need to package the encoded video to be streamed in playout formats like HLS and DASH. [GPAC](https://gpac.io/) again here. It seems to be a good option because it supports a wide range of formats.

We are going to start with a simple setup covering these components

Our goal is to have a POC with encoding, ABR and packaging working by the end of April.

## 2024-03-27

I've been pretty busy these days. I've been studying [GPAC](https://gpac.io/) and ffmpeg to see how we can use them to be our encoder and packager.

For now, I've been using a local server to stream RTMP video and pull it with GPAC to generate ABR and HLS streams.

## 2024-02-16

Welcome the first entry of this dev log. This is a place where we will keep track of the development of our project, documenting our progress and all the ideas and problems we encounter along the way.

*Why build a streaming platform?*

We believe that the available platforms don't let creators be free. We want to build a platform that don't lock creators into a single platform, or don't take a big cut of their earnings. We want to build a platform that is open and less expensive.

Fork the project if you disagree with any of our decisions.
