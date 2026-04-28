
package main 
import ("fmt")

type serverers interface {
	Privateserver(ser string)
}

type filsesystem struct {
	filestype string
}

func (self *filsesystem) Filesystemtype() {
	fmt.Println("The file system type is ", self.filestype)
}

/*
class storage extend filsesystem
{

}
*/

type storage struct {

	 filsesystem 
	 servername string
	 servconfig []string

	}

func (self *storage) Methods( ser string){
	fmt.Println("The server name is ", self.servername)
}

func (self *storage) Privateserver(ser string) {
	fmt.Println("The private server name is ", self.servername)
}



func main(){

	var ec2 storage

	ec2.servername = "aws"
	ec2.servconfig = []string{"t2.micro", "t2.small", "t2.medium"}
    ec2.filestype = "ext4"

    ec2.Methods("azzure")
    ec2.Filesystemtype()

    var corevalue serverers = 	&ec2
	corevalue.Privateserver("google cloud")

}