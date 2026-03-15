# root 디렉터리에서 path 시작
# 1. kind-config 작성 필요
kind create cluster --config ./infra/k8s/kind-config.yaml

# 2. nginx-ingress를 kind 내에 설치
sh ./infra/k8s/ingress_download.sh

# 3. 컨트롤러가 완전히 뜰 때까지 대기 (Running 상태 확인)
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

# 4. 클러스터 내 도커 이미지 전송
# kind load docker-image <이미지명>:<태그> --name <클러스터명>
kind load docker-image hanwoo-admin-service:latest --name kind && echo "admin-service" || (echo "failed" && exit 1)
sleep 0.5

kind load docker-image hanwoo-client-service:latest --name kind && echo "client-service" || (echo "failed" && exit 1)
sleep 0.5

kind load docker-image hanwoo-client-web:latest --name kind && echo "client-web" || (echo "failed" && exit 1)
sleep 0.5

kind load docker-image hanwoo-admin-web:latest --name kind && echo "admin-web" || (echo "failed" && exit 1)
sleep 0.5

kind load docker-image nginx:alpine --name kind && echo "resource-server" || (echo "failed" && exit 1)
sleep 0.5

# kind cluster 내부에서 직접 이미지 pull
docker exec -it kind-control-plane crictl pull mysql:8.2.0

# 5. Kompose로 Manifest 작성
kompose convert -f docker-compose.yml
# imagePullPolicy 설정 추가
sed -i '/image: /a \ 
imagePullPolicy: Never' ./infra/k8s/manifests/*-deployment.yaml

# 6. kind cluster에 manifests 적용
kubectl apply -f ./infra/k8s/manifests

# 7. 상태 확인
kubectl get pods -w