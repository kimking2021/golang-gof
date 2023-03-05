package proxy

// 代理模式 (Proxy Pattern)，可以为另一个对象提供一个替身或者占位符，以控制对这个对象的访问。

import "fmt"

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (s *Station) SetStock(stock int) {
	s.stock = stock
}

func (s *Station) sell(name string) {
	if s.stock > 0 {
		s.stock--
		fmt.Printf("%s sell a ticket, reserve: %d", name, s.stock)
	} else {
		fmt.Println("no ticket for sale")
	}
}

type StationProxy struct {
	station *Station
}

func (sp *StationProxy) sell(name string) {
	if sp.station.stock > 0 {
		sp.station.stock--
		fmt.Printf("%s sell a ticket, left: %d", name, sp.station.stock)
	} else {
		fmt.Println("no ticket for sale")
	}

}

func NewStationProxy(stock int) *StationProxy {
	return &StationProxy{
		station: &Station{stock: stock},
	}
}
