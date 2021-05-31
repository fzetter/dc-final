# User Guide

###### RUN MAIN PROJECT

  The project runs on port 8080 per defect, so it's only needed to run the server as usual.
  This will automatically start all the nodes: API, Controller and Scheduler.

    ```bash
    go run main.go
    ```

###### RUN WORKERS

    ```bash
    go run main.go --worker-name {{insert-worker-name}}
    ```
