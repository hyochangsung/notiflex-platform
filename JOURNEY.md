# Notiflex 실습 진행 기록

## 진행 현황
- ✅ 2.4 GitHub 저장소 생성 및 초기 구조 구성 — 2026-10-10
  - 비공개 저장소: https://github.com/hyochangsung/notiflex-platform
  - CLAUDE.md와 AGENTS.md에 GCP 진행 방향 기록
  - app/, k8s/smb/, .github/workflows/ 초기 구조 생성
  - 가이드의 GCP 프로젝트·리전·존 정보 기록은 설정 후 완료 필요

## 현재 상태
- 클라우드 선택: GCP
- 로컬 gcloud 프로젝트 ID, 리전, 존: 미설정 (2026-10-10 확인)
- 애플리케이션, CI, Kubernetes 매니페스트: 아직 작성하지 않음
- 이번 저장소 구성 단계에서 클라우드 리소스 생성 및 배포는 실행하지 않음
- 다른 단계의 완료 여부는 별도 확인 필요

## 변경 이력
- 2026-10-10: 초기 문서에 AWS 기준을 잘못 적용한 내용을 사용자 지시에 따라 GCP/GKE/Artifact Registry 기준으로 수정했다.
