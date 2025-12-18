# Homebrew 설정 가이드

commitgen을 Homebrew로 배포하기 위한 설정 가이드입니다.

## 1단계: homebrew-tap 저장소 생성

### GitHub에서 새 저장소 생성

1. GitHub에서 새 저장소 생성
   - 저장소 이름: `homebrew-tap`
   - 설명: "Homebrew tap for leehosu's projects"
   - Public
   - README 추가

2. 로컬에 클론
   ```bash
   cd /Users/lake/git/leehosu
   git clone https://github.com/leehosu/homebrew-tap.git
   cd homebrew-tap
   ```

3. Formula 디렉토리 생성
   ```bash
   mkdir -p Formula
   ```

4. README 업데이트
   ```bash
   cat > README.md << 'EOF'
   # Homebrew Tap for leehosu

   ## 사용 방법

   ```bash
   # Tap 추가
   brew tap leehosu/tap

   # commitgen 설치
   brew install commitgen
   ```

   ## 사용 가능한 Formula

   - **commitgen**: AI-powered Git commit message generator

   ## 개발자 정보

   더 많은 프로젝트는 [@leehosu](https://github.com/leehosu)를 확인하세요.
   EOF
   ```

5. 커밋 및 푸시
   ```bash
   git add .
   git commit -m "feat: initialize homebrew tap"
   git push origin main
   ```

## 2단계: commitgen 저장소 변경사항 커밋

```bash
cd /Users/lake/git/leehosu/commitgen
git add .
git commit -m "feat(brew): add Homebrew tap support"
git push
```

## 3단계: 첫 릴리즈 생성

```bash
cd /Users/lake/git/leehosu/commitgen

# 태그 생성
git tag -a v0.1.0 -m "Initial release

Features:
- AI-powered commit message generation (OpenAI & Claude)
- Conventional Commits format support
- Interactive UI (approve/edit/regenerate/cancel)
- Cross-platform support (Linux, macOS, Windows)
- Configuration management with YAML and environment variables"

# 태그 푸시
git push origin v0.1.0
```

## 4단계: GitHub Actions 확인

태그를 푸시하면 GitHub Actions가 자동으로:

1. ✅ 멀티 플랫폼 바이너리 빌드
2. ✅ GitHub Release 생성
3. ✅ homebrew-tap 저장소에 Formula 자동 생성/업데이트
4. ✅ Checksums 계산 및 추가

**GitHub Actions 진행 상황 확인:**
- https://github.com/leehosu/commitgen/actions

**Release 확인:**
- https://github.com/leehosu/commitgen/releases

**Homebrew Formula 확인:**
- https://github.com/leehosu/homebrew-tap/blob/main/Formula/commitgen.rb

## 5단계: 설치 테스트

릴리즈가 완료되면 테스트:

```bash
# Tap 추가
brew tap leehosu/tap

# 설치
brew install commitgen

# 확인
commitgen version

# 테스트
cd ~/test-project
echo "test" > test.txt
git add test.txt
commitgen --dry-run
```

## 트러블슈팅

### Formula가 생성되지 않음

**확인사항:**
1. homebrew-tap 저장소가 Public인가?
2. GitHub Token 권한이 있는가?
3. GoReleaser 설정이 올바른가?

**해결방법:**
```bash
# GoReleaser 로컬 테스트 (릴리즈 없이)
goreleaser release --snapshot --clean
```

### "Error: No available formula with the name"

**원인:** Formula가 아직 생성되지 않음

**해결방법:**
1. GitHub Release가 완료될 때까지 대기 (1-2분)
2. homebrew-tap 저장소 확인
3. brew tap 업데이트: `brew update`

### Formula 수동 생성 (비상시)

```bash
cd /Users/lake/git/leehosu/homebrew-tap

cat > Formula/commitgen.rb << 'EOF'
class Commitgen < Formula
  desc "AI-powered Git commit message generator"
  homepage "https://github.com/leehosu/commitgen"
  version "0.1.0"
  
  if OS.mac? && Hardware::CPU.arm?
    url "https://github.com/leehosu/commitgen/releases/download/v0.1.0/commitgen_0.1.0_Darwin_arm64.tar.gz"
    sha256 "CHECKSUM_HERE"
  elsif OS.mac? && Hardware::CPU.intel?
    url "https://github.com/leehosu/commitgen/releases/download/v0.1.0/commitgen_0.1.0_Darwin_x86_64.tar.gz"
    sha256 "CHECKSUM_HERE"
  elsif OS.linux? && Hardware::CPU.intel?
    url "https://github.com/leehosu/commitgen/releases/download/v0.1.0/commitgen_0.1.0_Linux_x86_64.tar.gz"
    sha256 "CHECKSUM_HERE"
  elsif OS.linux? && Hardware::CPU.arm?
    url "https://github.com/leehosu/commitgen/releases/download/v0.1.0/commitgen_0.1.0_Linux_arm64.tar.gz"
    sha256 "CHECKSUM_HERE"
  end

  depends_on "git"

  def install
    bin.install "commitgen"
  end

  test do
    system "#{bin}/commitgen", "version"
  end
end
EOF

git add .
git commit -m "feat: add commitgen formula"
git push
```

## 업데이트 방법

새 버전 릴리즈 시:

```bash
cd /Users/lake/git/leehosu/commitgen

# 새 버전 태그
git tag -a v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0
```

→ GitHub Actions가 자동으로 homebrew-tap 저장소의 Formula를 업데이트합니다!

## 참고 자료

- [Homebrew Formula Cookbook](https://docs.brew.sh/Formula-Cookbook)
- [GoReleaser Homebrew Documentation](https://goreleaser.com/customization/homebrew/)
- [GitHub Actions for Go](https://docs.github.com/en/actions/guides/building-and-testing-go)

## 완료! 🎉

이제 사용자들이 다음 명령어로 설치할 수 있습니다:

```bash
brew tap leehosu/tap
brew install commitgen
```
