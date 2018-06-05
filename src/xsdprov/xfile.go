package xsdprov

import (
	"archive/zip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	newFilePtr  *os.File
	fileInfo    os.FileInfo
	fileInfoPtr *os.FileInfo
	err         error
)

// DownloadFile ...
func DownloadFile(filepath string, w http.ResponseWriter) error {
	w.Header().Set("Content-type", "application/zip")
	fp, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	check(err)
	defer fp.Close()
	_, cerr := io.Copy(w, fp)
	check(cerr)
	return err
}

// CopyDirs ...
func CopyDirs(dest string, resdirectories map[string]string) {
	// Copy XML dependencies
	for _, d := range resdirectories {
		err = CopyDir(d, dest+d)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// CopyDir  ...
func CopyDir(source string, dest string) (err error) {
	// get properties of source dir
	//log.Println(source)
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}
	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()
		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}

// CopyFile  ...
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
func getFile(fname string) (fp *os.File, errr error) {
	fp, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		errr = err
		panic(err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			errr = err
			panic(err)
		}
	}()
	return fp, errr
}
func mkTempDir(dname string) string {
	tempDirPath, err := ioutil.TempDir("", dname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp dir created:", tempDirPath)
	return tempDirPath
}
func writeFile(fname string, data []byte) error {
	err := ioutil.WriteFile(fname, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	return err
}
func getDigests(fileslist map[string]string, dir string, algo string) map[string]string {
	log.Println("getDigests")
	var digests = map[string]string{}
	for r, res := range fileslist {
		var path = dir + res
		//log.Println(path)
		if fileExists(path) {
			digests[r] = GetHash(path, algo)
		}
	}
	return digests
}
func wgetFile(fpath string, urlstr string) error {
	log.Println("Wget Save To: " + fpath)
	// Create output dir
	p := filepath.Dir(fpath)
	os.MkdirAll(p, os.ModePerm)
	newFile, err := os.Create(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	// HTTP GET
	response, err := http.Get(urlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	return err
}
func hashFile(fname string, algo string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	switch algo {
	case "Md5":
		return fmt.Sprint(md5.Sum(data))
	case "Sha1":
		return fmt.Sprint(sha1.Sum(data))
	case "Sha256":
		return fmt.Sprint(sha256.Sum256(data))
	case "Sha512":
		return fmt.Sprint(sha512.Sum512(data))
	}
	return "Error"
}

//GetHash ...
func GetHash(fname string, algo string) string {
	file, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hasher := getHasher(algo)
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}
	sum := hasher.Sum(nil)
	//log.Println(fmt.Sprint(sum))
	return hex.EncodeToString(sum)
}
func getStrHash(data string, algo string) string {
	r := strings.NewReader(data)
	hasher := getHasher(algo)
	_, err = io.Copy(hasher, r)
	if err != nil {
		log.Fatal(err)
	}
	sum := hasher.Sum(nil)
	//log.Println(fmt.Sprint(sum))
	return hex.EncodeToString(sum)
}
func getHasher(algo string) hash.Hash {
	switch algo {
	case "Md5":
		return md5.New()
	case "Sha1":
		return sha1.New()
	case "Sha256":
		return sha256.New()
	case "Sha512":
		return sha512.New()
	}
	return sha256.New()
}
func ddgzip(dir string, dest string) {
	outFile, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	zipWriter := zip.NewWriter(outFile)
	addFiles(zipWriter, dir, "")
	if err != nil {
		fmt.Println(err)
	}

	err = zipWriter.Close()
	if err != nil {
		fmt.Println(err)
	}
}
func compress(dir string, dest string) error {
	log.Println("dir " + dir)
	log.Println("dest " + dest)
	p := filepath.Dir(dest)
	n := filepath.Base(dir)
	z := filepath.Base(dest)
	log.Println("runcmd " + "cd " + p + "&& /usr/bin/zip -r " + z + " " + n)
	err := runcmd("cd "+p+"&& zip -r "+z+" "+n, true)
	return err
	//log.Println(resultstring)
}
func runcmd(cmd string, shell bool) error {
	if shell {
		_, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return err
	}
	_, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, file.Name()+"/")
		}
	}
}
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
