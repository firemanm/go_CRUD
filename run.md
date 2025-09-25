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

curl -H "Host: arch.homework" http://$(minikube ip)/users
curl -H "Host: arch.homework" http://$(minikube ip)/health

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
