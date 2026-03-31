package certify

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestCertify(t *testing.T) {
	loadDotEnvIfPresent(t, "certify/.env", ".env")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("failed to load certification config: %v", err)
	}

	if !cfg.Run {
		t.Skip("set RUN_CREC_CERTIFICATION=true in certify/.env or the environment to run deployed CREC certification")
	}

	RunCertificationScenario(t, cfg)
}

func loadDotEnvIfPresent(t *testing.T, paths ...string) {
	t.Helper()

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			t.Fatalf("failed opening %s: %v", path, err)
		}

		func() {
			defer func() {
				if err := file.Close(); err != nil {
					t.Fatalf("failed closing %s: %v", path, err)
				}
			}()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}

				if strings.HasPrefix(line, "export ") {
					line = strings.TrimSpace(strings.TrimPrefix(line, "export "))
				}

				key, value, found := strings.Cut(line, "=")
				if !found {
					continue
				}

				key = strings.TrimSpace(key)
				value = strings.TrimSpace(value)
				if key == "" {
					continue
				}

				if len(value) >= 2 {
					if (value[0] == '"' && value[len(value)-1] == '"') || (value[0] == '\'' && value[len(value)-1] == '\'') {
						value = value[1 : len(value)-1]
					}
				}

				if _, exists := os.LookupEnv(key); exists {
					continue
				}

				if err := os.Setenv(key, value); err != nil {
					t.Fatalf("failed setting %s from %s: %v", key, path, err)
				}
			}

			if err := scanner.Err(); err != nil {
				t.Fatalf("failed scanning %s: %v", path, err)
			}
		}()
	}
}
