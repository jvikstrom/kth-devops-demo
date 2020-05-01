All resources required for the demo.

## Structure
* `docker`: Contains all docker files used for the demo.
* `sources`: Contains the source code for the GRPC server and client.
* * `proto`: Contains the protobuf definition and generated grpc stub.

## Prerequisites to run:
* A running minikube cluster ([Install instuctions](https://kubernetes.io/docs/tasks/tools/install-minikube/))
* Golang with minimum version 1.12
* Docker

## Prometheus

To get a nice CPU rate visualization of the "hello-server" pods, run this query: `rate(container_cpu_usage_seconds_total{container="hello-server-pod"}[5m])`.

## Running

* Start minikube.
* Run `./deploy.sh` this will deploy the client and server.
* Run `./connect-prometheus.sh` to connect to the prometheus server.

....
TODO: Add more instructions for adding linkerd as a side-car to the service.
## Screencast steps
* Show the load imbalance between the pods using prometheus.
* Install the linkerd CLI using: `curl -sL https://run.linkerd.io/install | sh`.
* Add linkerd to our path using: `export PATH=$PATH:$HOME/.linkerd2/bin`
* Verify linkerd is installed using: `linkerd version`
* Check and make sure kubernetes is configured correctly for linkerd: `linkerd check --pre`
* Install the linkerd control plane into the cluster: `linkerd install | kubectl apply -f -` (may take a while depending on internet connection speeds)
* Validate the installation using: `linkerd check`
* (show what deployments are deployed using `kubectl -n linkerd get deploy`)
* Show the linkerd dashboard using: `linkerd dashboard &`
* Check traffic we're generating using: `linkerd -n linkerd top deploy/linkerd-web`
* Inject linkerd into our deployment using:

