package config

import (
	"testing"
)

func TestResolveImages_Defaults(t *testing.T) {
	// Clear any env vars that might interfere
	t.Setenv(EnvValidatorImage, "")
	t.Setenv(EnvToolsImage, "")

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
	t.Setenv(EnvValidatorImage, "my-validator:v1")
	t.Setenv(EnvToolsImage, "my-tools:v2")

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
	t.Setenv(EnvValidatorImage, "")
	t.Setenv(EnvToolsImage, "custom-tools:latest")

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
