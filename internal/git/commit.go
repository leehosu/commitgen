package git

import (
	"bytes"
	"os"
	"os/exec"
)

// CommitOptions는 커밋 옵션을 담는 구조체입니다
type CommitOptions struct {
	NoVerify    bool
	GPGSign     bool  // true면 -S 플래그 추가
	NoGPGSign   bool  // true면 --no-gpg-sign 플래그 추가
}

// Commit은 주어진 메시지로 커밋을 생성합니다
func Commit(message string, opts CommitOptions) error {
	args := []string{"commit", "-m", message}
	
	if opts.NoVerify {
		args = append(args, "--no-verify")
	}
	
	// GPG 서명 옵션 처리
	if opts.GPGSign {
		args = append(args, "-S")
	} else if opts.NoGPGSign {
		args = append(args, "--no-gpg-sign")
	}
	
	cmd := exec.Command("git", args...)
	
	// TTY 연결로 GPG pinentry 지원
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

// CommitWithFile은 파일에서 커밋 메시지를 읽어 커밋합니다 (멀티라인 메시지용)
func CommitWithFile(messageFile string, opts CommitOptions) error {
	args := []string{"commit", "-F", messageFile}
	
	if opts.NoVerify {
		args = append(args, "--no-verify")
	}
	
	// GPG 서명 옵션 처리
	if opts.GPGSign {
		args = append(args, "-S")
	} else if opts.NoGPGSign {
		args = append(args, "--no-gpg-sign")
	}
	
	cmd := exec.Command("git", args...)
	
	// TTY 연결로 GPG pinentry 지원
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

// GetLastCommitMessage는 마지막 커밋 메시지를 반환합니다
func GetLastCommitMessage() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%B")
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		return "", err
	}
	
	return stdout.String(), nil
}
