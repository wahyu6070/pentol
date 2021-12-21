package main

import (
//"fmt"
"os"
"os/exec"
"syscall"
"PENTOL/lib"
"log"
"path/filepath"
)
func umount(src string) error{
	cmd := exec.Command("umount", src)
	
	err := cmd.Run()
	if err != nil {
		return err
		}
	return nil

	}
func BIND(INPUT string, OUTPUT string) error{
	cmd := exec.Command("mount", "-o bind", INPUT, OUTPUT)
	cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
		}
	
	return err
	}
func CONFIG() string{
	LISTF := map[string]string{
		".bashrc" : "root/.bashrc",
		".profile" : "root/.profile",
		"hostname" : "etc/hostname",
		"username" : "etc/username",
		"osinfo" : "etc/osinfo",
		}
	for A, B := range LISTF {
		IN := filepath.Join(BASE, "config", A)
		OUT := filepath.Join(PENTOL, B)
		shell.Copy(IN, OUT)
	}
	
	return ``
	}
func BIND_DIR() string{
	LIST_DIR := map[string]string{
		"/system" : "system",
		"/product" : "product",
		"/system_ext" : "system_ext",
		"/vendor"  : "vendor",
		"/system/bin" : "xbin",
		
		
		}
	for A, B := range LIST_DIR {
		if _, err := os.Stat(A); err == nil {
			TO := filepath.Join(PENTOL, B)
        	if _, err := os.Stat(TO); err != nil { 
				os.MkdirAll(TO, 0755)
			}
			if _, err := os.Stat(TO); err == nil { 
				umount(TO)
			}
			BIND(A, filepath.Join(PENTOL,B))
        }
			
	}
	return ` `
	}
var (
BASE string
PENTOL string
)
func main(){
	BASE = shell.IS_DIRNAME()
	PENTOL = filepath.Join(BASE, "root")
	BIND_DIR()
	CONFIG()
	err := syscall.Exec("/system/bin/sh", []string{
		"sh",
		filepath.Join(BASE, "start"),
	}, syscall.Environ())
	if err != nil {
		log.Fatal(err)
		}
	
	}
	