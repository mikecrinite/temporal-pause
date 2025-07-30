# temporal-pause

This project gives the most barebones possible example of temporarily pausing a sample workflow while some async task is completed in the background. For example, while waiting for a money transfer, you may need to get a manual approval. Once the manual approval is received, a signal can be sent to the service to complete the rest of the transfer tasks

I couldn't really find a good example of this, but I needed to learn how it works for my job, so might as well make a quick public repo out here

## Basic stuff, this is mostly for me cuz I'm new to temporal

### Start the temporal server

`temporal server start-dev`

- The Temporal Service will be available on localhost:7233.
- The Temporal Web UI will be available at http://localhost:8233.

### Start the workflow execution (hopefully that's the correct wording)

`go run start/main.go`

You should see a workflow pop up on the Temporal Web UI. You can open it for some details on the workflow execution. Nothing is happening yet, we don't have a worker.

### Start the worker

` go run worker/main.go`

Now the worker will poll the queue, see the workflow, run the activities, and you'll see the pause/resume steps happen. Right now, the workflow doesn't wait for the pause step to finish (obviously it can't in this scenario, because it will finish it itself. In real life, another service would resume the activity)
