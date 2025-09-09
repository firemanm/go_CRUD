kubectl apply -f storage.yaml
# check state
kubectl get pv,pvc,storageclass --all-namespaces

# create catalog on master node - minikube in my case
minikube ssh

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