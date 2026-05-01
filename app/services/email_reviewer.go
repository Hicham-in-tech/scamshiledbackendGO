package services

import (
	"strings"
)

type EmailSecurityReviewer struct{}

func NewEmailSecurityReviewer() *EmailSecurityReviewer {
	return &EmailSecurityReviewer{}
}

func (e *EmailSecurityReviewer) ReviewEmail(emailText string) (int, string) {
	score := 0
	riskLevel := "Safe"

	emailLower := strings.ToLower(emailText)

	// Check for phishing indicators
	phishingKeywords := []string{
		"verify your account",
		"confirm your identity",
		"urgent action required",
		"click here",
		"update payment",
		"confirm password",
		"suspicious activity",
		"unauthorized access",
		"urgent",
		"immediate",
		"limited time",
		"act now",
	}

	for _, keyword := range phishingKeywords {
		if strings.Contains(emailLower, keyword) {
			score += 10
		}
	}

	// Check for suspicious links
	if strings.Contains(emailLower, "http") {
		score += 5
	}

	// Check for urgency markers
	if strings.Contains(emailLower, "!") {
		count := strings.Count(emailLower, "!")
		if count > 2 {
			score += count * 2
		}
	}

	// Check for monetary requests
	monetaryKeywords := []string{"payment", "bitcoin", "wire", "transfer", "credit card", "bank account"}
	for _, keyword := range monetaryKeywords {
		if strings.Contains(emailLower, keyword) {
			score += 15
		}
	}

	// Check for generic greetings (common in phishing)
	if strings.Contains(emailLower, "dear user") || strings.Contains(emailLower, "dear customer") {
		score += 10
	}

	// Determine risk level
	if score >= 55 {
		riskLevel = "High Risk"
	} else if score >= 30 {
		riskLevel = "Suspicious"
	}

	return score, riskLevel
}
