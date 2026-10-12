# Notiflex 실습 진행 기록

## 진행 현황
- ✅ 2.4 GitHub 저장소 생성 및 초기 구조 구성 — 2026-10-10
  - 비공개 저장소: https://github.com/hyochangsung/notiflex-platform
  - CLAUDE.md와 AGENTS.md에 GCP 진행 방향 기록
  - app/, k8s/smb/, .github/workflows/ 초기 구조 생성
  - 실제 GCP 프로젝트·리전·존 정보를 문서에 반영 완료

- ✅ 2.5 GKE 클러스터 생성 및 검증 — 2026-10-10
  - 노드 2개 Ready, e2-medium, Spot VM, 디스크 각 30GB
  - Gateway API 활성화 및 GatewayClass 4개 Accepted=True
  - 컨텍스트 gke-sysnet4admin_book_gitaiops 연결 검증 완료

- ✅ 2.6 Notiflex 앱 생성·빌드·배포 — 2026-10-12
  - Go 1.25 표준 라이브러리 API (`GET /health`, `GET /id`), scratch 멀티스테이지 이미지
  - Artifact Registry `notiflex/api:v0.1.0` 푸시 (Cloud Build)
  - notiflex 네임스페이스, Deployment(replicas 2), Service(ClusterIP 80→8080) 배포
  - Pod 2개 Running, `/health` 200 OK, `/id` 순차 ID 확인

## 현재 버전
- Notiflex 이미지: `v0.1.0`
- Go: `1.25` (빌드 이미지 `golang:1.25-alpine`)

## 현재 상태
- 클라우드 선택: GCP
- 프로젝트 ID: my-gitaiops-project-261010
- 리전: asia-northeast3, 존: asia-northeast3-a
- 프로젝트 생성 및 결제 활성화 완료 (2026-10-10)
- 애플리케이션과 Kubernetes 매니페스트(k8s/smb/): 작성 및 배포 완료 (2.6). CI: 아직 작성하지 않음
- 2.5 GKE 클러스터 생성 완료, 노드 2개 Ready 및 Spot 확인; GatewayClass 4개 Accepted=True 확인
- 클러스터: notiflex-cluster, Kubernetes 1.35.8-gke.1380001 (2026-10-12 확인; 기존 기록 1225000에서 변경됨)
- 노드: e2-medium ×2, Spot, 디스크 각 30GB
- 컨텍스트: gke-sysnet4admin_book_gitaiops
- 로컬 kubectl: 1.37.1 (서버 1.35와 마이너 버전 차이 2; 후속 운영 전 호환 버전 사용 필요)
- 다른 단계의 완료 여부는 별도 확인 필요

## 변경 이력
- 2026-10-10: 초기 문서에 AWS 기준을 잘못 적용한 내용을 사용자 지시에 따라 GCP/GKE/Artifact Registry 기준으로 수정했다.

- 2026-10-10: 요청한 프로젝트 ID my-gitaiops-project가 이미 사용 중이므로 이름은 유지하고 ID my-gitaiops-project-261010으로 생성했다.

- 2026-10-12: 2.6 빌드 중 Cloud Build API를 새로 활성화했고, 활성화 직후 `gcloud builds submit`이 403 PERMISSION_DENIED로 실패했다. 계정은 프로젝트 owner였으며 약 1~2분 뒤 같은 명령이 성공했다(권한 반영 지연으로 추정, 원인 미확정).
