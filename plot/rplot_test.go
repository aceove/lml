package plot

import (
	"testing"
	"os/exec"
	"os"
	"syscall"
	"fmt"
)

func TestRPlot(t *testing.T) {


	/*
	binary, lookErr := exec.LookPath("R")
	if lookErr != nil {
		panic(lookErr)
	}
	//Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"--no-save"}
	//Exec 同样需要使用环境变量。这里我们仅提供当前的环境变量。
	env := os.Environ()
	//这里是 os.Exec 调用。如果这个调用成功，那么我们的进程将在这里被替换成 /bin/ls -a -l -h 进程。如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		//panic(execErr)
		fmt.Println(execErr)
	}

	*/
	/*
	cmd := exec.Command("Rcmd", "demo()")
	buf, err := cmd.Output()
	fmt.Printf("%s\n%s",buf,err)
	*/

	/*
	arg := []string{"--no-save"}
	grepCmd := exec.Command("R", arg...)
	grepIn, err := grepCmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	grepOut, err := grepCmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	grepCmd.Start()
	grepIn.Write([]byte("require('rgl') ;open3d() ;	x <- sort(rnorm(1000)) ;	y <- rnorm(1000) ;	z <- rnorm(1000) + atan2(x, y) ;plot3d(x, y, z, col = rainbow(1000));"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepBytes))
	*/

	/*
	grepCmd := exec.Command("grep", "hello")
	//这里我们明确的获取输入/输出管道，运行这个进程，写入一些输入信息，读取输出的结果，最后等待程序运行结束。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	*/

	/*
	c:="require('rgl') ;open3d() ;	x <- sort(rnorm(1000)) ;	y <- rnorm(1000) ;	z <- rnorm(1000) + atan2(x, y) ;	plot3d(x, y, z, col = rainbow(1000));"
	n, err := io.WriteString(stdin, c)
	fmt.Printf("cmd> %v", cmd)
	fmt.Printf("res> %v\n", n)
	stdin.Close()
	cmd.Wait()
	*/
	/*
	c="plot(sin, -pi, 2*pi) "
	n, err = io.WriteString(stdin, c)
	fmt.Printf("cmd> %v", cmd)
	fmt.Printf("res> %v\n", n)
	//stdin.Close()
	//cmd.Wait()
	*/

	/*
	arg := []string{"--no-save"}
	cmd := exec.Command("R", arg...)
	//会向 cmd.Stdout和cmd.Stderr写入信息,其实cmd.Stdout==cmd.Stderr,具体可见源码
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("callEXE2结果:\n%v\n\n%v\n\n%v", string(output), cmd.Stdout, cmd.Stderr)
	*/

	/*
	cmd := exec.Command("C:\\Program Files\\R\\R-3.4.3\\bin\\x64\\Rscript.exe", "demo()")
	buf, err := cmd.Output()
	fmt.Printf("%s\n%s",buf,err)
*/

/*
	fname := ""
	persist := false
	debug := true
	p, err := NewRPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.RClose()
	//p.RCheckedCmd("require('rgl') ")
	//p.RCheckedCmd("open3d()")
	//p.RCheckedCmd("x <- sort(rnorm(1000))")
	//p.RCheckedCmd("y <- rnorm(1000)")
	//p.RCheckedCmd("z <- rnorm(1000) + atan2(x, y)")
	//p.RCheckedCmd("plot3d(x, y, z, col = rainbow(1000))")
	p.RCheckedCmd("require('rgl') ;open3d() ;	x <- sort(rnorm(1000)) ;	y <- rnorm(1000) ;	z <- rnorm(1000) + atan2(x, y) ;	plot3d(x, y, z, col = rainbow(1000))")
	p.RCheckedCmd("q()")
*/

}
