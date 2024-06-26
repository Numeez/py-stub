package main

import (
	"os"
)
const four_spaces = "    "
func CreatePyFile(filename string) error {
	stub_code := []string{
		"def main():\n",
		four_spaces,
		"pass",
		"\n\n\n\n\n",
		`if __name__=="__main__":`,
		"\n",
		four_spaces,
		"main()",
	}
	file,err:= os.Create("stub.py")
	if err !=nil{
		return err
	}
	for _,code:=range stub_code{
		file.WriteString(code)
	}
	if err !=nil{
		return err
	}
	file.Close()
	return nil
}