package main

import (
    "io"
	"path/filepath"
	"fmt"
	"os"
	
	"golang.org/x/crypto/ssh"
	"github.com/urfave/cli"
    // "github.com/nsf/termbox-go"
)

var (
	cfgParser CfgParser
	gitCommit = ""
	app = NewApp(gitCommit, "Berith Test Manager")
)


func handleError(err error){
	if err != nil {
        if err == io.EOF{
            return;
        } 
        panic(err)
    }
}


func NewApp(gitCommit, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "jongyoungcha"
	app.Email = "jongyoungcha@gmail.com"
	app.Version = gitCommit

	return app
}


func jonitoringBerith(context *cli.Context) error {
	if cfgParser.Load() == false {
		fmt.Println("Error!!")
		os.Exit(1)
	}
	
    readBuff := make([]byte, 1000)
    fmt.Println(len(readBuff))
    
    for _, targetNode := range cfgParser.TargetNodes {
        _, session, err := connectToHost(targetNode)
        if err != nil {
            fmt.Println("Coudlnt connect to server...")
            fmt.Println(err)
            continue;
        }
        
        sshIn, err := session.StdinPipe()
        if err != nil {
            fmt.Println(err)
            return err
        }
        
        sshOut, err := session.StdoutPipe()
        if err != nil {
            fmt.Println(err)
            return err
        }
        
        err = session.Run("tail output")
        if err != nil {
            fmt.Println(err)
            return err
        }
        sshIn.Write([]byte("ls"))
        _, err = sshIn.Write([]byte("ls\n"))
        handleError(err)
        
        // for i := 0; i < 10; i++ {
        n, err := sshOut.Read(readBuff)
        if err != nil {
            fmt.Println(err)
            return err
        }

        fmt.Println(string(readBuff))
        fmt.Println("readed ", n)
    }
	
	return nil
}


func init() {
	app.Action = jonitoringBerith
	app.HideVersion = true
	app.Copyright = "Copyright 2019 by jongyoungcha"
	app.Commands = []cli.Command{
	}
}


func main() {
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}


func connectToHost(node Ethnode) (*ssh.Client, *ssh.Session, error){
    sshConfig := &ssh.ClientConfig{
        User: node.User,
        Auth: []ssh.AuthMethod{ssh.Password(node.Passwd)},
    }
    
    sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
    
    client, err := ssh.Dial("tcp", node.Host, sshConfig)
    if err != nil{
        return nil, nil, err
    }
    
    session, err := client.NewSession()
    if (err != nil) {
        client.Close()
        return nil, nil, err
    }
    
    return client, session, nil
}
