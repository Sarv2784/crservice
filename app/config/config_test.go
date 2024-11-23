package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tempDir := t.TempDir()
	validConfig := `server:
  host: "127.0.0.1"
  port: 8080
grpc:
  host: "127.0.0.1"
  port: 50051
database:
  host: "localhost"
  port: 5432
  username: "admin"
  password: "password"`
	invalidConfig := `server:
						  host: "127.0.0.1"
						  port: "not-a-number"  # Invalid port value
						`
	validFilePath := tempDir + "/config.yaml"
	err := os.WriteFile(validFilePath, []byte(validConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to write valid config file: %v", err)
	}

	invalidFilePath := tempDir + "/invalid_config.yaml"
	err = os.WriteFile(invalidFilePath, []byte(invalidConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid config file: %v", err)
	}

	// Test cases
	tests := []struct {
		name    string
		path    string
		want    Config
		wantErr bool
	}{
		{
			name: "Valid config file",
			path: tempDir,
			want: Config{
				App: App{
					Name:    "TestApp",
					Version: "1.0.0",
				},
				Server: HttpServer{
					Port: 8080,
					Host: "127.0.0.1",
				},
				GRPCServer: GRPCServer{
					Host: "127.0.0.1",
					Port: 50051,
				},
				Database: Database{
					User:     "admin",
					Password: "password",
					Name:     "testdb",
					Host:     "localhost",
					Port:     5432,
				},
				MarineForecast: MarineForecast{
					Url: "http://example.com/marine",
				},
			},
			wantErr: false,
		},
		{
			name:    "Invalid config file",
			path:    invalidFilePath,
			want:    Config{},
			wantErr: true,
		},
		{
			name:    "Missing config file",
			path:    tempDir + "/nonexistent",
			want:    Config{},
			wantErr: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
