package services

import (
	"net/url"
	"strings"
)

type URLRiskAnalyzer struct{}

func NewURLRiskAnalyzer() *URLRiskAnalyzer {
	return &URLRiskAnalyzer{}
}

func (u *URLRiskAnalyzer) AnalyzeURL(urlString string) (int, string) {
	// Validate URL format
	if !strings.Contains(urlString, ".") {
		return 100, "Invalid URL"
	}

	_, err := url.Parse(urlString)
	if err != nil {
		return 100, "Invalid URL"
	}

	// Simple scoring logic
	score := 0
	riskLevel := "Safe"

	// Check for suspicious patterns
	suspiciousKeywords := []string{"phishing", "malware", "scam", "suspicious", "verify", "confirm", "urgent", "immediate"}
	for _, keyword := range suspiciousKeywords {
		if strings.Contains(strings.ToLower(urlString), keyword) {
			score += 20
		}
	}

	// Check URL length (very long URLs can be suspicious)
	if len(urlString) > 100 {
		score += 15
	}

	// Check for IP instead of domain
	if strings.Contains(urlString, "http://") && !strings.Contains(urlString, "https://") {
		score += 10
	}

	// Determine risk level
	if score >= 70 {
		riskLevel = "High Risk"
	} else if score >= 40 {
		riskLevel = "Suspicious"
	} else if score < 20 {
		riskLevel = "Safe"
	}

	return score, riskLevel
}
