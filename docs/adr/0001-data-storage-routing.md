# Data Storage and Routing

## Status

Accepted

## Context

Currently, all encoded files are stored locally in a directory on each worker node. This approach suits our live streaming requirements due to its low latency and direct access. However, the system lacks a defined method for efficiently routing data requests to the appropriate worker, which is crucial for targeting specific live streams.

## Decision

To address this, we propose the following:
* Maintain local storage of data in each worker to leverage quick access and response times.
* Implement a monitoring and routing routine that starts with each new transcoding session.

This routine will:
* Periodically verify (every 5 seconds) if the worker is correctly transcoding the stream.
* Notify a central API with status updates and routing metadata necessary for efficiently directing incoming requests.
* Every key (stream) lasts for 30 seconds, after which it is deleted from the central API. This ensures the system always has the most up-to-date information.

## Consequences

* By avoiding the use of services like S3 for data storage, we anticipate significant cost savings without sacrificing performance for low latency streaming.
* Without S3, we need to store which worker is responsible for each stream in a central key-value database to ensure efficient routing keeping resources usage low - probably Redis.
