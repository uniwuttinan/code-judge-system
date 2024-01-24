# Deploy using GCP Autopilot

Initialize the cluster

```sh
gcloud beta container --project "[PROJECT_NAME]" \
    clusters create-auto "[CLUSTER_NAME]" --region "[PROJECT_REGION]" \
    --release-channel "regular" \
    --network "projects/[PROJECT_NAME]/global/networks/default" \
    --subnetwork "projects/[PROJECT_NAME]/regions/[PROJECT_REGION]/subnetworks/default" \
    --cluster-ipv4-cidr "/17" \
    --binauthz-evaluation-mode=DISABLED
```

Access the cluster

```sh
gcloud container clusters get-credentials [CLUSTER_NAME] --region [PROJECT_REGION]--project [PROJECT_NAME]
```

## Kubernetes setup

Create namespace
```sh
kubectl create namespace cjs
```

Use namespace
```sh
kubectl config set-context --current --namespace=cjs
```

Create configmap
```sh
kubectl create configmap cjs-configmap --from-env-file .env
```

Modify configmap
```sh
# roll out the changes
kubectl rollout restart deployment your-deployment-name
```
