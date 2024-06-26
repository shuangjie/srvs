package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"srvs/inventory_srv/proto"
	"sync"
)

var invClient proto.InventoryClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	invClient = proto.NewInventoryClient(conn)
}

func TestSetInv(goodsId, Num int32) {
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("SetInv success")
}

func TestInvDetail(goodsId int32) {
	resp, err := invClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("InvDetail goodsId:%d, num:%d\n", resp.GoodsId, resp.Num)
}

func TestSell(wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 1},
			//{GoodsId: 422, Num: 40},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Sell success")
}

func TestReBack() {
	_, err := invClient.ReBack(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 10},
			{GoodsId: 422, Num: 40},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("ReBack success")
}

func main() {
	Init()

	//TestSetInv(422, 40)
	//TestInvDetail(422)
	//TestSell()
	//TestReBack()

	//批量生成库存，goodsId范围是421-840
	//for i := 421; i <= 840; i++ {
	//	TestSetInv(int32(i), 100)
	//}

	// 模拟并发扣减库存
	var wg sync.WaitGroup
	//wg.Add(30)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go TestSell(&wg)
	}
	wg.Wait()

	conn.Close()
}
