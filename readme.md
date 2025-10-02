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

# next lab - prometheus metrics
# install prometheus

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/prometheus
kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-np

# check it running
kubectl get pods -l app.kubernetes.io/instance=prometheus

# open web
minikube service prometheus-server-np

# Install Grafana
helm repo add grafana https://grafana.github.io/helm-charts
helm install grafana grafana/grafana
kubectl expose service grafana --type=NodePort --target-port=3000 --name=grafana-np
# retreive grafana password (user: admin)
kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo 

# open it
minikube service grafana-np

# add app to scraping (modify deploy file)



# Configure Prometheus Datasource (Connections > Datasources)
# The URL for our Prometheus instance is the name of the service http://prometheus-server:80.

# Kubernetes Dashboard bootstrap
# Create (+) > Import section to Import via grafana.com and we get 15661 (K8S Dashboard)






















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