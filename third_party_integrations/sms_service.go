package third_party_integrations
import "fmt"

type SmsService struct{

}

func (s SmsService) Send(id string, msg string){
	fmt.Println("msg: " + msg)
	fmt.Println("sms sending ...")
}


func (s SmsService) BulkSend(ids []string, msg string){
	fmt.Println("msg: " + msg)
	fmt.Println("bulk sms sending ...")
}
