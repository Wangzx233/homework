package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)
// 需要先在终端输入docker-compose up -d启动docker里的kafka
func main(){
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"},nil)
	if err !=nil{
		log.Println(err)
		return
	}
	partitionList, err := consumer.Partitions("hello")// 根据topic取到所有的分区
	if err !=nil{
		log.Println(err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {// 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("hello", int32(partition), sarama.OffsetNewest)
		if err !=nil{
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer){
			for msg := range pc.Messages(){
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	select{} //阻塞进程
}