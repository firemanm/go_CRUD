# Once per minikube installation
minikube addons enable ingress
minikube dashboard

# Create data catalog for Postgre on control plane
minikube ssh
sudo mkdir -p /data/postgresql
sudo chown -R 999:999 /data/postgresql
sudo chmod -R 755 /data/postgresql

# Prepare infra - not always needed
kubectl apply -f storage.yaml 

# Partition for PostgreSQL
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml

# Install it with parameters
helm install my-postgresql oci://registry-1.docker.io/bitnamicharts/postgresql -f pg_values.yaml 

# Deploy app etc
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
kubectl apply -f deployment.yaml
kubectl apply -f svc.yaml
kubectl apply -f ingress.yaml

# to view from localhost - not always needed
minikube tunnel

# view logs in separate shell
kubectl logs -f deployment/go-crud-app --all-containers=true --prefix --timestamps

# test it
curl -H "Host: arch.homework" http://$(minikube ip)/users
curl -H "Host: arch.homework" http://$(minikube ip)/health

# create data, test methods
curl -X POST http://$(minikube ip)/users \
  -H "Content-Type: application/json" \
  -H "Host: arch.homework" \
  -d '{"name": "Mikhail Pavlov", "email": "firemanm@gmail.com", "age": 52}'

curl -X GET http://$(minikube ip)/users/1 \
  -H "Content-Type: application/json" \
  -H "Host: arch.homework"

curl -H "Host: arch.homework" http://$(minikube ip)/users

curl -X GET http://$(minikube ip)/users/7 \
  -H "Content-Type: application/json" \
  -H "Host: arch.homework"

curl -X POST http://$(minikube ip)/users \
  -H "Content-Type: application/json" \
  -H "Host: arch.homework" \
  -d '{"name": "Peter Nilopv", "email": "pnilov@gmail.com", "age": 30}'

curl -X POST http://$(minikube ip)/users \
  -H "Content-Type: application/json" \
  -H "Host: arch.homework" \
  -d '{"name": "Check Duplicate", "email": "firemanm@gmail.com", "age": 0}'

curl -X PUT http://$(minikube ip)/users/1 \
  -H "Content-Type: application/json" \
  -H "Host: arch.homework" \
  -d '{"name": "Check Put", "email": "firemanmCP@gmail.com", "age": 0}'

curl -X DELETE http://$(minikube ip)/users/1 \
  -H "Content-Type: application/json" \
  -H "Host: arch.homework"

curl -H "Host: arch.homework" http://$(minikube ip)/users






















---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - name: postgres
        image: postgres:15-alpine
        ports:
        - containerPort: 5432
        env:
        - name: POSTGRES_DB
          value: "crudapp"
        - name: POSTGRES_USER
          valueFrom:
            secretKeyRef:
              name: go-crud-app-secret
              key: DB_USER
        - name: POSTGRES_PASSWORD
          valueFrom:
            secretKeyRef:
              name: go-crud-app-secret
              key: DB_PASSWORD
        volumeMounts:
        - name: postgres-storage
          mountPath: /var/lib/postgresql/data
      volumes:
      - name: postgres-storage
        emptyDir: {}




        ---
apiVersion: v1
kind: Service
metadata:
  name: postgres-service
spec:
  selector:
    app: postgres
  ports:
  - port: 5432
    targetPort: 5432