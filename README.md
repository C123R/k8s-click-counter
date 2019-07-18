# **k8s-click-counter**

K8s-click-counter is demo application, can be used for the Kubernetes workshops.

This demo app runs across two containers:

- [backend](https://github.com/C123R/k8s-click-counter/blob/master/backend) - a Go REST API which serves the click's count.
- [fronend](https://github.com/C123R/k8s-click-counter/blob/master/frontend) - a simple Go web application calls the backend REST API /api/counter

## **Running**

### **Docker**

- click-counter-backend

```sh
cd backend

docker build -t click-counter-backend .

docker run -it -p 9001:9001 click-counter-backend
```

- click-counter-frontend

```sh
cd frontend

docker build -t click-counter-frontend .

docker run -it -p 9000:9000 --env BACKEND_SERVICE=http://localhost:9001/api/counter click-counter-frontend
```

You can access the frontend on: http://localhost:9000

### **Kubernetes**

Deploy it on Kubernetes cluster:

```sh
kubectl delete -f https://raw.githubusercontent.com/C123R/k8s-click-counter/master/k8s-click-counter.yaml -n $(whoami)

# it will create a namespace with your userid(whoami)
```

Since we are creating service as cluster type, we cant access it from outside our cluster. Still you can do it using `kubectl port-forward`.

```sh
kubectl port-forward service/click-counter-frontend-svc :80 -n $(whoami)
```
