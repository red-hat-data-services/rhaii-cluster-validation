package config

import (
	"os"
	"testing"
)

func TestResolveImages_Defaults(t *testing.T) {
	// Clear any env vars that might interfere
	os.Unsetenv(EnvValidatorImage)
	os.Unsetenv(EnvToolsImage)

	validator, tools, err := ResolveImages()
	if err != nil {
		t.Fatalf("ResolveImages() returned error: %v", err)
	}

	if validator == "" {
		t.Error("validator image must not be empty")
	}
	if tools == "" {
		t.Error("tools image must not be empty")
	}
}

func TestResolveImages_EnvOverrideBoth(t *testing.T) {
	os.Setenv(EnvValidatorImage, "my-validator:v1")
	os.Setenv(EnvToolsImage, "my-tools:v2")
	defer os.Unsetenv(EnvValidatorImage)
	defer os.Unsetenv(EnvToolsImage)

	validator, tools, err := ResolveImages()
	if err != nil {
		t.Fatalf("ResolveImages() returned error: %v", err)
	}

	if validator != "my-validator:v1" {
		t.Errorf("validator = %q, want %q", validator, "my-validator:v1")
	}
	if tools != "my-tools:v2" {
		t.Errorf("tools = %q, want %q", tools, "my-tools:v2")
	}
}

func TestResolveImages_EnvOverridePartial(t *testing.T) {
	// Only override tools image
	os.Unsetenv(EnvValidatorImage)
	os.Setenv(EnvToolsImage, "custom-tools:latest")
	defer os.Unsetenv(EnvToolsImage)

	validator, tools, err := ResolveImages()
	if err != nil {
		t.Fatalf("ResolveImages() returned error: %v", err)
	}

	// Validator should use the embedded default
	if validator == "" {
		t.Error("validator image must not be empty (should use embedded default)")
	}
	if validator == "custom-tools:latest" {
		t.Error("validator image should not be affected by tools env var")
	}

	if tools != "custom-tools:latest" {
		t.Errorf("tools = %q, want %q", tools, "custom-tools:latest")
	}
}
