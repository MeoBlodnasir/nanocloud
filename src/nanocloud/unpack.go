package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func unpackGo(sourcefile string, runningPlugins []string) {
	time.Sleep(1000 * time.Millisecond) // TODO, Delete that and see when file is fully copied
	file, err := os.Open(sourcefile)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {

			log.Println(err)
			os.Exit(1)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			os.Exit(1)
		}

		// get the individual filename and extract to the current directory
		filename := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
		case tar.TypeReg:
			// handle normal file
			log.Println("Untarring :", filename)
			if !strings.Contains(filename, "/") {
				writer, err := os.OpenFile("plugins/staging/"+filename, os.O_WRONLY|os.O_CREATE, os.FileMode(header.Mode))

				if err != nil {
					log.Println(err)
					os.Exit(1)
				}

				io.Copy(writer, tarBallReader)

				//err = os.Chmod(filename, os.FileMode(header.Mode))

				if err != nil {
					log.Println(err)
					os.Exit(1)
				}

				writer.Close()
				runningPlugins = createEvent(runningPlugins, filename, conf.StagDir+filename, sourcefile)
			}
		default:
			log.Println("Unable to untar type :", header.Typeflag, " in file", filename)
		}
	}

}

func deleteOldFront(sourcefile string) {
	file, err := os.Open(sourcefile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	var fileReader io.ReadCloser = file
	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {

			log.Println(err)
			os.Exit(1)
		}
		defer fileReader.Close()
	}
	tarBallReader := tar.NewReader(fileReader)
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			os.Exit(1)
		}
		filename := header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			log.Println("Deleting directory :", conf.FrontDir+filename)
			err = os.RemoveAll(conf.FrontDir + filename) // or use 0755 if you prefer
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		case tar.TypeReg:
		default:
			log.Println("Unable to delete type :", header.Typeflag, " in file", filename)
		}
	}
}

func unpackFront(sourcefile string) {
	file, err := os.Open(sourcefile)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {

			log.Println(err)
			os.Exit(1)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			os.Exit(1)
		}

		// get the individual filename and extract to the current directory
		filename := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			log.Println("Creating directory :", conf.FrontDir+filename)
			err = os.MkdirAll(conf.FrontDir+filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

		case tar.TypeReg:
			// handle normal file
			log.Println("Untarring :", filename)
			writer, err := os.OpenFile(conf.FrontDir+filename, os.O_WRONLY|os.O_CREATE, os.FileMode(header.Mode))

			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			io.Copy(writer, tarBallReader)

			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			writer.Close()
		default:
			log.Println("Unable to untar type :", header.Typeflag, " in file", filename)
		}
	}

}