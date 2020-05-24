package services

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	//database"
	//m "grabber/models"
	"grabber/database"
	pb "grabber/pb"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
)

type FacultyServiceGrpcImpl struct {
	ctx context.Context
}

func NewFacultyServiceGrpcImpl(ctx context.Context) *FacultyServiceGrpcImpl {
	return &FacultyServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *FacultyServiceGrpcImpl) GetFaculty(ctx context.Context, in *pb.GetFacultyRequest) (*pb.Faculty, error) {
	var faculties *pb.Faculty
	return faculties, nil
}

func (serviceImpl *FacultyServiceGrpcImpl) ListFaculties(ctx context.Context, in *pb.ListFacultiesRequest) (*pb.ListFacultiesResponse, error) {
	var faculties *pb.ListFacultiesResponse

	db := database.FromContext(serviceImpl.ctx)
	academy, err := db.GetAcademy(in.AcademyId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, academy.Endpoint, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.URL.Path = "/Dek/Default.aspx"
	q := req.URL.Query()
	q.Add("mode", "facultet")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	log.WithFields(log.Fields{
		"method": req.Method,
		"query":  req.URL.RawQuery,
		"body":   req.Body,
	}).Debugf("Sending a request to %s", req.URL.RequestURI())
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Fix charset
	utf8, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		log.Error("IO error:", err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	table := doc.Find("table#ContentPage_ucFacultets_Grid").First()
	table.Find(".TblText").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")
		name := cols.Eq(0).Text()
		log.Debug(name)
	})

	return faculties, nil
}
