# Deploy with helm chart


In this project, first, I deploy the MariaDB cluster with a helmchart. Second, create a helm chart for maxscale to use for the master/slave MariaDB cluster, and at the end, make a helm chart to deploy the usermanagement app and create GitLab CI/CD for it.



## MariaDB cluster

For it, I use [MariaDB packaged by Bitnami](https://artifacthub.io/packages/helm/bitnami/mariadb) and create two statefulsets for each master and slave through these steps:

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install master bitnami/MariaDB -f <master-values.yaml> --namespace=<namespace-name>
helm install slave bitnami/MariaDB -f <slave-values.yaml> --namespace=<namespace-name>

```
I add the configs I want to apply to the MariaDB cluster in *-values.yaml file.

### Challenges


- The PVC was pending . I had to create a storage class and set the storage class with `global.storageClass=<storageclassname>` ﻿. After that, the problem didn't solve. I found that the default PVC size was 8Gi, which was bigger than the  PV, so I changed  ﻿the primary.persistence.size value.
-  The pod is stuck in crashloopback off. I checked the logs, and  the error was : 
 `﻿mkdir: cannot create directory '/bitnami/MariaDB/data': Permission denied `
 The problem was solved whit these changes in this [link](https://github.com/bitnami/bitnami-docker-mariadb/issues/186#issuecomment-1128670750 ) and set these flags:
 primary.containerSecurityContext.enabled=true
primary.containerSecurityCon  text.runAsUser=0
primary.containerSecurityContext.runAsNonRoot=false

- This error: ﻿
`The MariaDB configuration file '/opt/bitnami/MariaDB/conf/my.cnf' is not writable. Configurations based on environment variables will not be applied for this file." ﻿it was because the pc host path was not writeable .`
 ﻿it was solved by changing the directory access mode on the node : ﻿
 ﻿`chmod 777 /mnt/data

## Maxscale 


I tried to use publish Kubernetes packages for maxscale, but all of them were for the Galera cluster for MariaDB.
So I create one for master/slave replication.
To create a new package, go through these steps :
`helm create maxscale.`
It has three templates : service, configmap, and deployment.
[There]() is its repository i create in github pages.

### How to deploy?

```
helm repo add maxscale https://mona-mp.github.io/maxscale-helmchart
helm upgrade -i maxscale https://mona-mp.github.io/maxscale-helmchart/maxscale-0.1.0.tgz --namespace=<namespace>`
```

### challenges

While config the maxscale, to connect to maxscale, this error occurred:

`Error: Could not connect to MaxScale`

It was about the conflict in maxscale config. I add a new config in maxscale. cnf and mounted it to /etc/maxscale.cnf.d/  but I think this config was appended to /etc/maxscale.cnf, so when I have the below part in /etc/maxscale.cnf.d/maxscale.cnf again, it makes conflict:

```SQL
[maxscale]
    threads=auto
    admin_enabled=false
```

So I deleted this part from my configmap, and it worked.

## Usermanagement-api

In this part, I deploy the user management-API, which I create in [this](https://github.com/mona-mp/mariadb-k8s) repo with helmchart.
I use this [doc](https://medium.com/geekculture/helm-create-helm-chart-7084666fab90) to create it.
The chart was created from ingress, service, deployment, and secret.
It exposes it on 80 port with the ingress.

### How to deploy?
```bash
helm repo add api https://mona-mp.github.io/api
helm upgrade -i api https://mona-mp.github.io/api/api-0.1.0.tgz --namespace=<namespace> `
```

### challenges

- for deploying ingress, when I use `kubectl apply -f ingress.yaml` command, I get the error:
```bash
# [Nginx Ingress: service "ingress-Nginx-controller-admission" not found](https://stackoverflow.com/questions/61365202/nginx-ingress-service-ingress-nginx-controller-admission-not-found)
```

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;    the solution is there:
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   i find the answer in this link.
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   https://stackoverflow.com/a/62044090/18421292
  
  - when I created the  Nginx ingress-controller, it was a deployment and design in one worker, so when the request sends to another worker which didn't have an ingress component, I failed, so I changed it to daemonset.



