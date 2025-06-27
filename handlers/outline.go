package handlers

import (
	"course-api/constants"
	"course-api/utils"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func OutlineHandler(c *gin.Context) {
	term, course := c.Param("term"), strings.ToUpper(c.Param("course"))
	unpublished := strings.ToLower(c.Query("unpublished")) == "true"

	fmt.Println("Received request for outline:", term, course, "Unpublished:", unpublished)

	outline_link := fmt.Sprintf(constants.OutlineUrl, term, course)

	if unpublished {
		outline_link = fmt.Sprintf("%s?unpt=t", outline_link)
	}

	roots, err := x509.SystemCertPool()

	if err != nil {
		log.Fatal(err)
	}

	certs, err := os.ReadFile("/app/server_cert.pem")

	if err != nil {
		log.Fatal(err)
	}

	if ok := roots.AppendCertsFromPEM(certs); !ok {
		log.Println("Warning: no certs appended, check PEM file")
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: roots,
			},
		},
	}

	resp, err := client.Get(outline_link)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.WriteError(c, fmt.Sprintf("Failed to read response body: %v", err))
		return
	}

	bodyLen := len(string(body))

	utils.WriteSuccess(c, gin.H{
		"link":    outline_link,
		"isValid": bodyLen > 2000,
	})
}
