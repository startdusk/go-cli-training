package metadata

import (
	"audiofile/extractors/tags"
	"audiofile/extractors/transcript"
	"audiofile/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func (m *MetadataService) getByID(w http.ResponseWriter, r *http.Request) {
	value, ok := r.URL.Query()["id"]
	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'id' is missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := string(value[0])
	fmt.Println("requesting audio by id: ", id)

	audio, err := m.Storage.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no such file or directory") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	audioString, err := audio.JSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, audioString)
}

func (m *MetadataService) list(w http.ResponseWriter, r *http.Request) {
	audioFiles, err := m.Storage.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(audioFiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(jsonData), "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, prettyJSON.String())
}

func (m *MetadataService) upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("req: ", r)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		fmt.Println("error creating formfile: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		err = os.Remove(fileHeader.Filename)
		if err != nil {
			fmt.Println("error opening file: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		f.Close()
	}()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println("error copying file to buffer: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, audioFilePath, err := m.Storage.Upload(buf.Bytes(), fileHeader.Filename)
	audio := &models.Audio{
		Id:   id,
		Path: audioFilePath,
	}
	err = m.Storage.SaveMetadata(audio)
	if err != nil {
		fmt.Println("error saving metadata: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	audio.Status = "Initiating"

	io.WriteString(w, id)
	go func() {
		var errors []string

		audio.Status = "Complete"

		// tags
		err = tags.Extract(audio)
		if err != nil {
			fmt.Println("error extracting tags metadata: ", err)
			errors = append(errors, err.Error())
		}
		err = m.Storage.SaveMetadata(audio)
		if err != nil {
			fmt.Println("error saving metadata: ", err)
			errors = append(errors, err.Error())
		}

		// transcript
		err = transcript.Extract(audio)
		if err != nil {
			fmt.Println("error extracting transcript metadata: ", err)
			errors = append(errors, err.Error())
		}

		audio.Error = errors
		audio.Status = "Complete"
		err = m.Storage.SaveMetadata(audio)
		if err != nil {
			fmt.Println("error saving metadata: ", err)
			errors = append(errors, err.Error())
		}

		if len(errors) > 0 {
			fmt.Println("errors occurred extracting metadata: ")
			for i := 0; i < len(errors); i++ {
				fmt.Printf("\terror[%d]: %s\n", i, errors[i])
			}
		} else {
			fmt.Println("successfully extracted and saved audio metadata: ", audio)
		}
	}()
}
