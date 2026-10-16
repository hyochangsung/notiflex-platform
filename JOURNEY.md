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

- ✅ 2.7 GitHub에 첫 커밋 — 2026-10-12
  - 커밋 df6dd7c `Initial commit: Notiflex app, K8s manifests, CLAUDE.md, JOURNEY.md`
  - origin/main에 푸시 (https://github.com/hyochangsung/notiflex-platform)

- ✅ 3.2 ArgoCD 설치 및 GitOps 연결 — 2026-10-16
  - ArgoCD v3.5.5 (stable manifest, server-side apply), argocd 네임스페이스 Pod 7개 Running
  - Application `notiflex-smb` (`argocd/notiflex-smb.yaml`): `k8s/smb` → notiflex 네임스페이스, automated(prune, selfHeal)
  - Sync Status: Synced, Health: Healthy (revision 817d8bb)
  - UI 포트포워딩 https://localhost:8443 응답 200 확인
  - 저장소 연결: Settings > Repositories에 공개 저장소 등록(인증 없음), 연결 상태 Successful. Application이 origin/main 최신 커밋(817d8bb)과 동기화됨을 확인

## 도구 선택 기록
- **GitOps(CD) 도구: ArgoCD** — 검토한 대안: Flux, Jenkins X, Spinnaker
  - 선택 이유: Self-Heal로 kubectl 직접 수정을 Git 상태로 되돌림, 메모리 약 500MB·CPU 약 500m로 e2-medium 2대에 수용 가능, Web UI로 배포 상태를 눈으로 확인, Application으로 `k8s/smb` 경로를 배포 단위로 지정
  - Flux는 가볍지만(약 100MB) UI가 없어 학습 단계에서는 차선으로 판단

## 현재 버전
- Notiflex 이미지: `v0.1.0`
- Go: `1.25` (빌드 이미지 `golang:1.25-alpine`)
- ArgoCD: `v3.5.5` (quay.io/argoproj/argocd:v3.5.5, 2026-10-16 설치. 책 검증 버전은 3.3.6)

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
- 로컬 argocd CLI: v3.5.5 (Homebrew, 2026-10-16 설치; 서버 v3.5.5와 동일)
- 다른 단계의 완료 여부는 별도 확인 필요

## 변경 이력
- 2026-10-10: 초기 문서에 AWS 기준을 잘못 적용한 내용을 사용자 지시에 따라 GCP/GKE/Artifact Registry 기준으로 수정했다.

- 2026-10-10: 요청한 프로젝트 ID my-gitaiops-project가 이미 사용 중이므로 이름은 유지하고 ID my-gitaiops-project-261010으로 생성했다.

- 2026-10-12: 2.6 빌드 중 Cloud Build API를 새로 활성화했고, 활성화 직후 `gcloud builds submit`이 403 PERMISSION_DENIED로 실패했다. 계정은 프로젝트 owner였으며 약 1~2분 뒤 같은 명령이 성공했다(권한 반영 지연으로 추정, 원인 미확정).

- 2026-10-16: 3.2에서 Application 생성 직후 `Sync: Unknown` (`authentication required: Repository not found`). 저장소가 비공개라 ArgoCD가 읽지 못한 것이 원인. 토큰 등록 대신 사용자가 저장소를 **공개(PUBLIC)로 전환**했고, hard refresh 후 Synced/Healthy 확인. 이후 저장소에 토큰·비밀번호·kubeconfig를 올리지 않도록 특히 주의.
- 2026-10-16: 가이드 트러블슈팅의 NetworkPolicy egress 차단 문제는 해당 없음을 확인. v3.5.5의 repo-server 정책은 Ingress만 제한하고, 클러스터의 NetworkPolicy 적용 기능도 비활성(disabled)이라 정책을 삭제하지 않았다.
- 2026-10-16: 전체 rollout restart 후 기본 AppProject `default` 유실 문제는 발생하지 않았다. restart 중 실행 중이던 port-forward가 끊겨 다시 연결했다.
- 2026-10-16: ArgoCD admin 비밀번호를 UI에서 변경(07:21 UTC)한 것을 확인한 뒤 `argocd-initial-admin-secret`을 삭제했다. 초기 비밀번호가 대화 기록에 노출됐기 때문이다.
