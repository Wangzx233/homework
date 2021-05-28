package main
import(
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)
// 需要先在终端输入docker-compose up -d启动docker里的kafka
func main(){
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, nil)
	if err !=nil{
		log.Println(err)
		return
	}
	defer client.Close()
	// 消息
	msg :=&sarama.ProducerMessage{}
	msg.Topic="hello"
	msg.Value= sarama.StringEncoder("Hello world")
	pid, offset, err := client.SendMessage(msg)
	if err !=nil{
		log.Println(err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}