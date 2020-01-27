package third_party_integrations
import "fmt"

type FirebaseService struct{

}

func (s FirebaseService) Send(id string, msg string){
	fmt.Println("msg: " + msg)
	fmt.Println("push notification sending ...")
}


func (s FirebaseService) BulkSend(ids []string, msg string){
	fmt.Println("msg: " + msg)
	fmt.Println("bulk push notification  sending ...")
}
