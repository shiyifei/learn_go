package practice

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	cmd0 := exec.Command("echo", "-n", "What are you doing now? shiyf. I am learning golang")
	stdout0,err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command No.0:%s\n", err)
	}

	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error:The command No.0 can not be startup:%s\n", err)
	}

	var outputBuf0 bytes.Buffer
	//内容过多，需要多次读取，保证读取完管道中的所有内容才行
	for {
		output := make([]byte, 20)
		n,err := stdout0.Read(output)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error: Can not read data from th pipe: %s\n", err)
				return
			}
		}
		if n>0 {
			outputBuf0.Write(output[:n])  //将每次读取的内容写入到缓冲区里
		}
	}
	fmt.Printf("%s\n", outputBuf0.String()) //将缓冲区字节转化为字符串

	/*//一个更加方便的方法是，一开始就使用带缓冲的读取器从输出管道中读取数据，缓冲区长度默认：4096 如下代码
	outputBuf1 := bufio.NewReader(stdout0)
	output0,result,err := outputBuf1.ReadLine()
	if err != nil {
		fmt.Printf("Error: Can not read data from the pipe:%s\n", err)
		return
	}
	fmt.Println(result) //result表示是否还未被读完
	fmt.Printf("%s\n", string(output0))*/

}
