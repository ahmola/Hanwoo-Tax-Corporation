# 한우세무법인 동대문점 웹사이트 및 관리자 페이지

본 프로젝트는 GPL-3.0 라이센스를 따릅니다.

## Contract

- [SW종사자 표준도급계약서](https://spri.kr/posts/view/23132?code=information)

- 과업내용서(따로 양식 생성 필요)

- [표준비밀유지계약서](https://mss.go.kr/site/smba/ex/bbs/View.do?cbIdx=81&bcIdx=1031902)

- [개인정보처리위탁 계약서](https://share.google/JyGWqBi9g8lEebgE3)

- 검수확인서(따로 양식 생성 필요)

## Architecture
![img](architecture_diagram.png)

## Tech Stack

- Front-end: Next.js (App Router), TypeScript, Tailwind CSS

- Back-end: Go (Gin Framework)

- Database: MySQL

- Infra: Docker, Nginx, Kubernetes

- External API: Kakao Map API, Resend API

- Cloud: AWS EC2

P.S.
추후, Terraform 도입과 EC2 -> AWS EKS, MySQL -> RDS Aurora, Nginx -> S3 등으로 대체할 예정

## 인프라 상세 설정

- Storage: Nginx에서 resource를 persistan volume claim으로 저장하여 관리함

- network: Ingress-nginx controller 기반 라우팅

- External API: 카카오맵 API, Resend API

## 프로젝트 구조

admin-web, client-web, server/admin, server/client에 각각 Dockerfile 존재

    ├── admin-web/       # 관리자 페이지 (Next.js)
    ├── client-web/      # 고객 페이지 (Next.js)
    ├── server/          # Gin API 서버
    ├── infra/           
    │   ├── k8s/         # Manifests (Deployment, Service, Ingress)
    │   └── nginx/       # Static Resource Server 설정
    └── scripts/         # 쉘, SQL 등의 스크립트

## 설치 및 실행 방법(How to Run)

- Prerequisties: Docker, kubectl, kind 설치가 필요하며 kakao api key와 resend api key가 필요합니다.

- docker-compose 사용시 루트 디렉토리의 build_compose.sh을 사용하면 됩니다.

- kind 사용 시 infra/k8s에서 init_kind.sh을 실행

## Entity Model
![img](db_diagram.png)

BaseEntity는 MySQL에서 사용되는 Table이 아닌, ORM에서 기본 테이블로 정의하여 사용

## API

### admin

- notice

        POST /notices 

        DELETE /notices?id={id}


- document 

        POST /documents 

        DELETE /documents?id={id}

- contact 

        GET /contacts/all

        GET /contacts

### client

- notice

        GET /notices 

        GET /notices?id={id}

- document

        GET /documents 

        GET /documents?id={id}

- contact

        POST /contacts

---

## Services

### Front-end

브라우저를 통해 접근할 웹페이지입니다.

admin web은 관리자만 접근 가능한 제한된 웹페이지입니다.

client web은 Gabia에서 도메인을 구매하여 공개 웹페이지입니다.

#### Admin Web
관리자용 웹사이트입니다. Next.js로 개발되었으며, 자료실 문서 업로드 등 웹사이트 콘텐츠를 관리하는 기능을 제공합니다.

#### Client Web
고객에게 보여지는 웹사이트입니다. Next.js로 개발되었으며, 세무법인 소개, 서비스 안내, 자료실 등의 정보를 제공합니다. 카카오맵 API를 활용하여 오시는 길 안내 기능이 포함되어 있습니다.
상담 신청 폼에서 정보를 입력하면 이메일을 관리자에게 전달합니다.

---

### Back-end
Go의 Gin프레임워크 기반의 API 서버입니다. Front-end 애플리케이션에 필요한 데이터와 파일 업로드 처리 등의 비즈니스 로직을 담당합니다.

---

### Infra
인프라에 필요한 설정 파일들을 관리합니다.

k8s는 kind설정과 쿠버네티스 클러스터에서 파드에 적용될 manifest와 ingress 설정 등을 관리합니다.

nginx는 static resource server로 활용될 nginx 설정과 dockerfile을 관리합니다.

Container -> Docker Compose -> k8s Kind(현재 단계) -> k8s cluster 순으로 진행 중