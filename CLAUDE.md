# Notiflex 실습 프로젝트

## 프로젝트 개요
Notiflex — B2B 알림 SaaS 플랫폼. GitAIOps 책의 실습을 GCP 환경에서 진행한다.
대화와 문서는 한국어로 작성한다.

## 기술 스택과 진행 방향
- 애플리케이션: Go 표준 라이브러리
- 컨테이너: scratch 베이스 이미지
- 클라우드: Google Cloud (GCP)
- Kubernetes: GKE Standard (Zonal), Spot VM 사용 예정
- 이미지 레지스트리: Artifact Registry
- GitOps: ArgoCD 사용 예정

## GCP 설정
2026-10-10 로컬 `gcloud config list` 확인 결과:
- 프로젝트 ID: 미설정
- 리전: 미설정 (책의 서울 리전 예시: asia-northeast3)
- 존: 미설정
- Artifact Registry 경로 형식: `asia-northeast3-docker.pkg.dev/<PROJECT_ID>/notiflex`
- 실제 프로젝트와 리전, 존을 설정한 뒤 위 정보를 갱신한다.
- 실제 Artifact Registry 저장소와 GKE 클러스터의 생성 여부는 아직 확인하지 않았다.

## 행동 규칙
- 항상 확인 후 실행하며, 변경 전 현재 상태와 대상 프로젝트, 리전, 저장소를 확인한다.
- 모든 kubectl 명령에는 책의 규칙대로 `--context gke-sysnet4admin_book_gitaiops`를 지정한다.
- 실행 전에 해당 컨텍스트가 실제 실습용 GKE 클러스터를 가리키는지 확인한다. 컨텍스트가 없거나 대상이 다르면 실행하지 않고 설정을 먼저 확인한다.
- 인증 정보와 토큰을 Git에 저장하지 않는다.
- 확인한 구현과 예정 작업을 구분하여 기록한다.
- 실습 완료 후 JOURNEY.md에 진행 상태를 기록한다.
