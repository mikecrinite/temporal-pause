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

`go run worker/main.go`

Now the worker will poll the queue, see the workflow, and start running the activities. The "DoWork" activity just has a 5s sleep right now to simulate something happening. This isn't really that important, you can remove this or replace it with whatever you need. 

Once it finishes waiting, it will move on to the Pause step which should wait indefinitely.

### Resume the workflow by completing the activity

`go run start/resume.go`

This will read the taskToken from the `./tmp/tasktoken` file and dial up a client to call `client.CompleteActivity()`. You should be able to see the activity finish immediately if you have the UI up, and the workflow will be marked as "Completed"
