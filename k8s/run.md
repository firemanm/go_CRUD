minikube addons enable ingress
minikube dashboard

minikube ssh
sudo mkdir -p /data/postgresql
sudo chown -R 999:999 /data/postgresql
sudo chmod -R 755 /data/postgresql

kubectl apply -f storage.yaml

kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml


helm install my-postgresql oci://registry-1.docker.io/bitnamicharts/postgresql -f pg_values.yaml 
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
kubectl apply -f deployment.yaml
kubectl apply -f svc.yaml
kubectl apply -f ingress.yaml

minikube tunnel

kubectl logs -f deployment/go-crud-app --all-containers=true --prefix --timestamps
