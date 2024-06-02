package file_manager

import (
	"os"
	"strings"
	"testing"
)

func TestIsUrl_Positive(t *testing.T) {
	url := "https://www.example.com"
	result := IsUrl(url)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

func TestIsUrl_Negative(t *testing.T) {
	url := "not a valid url"
	result := IsUrl(url)
	if result {
		t.Errorf("Expected false, got true")
	}
}

func TestReadUrl_Positive(t *testing.T) {
	fileName := "test.txt"
	urls := []string{"https://www.example1.com", "https://www.example2.com"}
	file, _ := os.Create(fileName)
	defer file.Close()
	file.WriteString(strings.Join(urls, "\r\n"))

	result := ReadUrl(fileName)

	if len(result) != len(urls) {
		t.Errorf("Expected %d urls, got %d", len(urls), len(result))
	}

	for i, url := range result {
		if url != urls[i] {
			t.Errorf("Expected %s, got %s", urls[i], url)
		}
	}
	os.Remove(fileName)
}

func TestReadUrl_Negative(t *testing.T) {
	fileName := "non_existent_file.txt"

	result := ReadUrl(fileName)

	if len(result) != 0 {
		t.Errorf("Expected 0 urls, got %d", len(result))
	}
}

func TestSaveResult(t *testing.T) {
	// Positive test case
	respResult := map[string]string{"key": "value"}
	format := "JSON"
	SaveResult(respResult, format)
	// Add assertion here to check if file "results.json" is created

	// Negative test case
	// Test for unsupported format
	respResult = map[string]string{"key": "value"}
	format = "XML"
	SaveResult(respResult, format)
	// Add assertion here to check if log message is printed
}
