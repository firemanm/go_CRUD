
kubectl apply -f storage.yaml
# check state
kubectl get pv,pvc,storageclass --all-namespaces
# helm install
sudo dnf install helm

# create catalog on master node - minikube in my case
minikube ssh
sudo mkdir -p /data/postgresql

# Set appropriate permissions (PostgreSQL typically runs as user 999)
sudo chown -R 999:999 /data/postgresql
sudo chmod -R 755 /data/postgresql

# apply PV spec
kubectl apply -f pv.yaml

# apply PV spec
kubectl apply -f pvc.yaml

# custom values file for PostgreSQL in pg_values.yaml
# install it on custom namespace
# kubectl create namespace pg-dbs
helm install my-postgresql oci://registry-1.docker.io/bitnamicharts/postgresql -f pg_values.yaml 
# Add -n pg-dbs to install in namespace

# manipulating
helm uninstall my-postgresql
helm upgrade my-postgresql

# check
kubectl get pods
kubectl get pvc
kubectl get svc
helm list

# check
# Get the password
kubectl get secret my-postgresql -o jsonpath='{.data.password}' | base64 -d

# Port forward for local access
kubectl port-forward svc/my-postgresql 5432:5432

# Connect using psql
PGPASSWORD=$(kubectl get secret my-postgresql -o jsonpath='{.data.password}' | base64 -d) \
psql -h localhost -U user00 -d mydb1

# build and push
docker build -t firemanm/go-crud-app .
docker push firemanm/go-crud-app 

# deploy app
kubectl apply -f configmap.yaml, -f secret.yaml, -f deployment.yaml, apply -f svc.yaml, -f ingress.yaml

# enable before
minikube addons enable ingress

minikube tunnel

# monitor logs
kubectl logs -f deployment/go-crud-app --all-containers=true --prefix --timestamps

# create user

























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