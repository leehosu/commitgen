package prompt

import "fmt"

// GetSystemPrompt는 AI 모델에게 전달할 시스템 프롬프트를 생성합니다
func GetSystemPrompt(language, template string) string {
	if template == "conventional" {
		return getConventionalCommitsPrompt(language)
	}
	return getSimplePrompt(language)
}

// GetRegenerateSystemPrompt는 재생성 시 사용할 시스템 프롬프트를 생성합니다
func GetRegenerateSystemPrompt(language, template, previousMessage string) string {
	basePrompt := GetSystemPrompt(language, template)
	
	if language == "ko" {
		return fmt.Sprintf(`%s

중요: 이전에 다음과 같은 메시지를 생성했습니다:
"%s"

이번에는 동일한 변경사항에 대해 다른 관점이나 표현으로 새로운 커밋 메시지를 생성하세요.
- 다른 타입을 사용할 수 있다면 시도하세요
- 다른 단어나 표현을 사용하세요
- 강조하는 부분을 다르게 할 수 있습니다
- 절대 이전 메시지와 동일한 메시지를 생성하지 마세요`, basePrompt, previousMessage)
	}
	
	return fmt.Sprintf(`%s

IMPORTANT: Previously generated message was:
"%s"

Generate a NEW commit message with a different perspective or expression for the same changes:
- Try a different type if applicable
- Use different words or expressions
- Emphasize different aspects
- NEVER generate the same message as before`, basePrompt, previousMessage)
}

// getConventionalCommitsPrompt는 Conventional Commits 형식의 프롬프트를 반환합니다
func getConventionalCommitsPrompt(language string) string {
	if language == "ko" {
		return `당신은 Git 커밋 메시지 생성 전문가입니다. git diff를 분석하여 Conventional Commits 형식에 맞는 간결하고 의미있는 커밋 메시지를 생성하세요.

형식: <type>(<scope>): <subject>

타입:
- feat: 새로운 기능
- fix: 버그 수정
- docs: 문서 변경
- style: 코드 포맷팅 (기능 변경 없음)
- refactor: 리팩토링
- test: 테스트 추가/수정
- chore: 빌드, 설정 등 기타 변경
- perf: 성능 개선
- ci: CI 설정 변경
- build: 빌드 시스템 변경
- revert: 이전 커밋 되돌리기

규칙:
- subject는 소문자로 시작, 마침표 없음
- 명령형으로 작성 ("add" not "added" or "adds")
- 최대 72자
- 구체적이면서도 간결하게
- 한국어로 작성

커밋 메시지만 반환하고 다른 설명은 하지 마세요.`
	}

	return `You are a Git commit message generator. Analyze the git diff and generate a concise, meaningful commit message following Conventional Commits format.

Format: <type>(<scope>): <subject>

Types:
- feat: new feature
- fix: bug fix
- docs: documentation changes
- style: code formatting (no functional changes)
- refactor: code refactoring
- test: adding/modifying tests
- chore: build, config, and other changes
- perf: performance improvements
- ci: CI configuration changes
- build: build system changes
- revert: revert previous commit

Rules:
- Subject must be lowercase, no period at end
- Use imperative mood ("add" not "added" or "adds")
- Maximum 72 characters for subject
- Be specific but concise

Return only the commit message without any additional explanation.`
}

// getSimplePrompt는 간단한 형식의 프롬프트를 반환합니다
func getSimplePrompt(language string) string {
	if language == "ko" {
		return `당신은 Git 커밋 메시지 생성 전문가입니다. git diff를 분석하여 간결하고 의미있는 커밋 메시지를 한국어로 생성하세요.

규칙:
- 명령형으로 작성
- 최대 72자
- 구체적이면서도 간결하게
- 한국어로 작성

커밋 메시지만 반환하고 다른 설명은 하지 마세요.`
	}

	return `You are a Git commit message generator. Analyze the git diff and generate a concise, meaningful commit message.

Rules:
- Use imperative mood
- Maximum 72 characters
- Be specific but concise

Return only the commit message without any additional explanation.`
}

// GetUserPrompt는 git diff를 포함한 사용자 프롬프트를 생성합니다
func GetUserPrompt(diff string) string {
	return fmt.Sprintf("다음 git diff를 분석하여 커밋 메시지를 생성하세요:\n\n%s", diff)
}
