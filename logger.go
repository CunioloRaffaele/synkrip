package main

import (
	"log"
	"os"
	"synkrip/fsHandler"
)

func loggerSetup () {
	path,_ := fsHandler.GetAppDataPath()

	// Create a txt file inside the directory
	if _, err := os.Stat(path + "/log.txt"); os.IsNotExist(err) {
		// File doesn't exist, create it
		txtFile, err := os.Create(path + "/log.txt")
		if err != nil {
			log.Panic(err)
		}
		txtFile.Close()
	}


    // open log file
    logFile, err := os.OpenFile(path + "/log.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Panic(err)
    }
    //defer logFile.Close()

    // Set log out put and enjoy :)
    log.SetOutput(logFile)

    // optional: log date-time, filename, and line number
    log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("--------------------START-----------------------")
    log.Println("Starting program with logger set to custom file")
}