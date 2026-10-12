# Notiflex 실습 프로젝트

## 프로젝트 개요
Notiflex — B2B 알림 SaaS 플랫폼. GitAIOps 책의 실습을 GCP 환경에서 진행한다.
대화와 문서는 한국어로 작성한다.

## 기술 스택
- (애플리케이션: Go 표준 라이브러리만 사용, 외부 프레임워크 없음)
- 컨테이너: scratch 베이스 이미지
- 인프라: GKE Standard (Zonal)
- CI/CD: Github Actions + ArgoCD 사용 예정

## GCP 설정
- 프로젝트 ID: my-gitaiops-project-261010
- 리전: asia-northeast3
- 존: asia-northeast3-a
- Artifact Registry 경로 형식: `asia-northeast3-docker.pkg.dev/my-gitaiops-project-261010/notiflex`
- 프로젝트와 서울 리전·존 설정 완료 (2026-10-10).
- GKE 클러스터: notiflex-cluster (asia-northeast3-a), 노드 2개 Ready 확인
- 노드: e2-medium ×2, Spot VM, 부팅 디스크 각 30GB
- Kubernetes 버전: 1.35.8-gke.1380001
- 접속 컨텍스트: gke-sysnet4admin_book_gitaiops
- Gateway API: CHANNEL_STANDARD 활성화, GatewayClass 4개 Accepted=True 확인
- Artifact Registry 저장소 notiflex 생성 완료, 이미지 api:v0.1.0 푸시 (2026-10-12).

## 행동 규칙
1. 명령 실행 전 현재 상태를 확인한다. (`kubectl get`, `gcloud config list 등)
2. 파일 변경 전 기존 내용을 먼저 읽는다.
3. 에러 발생 시 원인을 분석하고 해결 방안을 제시한 뒤 진행한다.
4. 매니페스트 작성 시 네임스페이즈(`notiflex)를 명시한다.
5. **모든 `kubectl` 명령에 `--context gke-sysnet4admin_book_gitaiops`를 반드시 지정한다** (잘못된 클러스터 대상 실행 방지)
6. 리소스 생성/삭제 전에는 영향 범위를 먼저 설명한다.
7. 이미지 태그는 `latest`를 쓰지 않고 명시적 버전(`v0.1.0` 등)을 사용한다.
8. 토큰, 키 그리고 비밀번호는 코드/매니페스트에 하드코딩하지 않는다. (환경 변수, Github Secret, Secret Manager 사용)
