# Notiflex 실습 프로젝트

## 프로젝트 개요
Notiflex — B2B 알림 SaaS 플랫폼. GitAIOps 책의 실습을 AWS 환경에서 진행한다.
대화와 문서는 한국어로 작성한다.

## 기술 스택과 진행 방향
- 애플리케이션: Go 표준 라이브러리
- 컨테이너: scratch 베이스 이미지
- 클라우드: AWS (책의 GCP/GKE 명령은 그대로 실행하지 않고 AWS에 맞게 변환)
- Kubernetes: EKS 사용 예정, 아직 생성 여부를 확인하지 않음
- 이미지 레지스트리: Amazon ECR
- GitOps: ArgoCD 사용 예정

## AWS 설정
- 로컬 AWS CLI 기본 리전: ap-northeast-2 (서울), 2026-10-10 확인
- 계정 ID와 역할: 작업 직전 `aws sts get-caller-identity`로 확인
- ECR 레지스트리 형식: `<AWS_ACCOUNT_ID>.dkr.ecr.ap-northeast-2.amazonaws.com`
- ECR 저장소 이름과 실제 이미지 경로: 아직 확인하지 않음
- 이 단계에서는 AWS 클라우드 리소스를 생성하지 않았다.

## 행동 규칙
- 변경 전 현재 상태와 대상 계정, 리전, 저장소를 확인한다.
- 모든 kubectl 명령에는 실제로 검증한 EKS 컨텍스트를 명시한다.
- 책의 GKE 컨텍스트를 AWS 실습에 사용하지 않는다.
- 인증 정보와 토큰을 Git에 저장하지 않는다.
- 확인한 구현과 예정 작업을 구분하여 기록한다.
- 실습 완료 후 JOURNEY.md에 진행 상태를 기록한다.
